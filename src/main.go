package main

import (
	"net/http"

	routes "OSTB_Go/src/router"
)

func main() {
	router := routes.NewRouter()
	if err := http.ListenAndServe(":3000", router); err != nil {
		panic(err)
	}
}
