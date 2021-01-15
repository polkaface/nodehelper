package main

import (
	"net/http"
	"time"

	"github.com/cz-theng/czkit-go/log"
)

var server APIServer

type APIServer struct {
	httpd     http.Server
	router    *Router
	secTicker *time.Ticker
}

func (svr *APIServer) Init() {
	svr.httpd = http.Server{
		Addr:           "0.0.0.0:10080",
		ReadTimeout:    10 * time.Second,
		WriteTimeout:   10 * time.Second,
		MaxHeaderBytes: 1 << 20,
		Handler:        svr,
	}
	svr.router = router
}

func (svr *APIServer) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	svr.router.DealRaw(r.URL.Path, w, r)
}

func (svr *APIServer) Start() {
	err := svr.httpd.ListenAndServe()
	if err != nil {
		log.Error("ListenAndServe Error:%s", err.Error())
	}
}
