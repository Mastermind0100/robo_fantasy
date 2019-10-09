package main

import "testing"

func TestAddUserPrediction(t *testing.T) {
	s.InitServer("robowarsTest")
	err := AddUserPrediction("anshu2", UserPrediction{
		MatchID:    1,
		Prediction: "blue",
		Result:     false,
	})
	if err != nil {
		t.Fail()
	}
}

func TestReadUserPredictions(t *testing.T) {
	s.InitServer("robowarsTest")
	d := ReadUserPredictions("anshu2")
	if d == nil {
		t.Fail()
		return
	}
	t.Log(d)
}

func TestReadAllUsersPredictions(t *testing.T) {
	s.InitServer("robowarsTest")
	d := ReadAllUsersPredictions()
	if d == nil {
		t.Fail()
		return
	}
	t.Log(d)
}
