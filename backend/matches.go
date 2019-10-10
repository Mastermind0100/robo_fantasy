package main

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/gorilla/mux"
	log "github.com/sirupsen/logrus"
	"go.mongodb.org/mongo-driver/bson"
	"net/http"
	"strconv"
)

// MatchStatus return different codes for different match status
type MatchStatus int

const (
	StatusUpcoming MatchStatus = 0 + iota
	StatusRed
	StatusBlue
	StatusDraw
)

// Match struct for matches
type Match struct {
	MatchID  int         `json:"id"`
	BlueTeam string      `json:"blue"`
	RedTeam  string      `json:"red"`
	Category Categories  `json:"category"`
	Status   MatchStatus `json:"status"`
}

func UpdateMatchCount() {

	s.mux.Lock()
	defer s.mux.Unlock()

	configCollection := s.db.Collection("config")
	update := bson.D{{
		"$set", bson.D{
			{"matchcount", s.config.MatchCount},
		}}}
	_, err := configCollection.UpdateOne(context.Background(), bson.D{}, update)
	if err != nil {
		log.WithField("err", err).Error("Can't update match count in database")
	}
}

func AddMatch(m Match) error {
	matchCollection := s.db.Collection("matches")
	s.mux.Lock()
	s.config.MatchCount++
	m.MatchID = s.config.MatchCount
	s.mux.Unlock()
	UpdateMatchCount()
	_, err := matchCollection.InsertOne(context.Background(), &m)
	if err != nil {
		log.WithField("err", err).Error("Can't insert Match document")
		return err
	}
	return nil
}

func ReadMatch(id int) *Match {
	matchCollection := s.db.Collection("matches")
	filter := bson.D{{"matchid", id}}
	res := matchCollection.FindOne(context.Background(), filter)
	if res.Err() != nil {
		log.WithFields(log.Fields{"err": res.Err(), "id": id}).Error("Can't read Match document")
		return nil
	}
	var m Match
	err := res.Decode(&m)
	if err != nil {
		log.WithField("err", err).Error("Can't decode match document")
	}
	return &m
}

func UpdateMatch(m Match) error {
	matchCollection := s.db.Collection("matches")
	filter := bson.D{{"matchid", m.MatchID}}
	update := bson.D{{
		"$set", bson.D{
			{"matchid", m.MatchID},
			{"blueteam", m.BlueTeam},
			{"readteam", m.RedTeam},
			{"status", m.Status},
			{"category", m.Category},
		}}}

	_, err := matchCollection.UpdateOne(context.Background(), filter, update)
	if err != nil {
		log.WithFields(log.Fields{"err": err, "id": m.MatchID}).Error("Cant update match")
		return err
	}
	return nil
}

func DeleteMatch(id int) error {
	matchCollection := s.db.Collection("matches")
	filter := bson.D{{"matchid", id}}

	_, err := matchCollection.DeleteOne(context.Background(), filter)
	if err != nil {
		log.WithFields(log.Fields{"err": err, "id": id}).Error("Cant update match")
		return err
	}
	return nil
}

func GetAllMatch() *[]Match {
	matchCollection := s.db.Collection("matches")
	matches := make([]Match, 10)

	res, err := matchCollection.Find(context.Background(), bson.D{})
	if err != nil {
		log.WithField("err", err).Error("Cant read the matches")
		return nil
	}

	err = res.All(context.Background(), &matches)
	if err != nil {
		log.WithField("err", err).Error("Cant read decode all  matches")
		return nil
	}

	return &matches
}

func GetAllMatchMap() map[int]Match {
	matchCollection := s.db.Collection("matches")
	matches := make(map[int]Match)

	res, err := matchCollection.Find(context.Background(), bson.D{})
	if err != nil {
		log.WithField("err", err).Error("Cant read the matches")
		return nil
	}

	var match Match
	for res.Next(context.Background()) {
		err := res.Decode(&match)
		if err != nil {
			log.WithField("err", err).Error("Cant read decode all  matches")
			return nil
		}
		matches[match.MatchID] = match
	}
	if err != nil {
		log.WithField("err", err).Error("Cant read decode all  matches")
		return nil
	}

	return matches
}

func UpdateStatus(id int, status MatchStatus) error {
	m := ReadMatch(id)
	m.Status = status
	err := UpdateMatch(*m)
	if err != nil {
		log.WithFields(log.Fields{"err": err, "id": id}).Error("Can't update the match status")
		return err
	}
	return nil
}

type MatchResponse struct {
	Status int `json:"status"`
}

func PostMatchNew(w http.ResponseWriter, r *http.Request) {
	cat := r.FormValue("category")
	category, err := strconv.Atoi(cat)

	var res MatchResponse

	if err != nil {
		res.Status = 1
		p, _ := json.Marshal(&res)
		_, _ = w.Write(p)
		return
	}

	match := Match{
		BlueTeam: r.FormValue("blue"),
		RedTeam:  r.FormValue("red"),
		Category: Categories(category),
		Status:   StatusUpcoming,
	}

	err = AddMatch(match)
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

func GetMatchDelete(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	i := vars["id"]
	id, err := strconv.Atoi(i)

	var res MatchResponse

	if err != nil {
		res.Status = 1
		p, _ := json.Marshal(&res)
		_, _ = w.Write(p)
		return
	}

	err = DeleteMatch(id)

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

func PostMatchStatus(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	i := vars["id"]
	id, err := strconv.Atoi(i)

	var res MatchResponse

	if err != nil {
		res.Status = 1
		p, _ := json.Marshal(&res)
		_, _ = w.Write(p)
		return
	}

	d := ReadMatch(id)
	if d == nil {
		res.Status = 1
		p, _ := json.Marshal(&res)
		_, _ = w.Write(p)
		return
	}
	st := r.FormValue("status")
	status, err := strconv.Atoi(st)
	if err != nil {
		res.Status = 1
		p, _ := json.Marshal(&res)
		_, _ = w.Write(p)
		return
	}
	err = UpdateStatus(id, MatchStatus(status))

	if err != nil {
		res.Status = 1
		p, _ := json.Marshal(&res)
		_, _ = w.Write(p)
		return
	}

	res.Status = 0
	p, _ := json.Marshal(&res)
	_, _ = w.Write(p)

	s.mux.Lock()
	PerformPostStatusTasks()
	s.mux.Unlock()
	return
}

func PerformPostStatusTasks() {
	s.Matches = GetAllMatchMap()
	CalculateUserPoints()
	CalculateLeaderBoard()
}

func CalculateLeaderBoard() {
	type pointTable struct {
		username string
		point    int
	}
	leader := make([]pointTable, 11)
	for k, v := range s.UserPointMap {
		leader[10] = pointTable{
			username: k,
			point:    v,
		}
		for i := 10; i > 0; i-- {
			if leader[i].point >= leader[i-1].point {
				leader[i], leader[i-1] = leader[i-1], leader[i]

			}
		}
	}

	fmt.Print(leader)

	leaderBoard := make([]UserPoints, 10)
	for i, _ := range leaderBoard {
		u := GetUserByUsername(leader[i].username)
		if u != nil {
			leaderBoard[i] = UserPoints{
				UserDetails: u.UserDetails,
				Points:      leader[i].point,
			}
		}
	}
	s.leaderboard = leaderBoard
}

func CalculateUserPoints() {
	u := *ReadAllUsersPredictions()
	points := make(map[string]int)
	m := GetAllMatchMap()

	for _, user := range u {
		point := 0
		for _, prediction := range user.Predictions {
			res := m[prediction.MatchID].Status
			if (res == 1 && prediction.Prediction == "red") {
				point++
			}
			if (res == 2 && prediction.Prediction == "blue") {
				point++
			}
		}
		points[user.Username] = point
	}
	s.UserPointMap = points
}

type SingleMatchResponse struct {
	Status int   `json:"status"`
	Data   Match `json:"data"`
}

func GetMatchIdDetails(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	i := vars["id"]
	id, err := strconv.Atoi(i)

	var res SingleMatchResponse

	if err != nil {
		res.Status = 1
		p, _ := json.Marshal(&res)
		_, _ = w.Write(p)
		return
	}

	d := ReadMatch(id)
	if d == nil {
		res.Status = 1
		p, _ := json.Marshal(&res)
		_, _ = w.Write(p)
		return
	}

	res.Status = 0
	res.Data = *d

	p, _ := json.Marshal(&res)
	_, _ = w.Write(p)
	return
}

type AllMatchResponse struct {
	Status int     `json:"status"`
	Data   []Match `json:"data"`
}

func GetMatchAll(w http.ResponseWriter, r *http.Request) {

	var res AllMatchResponse

	d := GetAllMatch()
	if d == nil {
		res.Status = 1
		p, _ := json.Marshal(&res)
		_, _ = w.Write(p)
		return
	}

	res.Status = 0
	res.Data = *d

	p, _ := json.Marshal(&res)
	_, _ = w.Write(p)
	return
}
