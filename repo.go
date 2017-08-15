package main

import "fmt"

var currentId int

var people Persons

// Give us some seed data
func init() {
	RepoCreatePerson(Person{Firstname: "Carlos", Lastname: "Ruiz", Address: &Address{City: "Guatemala", State: "Guatemala"}})
	RepoCreatePerson(Person{Firstname: "Diego", Lastname: "Jacobs", Address: &Address{City: "Guatemala", State: "Guatemala"}})
}

func RepoFindPerson(id int) Person {
	for _, p := range people {
		if p.Id == id {
			return p
		}
	}
	// return empty Person if not found
	return Person{}
}

func RepoCreatePerson(p Person) Person {
	currentId++
	p.Id = currentId
	people = append(people, p)
	return p
}

func RepoDestroyPerson(id int) error {
	for i, p := range people {
		if p.Id == id {
			people = append(people[:i], people[i+1:]...)
			return nil
		}
	}
	return fmt.Errorf("Could not find a Person with id of %d to delete", id)
}
