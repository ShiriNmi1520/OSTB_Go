package main

import (
	"net/http"

	routes "OSTB_Go/src/router"
)

func main() {

	router := routes.NewRouter()
	http.ListenAndServe(":3000", router)

}
