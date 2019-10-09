package main

import (
	"bytes"
	"context"
	"encoding/json"
	"errors"
	"github.com/dgrijalva/jwt-go"
	"github.com/gorilla/mux"
	log "github.com/sirupsen/logrus"
	"go.mongodb.org/mongo-driver/bson"
	"golang.org/x/crypto/bcrypt"
	"net/http"
)

// User struct for the user
type User struct {
	FirstName  string `json:"firstName"`
	LastName   string `json:"lastName"`
	Username   string `json:"username"`
	Email      string `json:"email"`
	CIUsername string `json:"ci-username"`
}

func (u *User) Valid() error {
	if u.FirstName == "" || u.LastName == "" || u.Username == "" || u.Email == "" {
		return errors.New("Illegal user")
	}
	return nil
}

type AuthUser struct {
	UserDetails User
	Password    string
}

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

func GetUserByUsername(username string) *AuthUser {
	ciUsername := string(bytes.ToUpper([]byte(username)))

	filter := bson.D{{"userdetails.ciusername", ciUsername}}

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

func GetUserByEmail(email string) *AuthUser {
	//ciUsername := string(bytes.ToUpper([]byte(username)))

	filter := bson.D{{"userdetails.email", email}}

	userCollection := s.db.Collection("users")

	res := userCollection.FindOne(context.Background(), filter)

	if res.Err() != nil {
		log.WithField("err", res.Err()).Error("Error in getting details for email: ", email)
		return nil
	}
	au := AuthUser{}

	err := res.Decode(&au)
	if err != nil {
		log.WithField("err", err).Error("Error in decoding details for email: ", email)
		return nil
	}
	return &au
}

func UpdateUser(au *AuthUser) error {

	userCollection := s.db.Collection("users")
	filter := bson.D{{"userdetails.ciusername", au.UserDetails.CIUsername}}
	print(au.UserDetails.CIUsername)
	update := bson.D{{"$set", bson.D{
		{"userdetails.firstname", au.UserDetails.FirstName},
		{"userdetails.lastname", au.UserDetails.LastName},
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

type LoginResponse struct {
	Status int    `json:"status"`
	Token  string `json:"token"`
}

type UserDetailsResponse struct {
	Status    int    `json:"status"`
	FirstName string `json:"firstname"`
	LastName  string `json:"lastname"`
	Email     string `json:"email"`
	Points    int    `json:"points"`
}

func PostUserNew(w http.ResponseWriter, r *http.Request) {
	firstName := r.FormValue("firstname")
	lastName := r.FormValue("lastname")
	email := r.FormValue("email")
	username := r.FormValue("username")
	password := r.FormValue("password")

	var res LoginResponse

	if firstName == "" || lastName == "" || email == "" ||
		username == "" || password == "" {
		res.Status = 2
		res.Token = ""
		p, _ := json.Marshal(&res)
		_, _ = w.Write(p)
		return
	}

	au := AuthUser{
		UserDetails: User{
			FirstName: firstName,
			LastName:  lastName,
			Username:  username,
			Email:     email,
		},
		Password: password,
	}
	already1 := GetUserByUsername(username)
	already2 := GetUserByEmail(email)
	if already1 != nil || already2 != nil {
		res.Status = 1
		res.Token = ""
		p, _ := json.Marshal(&res)
		_, _ = w.Write(p)
		return
	}

	err := AddUser(&au.UserDetails, password)
	if err != nil {
		res.Status = 2
		res.Token = ""
		p, _ := json.Marshal(&res)
		_, _ = w.Write(p)
		return
	}

	token := GetToken(au.UserDetails)
	if token == "" {
		res.Status = 2
		res.Token = ""
		p, _ := json.Marshal(&res)
		_, _ = w.Write(p)
		return
	}

	res.Status = 0
	res.Token = token
	p, _ := json.Marshal(&res)
	_, _ = w.Write(p)
	return
}

func PostUserLogin(w http.ResponseWriter, r *http.Request) {
	email := r.FormValue("email")
	username := r.FormValue("username")
	password := r.FormValue("password")
	var res LoginResponse

	if email == "" && username == "" {
		res.Status = 2
		res.Token = ""
		p, _ := json.Marshal(&res)
		_, _ = w.Write(p)
		return
	}

	var au *AuthUser
	if email != "" {
		au = GetUserByEmail(email)
	}
	if username != "" {
		au = GetUserByUsername(username)
	}
	if au == nil {
		res.Status = 2
		res.Token = ""
		log.Error("No users found")
		p, _ := json.Marshal(&res)
		_, _ = w.Write(p)
		return
	}
	err := bcrypt.CompareHashAndPassword([]byte(au.Password), []byte(password))
	if err != nil {
		res.Status = 1
		res.Token = ""
		p, _ := json.Marshal(&res)
		_, _ = w.Write(p)
		return
	}

	token := GetToken(au.UserDetails)
	if token == "" {
		res.Status = 2
		res.Token = ""
		log.Error("Error in getting token")
		p, _ := json.Marshal(&res)
		_, _ = w.Write(p)
		return
	}

	res.Status = 0
	res.Token = token
	p, _ := json.Marshal(&res)
	_, _ = w.Write(p)
	return
}

func GetUserDetails(w http.ResponseWriter, req *http.Request) {
	vars := mux.Vars(req)
	username := vars["username"]
	au := GetUserByUsername(username)
	var res UserDetailsResponse
	if au == nil {
		res.Status = 1
		p, _ := json.Marshal(&res)
		_, _ = w.Write(p)
		return
	}
	res.Email = au.UserDetails.Email
	res.FirstName = au.UserDetails.FirstName
	res.LastName = au.UserDetails.LastName
	res.Points = GetPoints()

	res.Status = 0
	p, _ := json.Marshal(&res)
	_, _ = w.Write(p)
	return
}

func GetToken(user User) string {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, &user)
	ss, err := token.SignedString(s.signingKey)
	if err != nil {
		log.WithField("err", err).Error("Error in signing jwt token")
		return ""
	}
	return ss
}
