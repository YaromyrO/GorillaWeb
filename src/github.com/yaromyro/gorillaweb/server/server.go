package main

import (
	"../myserver"
	"net/http"
	"os"
	"log"
	"../client"
)

func main() {
	r := myserver.MakeRouter()
	os.Setenv("Address", "localhost:8080")
	srv := &http.Server{
		Handler: r,
		Addr:    os.Getenv("Address"),
	}
	client.Listener()
	log.Fatal(srv.ListenAndServe())
}

