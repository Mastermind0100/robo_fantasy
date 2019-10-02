package main

import (
	"github.com/gorilla/handlers"
	"github.com/gorilla/mux"
	"net/http"
)

//getRouter returns the router with all handlers attached
func getHandler() http.Handler {
	router := mux.NewRouter()

	// Various Handling will go here

	headersOk := handlers.AllowedHeaders([]string{"X-Requested-With", "Content-Type", "Origin"})
	originsOk := handlers.AllowedOrigins([]string{"*"})
	methodsOk := handlers.AllowedMethods([]string{"GET", "POST", "PUT", "DELETE"})

	handler := handlers.CORS(headersOk, originsOk, methodsOk)(router)
	return handler
}
