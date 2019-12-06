package main

type player struct {
	name         string
	cards        []card
	won          int
	roundBet     int
	initialStack int
	hasSpoken    bool
	hasFolded    bool
	isAllIn      bool
	isBot        bool
	decision     float64
}

type person struct {
	name  string
	stack int
	isBot bool
}
