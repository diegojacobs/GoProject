package Repositories

import (
	"fmt"

	"github.com/diegojacobs/GoProject/Structures"
)

var currentID int

var people Structures.Persons

// Give us some seed data
func init() {
	RepoCreatePerson(Structures.Person{Firstname: "Carlos", Lastname: "Ruiz", Address: &Structures.Address{City: "Guatemala", Country: "Guatemala"}})
	RepoCreatePerson(Structures.Person{Firstname: "Diego", Lastname: "Jacobs", Address: &Structures.Address{City: "Guatemala", Country: "Guatemala"}})
}

func RepoFindPerson(id int) Structures.Person {
	for _, p := range people {
		if p.ID == id {
			return p
		}
	}
	// return empty Person if not found
	return Structures.Person{}
}

func RepoCreatePerson(p Structures.Person) Structures.Person {
	currentID++
	p.ID = currentID
	people = append(people, p)
	return p
}

func RepoDestroyPerson(id int) error {
	for i, p := range people {
		if p.ID == id {
			people = append(people[:i], people[i+1:]...)
			return nil
		}
	}
	return fmt.Errorf("Could not find a Person with id of %d to delete", id)
}

func RepoFindPeople() Structures.Persons {
	return people
}
