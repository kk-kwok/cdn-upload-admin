package http

import (
	"net/http"

	"cdn-upload-admin/cdnadmin_apiserver/config"
)

// Index : default index
func Index(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Welcome"))
}

// Health : check health
func Health(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("OK"))
}

// Version : get version
func Version(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte(config.VERSION))
}
