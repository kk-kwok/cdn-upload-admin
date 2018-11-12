package models

import (
	"cdn-upload-admin/cdnadmin_apiserver/config"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
	"strconv"
	"strings"

	jsoniter "github.com/json-iterator/go"
)

// CDNPush : CDN FILE PUSH
func CDNFilePush(cfId int) (content []byte, err error) {
	var push CDNPush
	db := config.GetDB()
	sql := "SELECT a.`cdn_id`, a.`domain`, a.`path`, b.`file_name`, b.`file_size`," +
		" c.`name`, c.`secret_id`, c.`secret_key`, c.`api_url`, c.`action`" +
		"FROM (t_project a INNER JOIN t_cdn_file b ON a.`id` = b.`project_id`)" +
		"INNER JOIN t_cdn_platform c ON a.`platform_id` = c.`id` WHERE b.`id`=?;"
	err = db.QueryRow(sql, cfId).Scan(&push.CDNId, &push.Domain, &push.Path,
		&push.FileName, &push.FileSize, &push.CDNName, &push.SecretID, &push.SecretKey,
		&push.APIURL, &push.Action)
	if err != nil {
		return
	}
	if push.CDNName == "越南资源服" || push.CDNName == "泰国资源服" {
		content, err = jsoniter.Marshal(Response{Code: 200, Msg: "IDC资源服无需推送CDN"})
	} else if push.SecretID == "" || push.SecretKey == "" || push.APIURL == "" || push.Action == "" {
		content, err = jsoniter.Marshal(Response{Code: 4000, Msg: "该平台未填写验证信息"})
	} else {
		if push.CDNName == "CDN77" {
			content, err = CDN77Push(&push)
		} else if push.CDNName == "QCLOUD" {
			content, err = QcloudPush(&push)
		} else {
			content, err = jsoniter.Marshal(Response{Code: 4000, Msg: "未知平台"})
		}
	}
	return
}

// CDN77Push : CDN77 PUSH
func CDN77Push(push *CDNPush) (content []byte, err error) {
	params := make(url.Values)
	params["cdn_id"] = []string{strconv.Itoa(push.CDNId)}
	params["login"] = []string{push.SecretID}
	params["passwd"] = []string{push.SecretKey}
	params["url[]"] = []string{"/" + push.Path + "/" + push.FileName}

	res, err := http.PostForm(push.APIURL+"/"+push.Action, params)
	if err != nil {
		return
	}
	defer res.Body.Close()
	body, _ := ioutil.ReadAll(res.Body)

	var data map[string]interface{}
	if err := jsoniter.Unmarshal(body, &data); err == nil {
		if data["status"] == "ok" {
			content, err = jsoniter.Marshal(Response{Code: 200, Msg: "推送CDN文件成功, 视文件大小大概需要30分钟生效"})
		} else {
			content, err = jsoniter.Marshal(Response{Code: 4000, Msg: string(body)})
		}
	} else {
		content, err = jsoniter.Marshal(Response{Code: 4000, Msg: "error response body"})
	}
	fmt.Println(string(body))
	return
}

// QcloudPush : QCLOUD PUSH
func QcloudPush(push *CDNPush) (content []byte, err error) {
	var Requesturl string = push.APIURL
	var SecretKey string = push.SecretKey
	var Method string = "POST"

	params := make(map[string]interface{})
	params["SecretId"] = push.SecretID
	params["Action"] = push.Action
	if !strings.HasPrefix(push.Domain, "http://") {
		push.Domain = "http://" + push.Domain
	}
	params["urls.0"] = push.Domain + "/" + push.Path + "/" + push.FileName

	signature, request_params := Signature(SecretKey, params, Method, Requesturl)
	fmt.Println("signature : ", signature)
	fmt.Println("request_params : ", request_params)
	res := SendRequest(Requesturl, request_params, Method)

	var data map[string]interface{}
	if err := jsoniter.Unmarshal([]byte(res), &data); err == nil {
		if data["codeDesc"] == "Success" {
			content, err = jsoniter.Marshal(Response{Code: 200, Msg: "推送CDN文件成功, 视文件大小大概需要30分钟生效"})
		} else {
			content, err = jsoniter.Marshal(Response{Code: 4000, Msg: res})
		}
	} else {
		content, err = jsoniter.Marshal(Response{Code: 4000, Msg: "error response body"})
	}
	fmt.Println(res)
	return
}
