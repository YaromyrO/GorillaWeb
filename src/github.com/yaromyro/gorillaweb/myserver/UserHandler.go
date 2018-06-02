package myserver

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
	ID      int    `json:"id"`
	Name    string `json:"name"`
	Surname string `json:"surname"`
	Age     int    `json:"age"`
}

var users []user

func getFromFile() {
	response, _ := ioutil.ReadFile("users.json")
	json.Unmarshal(response, &users)
}

func createUser(w http.ResponseWriter, r *http.Request) {
	getFromFile()
	r.ParseForm()
	rand.Seed(time.Now().Unix())
	id := rand.Intn(99)
	name := r.Form.Get("name")
	surname := r.Form.Get("surname")
	age, _ := strconv.Atoi(r.Form.Get("age"))
	users = append(users, user{id, name, surname, age})
	writeToFile()
}

func getUser(w http.ResponseWriter, r *http.Request) {
	getFromFile()
	vars := mux.Vars(r)
	id, err := strconv.Atoi(vars["id"])
	if err != nil {
		fmt.Println(err)
	}
	for _, usr := range users {
		if usr.ID == id {
			fmt.Fprint(w, "||ID: ", usr.ID, "|| ")
			fmt.Fprint(w, "||Name: ", usr.Name, "|| ")
			fmt.Fprint(w, "||Surname: ", usr.Surname, "|| ")
			fmt.Fprint(w, "||Age: ", usr.Age, "|| ")
			fmt.Fprintln(w)
		}
	}
}

func getUsers(w http.ResponseWriter, r *http.Request) {
	getFromFile()
	for _, usr := range users {
		fmt.Fprint(w, "||ID: ", usr.ID, "|| ")
		fmt.Fprint(w, "||Name: ", usr.Name, "|| ")
		fmt.Fprint(w, "||Surname: ", usr.Surname, "|| ")
		fmt.Fprint(w, "||Age: ", usr.Age, "|| ")
		fmt.Fprintln(w)
	}
}

func updateUser(w http.ResponseWriter, r *http.Request) {
	getFromFile()
	vars := mux.Vars(r)
	id, err := strconv.Atoi(vars["id"])
	if err != nil {
		fmt.Println(err)
	}
	r.ParseForm()
	name := r.Form.Get("name")
	surname := r.Form.Get("surname")
	age, _ := strconv.Atoi(r.Form.Get("age"))

	for index, usr := range users {
		if usr.ID == id {
			updatedUser := user{id, name, surname, age}
			users[index] = updatedUser
		}
	}
	writeToFile()
}

func deleteUser(w http.ResponseWriter, r *http.Request) {
	getFromFile()
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
	writeToFile()
}

func writeToFile() {
	response, _ := json.Marshal(users)
	err := ioutil.WriteFile("users.json", response, os.ModeDir)
	if err != nil {
		fmt.Println(err)
	}
}
