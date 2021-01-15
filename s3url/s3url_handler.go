package main

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/cz-theng/czkit-go/log"
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
	w.Header().Set("content-type", "application/json")

	buf, err := json.Marshal(&files)
	if err != nil {
		log.Error("Marshal Error")
		return
	}
	_, err = fmt.Fprintf(w, string(buf))
	if err != nil {
		log.Error("WriteMessage Error:%s", err.Error())
	}
	return
}
