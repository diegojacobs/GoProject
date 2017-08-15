package main

import (
	"encoding/json"
	"io"
	"io/ioutil"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

// @Title GetPersonEndpoint
// @Description retrieves person with the id received
// @Accept  json
// @Param   id
// @Success 200 {object}  my_api.model.OrderRow
// @Failure 400 {object} my_api.ErrorResponse    "Customer ID must be specified"
// @Resource /people
// @Router /people/{id} [get]
func GetPersonEndpoint(w http.ResponseWriter, req *http.Request) {
	params := mux.Vars(req)
	number, _ := strconv.Atoi(params["id"])
	json.NewEncoder(w).Encode(RepoFindPerson(number))
}

func GetPeopleEndpoint(w http.ResponseWriter, req *http.Request) {
	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	w.WriteHeader(http.StatusOK)
	if err := json.NewEncoder(w).Encode(people); err != nil {
		panic(err)
	}
}

func CreatePersonEndpoint(w http.ResponseWriter, req *http.Request) {
	var person Person
	body, err := ioutil.ReadAll(io.LimitReader(req.Body, 1048576))

	if err != nil {
		panic(err)
	}

	if err := req.Body.Close(); err != nil {
		panic(err)
	}

	if err := json.Unmarshal(body, &person); err != nil {
		w.Header().Set("Content-Type", "application/json; charset=UTF-8")
		w.WriteHeader(422) // unprocessable entity
		if err := json.NewEncoder(w).Encode(err); err != nil {
			panic(err)
		}
	}

	t := RepoCreatePerson(person)
	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	w.WriteHeader(http.StatusCreated)
	if err := json.NewEncoder(w).Encode(t); err != nil {
		panic(err)
	}
}

func DeletePersonEndpoint(w http.ResponseWriter, req *http.Request) {
	params := mux.Vars(req)
	number, _ := strconv.Atoi(params["id"])
	RepoDestroyPerson(number)
	json.NewEncoder(w).Encode(people)
}
