package main

import (
	"testing"
)

func TestNewGame(t *testing.T) {
	person1 := person{
		name:  "Simon",
		stack: 1000,
	}
	person2 := person{
		name:  "Orane",
		stack: 1000,
	}

	persons := []person{person1, person2}

	g := newGame(persons)

	if !g.isGameOn() {
		t.Errorf("Expected the game to be on")
	}

}

func TestNewRound(t *testing.T) {
	person1 := person{
		name:  "Simon",
		stack: 1000,
	}
	person2 := person{
		name:  "Orane",
		stack: 1000,
	}
	person3 := person{
		name:  "Theo",
		stack: 0,
	}
	persons := []person{person1, person2, person3}

	g := newGame(persons)
	g.newRound()
	if g.getNumberOfPlayersAlive() != 2 {
		t.Errorf("Expected to find two players alive but got %v", g.getNumberOfPlayersAlive())
	}
	if len(g.rounds) != 1 {
		t.Errorf("Expected to find 1 round but got %v", len(g.rounds))
	} else {
		if len(g.rounds[0].players) != 2 {
			t.Errorf("Expected to find 2 players in round 0 but got %v", len(g.rounds[0].players))
		}
	}

}

func TestBet(t *testing.T) {
	person1 := person{
		name:  "Simon",
		stack: 1000,
	}
	person2 := person{
		name:  "Orane",
		stack: 1000,
	}
	person3 := person{
		name:  "Theo",
		stack: 0,
	}
	persons := []person{person1, person2, person3}

	g := newGame(persons)
	g.newRound()
	g.newRound()

	g.rounds[1].bet(0, 200)
	g.rounds[1].bet(1, 300)
	g.rounds[1].bet(0, 150)
	if g.rounds[1].maxBet != 350 {
		t.Errorf("Expected to find a max bet of 350 but got %v", g.rounds[1].maxBet)
	}
	if g.rounds[1].pot != 650 {
		t.Errorf("Expected to find a pot of 650 but got %v", g.rounds[1].pot)
	}
	if g.rounds[1].players[0].roundBet != 350 {
		t.Errorf("Expected first player to have betted 350 but got %v", g.rounds[1].players[0].roundBet)
	}
	if g.rounds[1].players[0].isAllIn {
		t.Errorf("Expected first player to not be all in but is")
	}
	if g.rounds[1].players[0].initialStack != 1000 {
		t.Errorf("Expected initial stack of first player to still be 1000 but got %v", g.rounds[1].players[0].initialStack)
	}

}

func TestGetResult(t *testing.T) {
	person1 := person{
		name:  "Simon",
		stack: 1000,
	}
	person2 := person{
		name:  "Orane",
		stack: 1000,
	}
	person3 := person{
		name:  "Theo",
		stack: 100,
	}
	persons := []person{person1, person2, person3}

	g := newGame(persons)
	g.newRound()
	g.newRound()

	g.rounds[1].bet(0, 500)
	g.rounds[1].bet(1, 1000)
	g.rounds[1].bet(2, 100)

	deck := deck{
		createCard(6, "Spades"),
		createCard(7, "Hearts"),
		createCard(14, "Diamonds"),
		createCard(10, "Spades"),
		createCard(8, "Spades"),
	}
	g.rounds[1].sharedCards = deck

	g.rounds[1].players[0].cards = []card{
		createCard(14, "Hearts"),
		createCard(3, "Spades"),
	}
	g.rounds[1].players[1].cards = []card{
		createCard(13, "Spades"),
		createCard(11, "Spades"),
	}
	g.rounds[1].players[2].cards = []card{
		createCard(14, "Clubs"),
		createCard(2, "Hearts"),
	}
	g.rounds[1].players[2].isAllIn = true
	g.rounds[1].players[1].isAllIn = true

	results := g.rounds[1].getResults()
	if results[0][0].playerIndex != 1 || results[1][0].playerIndex != 0 || results[1][1].playerIndex != 2 {
		t.Errorf("Expected a different result")
	}
}

func TestEndRound(t *testing.T) {
	person1 := person{
		name:  "Simon",
		stack: 1000,
	}
	person2 := person{
		name:  "Orane",
		stack: 1000,
	}
	person3 := person{
		name:  "Theo",
		stack: 100,
	}
	persons := []person{person1, person2, person3}

	g := newGame(persons)
	g.newRound()

	g.rounds[0].bet(0, 500)
	g.rounds[0].bet(1, 1000)
	g.rounds[0].bet(2, 100)

	deck := deck{
		createCard(6, "Spades"),
		createCard(7, "Hearts"),
		createCard(14, "Diamonds"),
		createCard(10, "Spades"),
		createCard(8, "Spades"),
	}
	g.rounds[0].sharedCards = deck

	g.rounds[0].players[0].cards = []card{
		createCard(14, "Hearts"),
		createCard(3, "Spades"),
	}
	g.rounds[0].players[2].cards = []card{
		createCard(13, "Spades"),
		createCard(11, "Spades"),
	}
	g.rounds[0].players[1].cards = []card{
		createCard(14, "Clubs"),
		createCard(2, "Hearts"),
	}
	g.endRound()

	if g.players[0].stack != 900 || g.players[1].stack != 900 || g.players[2].stack != 300 {
		t.Errorf("Expected different stacks after the first round")
	}
}

func TestBotMode(t *testing.T) {
	botMode()
}
