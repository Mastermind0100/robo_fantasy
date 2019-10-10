package main

import "testing"

func TestAddMatch(t *testing.T) {
	s.InitServer("robowarsTest")
	err := AddMatch(Match{
		MatchID:  0,
		BlueTeam: "TestBlue",
		RedTeam:  "TestRed",
		Category: Sixty,
		Status:   StatusUpcoming,
	})
	err = AddMatch(Match{
		MatchID:  0,
		BlueTeam: "TestBlue1",
		RedTeam:  "TestRed1",
		Category: Thirty,
		Status:   StatusUpcoming,
	})
	if err != nil {
		s.Shutdown()
		t.Fail()
	}
	s.Shutdown()
}

func TestUpdateMatchCount(t *testing.T) {
	s.InitServer("robowarsTest")
	s.config.MatchCount = 0
	UpdateMatchCount()
}

func TestReadMatch(t *testing.T) {
	s.InitServer("robowarsTest")
	m := ReadMatch(1)
	t.Log(m)
	s.Shutdown()
}

func TestGetAllMatch(t *testing.T) {
	s.InitServer("robowarsTest")
	m := GetAllMatch()
	if m == nil {
		t.Fail()
	}
	t.Log(m)
	s.Shutdown()
}

func TestUpdateStatus(t *testing.T) {
	s.InitServer("robowarsTest")
	err := UpdateStatus(1, StatusBlue)
	if err != nil {
		t.Fail()
	}
	s.Shutdown()
}

func TestDeleteMatch(t *testing.T) {
	s.InitServer("robowarsTest")
	_ = DeleteMatch(1)
	err := DeleteMatch(2)
	if err != nil {
		t.Fail()
	}
	s.Shutdown()
}

func TestGetAllMatchMap(t *testing.T) {
	s.InitServer("robowarsTest")
	m := GetAllMatchMap()

	if m == nil {
		t.Fail()
	}
	t.Log(m)
}

func TestCalculateLeaderBoard(t *testing.T) {
	CalculateLeaderBoard()
}
