package main

import "github.com/gorilla/mux"
import "net/http"
import "log"
import "encoding/json"

type Person struct {
	Id 			string 	`json:"id,omitempty"`
	Firstname 	string 	`json:"firstname,omitempty"`
	Lastname 	string 	`json:"lastname,omitempty"`
	Address 	*Address `json:"address,omitempty"`
}

type Address struct {
	City 	string `json:"city,omitempty"`
	State 	string `json:"state,omitempty"`
}

var people []Person

func GetPersonEndpoint(w http.ResponseWriter, req *http.Request) {
	params := mux.Vars(req)
	for _, item := range people {
		if item.Id == params["id"]{
			json.NewEncoder(w).Encode(item)
			return
		}
	}
	json.NewEncoder(w).Encode(&Person{})
}

func GetPeopleEndpoint(w http.ResponseWriter, req *http.Request) {
	json.NewEncoder(w).Encode(people)
}

func CreatePersonEndpoint(w http.ResponseWriter, req *http.Request) {
	params := mux.Vars(req)
	var person Person
	_ = json.NewDecoder(req.Body).Decode(&person)
	person.Id = params["id"]
	people = append(people, person)
	json.NewEncoder(w).Encode(people)
}

func DeletePersonEndpoint(w http.ResponseWriter, req *http.Request) {
	params := mux.Vars(req)
	for index, item := range people {
		if item.Id == params["id"] {
			people = append(people[:index], people[index+1:]...)
			break
		}
	}
	json.NewEncoder(w).Encode(people)
}

func main(){
	router := mux.NewRouter()
	people = append(people, Person{Id:"1", Firstname: "Carlos", Lastname: "Ruiz", Address: &Address{City: "Guatemala", State: "Guatemala"}})
	people = append(people, Person{Id:"2", Firstname: "Diego", Lastname: "Jacobs", Address: &Address{City: "Guatemala", State: "Guatemala"}})
	
	router.HandleFunc("/people", GetPeopleEndpoint).Methods("GET")
	router.HandleFunc("/people/{id}", GetPersonEndpoint).Methods("GET")
	router.HandleFunc("/people/{id}", CreatePersonEndpoint).Methods("POST")
	router.HandleFunc("/people/{id}", GetPeopleEndpoint).Methods("DELETE")
	
	log.Fatal(http.ListenAndServe(":12345", router))	
}