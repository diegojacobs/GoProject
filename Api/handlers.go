package Api

import (
	"encoding/json"
	"io"
	"io/ioutil"
	"net/http"
	"strconv"

	"github.com/diegojacobs/GoProject/Repositories"
	"github.com/diegojacobs/GoProject/Structures"
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
	json.NewEncoder(w).Encode(Repositories.RepoFindPerson(number))
}

func GetPeopleEndpoint(w http.ResponseWriter, req *http.Request) {
	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	w.WriteHeader(http.StatusOK)

	err := json.NewEncoder(w).Encode(Repositories.RepoFindPeople())
	if err != nil {
		panic(err)
	}
}

func CreatePersonEndpoint(w http.ResponseWriter, req *http.Request) {
	var person Structures.Person
	body, limitError := ioutil.ReadAll(io.LimitReader(req.Body, 1048576))

	if limitError != nil {
		panic(limitError)
	}

	error := req.Body.Close()
	if error != nil {
		panic(error)
	}

	encodeError := json.Unmarshal(body, &person)
	if encodeError != nil {
		w.Header().Set("Content-Type", "application/json; charset=UTF-8")
		w.WriteHeader(422) // unprocessable entity
		error := json.NewEncoder(w).Encode(encodeError)
		if error != nil {
			panic(error)
		}
	}

	t := Repositories.RepoCreatePerson(person)
	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	w.WriteHeader(http.StatusCreated)
	err := json.NewEncoder(w).Encode(t)
	if err != nil {
		panic(err)
	}
}

func DeletePersonEndpoint(w http.ResponseWriter, req *http.Request) {
	params := mux.Vars(req)
	number, _ := strconv.Atoi(params["id"])
	Repositories.RepoDestroyPerson(number)
	json.NewEncoder(w).Encode(Repositories.RepoFindPeople())
}
