package main

import (
	"log"
	"net/http"

	"github.com/andriassundskard/rest/handlers"
)

func main() {
	log.Print("Starting server on port 3000")

	// Fire up the server
	http.ListenAndServe("localhost:3000", api.Handlers())
}
