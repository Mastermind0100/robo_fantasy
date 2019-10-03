package main

import (
	"context"
	log "github.com/sirupsen/logrus"
	"go.mongodb.org/mongo-driver/bson"
)

// MatchStatus return different codes for different match status
type MatchStatus int

const (
	Upcoming MatchStatus = 0 + iota
	Red
	Blue
	Draw
)

// Match struct for matches
type Match struct {
	MatchID  int         `json:"id"`
	BlueTeam string      `json:"blue"`
	RedTeam  string      `json:"red"`
	Category Categories  `json:"category"`
	Status   MatchStatus `json:"status"`
}

// MatchCount stores the information about the matches.
type MatchCount struct {
	Count        int
	CurrentMatch int
}

func GetLatestMatchCount() {
	matchCountCollection := s.db.Collection("MatchCount")
	var mc MatchCount
	res := matchCountCollection.FindOne(context.Background(), bson.D{})
	if res.Err() != nil {
		_, err := matchCountCollection.InsertOne(context.Background(), &mc)
		if err != nil {
			log.WithField("err", err).Error("Can't insert match count in database")
		}
		return
	}
	err := res.Decode(&mc)
	if err != nil {
		log.WithField("err", err).Error("Can't load read document from the database")
	}
	s.mc = mc
}

func UpdateMatchCount() {

	s.mux.Lock()
	defer s.mux.Unlock()

	matchCountCollection := s.db.Collection("MatchCount")
	update := bson.D{{
		"$set", bson.D{
			{"count", s.mc.Count},
			{"currentmatch", s.mc.CurrentMatch},
		}}}
	_, err := matchCountCollection.UpdateOne(context.Background(), bson.D{}, update)
	if err != nil {
		log.WithField("err", err).Error("Can't update match count in database")
	}
}

func AddMatch(m Match) error {
	s.mux.Lock()
	defer s.mux.Unlock()

	matchCollection := s.db.Collection("match")
	s.mc.Count++
	m.MatchID = s.mc.Count

	_, err := matchCollection.InsertOne(context.Background(), &m)
	if err != nil {
		log.WithField("err", err).Error("Can't insert Match document")
		return err
	}
	return nil
}

