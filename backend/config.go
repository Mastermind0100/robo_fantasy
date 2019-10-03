package main

import (
	"context"
	log "github.com/sirupsen/logrus"
	"go.mongodb.org/mongo-driver/bson"
)

// MatchCount stores the information about the matches.
type Config struct {
	MatchCount   int
	BotCount     int
	CurrentMatch int
}

func GetConfig() {
	configCollection := s.db.Collection("config")
	var conf Config
	res := configCollection.FindOne(context.Background(), bson.D{})
	if res.Err() != nil {
		_, err := configCollection.InsertOne(context.Background(), &conf)
		if err != nil {
			log.WithField("err", err).Error("Can't insert config in database")
		}
		return
	}
	err := res.Decode(&conf)
	if err != nil {
		log.WithField("err", err).Error("Can't read config document from the database")
	}
	s.config = conf
}
