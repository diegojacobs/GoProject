package main

import (
	"log"
	"net/http"
)

func main() {
	router := NewRouter()

	people = append(people, Person{Id: "1", Firstname: "Carlos", Lastname: "Ruiz", Address: &Address{City: "Guatemala", State: "Guatemala"}})
	people = append(people, Person{Id: "2", Firstname: "Diego", Lastname: "Jacobs", Address: &Address{City: "Guatemala", State: "Guatemala"}})

	log.Fatal(http.ListenAndServe(":12345", router))
}
