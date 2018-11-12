package models

import (
	"cdn-upload-admin/cdnadmin_apiserver/config"
	"encoding/json"
	"net/http"
	"time"

	"github.com/dgrijalva/jwt-go"
	jsoniter "github.com/json-iterator/go"
)

// GetToken : GET TOKEN
func GetToken(user *UserCredentials) (content []byte, err error) {
	// query user
	db := config.GetDB()
	var queryUser UserCredentials
	var u UserAuthRes
	sql := "SELECT `username`, `password`, `is_admin` FROM t_user WHERE `username` = ?"
	err = db.QueryRow(sql, user.UserName).Scan(&u.UserName, &queryUser.Password, &u.IsAdmin)
	if err != nil {
		return
	}
	if user.Password != queryUser.Password {
		content, err = jsoniter.Marshal(Response{Code: 4000, Msg: "password is wrong"})
		return
	}
	// get token
	token := jwt.New(jwt.SigningMethodHS256)
	claims := make(jwt.MapClaims)
	claims["username"] = user.UserName
	claims["exp"] = time.Now().Add(time.Hour * time.Duration(6)).Unix()
	claims["iat"] = time.Now().Unix()
	token.Claims = claims
	u.Token, err = token.SignedString([]byte(TokenSecretKey))
	if err != nil {
		return
	}
	content, err = jsoniter.Marshal(Response{Code: 200, Msg: "user auth success", Data: u})
	return
}

// ResponseWithJSON : RESPONSE WITH JSON
func ResponseWithJSON(w http.ResponseWriter, code int, payload interface{}) {
	response, _ := json.Marshal(payload)
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(code)
	w.Write(response)
}
