package v1

import (
	"net/http"
	"strconv"

	"cdn-upload-admin/cdnadmin_apiserver/models/v1"
)

// CDNPush : CDN FILE PUSH
func CDNPush(w http.ResponseWriter, r *http.Request) {
	// set response header json
	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	r.ParseForm()
	cfID, _ := strconv.Atoi(r.Form.Get("id"))
	if cfID < 1 {
		models.ResponseWithJSON(w, http.StatusOK,
			models.Response{Code: 4000, Msg: "bad param"})
		return
	}

	data, err := models.CDNFilePush(cfID)
	if err != nil {
		models.ResponseWithJSON(w, http.StatusOK,
			models.Response{Code: 4000, Msg: err.Error()})
		return
	}
	w.Write([]byte(data))
}
