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
	mongoClient  *mongo.Client
	db           *mongo.Database
	config       Config
	mux          sync.Mutex
	signingKey   []byte
	leaderboard  []UserPoint
	UserPointMap map[string]int
	Matches      map[int]Match
	ranking      []UserPoint
	latestMatch  Match
}

func DemoFuncHandler(w http.ResponseWriter, _ *http.Request) {
	_, _ = w.Write([]byte(`hello, this functions hasn't been implemented` +
		`but we plan to implement it in the future`))
}

//getRouter returns the router with all handlers attached
func getHandler() http.Handler {
	router := mux.NewRouter()
	// Various Handling will go here

	headersOk := handlers.AllowedHeaders([]string{"X-Requested-With", "Content-Type", "Origin"})
	originsOk := handlers.AllowedOrigins([]string{"*"})
	methodsOk := handlers.AllowedMethods([]string{"GET", "POST", "PUT", "DELETE"})

	//Operation Related to user
	router.HandleFunc("/user/new", PostUserNew).Methods("POST")
	router.HandleFunc("/user/change", DemoFuncHandler).Methods("POST")
	router.HandleFunc("/user/change/password", DemoFuncHandler).Methods("POST")
	router.HandleFunc("/user/login", PostUserLogin).Methods("POST")
	router.HandleFunc("/user/{username}/details", GetUserDetails).Methods("GET")
	router.HandleFunc("/user/{username}/matches", GetUserMatchDetails).Methods("GET")
	router.HandleFunc("/user/{username}/leaderboard", GetUserLeaderBoard).Methods("GET")

	//Operations Related to matches
	router.HandleFunc("/match/new", PostMatchNew).Methods("POST")
	router.HandleFunc("/match/{id}/edit", DemoFuncHandler).Methods("POST")
	router.HandleFunc("/match/{id}/delete", GetMatchDelete).Methods("GET")
	router.HandleFunc("/match/{id}/status", PostMatchStatus).Methods("POST")
	router.HandleFunc("/match/{id}/details", GetMatchIdDetails).Methods("GET")
	router.HandleFunc("/match/all", GetMatchAll).Methods("GET")

	//Operations Related to Bots
	router.HandleFunc("/bot/new", PostBotNew).Methods("POST")
	router.HandleFunc("/bot/{id}/edit", DemoFuncHandler).Methods("POST")
	router.HandleFunc("/bot/{id}/delete", GetBotDelete).Methods("GET")
	router.HandleFunc("/bot/all", GetBotAll).Methods("GET")
	router.HandleFunc("/bot/{id}", GetBotID).Methods("GET")

	/**
	 TODO when the match status update the calculate points for every user, update current matchCount to get the current match, calculate the leaderboard
	TODO add leser board to server
	*/

	//Operations Related to Betting
	router.HandleFunc("/bet", PostBet).Methods("POST")

	//For LiveLeader board websocket
	router.HandleFunc("/leaderboard", GetLeaderBoard).Methods("GET")

	handler := handlers.CORS(headersOk, originsOk, methodsOk)(router)
	return handler
}

func (s *Server) InitServer(dbName string) {
	s.mongoClient = GetClient("mongodb://localhost:27017")
	s.db = s.mongoClient.Database(dbName)
	s.mux = sync.Mutex{}
	s.signingKey = []byte("achhapolia10")
	GetConfig()
	UpdateMatchCount()
	PerformPostStatusTasks()
}

func (s *Server) Shutdown() {
	DisconnectClient(s.mongoClient)
	log.Info("Shutting down the server")
}
