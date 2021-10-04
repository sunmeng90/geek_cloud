package main

import (
	"fmt"
	"github.com/golang/glog"
	"log"
	"net/http"
	"os"
	"time"
)

func main() {
	http.HandleFunc("/echo", echo)
	http.HandleFunc("/healthz", healthz)
	if err := http.ListenAndServe(":8090", nil); err != nil {
		log.Fatal("Start up failed", err)
	}
}

func healthz(writer http.ResponseWriter, request *http.Request) {
	writer.WriteHeader(200)
}

func echo(rw http.ResponseWriter, r *http.Request) {
	glog.Infof("Request from: %s", r.RemoteAddr)
	glog.Infof("Request Header, Request-From: %s", r.Header.Get("Request-From"))
	rw.Header().Set("Response-From", "go-server")
	env, _ := os.LookupEnv("VERSION")
	rw.Header().Set("Version", env)
	rw.Header().Set("Content-Type", "text/plain; charset=utf-8") // normal header
	rw.WriteHeader(http.StatusOK)
	glog.Infof("Response Code: %d", http.StatusOK)
	rw.Header().Set("Response-Trailer1", "val1") // will not sent to client
	fmt.Fprintf(rw, "Greeting: %s. Time: %s", r.URL.Query().Get("name"), time.Now().Format("2006-01-02 15:04:05 -0700 MST"))
	rw.Header().Set("Response-Trailer2", "val2") // will not sent to client
}
