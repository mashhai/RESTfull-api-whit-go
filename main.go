package main

import (
	"log"
	"net/http"

	"encoding/json"

	"github.com/gorilla/mux"
)

// Person represents a person to run a command.
type Person struct {
	ID        string `json:"id,omitempty"`
	FirstName string `json:"firstname,omitempty"`
	LastName  string `json:"lastname,omitempty"`
	Address   *Address
}

//Address represents a address to run comand.
type Address struct {
	City  string
	State string
}

var people []Person

// GetPersonEndPoint represents a getPersonEndPoint to run a command.
func GetPersonEndPoint(w http.ResponseWriter, req *http.Request) {
	params := mux.Vars(req)
	for _, item := range people {
		if item.ID == params["id"] {
			json.NewEncoder(w).Encode(item)
			return
		}
	}
	json.NewEncoder(w).Encode(&Person{})
}

// GetPeopleEndPoint represents a GetPeopleEndPoint to run a command.
func GetPeopleEndPoint(w http.ResponseWriter, req *http.Request) {
	json.NewEncoder(w).Encode(people)
}

func CreatePersonEndPoint(w http.ResponseWriter, req *http.Request) {
	params := mux.Vars(req)
	var person Person
	_ = json.NewDecoder(req.Body).Decode(&person)
	person.ID = params["id"]
	people = append(people, person)
	json.NewEncoder(w).Encode(people)
}

// DeletePersonEndPoint represents a DeletePersonEndPoint to run a command.
func DeletePersonEndPoint(w http.ResponseWriter, req *http.Request) {
	params := mux.Vars(req)
	for index, item := range people {
		if item.ID == params["id"] {
			people = append(people[:index], people[index+1:]...)
			break
		}
	}
	json.NewEncoder(w).Encode(people)
}
func main() {
	router := mux.NewRouter()
	people = append(people, Person{ID: "1", FirstName: "Mohsen", LastName: "Farahani", Address: &Address{City: "Tehrean", State: "TehranPars"}})
	people = append(people, Person{ID: "2", FirstName: "Mohsen", LastName: "Farahani"})
	router.HandleFunc("/people", GetPeopleEndPoint).Methods("GET")
	router.HandleFunc("/people/{id}", GetPersonEndPoint).Methods("GET")
	router.HandleFunc("/people/{id}", CreatePersonEndPoint).Methods("POST")
	router.HandleFunc("/people/{id}", DeletePersonEndPoint).Methods("DELETE")
	log.Fatal(http.ListenAndServe(":12345", router))
}
