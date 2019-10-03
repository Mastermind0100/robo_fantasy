package main

import (
	"github.com/gorilla/handlers"
	"github.com/gorilla/mux"
	log "github.com/sirupsen/logrus"
	"go.mongodb.org/mongo-driver/mongo"
	"net/http"
	"sync"
)

type Server struct {
	mongoClient *mongo.Client
	db          *mongo.Database
	config      Config
	mux         sync.Mutex
}

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

func (s *Server) InitServer(dbName string) {
	s.mongoClient = GetClient("mongodb://localhost:27017")
	s.db = s.mongoClient.Database(dbName)
	s.mux = sync.Mutex{}
	GetConfig()
	UpdateMatchCount()
}

func (s *Server) Shutdown() {
	DisconnectClient(s.mongoClient)
	log.Info("Shutting down the server")
}
