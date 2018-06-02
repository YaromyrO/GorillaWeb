package main

import (
	"./myserver"
	"log"
	"net/http"
	"os"
)

func main() {
	r := myserver.MakeRouter()
	os.Setenv("Address", "localhost:8080")
	srv := &http.Server{
		Handler: r,
		Addr:    os.Getenv("Address"),
	}
	log.Fatal(srv.ListenAndServe())
}
