package myServer

import (
	"github.com/gorilla/mux"
)

func MakeRouter() *mux.Router {
	r := mux.NewRouter()
	r.HandleFunc("/user/{name}/{age}/", createUser).Methods("POST")
	r.HandleFunc("/user/{id}/", getUser).Methods("GET")
	r.HandleFunc("/user/", getUsers).Methods("GET")
	r.HandleFunc("/user/{id}", deleteUser).Methods("DELETE")
	r.HandleFunc("/user/{id}/{name}/{age}", updateUser).Methods("POST")
	return r
}
