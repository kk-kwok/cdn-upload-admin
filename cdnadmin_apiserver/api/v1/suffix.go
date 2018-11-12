package v1

import (
	"net/http"
	"strconv"

	"cdn-upload-admin/cdnadmin_apiserver/models/v1"
)

// GetFileSuffixList : get file suffix list
func GetFileSuffixList(w http.ResponseWriter, r *http.Request) {
	// set response header json
	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	r.ParseForm()
	num, _ := strconv.Atoi(r.Form.Get("page_num"))
	size, _ := strconv.Atoi(r.Form.Get("page_size"))
	if num < 1 || size < 1 {
		num = 0
		size = 0
	}

	data, err := models.GetFileSuffixList(num, size)
	if err != nil {
		models.ResponseWithJSON(w, http.StatusOK,
			models.Response{Code: 4000, Msg: err.Error()})
		return
	}
	w.Write([]byte(data))
}

// AddFileSuffix : add file suffix
func AddFileSuffix(w http.ResponseWriter, r *http.Request) {
	var fsuffix models.FileSuffix
	// set response header json
	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	r.ParseForm()
	fsuffix.FileSuffix = r.Form.Get("file_suffix")

	if fsuffix.FileSuffix == "" {
		models.ResponseWithJSON(w, http.StatusOK,
			models.Response{Code: 4000, Msg: "bad params"})
		return
	}
	data, err := models.AddFileSuffix(&fsuffix)
	if err != nil {
		models.ResponseWithJSON(w, http.StatusOK,
			models.Response{Code: 4000, Msg: err.Error()})
		return
	}
	w.Write([]byte(data))
}

// UpdateFileSuffix : update file suffix
func UpdateFileSuffix(w http.ResponseWriter, r *http.Request) {
	var fsuffix models.FileSuffix
	// set response header json
	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	r.ParseForm()
	fsuffix.ID, _ = strconv.Atoi(r.Form.Get("id"))
	fsuffix.FileSuffix = r.Form.Get("file_suffix")

	if fsuffix.ID < 1 || fsuffix.FileSuffix == "" {
		models.ResponseWithJSON(w, http.StatusOK,
			models.Response{Code: 4000, Msg: "bad params"})
		return
	}
	data, err := models.UpdateFileSuffix(&fsuffix)
	if err != nil {
		models.ResponseWithJSON(w, http.StatusOK,
			models.Response{Code: 4000, Msg: err.Error()})
		return
	}
	w.Write([]byte(data))
}

// DeleteFileSuffix : delete file suffix
func DeleteFileSuffix(w http.ResponseWriter, r *http.Request) {
	var fsuffix models.FileSuffix
	// set response header json
	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	r.ParseForm()
	fsuffix.ID, _ = strconv.Atoi(r.Form.Get("id"))
	fsuffix.FileSuffix = r.Form.Get("file_suffix")

	if fsuffix.ID < 1 || fsuffix.FileSuffix == "" {
		models.ResponseWithJSON(w, http.StatusOK,
			models.Response{Code: 4000, Msg: "bad params"})
		return
	}
	data, err := models.DeleteFileSuffix(&fsuffix)
	if err != nil {
		models.ResponseWithJSON(w, http.StatusOK,
			models.Response{Code: 4000, Msg: err.Error()})
		return
	}
	w.Write([]byte(data))
}
