package main

import (
	"log"
	"net/http"
)

const listenAddr = ":8080"

func main() {
	http.HandleFunc("/", rootHandler)
	log.Printf("start webserver on %s", listenAddr)
	panic(http.ListenAndServe(listenAddr, nil))
}

func rootHandler(w http.ResponseWriter, r *http.Request) {
	//w.WriteHeader(http.StatusForbidden)
	_, err := w.Write([]byte("Hello, world!"))
	if err != nil {
		log.Print(err)
	}
}
