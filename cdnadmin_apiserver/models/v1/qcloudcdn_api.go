package models

import (
	"crypto/hmac"
	"crypto/sha1"
	"crypto/tls"
	"encoding/base64"
	"fmt"
	"io/ioutil"
	"math/rand"
	"net/http"
	"net/url"
	"sort"
	"strconv"
	"strings"
	"time"
)

/***
	qcloud cdn openapi
	author:evincai@tencent.com
***/

/**
	*@brief        qcloud cdn openapi signature
	*@param        secretKey    secretKey to log in qcloud
	*@param        params       params of qcloud openapi interface
	*@param        method       http method
	*@param        requesturl   url
	*@return       Signature    signature
	*@return       params       params of qcloud openapi interfac include Signature
**/

// Signature Signature
func Signature(secretKey string, params map[string]interface{}, method string, requesturl string) (string, map[string]interface{}) {
	/*add common params*/
	timestamp := time.Now().Unix()
	rd := rand.New(rand.NewSource(time.Now().UnixNano())).Intn(1000)
	params["Timestamp"] = timestamp
	params["Nonce"] = rd
	/**sort all the params to make signPlainText**/
	sigURL := method + requesturl + "?"
	var keys []string
	for k := range params {
		keys = append(keys, k)
	}
	sort.Strings(keys)
	isfirst := true
	for _, key := range keys {
		if !isfirst {
			sigURL = sigURL + "&"
		}
		isfirst = false
		if strings.Contains(key, "_") {
			strings.Replace(key, ".", "_", -1)
		}
		value := typeSwitcher(params[key])
		sigURL = sigURL + key + "=" + value
	}
	fmt.Println("signPlainText: ", sigURL)
	unencodeSign, _sign := sign(sigURL, secretKey)
	params["Signature"] = unencodeSign
	fmt.Println("unencoded signature: ", unencodeSign)
	return _sign, params
}

/**
	*@brief        send request to qcloud
	*@param        params       params of qcloud openapi interface include signature
	*@param        method       http method
	*@param        requesturl   url
	*@return       Signature    signature
	*@return       params       params of qcloud openapi interfac include Signature
**/

// SendRequest : send request
func SendRequest(requesturl string, params map[string]interface{}, method string) string {
	requesturl = "https://" + requesturl
	var response string
	if method == "GET" {
		paramsStr := "?" + ParamsToStr(params)
		requesturl = requesturl + paramsStr
		response = httpGet(requesturl)
	} else if method == "POST" {
		response = httpPost(requesturl, params)
	} else {
		fmt.Println("unsuppported http method")
		return "unsuppported http method"
	}
	return response
}

func typeSwitcher(t interface{}) string {
	switch v := t.(type) {
	case int:
		return strconv.Itoa(v)
	case string:
		return v
	case int64:
		return strconv.Itoa(int(v))
	default:
		return ""
	}
}

// ParamsToStr : params to str
func ParamsToStr(params map[string]interface{}) string {
	isfirst := true
	requesturl := ""
	for k, v := range params {
		if !isfirst {
			requesturl = requesturl + "&"
		}
		isfirst = false
		if strings.Contains(k, "_") {
			strings.Replace(k, ".", "_", -1)
		}
		v := typeSwitcher(v)
		requesturl = requesturl + k + "=" + url.QueryEscape(v)
	}
	return requesturl
}

func sign(signPlainText string, secretKey string) (string, string) {
	key := []byte(secretKey)
	hash := hmac.New(sha1.New, key)
	hash.Write([]byte(signPlainText))
	sig := base64.StdEncoding.EncodeToString([]byte(string(hash.Sum(nil))))
	encdSig := url.QueryEscape(sig)
	return sig, encdSig
}

func httpGet(url string) string {
	tr := &http.Transport{
		TLSClientConfig: &tls.Config{InsecureSkipVerify: true},
	}
	client := &http.Client{Transport: tr, Timeout: time.Duration(3) * time.Second}
	fmt.Println(url)
	resp, err := client.Get(url)
	if err != nil {
		fmt.Println(err)
		return err.Error()
	}
	defer resp.Body.Close()
	body, erro := ioutil.ReadAll(resp.Body)
	if erro != nil {
		fmt.Println("http wrong erro")
		return erro.Error()
	}
	return string(body)
}

func httpPost(requesturl string, params map[string]interface{}) string {
	req, err := http.NewRequest("POST", requesturl, strings.NewReader(ParamsToStr(params)))
	req.Header.Set("Content-Type", "application/x-www-form-urlencoded; param=value")
	/*
		req, err := http.NewRequest("POST", requesturl, strings.NewReader(form.Encode()))
		fmt.Println(strings.NewReader(form.Encode()))
	*/
	tr := &http.Transport{
		TLSClientConfig: &tls.Config{InsecureSkipVerify: true},
	}
	client := &http.Client{Transport: tr, Timeout: time.Duration(3) * time.Second}
	resp, err := client.Do(req)
	if err != nil {
		fmt.Println(err)
		return err.Error()
	}
	defer resp.Body.Close()
	body, erro := ioutil.ReadAll(resp.Body)
	if erro != nil {
		fmt.Println("http wrong erro")
		return erro.Error()
	}
	return string(body)
}
