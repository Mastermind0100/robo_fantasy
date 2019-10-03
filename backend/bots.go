package main

import (
	"context"
	log "github.com/sirupsen/logrus"
	"go.mongodb.org/mongo-driver/bson"
)

// Bot struct for the bots in the robowars
type Bot struct {
	BotID       int        `json:"id"`
	Name        string     `json:"name"`
	Team        string     `json:"name"`
	Category    Categories `json:"category"`
	Description string     `json:"description"`
	Youtube     string     `json:"video"`
	Image       []byte     `json:"img"`
}

// Categories define the type of Categories
type Categories int

// Enum for Different Categories
const (
	Fifteen Categories = 15
	Thirty             = 30
	Sixty              = 60
)

func UpdateBotCount() {
	s.mux.Lock()
	defer s.mux.Unlock()
	configCollection := s.db.Collection("config")

	update := bson.D{{
		"$set", bson.D{
			{"botcount", s.config.BotCount},
		}}}

	_, err := configCollection.UpdateOne(context.Background(), bson.D{}, update)
	if err != nil {
		log.WithField("err", err).Error("Can't update the bot count")
	}
}

func AddBot(bot Bot) error {

	botsCollection := s.db.Collection("bots")

	s.mux.Lock()
	s.config.BotCount++
	bot.BotID = s.config.BotCount
	s.mux.Unlock()
	UpdateBotCount()

	_, err := botsCollection.InsertOne(context.Background(), bot)
	if err != nil {
		log.WithField("err", err).Error("Can't add bot to database")
		return err
	}

	return nil
}

func ReadBot(id int) *Bot {
	var bot Bot

	botsCollection := s.db.Collection("bots")
	filter := bson.D{{"botid", id}}

	res := botsCollection.FindOne(context.Background(), filter)
	if res.Err() != nil {
		log.WithFields(log.Fields{"err": res.Err(), "id": id}).Error("Can't find ID")
		return nil
	}

	err := res.Decode(&bot)
	if err != nil {
		log.WithFields(log.Fields{"err": err, "id": id}).Error("Can't Decode the bot")
		return nil
	}
	return &bot
}

func GetAllBots() *[]Bot {
	var bots *[]Bot

	botsCollection := s.db.Collection("bots")

	res, _ := botsCollection.Find(context.Background(), bson.D{})
	if res.Err() != nil {
		log.WithField("err", res.Err()).Error("Can't read the bots in the bot collection")
		return nil
	}

	err := res.All(context.Background(), res)
	if err != nil {
		log.WithField("err", err).Error("Can't decode the bots in the bot collection")
		return nil
	}

	return bots
}

func UpdateBot(bot Bot) error {

	botsCollection := s.db.Collection("bots")

	filter := bson.D{{"botid", bot.BotID}}
	update := bson.D{{"$set", bson.D{
		{"description", bot.Description},
		{"name", bot.Name},
		{"team", bot.Team},
		{"youtube", bot.Youtube},
		{"category", bot.Category},
		{"image", bot.Image},
	}}}

	_, err := botsCollection.UpdateOne(context.Background(), filter, update)

	if err != nil {
		log.WithFields(log.Fields{"err": err, "id": bot.BotID}).Error("Cant update the bot")
		return err
	}
	return nil
}

func DeleteBot(id int) error {

	botsCollection := s.db.Collection("bots")

	filter := bson.D{{"botid", id}}

	_, err := botsCollection.DeleteOne(context.Background(), filter)
	if err != nil {
		log.WithFields(log.Fields{"err": err, "id": id}).Error("Couldn't delete the bot")
		return err
	}
	return nil
}
