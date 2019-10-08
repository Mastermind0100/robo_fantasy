package main

import (
	"context"
	"encoding/json"
	"github.com/gorilla/mux"
	log "github.com/sirupsen/logrus"
	"go.mongodb.org/mongo-driver/bson"
	"net/http"
	"strconv"
)

// Bot struct for the bots in the robowars
type Bot struct {
	BotID       int        `json:"id"`
	Name        string     `json:"name"`
	Team        string     `json:"team"`
	Category    Categories `json:"category"`
	Description string     `json:"description"`
	Youtube     string     `json:"video"`
	Image       string     `json:"img"`
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
	bots := make([]Bot, 10)

	botsCollection := s.db.Collection("bots")

	res, _ := botsCollection.Find(context.Background(), bson.D{})
	if res.Err() != nil {
		log.WithField("err", res.Err()).Error("Can't read the bots in the bot collection")
		return nil
	}

	err := res.All(context.Background(), &bots)
	if err != nil {
		log.WithField("err", err).Error("Can't decode the bots in the bot collection")
		return nil
	}

	return &bots
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

type BotResponse struct {
	Status int `json:"status"`
}

func PostBotNew(w http.ResponseWriter, r *http.Request) {
	name := r.FormValue("name")
	team := r.FormValue("team")
	cat := r.FormValue("category")
	category, _ := strconv.Atoi(cat)
	desc := r.FormValue("description")
	youtube := r.FormValue("video")
	img := r.FormValue("image")
	bot := Bot{
		BotID:       0,
		Name:        name,
		Team:        team,
		Category:    Categories(category),
		Description: desc,
		Youtube:     youtube,
		Image:       img,
	}
	err := AddBot(bot)
	var res BotResponse
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

func GetBotDelete(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	var res BotResponse
	id := vars["id"]

	if id == "" {
		res.Status = 1
		p, _ := json.Marshal(&res)
		_, _ = w.Write(p)
		return
	}

	idNum, err := strconv.Atoi(id)
	if err != nil {
		res.Status = 1
		p, _ := json.Marshal(&res)
		_, _ = w.Write(p)
		return
	}

	err = DeleteBot(idNum)
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

type AllBotResponse struct {
	Status int    `json:"status"`
	Bots   *[]Bot `json:"data"`
}

type SingleBotResponse struct {
	Status int  `json:"status"`
	Bots   *Bot `json:"data"`
}

func GetBotAll(w http.ResponseWriter, r *http.Request) {
	b := GetAllBots()
	var res AllBotResponse
	if b == nil {
		res.Status = 1
		p, _ := json.Marshal(&res)
		_, _ = w.Write(p)
		return
	}
	res.Status = 0
	res.Bots = b
	res.Status = 0
	p, _ := json.Marshal(&res)
	_, _ = w.Write(p)
	return
}

func GetBotID(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	var res SingleBotResponse
	id := vars["id"]

	if id == "" {
		res.Status = 1
		p, _ := json.Marshal(&res)
		_, _ = w.Write(p)
		return
	}

	idNum, err := strconv.Atoi(id)
	if err != nil {
		res.Status = 1
		p, _ := json.Marshal(&res)
		_, _ = w.Write(p)
		return
	}

	b := ReadBot(idNum)

	if b == nil {
		res.Status = 1
		p, _ := json.Marshal(&res)
		_, _ = w.Write(p)
		return
	}
	res.Status = 0
	res.Bots = b
	p, _ := json.Marshal(&res)
	_, _ = w.Write(p)
	return
}
