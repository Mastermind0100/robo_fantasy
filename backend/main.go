package main

import (
	log "github.com/sirupsen/logrus"
	"net/http"
	"os"
)

func main() {
	initLog()
	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}

	httpHandler := getHandler()

	log.Info("Starting the server on port ", port)
	if err := http.ListenAndServe(":"+port, httpHandler); err != nil {
		log.WithField("error",err).Error("Can't listen on port ", port)
	}
}

func initLog() {
	log.SetOutput(os.Stdout)
	log.SetFormatter(&log.JSONFormatter{
		DisableTimestamp: false,
		TimestampFormat:  "02-01-2014 15:04:05",
	})
}
