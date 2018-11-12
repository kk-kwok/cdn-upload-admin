package v1

import (
	"net/http"
	"strconv"

	"cdn-upload-admin/cdnadmin_apiserver/models/v1"
)

// GetPlatformList : GET PLATFORM LIST
func GetPlatformList(w http.ResponseWriter, r *http.Request) {
	var plat models.Platform
	// set response header json
	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	r.ParseForm()
	num, _ := strconv.Atoi(r.Form.Get("page_num"))
	size, _ := strconv.Atoi(r.Form.Get("page_size"))
	plat.Name = r.Form.Get("name")
	plat.Status = r.Form.Get("status")
	if num < 1 || size < 1 {
		num = 0
		size = 0
	}

	data, err := models.GetPlatformList(num, size, &plat)
	if err != nil {
		models.ResponseWithJSON(w, http.StatusOK,
			models.Response{Code: 4000, Msg: err.Error()})
		return
	}
	w.Write([]byte(data))
}

// AddPlatform : CREATE PLATFORM
func AddPlatform(w http.ResponseWriter, r *http.Request) {
	var plat models.Platform
	// set response header json
	w.Header().Set("Content-Type", "application/json; charset=UTF-8")

	r.ParseForm()
	plat.Name = r.Form.Get("name")
	plat.SecretID = r.Form.Get("secret_id")
	plat.SecretKey = r.Form.Get("secret_key")
	plat.APIUrl = r.Form.Get("api_url")
	plat.Action = r.Form.Get("action")
	plat.Status = r.Form.Get("status")

	if plat.Name == "" {
		models.ResponseWithJSON(w, http.StatusOK,
			models.Response{Code: 4000, Msg: "bad params"})
		return
	}
	data, err := models.AddPlatform(&plat)
	if err != nil {
		models.ResponseWithJSON(w, http.StatusOK,
			models.Response{Code: 4000, Msg: err.Error()})
		return
	}
	w.Write([]byte(data))
}

// UpdatePlatform : UPDATE PLATFORM
func UpdatePlatform(w http.ResponseWriter, r *http.Request) {
	var plat models.Platform
	// set response header json
	w.Header().Set("Content-Type", "application/json; charset=UTF-8")

	r.ParseForm()
	plat.ID, _ = strconv.Atoi(r.Form.Get("id"))
	plat.Name = r.Form.Get("name")
	plat.SecretID = r.Form.Get("secret_id")
	plat.SecretKey = r.Form.Get("secret_key")
	plat.APIUrl = r.Form.Get("api_url")
	plat.Action = r.Form.Get("action")
	plat.Status = r.Form.Get("status")

	if plat.ID < 1 || plat.Name == "" {
		models.ResponseWithJSON(w, http.StatusOK,
			models.Response{Code: 4000, Msg: "bad params"})
		return
	}
	data, err := models.UpdatePlatform(&plat)
	if err != nil {
		models.ResponseWithJSON(w, http.StatusOK,
			models.Response{Code: 4000, Msg: err.Error()})
		return
	}
	w.Write([]byte(data))
}

// DeletePlatform : DELETE PLATFORM
func DeletePlatform(w http.ResponseWriter, r *http.Request) {
	var plat models.Platform
	// set response header json
	w.Header().Set("Content-Type", "application/json; charset=UTF-8")

	r.ParseForm()
	plat.ID, _ = strconv.Atoi(r.Form.Get("id"))
	plat.Name = r.Form.Get("name")

	if plat.ID < 1 || plat.Name == "" {
		models.ResponseWithJSON(w, http.StatusOK,
			models.Response{Code: 4000, Msg: "bad params"})
		return
	}
	data, err := models.DeletePlatform(&plat)
	if err != nil {
		models.ResponseWithJSON(w, http.StatusOK,
			models.Response{Code: 4000, Msg: err.Error()})
		return
	}
	w.Write([]byte(data))
}
