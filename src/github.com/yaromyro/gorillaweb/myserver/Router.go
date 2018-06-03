package myserver

import (
	"github.com/gorilla/mux"
)

func MakeRouter() *mux.Router {
	r := mux.NewRouter()
	r.HandleFunc("/user", createUser).Methods("POST")
	r.HandleFunc("/user/{id}", getUser).Methods("GET")
	r.HandleFunc("/users", getUsers).Methods("GET")
	r.HandleFunc("/user/{id}", deleteUser).Methods("DELETE")
	r.HandleFunc("/user/{id}", updateUser).Methods("PUT")
	r.HandleFunc("/user/{stop}/", stopServer).Methods("GET")
	return r
}
