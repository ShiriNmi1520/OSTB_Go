package main

import (
	"net/http"

	routes "OSTB_Go/src/router"
)

func main() {
	router := routes.NewRouter()
	err := http.ListenAndServe(":3000", router)
	if err != nil {
		panic(err)
	}
}
