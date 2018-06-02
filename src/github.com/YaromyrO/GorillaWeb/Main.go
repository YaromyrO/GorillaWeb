package main

import (
	"./myServer"
	"log"
	"net/http"
)

func main() {
	r := myServer.MakeRouter()
	srv := &http.Server{
		Handler: r,
		Addr:    "localhost:8080",
	}

	log.Fatal(srv.ListenAndServe())
}
