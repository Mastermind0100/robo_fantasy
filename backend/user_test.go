package main

import (
	"testing"
)

func TestAddUser(t *testing.T) {
	s.InitServer("robowarsTest")
	au := AuthUser{
		UserDetails: User{
			FirstName:  "Test",
			LastName:   "Test",
			Username:   "Test",
			Email:      "test@test.com",
			CIUsername: "",
		},
		Password: "hi12",
	}
	res := AddUser(&au.UserDetails, au.Password)
	if res != nil {
		t.Fail()
	}
}

func TestGetUserByEmail(t *testing.T) {
	s.InitServer("robowarsTest")

	au := GetUserByEmail("test@test.com")
	if au == nil {
		t.Fail()
	}
	t.Log(au)

	au = GetUserByEmail("asdf")
	if au != nil {
		t.Fail()
	}
}

func TestGetUserByUsername(t *testing.T) {
	s.InitServer("robowarsTest")

	au := GetUserByUsername("test")
	if au == nil {
		t.Fail()
	}
	t.Log(au)

	au = GetUserByUsername("23")
	if au != nil {
		t.Fail()
	}
}

func TestUpdateUser(t *testing.T) {
	s.InitServer("robowarsTest")

	au := GetUserByUsername("test")
	au.UserDetails.LastName = "update"
	au.Password = "update"

	err := UpdateUser(au)
	if err != nil {
		t.Fail()
	}
}
