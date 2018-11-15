package v1

import (
	"crypto/md5"
	"encoding/hex"
	"io"
	"net/http"
	"os"
	"os/exec"
	"path/filepath"
	"strconv"
	"strings"

	"cdn-upload-admin/cdnadmin_apiserver/models/v1"
)

// GetCDNFileList : GET CDN FILE LIST
func GetCDNFileList(w http.ResponseWriter, r *http.Request) {
	var cfile models.CDNFile
	// set response header json
	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	r.ParseForm()
	num, _ := strconv.Atoi(r.Form.Get("page_num"))
	size, _ := strconv.Atoi(r.Form.Get("page_size"))
	cfile.ProjectName = r.Form.Get("project_name")
	suffix := r.Form.Get("file_suffix")
	if num < 1 || size < 1 {
		num = 0
		size = 0
	}

	data, err := models.GetCDNFileList(num, size, &cfile, suffix)
	if err != nil {
		models.ResponseWithJSON(w, http.StatusOK,
			models.Response{Code: 4000, Msg: err.Error()})
		return
	}
	w.Write([]byte(data))
}

// GetCurrentPath : GET CURRENT PATH
func GetCurrentPath() string {
	file, _ := exec.LookPath(os.Args[0])
	path, _ := filepath.Abs(file)
	path = string(path[0:strings.LastIndex(path, "/")])
	return path
}

// Sizer : define size interface
type Sizer interface {
	Size() int64
}

// Stat : define stat interface
type Stat interface {
	Stat() (os.FileInfo, error)
}

// AddCDNFile : CREATE CDNFILE
func AddCDNFile(w http.ResponseWriter, r *http.Request) {
	var cfile models.CDNFile
	// set response header json
	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	r.ParseMultipartForm(32 << 20)
	file, _, err := r.FormFile("file")
	if err != nil {
		models.ResponseWithJSON(w, http.StatusOK,
			models.Response{Code: 4000, Msg: err.Error()})
		return
	}
	defer file.Close()

	cfile.FileName = r.MultipartForm.Value["file_name"][0]
	cfile.ProjectName = r.MultipartForm.Value["project_name"][0]
	cfile.Comment = r.MultipartForm.Value["comment"][0]
	cfile.UserID, _ = strconv.Atoi(r.MultipartForm.Value["user_id"][0])
	if cfile.FileName == "" || cfile.ProjectName == "" {
		models.ResponseWithJSON(w, http.StatusOK,
			models.Response{Code: 4000, Msg: "bad param"})
		return
	}
	// save file
	tmpPath := GetCurrentPath() + "/uploads"
	os.MkdirAll(tmpPath, os.ModePerm)
	tmpFile := tmpPath + "/" + cfile.FileName
	f, err := os.OpenFile(tmpFile, os.O_WRONLY|os.O_CREATE, 0644)
	if err != nil {
		models.ResponseWithJSON(w, http.StatusOK,
			models.Response{Code: 4000, Msg: err.Error()})
		return
	}
	defer f.Close()
	io.Copy(f, file)
	defer file.Close()
	// delete upload tmp file
	r.MultipartForm.RemoveAll()

	// get file size
	file, _ = os.Open(tmpFile)
	if statInterface, ok := file.(Stat); ok {
		fileInfo, statErr := statInterface.Stat()
		if statErr != nil {
			models.ResponseWithJSON(w, http.StatusOK,
				models.Response{Code: 4000, Msg: err.Error()})
			return
		}
		cfile.FileSize = strconv.FormatInt(fileInfo.Size(), 10)
	}
	// get file md5
	md5h := md5.New()
	io.Copy(md5h, file)
	cfile.FileMD5 = hex.EncodeToString(md5h.Sum(nil))
	defer file.Close()

	data, err := models.AddCDNFile(tmpFile, &cfile)
	if err != nil {
		models.ResponseWithJSON(w, http.StatusOK,
			models.Response{Code: 4000, Msg: err.Error()})
		return
	}
	w.Write([]byte(data))
}

// UpdateCDNFile : UPDATE CDN FILE
func UpdateCDNFile(w http.ResponseWriter, r *http.Request) {
	var cfile models.CDNFile
	// set response header json
	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	r.ParseForm()
	cfile.ID, _ = strconv.Atoi(r.Form.Get("id"))
	cfile.FileName = r.Form.Get("file_name")
	cfile.Comment = r.Form.Get("comment")
	if cfile.ID < 1 || cfile.FileName == "" {
		models.ResponseWithJSON(w, http.StatusOK,
			models.Response{Code: 4000, Msg: "bad param"})
		return
	}

	data, err := models.UpdateCDNFile(&cfile)
	if err != nil {
		models.ResponseWithJSON(w, http.StatusOK,
			models.Response{Code: 4000, Msg: err.Error()})
		return
	}
	w.Write([]byte(data))
}

// DeleteCDNFile : DELETE CDN FILE
func DeleteCDNFile(w http.ResponseWriter, r *http.Request) {
	var cfile models.CDNFile
	// set response header json
	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	r.ParseForm()
	cfile.ID, _ = strconv.Atoi(r.Form.Get("id"))
	cfile.FileName = r.Form.Get("file_name")
	if cfile.ID < 1 || cfile.FileName == "" {
		models.ResponseWithJSON(w, http.StatusOK,
			models.Response{Code: 4000, Msg: "bad param"})
		return
	}

	data, err := models.DeleteCDNFile(&cfile)
	if err != nil {
		models.ResponseWithJSON(w, http.StatusOK,
			models.Response{Code: 4000, Msg: err.Error()})
		return
	}
	w.Write([]byte(data))
}
