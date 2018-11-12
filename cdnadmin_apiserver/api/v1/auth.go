package v1

import (
	"fmt"
	"net/http"

	"cdn-upload-admin/cdnadmin_apiserver/models/v1"

	"github.com/dgrijalva/jwt-go"
)

// UserAuth : user login auth
func UserAuth(w http.ResponseWriter, r *http.Request) {
	// set response header json
	w.Header().Set("Content-Type", "application/json; charset=UTF-8")

	var user models.UserCredentials
	r.ParseForm()
	user.UserName = r.Form.Get("username")
	user.Password = r.Form.Get("password")
	// err := json.NewDecoder(r.Body).Decode(&user)
	if user.UserName == "" || user.Password == "" {
		models.ResponseWithJSON(w, http.StatusOK,
			models.Response{Code: http.StatusBadRequest, Msg: "bad params"})
		return
	}
	data, err := models.GetToken(&user)
	if err != nil {
		models.ResponseWithJSON(w, http.StatusOK,
			models.Response{Code: http.StatusBadRequest, Msg: err.Error()})
	}
	w.Write([]byte(data))
}

// TokenMiddleware : CHECK AUTH TOKEN
func TokenMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		tokenStr := r.Header.Get("authorization")
		if tokenStr == "" {
			models.ResponseWithJSON(w, http.StatusOK,
				models.Response{Code: http.StatusUnauthorized, Msg: "not authorized"})
		} else {
			token, _ := jwt.Parse(tokenStr, func(token *jwt.Token) (interface{}, error) {
				if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
					models.ResponseWithJSON(w, http.StatusOK,
						models.Response{Code: http.StatusUnauthorized, Msg: "not authorized"})
					return nil, fmt.Errorf("not authorization")
				}
				return []byte(models.TokenSecretKey), nil
			})
			if !token.Valid {
				models.ResponseWithJSON(w, http.StatusOK,
					models.Response{Code: http.StatusUnauthorized, Msg: "not authorized"})
			} else {
				next.ServeHTTP(w, r)
			}
		}
	})
}
