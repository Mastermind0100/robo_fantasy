package main

import "testing"

func TestAddUserPrediction(t *testing.T) {
	s.InitServer("robowarsTest")
	err := AddUserPrediction("anshu1", UserPrediction{
		MatchID:    19,
		Prediction: "red",
		Result:     false,
	})
	if err != nil {
		t.Fail()
	}
}

func TestReadUserPredictions(t *testing.T) {
	s.InitServer("robowarsTest")
	d := ReadUserPredictions("anshu1")
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

func TestReadUserPredictionsMap(t *testing.T) {
	s.InitServer("robowarsTest")
	d := ReadUserPredictionsMap("anshu1")
	if d == nil {
		t.Fail()
		return
	}
	t.Log(d)
}
