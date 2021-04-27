package main

import (
	"log"
	"net/http"
)

const listenAddr = ":8080"

var (
	version = "na"
	commit  = "na"
	date    = "na"
)

func main() {
	http.HandleFunc("/", rootHandler)
	log.Printf("Start Hello, world! (version=%s commit=%s date=%s) on %s", version, commit, date, listenAddr)
	panic(http.ListenAndServe(listenAddr, nil))
}

func rootHandler(w http.ResponseWriter, _ *http.Request) {
	//w.WriteHeader(http.StatusForbidden)
	_, err := w.Write([]byte("Hello, world!"))
	if err != nil {
		log.Print(err)
	}
}
