package v1

import (
	"net/http"
	"strconv"

	"cdn-upload-admin/cdnadmin_apiserver/models/v1"
)

// GetUserList : get user list
func GetUserList(w http.ResponseWriter, r *http.Request) {
	var u models.User
	// set response header json
	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	r.ParseForm()
	num, _ := strconv.Atoi(r.Form.Get("page_num"))
	size, _ := strconv.Atoi(r.Form.Get("page_size"))
	u.UserName = r.Form.Get("username")
	if num < 1 || size < 1 {
		num = 0
		size = 0
	}

	data, err := models.GetUserList(num, size, &u)
	if err != nil {
		models.ResponseWithJSON(w, http.StatusOK,
			models.Response{Code: 4000, Msg: err.Error()})
		return
	}
	w.Write([]byte(data))
}

// AddUser : create new user
func AddUser(w http.ResponseWriter, r *http.Request) {
	var u models.User
	// set response header json
	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	r.ParseForm()
	u.Name = r.Form.Get("name")
	u.UserName = r.Form.Get("username")
	u.Password = r.Form.Get("password")
	u.Status = r.Form.Get("status")

	if u.Name == "" || u.UserName == "" || u.Password == "" || u.Status == "" {
		models.ResponseWithJSON(w, http.StatusOK,
			models.Response{Code: 4000, Msg: "bad params"})
		return
	}
	data, err := models.AddUser(&u)
	if err != nil {
		models.ResponseWithJSON(w, http.StatusOK,
			models.Response{Code: 4000, Msg: err.Error()})
		return
	}
	w.Write([]byte(data))
}

// UpdateUser : update user info
func UpdateUser(w http.ResponseWriter, r *http.Request) {
	var u models.User
	// set response header json
	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	r.ParseForm()
	u.ID, _ = strconv.Atoi(r.Form.Get("id"))
	u.Name = r.Form.Get("name")
	u.Status = r.Form.Get("status")

	if u.ID < 1 || u.Name == "" || u.Status == "" {
		models.ResponseWithJSON(w, http.StatusOK,
			models.Response{Code: 4000, Msg: "bad params"})
		return
	}
	data, err := models.UpdateUser(&u)
	if err != nil {
		models.ResponseWithJSON(w, http.StatusOK,
			models.Response{Code: 4000, Msg: err.Error()})
		return
	}
	w.Write([]byte(data))
}

// UpdateUserPwd : update user password
func UpdateUserPwd(w http.ResponseWriter, r *http.Request) {
	var u models.User
	// set response header json
	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	r.ParseForm()
	u.ID, _ = strconv.Atoi(r.Form.Get("id"))
	u.OldPassword = r.Form.Get("old_password")
	u.Password = r.Form.Get("password")

	if u.ID < 1 || u.OldPassword == "" || u.Password == "" {
		models.ResponseWithJSON(w, http.StatusOK,
			models.Response{Code: 4000, Msg: "bad params"})
		return
	}
	data, err := models.UpdateUserPwd(&u)
	if err != nil {
		models.ResponseWithJSON(w, http.StatusOK,
			models.Response{Code: 4000, Msg: err.Error()})
		return
	}
	w.Write([]byte(data))
}

// ResetUserPwd : reset user password
func ResetUserPwd(w http.ResponseWriter, r *http.Request) {
	var u models.User
	// set response header json
	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	r.ParseForm()
	u.ID, _ = strconv.Atoi(r.Form.Get("id"))

	if u.ID < 1 {
		models.ResponseWithJSON(w, http.StatusOK,
			models.Response{Code: 4000, Msg: "bad params"})
		return
	}
	data, err := models.ResetUserPwd(&u)
	if err != nil {
		models.ResponseWithJSON(w, http.StatusOK,
			models.Response{Code: 4000, Msg: err.Error()})
		return
	}
	w.Write([]byte(data))
}

// SetUserAdmin : set user is admin or not
func SetUserAdmin(w http.ResponseWriter, r *http.Request) {
	var u models.User
	// set response header json
	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	r.ParseForm()
	u.ID, _ = strconv.Atoi(r.Form.Get("id"))
	u.IsAdmin, _ = strconv.Atoi(r.Form.Get("is_admin"))
	if u.ID < 1 || u.IsAdmin < 0 {
		models.ResponseWithJSON(w, http.StatusOK,
			models.Response{Code: 4000, Msg: "bad params"})
		return
	}
	data, err := models.SetUserAdmin(&u)
	if err != nil {
		models.ResponseWithJSON(w, http.StatusOK,
			models.Response{Code: 4000, Msg: err.Error()})
		return
	}
	w.Write([]byte(data))
}

// DeleteUser : delete user
func DeleteUser(w http.ResponseWriter, r *http.Request) {
	var u models.User
	// set response header json
	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	r.ParseForm()
	u.ID, _ = strconv.Atoi(r.Form.Get("id"))

	if u.ID < 1 {
		models.ResponseWithJSON(w, http.StatusOK,
			models.Response{Code: 4000, Msg: "bad params"})
		return
	}
	data, err := models.DeleteUser(&u)
	if err != nil {
		models.ResponseWithJSON(w, http.StatusOK,
			models.Response{Code: 4000, Msg: err.Error()})
		return
	}
	w.Write([]byte(data))
}
