package v1

import (
	"net/http"
	"strconv"

	"cdn-upload-admin/cdnadmin_apiserver/models/v1"
)

// GetProjectList : GET PROJECT LIST
func GetProjectList(w http.ResponseWriter, r *http.Request) {
	var p models.Project
	// set response header json
	w.Header().Set("Content-Type", "application/json; charset=UTF-8")

	r.ParseForm()
	num, _ := strconv.Atoi(r.Form.Get("page_num"))
	size, _ := strconv.Atoi(r.Form.Get("page_size"))
	p.Name = r.Form.Get("name")
	p.Status = r.Form.Get("status")
	platname := r.Form.Get("platform_name")
	if num < 1 || size < 1 {
		num = 0
		size = 0
	}

	data, err := models.GetProjectList(num, size, &p, platname)
	if err != nil {
		models.ResponseWithJSON(w, http.StatusOK,
			models.Response{Code: 4000, Msg: err.Error()})
		return
	}
	w.Write([]byte(data))
}

// AddProject : CREATE PROJECT
func AddProject(w http.ResponseWriter, r *http.Request) {
	var p models.Project
	// set response header json
	w.Header().Set("Content-Type", "application/json; charset=UTF-8")

	r.ParseForm()
	p.ID, _ = strconv.Atoi(r.Form.Get("id"))
	p.Name = r.Form.Get("name")
	p.PlatformName = r.Form.Get("platform_name")
	p.CdnID, _ = strconv.Atoi(r.Form.Get("cdn_id"))
	p.Path = r.Form.Get("path")
	p.Domain = r.Form.Get("domain")
	p.Status = r.Form.Get("status")

	if p.ID < 1 || p.Name == "" || p.Path == "" ||
		p.Domain == "" || p.Status == "" {
		models.ResponseWithJSON(w, http.StatusOK,
			models.Response{Code: 4000, Msg: "bad params"})
		return
	}
	data, err := models.AddProject(&p)
	if err != nil {
		models.ResponseWithJSON(w, http.StatusOK,
			models.Response{Code: 4000, Msg: err.Error()})
		return
	}
	w.Write([]byte(data))
}

// UpdateProject : UPDATE PROJECT
func UpdateProject(w http.ResponseWriter, r *http.Request) {
	var p models.Project
	// set response header json
	w.Header().Set("Content-Type", "application/json; charset=UTF-8")

	r.ParseForm()
	p.ID, _ = strconv.Atoi(r.Form.Get("id"))
	p.Name = r.Form.Get("name")
	p.PlatformName = r.Form.Get("platform_name")
	p.CdnID, _ = strconv.Atoi(r.Form.Get("cdn_id"))
	p.Path = r.Form.Get("path")
	p.Domain = r.Form.Get("domain")
	p.Status = r.Form.Get("status")

	if p.ID < 1 || p.Name == "" ||
		p.Path == "" || p.Domain == "" || p.Status == "" {
		models.ResponseWithJSON(w, http.StatusOK,
			models.Response{Code: 4000, Msg: "bad params"})
		return
	}
	data, err := models.UpdateProject(&p)
	if err != nil {
		models.ResponseWithJSON(w, http.StatusOK,
			models.Response{Code: 4000, Msg: err.Error()})
		return
	}
	w.Write([]byte(data))
}

// DeleteProject : DELETE PROJECT
func DeleteProject(w http.ResponseWriter, r *http.Request) {
	var p models.Project
	// set response header json
	w.Header().Set("Content-Type", "application/json; charset=UTF-8")

	r.ParseForm()
	p.ID, _ = strconv.Atoi(r.Form.Get("id"))
	p.Name = r.Form.Get("name")

	if p.ID < 1 || p.Name == "" {
		models.ResponseWithJSON(w, http.StatusOK,
			models.Response{Code: 4000, Msg: "bad params"})
		return
	}
	data, err := models.DeleteProject(&p)
	if err != nil {
		models.ResponseWithJSON(w, http.StatusOK,
			models.Response{Code: 4000, Msg: err.Error()})
		return
	}
	w.Write([]byte(data))
}
