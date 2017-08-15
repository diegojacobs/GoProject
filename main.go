// @APIVersion 1.0.0
// @APITitle GO Project
// @APIDescription My API usually works as expected.
// @Contact diegojacobs9595@gmail.com
// @BasePath http://localhost:12345/
package main

import (
	"log"
	"net/http"
)

func main() {
	router := NewRouter()

	log.Fatal(http.ListenAndServe(":12345", router))
}
