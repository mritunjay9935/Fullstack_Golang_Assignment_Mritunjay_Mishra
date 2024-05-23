package main

import (
	"apica-assignment/api"
	"log"
	"net/http"
)

func main() {
	router := api.NewRouter()

	log.Fatal(http.ListenAndServe(":8000", router))
}
