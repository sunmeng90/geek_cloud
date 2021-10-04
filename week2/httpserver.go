package main

import (
	"fmt"
	"log"
	"net/http"
	"time"
)

func main() {
	http.HandleFunc("/echo", func(rw http.ResponseWriter, r *http.Request) {
		log.Printf("Request from: %s", r.Header.Get("Request-From"))
		rw.Header().Set("Response-From", "go-server")
		rw.Header().Set("Content-Type", "text/plain; charset=utf-8") // normal header
		rw.WriteHeader(200)
		rw.Header().Set("Response-Trailer1", "val1") // will not sent to client
		fmt.Fprintf(rw, "Greeting: %s. Time: %s", r.URL.Query().Get("name"), time.Now().Format("2006-01-02 15:04:05 -0700 MST"))
		rw.Header().Set("Response-Trailer2", "val2") // will not sent to client

	})
	if err := http.ListenAndServe(":8090", nil); err != nil {
		log.Fatal("Start up failed", err)
	}
}
