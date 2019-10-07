package main

import (
	log "github.com/sirupsen/logrus"
	"net/http"
	"os"
	"os/signal"
	"syscall"
)

var s Server

func main() {
	initLog()
	s.InitServer("robowars")

	go func() {
		port := os.Getenv("PORT")
		if port == "" {
			port = "8080"
		}

		httpHandler := getHandler()

		log.Info("Starting the server on port ", port)
		if err := http.ListenAndServe(":"+port, httpHandler); err != nil {
			log.WithField("error", err).Error("Can't listen on port ", port)
		}
	}() //Starts the server

	signalChannel := make(chan os.Signal)

	signal.Notify(signalChannel, os.Interrupt, os.Kill, syscall.SIGTERM)
	<-signalChannel
	log.Info("Gracefully Shutting down the server")
	s.Shutdown()
}

func initLog() {
	log.SetOutput(os.Stdout)
	log.SetFormatter(&log.JSONFormatter{
		DisableTimestamp: false,
		TimestampFormat:  "02-01-2014 15:04:05",
	})
}
