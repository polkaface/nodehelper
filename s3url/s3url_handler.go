package main

import (
	"encoding/json"
	"net/http"
)

func init() {
}

type S3URLHdl struct {
	URLHdl
}

//Post is POST
func (hdl *S3URLHdl) Post(w http.ResponseWriter, r *http.Request) {
}

//Get is GET
func (hdl *S3URLHdl) Get(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	encoder := json.NewEncoder(w)
	encoder.SetEscapeHTML(false)
	encoder.Encode(files)

	return
}
