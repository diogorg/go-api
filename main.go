package main

import (
	"api/routes"
	"log"
	"net/http"
)

func main() {
	r := routes.GetRoutes()
	log.Fatal(http.ListenAndServe(":8000", r))
}
