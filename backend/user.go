package main

import (
	"bytes"
	"context"
	log "github.com/sirupsen/logrus"
	"go.mongodb.org/mongo-driver/bson"
	"golang.org/x/crypto/bcrypt"
)

// User struct for the user
type User struct {
	FirstName  string `json:"firstName"`
	LastName   string `json:"lastName"`
	Username   string `json:"username"`
	Email      string `json:"email"`
	CIUsername string `json:"ci-username"`
}

type UserPredictionData struct {
	Username    string           `json:"username"`
	Predictions []UserPrediction `json:"predictions"`
}
type AuthUser struct {
	UserDetails User
	Password    string
}

type UserPrediction struct {
	MatchID    string      `json:"matchID"`
	Prediction Predictions `json:"prediction"`
	Result     bool        `json:"result"`
}

type Predictions string

const (
	BLUE Predictions = "blue"
	RED  Predictions = "red"
	NONE Predictions = "none"
)

func AddUser(u *User, p string) error {
	u.CIUsername = string(bytes.ToUpper([]byte(u.Username)))
	au := AuthUser{
		UserDetails: *u,
		Password:    p,
	}

	gh, err := bcrypt.GenerateFromPassword([]byte(p), 10)
	if err != nil {
		log.WithField("err", err).Error("Error in hashing the password")
		return err
	}
	au.Password = string(gh)

	userColleciton := s.db.Collection("users")

	_, err = userColleciton.InsertOne(context.Background(), au)
	if err != nil {
		log.WithField("err", err).Error("Error in inserting user document")
		return err
	}
	return nil
}

func getUserByUsername(username string) *AuthUser {
	ciUsername := string(bytes.ToUpper([]byte(username)))

	filter := bson.D{{"userdetails", bson.D{{"ciusername", ciUsername}}}}

	userCollection := s.db.Collection("users")

	res := userCollection.FindOne(context.Background(), filter)

	if res.Err() != nil {
		log.WithField("err", res.Err()).Error("Error in getting details for username: ", username)
		return nil
	}
	au := AuthUser{}

	err := res.Decode(&au)
	if err != nil {
		log.WithField("err", err).Error("Error in decoding details for username: ", username)
		return nil
	}
	return &au
}

func updateUser(au *AuthUser) error {

	userCollection := s.db.Collection("users")
	filter := bson.D{{"userdetails", bson.D{{"username", au.UserDetails.CIUsername}}}}

	update := bson.D{{"$set", bson.D{
		{"userdetails", bson.D{
			{"firstname", au.UserDetails.FirstName},
			{"lastname", au.UserDetails.LastName},
		}},
		{"password", au.Password},
	}}}

	_, err := userCollection.UpdateOne(context.Background(), filter, update)

	if err != nil {
		log.WithField("err", err).Error("Error in updating a user")
		return err
	}
	return nil
}

// No Delete Functionality for the user
// We don't want users to get Frustrated and delete their accounts
