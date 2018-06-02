package myServer

import (
	"encoding/json"
	"fmt"
	"github.com/gorilla/mux"
	"io/ioutil"
	"math/rand"
	"net/http"
	"os"
	"strconv"
	"time"
)

type user struct {
	ID   int
	Name string
	Age  int
}

var users []user

func get() {
	response, _ := ioutil.ReadFile("users.json")
	json.Unmarshal(response, &users)
}

func createUser(w http.ResponseWriter, r *http.Request) {
	get()
	rand.Seed(time.Now().Unix())
	vars := mux.Vars(r)
	id := rand.Intn(99)
	name := vars["name"]
	age, err := strconv.Atoi(vars["age"])
	if err != nil {
		fmt.Println(err)
	}

	users = append(users, user{id, name, age})
	if err != nil {
		fmt.Println(err)
	}
	write()
}

func getUser(w http.ResponseWriter, r *http.Request) {
	get()
	vars := mux.Vars(r)
	id, err := strconv.Atoi(vars["id"])
	if err != nil {
		fmt.Println(err)
	}

	for _, usr := range users {
		if usr.ID == id {
			fmt.Fprint(w, "ID: ", usr.ID, "; ")
			fmt.Fprint(w, "Name: ", usr.Name, "; ")
			fmt.Fprint(w, "Age: ", usr.Age, "; ")
			fmt.Fprintln(w)
		}
	}
}

func getUsers(w http.ResponseWriter, r *http.Request) {
	get()
	for _, usr := range users {
		fmt.Fprint(w, "ID: ", usr.ID, "; ")
		fmt.Fprint(w, "Name: ", usr.Name, "; ")
		fmt.Fprint(w, "Age: ", usr.Age, "; ")
		fmt.Fprintln(w)
	}
}

func updateUser(w http.ResponseWriter, r *http.Request) {
	get()
	vars := mux.Vars(r)
	id, err := strconv.Atoi(vars["id"])
	name := vars["name"]
	age, err := strconv.Atoi(vars["age"])
	if err != nil {
		fmt.Println(err)
	}
	for index, usr := range users {
		if usr.ID == id {
			updatedUser := user{id, name, age}
			users[index] = updatedUser
		}
	}
	write()
}

func deleteUser(w http.ResponseWriter, r *http.Request) {
	get()
	vars := mux.Vars(r)
	id, err := strconv.Atoi(vars["id"])
	if err != nil {
		fmt.Println(err)
	}
	for index, user := range users {
		if user.ID == id {
			users = append(users[:index], users[index+1:]...)
		}
	}
	write()
}

func write() {
	response, _ := json.Marshal(users)
	err := ioutil.WriteFile("users.json", response, os.ModeDir)
	if err != nil {
		fmt.Println(err)
	}
}
