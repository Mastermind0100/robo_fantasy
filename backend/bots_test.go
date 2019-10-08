package main

import "testing"

func TestAddBot(t *testing.T) {
	s.InitServer("robowarsTest")
	bot1 := Bot{
		BotID:       0,
		Name:        "bot1",
		Team:        "team1",
		Category:    Sixty,
		Description: "test bot 1",
		Youtube:     "no",
		Image:       "no",
	}
	bot2 := Bot{
		BotID:       0,
		Name:        "bot2",
		Team:        "team2",
		Category:    Thirty,
		Description: "test bot 2",
		Youtube:     "no",
		Image:       "no",
	}
	_ = AddBot(bot1)
	err := AddBot(bot2)
	if err != nil {
		t.Fail()
	}
	s.Shutdown()
}

func TestReadBot(t *testing.T) {
	s.InitServer("robowarsTest")
	bot := ReadBot(1)
	if bot == nil {
		t.Fail()
	}
	t.Log(bot)
	s.Shutdown()
}

func TestGetAllBots(t *testing.T) {
	s.InitServer("robowarsTest")
	bots := GetAllBots()
	if bots == nil {
		t.Fail()
	}
	t.Log(bots)
	s.Shutdown()
}

func TestUpdateBot(t *testing.T) {
	s.InitServer("robowarsTest")
	bot := ReadBot(1)
	bot.Description = "Edited"
	err := UpdateBot(*bot)
	if err != nil {
		t.Fail()
	}
	s.Shutdown()
}

func TestUpdateBotCount(t *testing.T) {
	s.InitServer("robowarsTest")
	s.config.BotCount = 0
	UpdateBotCount()
	s.Shutdown()
}

func TestDeleteBot(t *testing.T) {
	s.InitServer("robowarsTest")
	_ = DeleteBot(1)
	err := DeleteBot(2)
	if err != nil {
		t.Fail()
	}
	s.Shutdown()
}
