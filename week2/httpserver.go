package main

import (
	"fmt"
	"log"
	"net/http"
	"time"
)

func main() {
	http.HandleFunc("/echo", func(rw http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(rw, "Greeting: %s. Time: %s", r.URL.Query().Get("name"), time.Now().Format("2006-01-02 15:04:05 -0700 MST"))
	})
	if err := http.ListenAndServe(":8090", nil); err != nil {
		log.Fatal("Start up failed", err)
	}
}
