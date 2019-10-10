package main

import (
	"context"
	"encoding/json"
	"errors"
	log "github.com/sirupsen/logrus"
	"go.mongodb.org/mongo-driver/bson"
	"net/http"
	"strconv"
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

		filter1 := bson.D{
			{"username", username},
			{"predictions.matchid", p.MatchID},
		}

		err1 := predictionCollection.FindOne(context.Background(), filter1)
		if err1.Err() == nil {
			log.WithField("err", "Prediction for this match already made").Error("Already made a prediction")
			return errors.New("Prediction for this match already made")
		}

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
func ReadUserPredictionsMap(username string) map[int]UserPrediction {
	var ud UserPredictionData
	res := make(map[int]UserPrediction)
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

	for _, b := range ud.Predictions {
		res[b.MatchID] = b
	}

	return res
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

type BetResponse struct {
	Status int `json:"status"`
}

func PostBet(w http.ResponseWriter, r *http.Request) {
	username := r.FormValue("username")
	id := r.FormValue("matchid")

	matchID, err := strconv.Atoi(id)

	var res BetResponse

	if err != nil {
		res.Status = 1
		p, _ := json.Marshal(&res)
		_, _ = w.Write(p)
		return
	}

	team := r.FormValue("team")

	err = AddUserPrediction(username, UserPrediction{
		MatchID:    matchID,
		Prediction: Predictions(team),
	})

	if err != nil {
		res.Status = 1
		p, _ := json.Marshal(&res)
		_, _ = w.Write(p)
		return
	}

	res.Status = 0
	p, _ := json.Marshal(&res)
	_, _ = w.Write(p)
	return
}
