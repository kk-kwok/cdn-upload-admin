package http

import (
	"log"
	"net/http"
	_ "net/http/pprof"

	"cdn-upload-admin/cdnadmin_apiserver/config"
)

// Start : http start
func Start() {
	router := NewRouter()
	addr := config.Config().HTTP.Listen
	if addr == "" {
		return
	}
	log.Println("http listening", addr)
	log.Fatalln(http.ListenAndServe(addr, router))
}
