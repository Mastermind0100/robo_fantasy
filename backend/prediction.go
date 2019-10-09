package main

import (
	"context"
	log "github.com/sirupsen/logrus"
	"go.mongodb.org/mongo-driver/bson"
)

type UserPredictionData struct {
	Username    string           `json:"username"`
	Predictions []UserPrediction `json:"predictions"`
}

type UserPrediction struct {
	MatchID    int         `json:"matchID"`
	Prediction Predictions `json:"prediction"`
	Result     bool        `json:"result"`
}

type Predictions string

const (
	BLUE Predictions = "blue"
	RED  Predictions = "red"
	NONE Predictions = "none"
)

func AddUserPrediction(username string, p UserPrediction) error {

	filter := bson.D{
		{"username", username},
	}

	predictionCollection := s.db.Collection("predictions")

	f := predictionCollection.FindOne(context.Background(), filter)
	if f.Err() == nil {

		update := bson.D{
			{"$push", bson.D{
				{"predictions", bson.D{
					{"matchid", p.MatchID},
					{"prediction", p.Prediction},
					{"result", p.Result},
				}}}}}
		_, err := predictionCollection.UpdateOne(context.Background(), filter, update)

		if err != nil {
			log.WithField("err", err).Error("Error in adding prediction")
			return err
		}
		return nil
	}

	insertD := bson.D{
		{"username", username},
		{"predictions", bson.A{
			bson.D{
				{"matchid", p.MatchID},
				{"prediction", p.Prediction},
				{"result", p.Result},
			},
		},
		},
	}

	_, err := predictionCollection.InsertOne(context.Background(), insertD)
	if err != nil {
		log.WithField("err", err).Error("Error in adding prediction")
		return err
	}
	return nil
}

func ReadUserPredictions(username string) *UserPredictionData {
	var ud UserPredictionData

	filter := bson.D{
		{"username", username},
	}

	predictionCollection := s.db.Collection("predictions")

	f := predictionCollection.FindOne(context.Background(), filter)
	if f.Err() != nil {
		log.WithField("err", f.Err()).Error("Error in getting prediction")
		return nil
	}

	err := f.Decode(&ud)
	if err != nil {
		log.WithField("err", err).Error("Error decoding in getting prediction")
		return nil
	}

	return &ud
}

func ReadAllUsersPredictions() *[]UserPredictionData {
	var p []UserPredictionData

	predictionCollection := s.db.Collection("predictions")

	f, _ := predictionCollection.Find(context.Background(), bson.D{})

	if f.Err() != nil {
		log.WithField("err", f.Err()).Error("Error in getting all prediction")
		return nil
	}

	err := f.All(context.Background(), &p)

	if err != nil {
		log.WithField("err", err).Error("Error decoding all prediction")
		return nil
	}
	return &p
}

func GetPoints() int { //TODO: Implement the points system, Write crud for predictions
	return 0
}