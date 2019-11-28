package main

import (
	"math"
	"testing"
)

func TestSpecialStraight(t *testing.T) {
	hand1 := []card{card{value: 13, suit: "Hearts"}, card{value: 5, suit: "Diamonds"}}
	hand2 := []card{card{value: 7, suit: "Clubs"}, card{value: 4, suit: "Hearts"}}

	sharedCards := []card{
		card{
			value: 8,
			suit:  "Hearts",
		},
		card{
			value: 9,
			suit:  "Hearts",
		},
		card{
			value: 10,
			suit:  "Diamonds",
		},
		card{
			value: 11,
			suit:  "Clubs",
		},
		card{
			value: 6,
			suit:  "Hearts",
		},
	}

	hs1 := calculateFiveBestCards(append(sharedCards, hand1...))
	hs2 := calculateFiveBestCards(append(sharedCards, hand2...))
	i := isBetterHand(hs1, hs2)

	if i != 1 {
		t.Errorf("Expected player 2 to win with straight but only got %v", hs2.toString())
	}
}

func TestStraightFlushOdds(t *testing.T) {
	myHand := []card{card{value: 13, suit: "Hearts"}, card{value: 5, suit: "Diamonds"}}

	commonCards := []card{
		card{
			value: 8,
			suit:  "Hearts",
		},
		card{
			value: 9,
			suit:  "Hearts",
		},
		card{
			value: 10,
			suit:  "Diamonds",
		},
		card{
			value: 11,
			suit:  "Clubs",
		},
		card{
			value: 6,
			suit:  "Hearts",
		},
	}

	odds := straightFlushOdds(commonCards, myHand)

	if odds != 4.0/1980.0 {
		t.Errorf("1st case: We expected a odd to lose of 0.0020 but got %v", odds)
	}

	commonCards = []card{
		card{
			value: 8,
			suit:  "Hearts",
		},
		card{
			value: 9,
			suit:  "Hearts",
		},
		card{
			value: 10,
			suit:  "Diamonds",
		},
		card{
			value: 11,
			suit:  "Clubs",
		},
		card{
			value: 7,
			suit:  "Hearts",
		},
	}

	odds = straightFlushOdds(commonCards, myHand)

	if odds != 6.0/1980.0 {
		t.Errorf("2nd case: We expected a odd to lose of 0.0030 but got %v", odds)
	}

	myHand = []card{card{value: 6, suit: "Hearts"}, card{value: 5, suit: "Hearts"}}

	commonCards = []card{
		card{
			value: 8,
			suit:  "Hearts",
		},
		card{
			value: 9,
			suit:  "Hearts",
		},
		card{
			value: 10,
			suit:  "Diamonds",
		},
		card{
			value: 11,
			suit:  "Clubs",
		},
		card{
			value: 7,
			suit:  "Hearts",
		},
	}

	odds = straightFlushOdds(commonCards, myHand)

	if odds != 2.0/1980.0 {
		t.Errorf("3rd case: We expected a odd to lose of 0.0010 but got %v", odds)
	}

	myHand = []card{card{value: 6, suit: "Diamonds"}, card{value: 5, suit: "Diamonds"}}

	commonCards = []card{
		card{
			value: 8,
			suit:  "Hearts",
		},
		card{
			value: 9,
			suit:  "Clubs",
		},
		card{
			value: 10,
			suit:  "Hearts",
		},
		card{
			value: 11,
			suit:  "Clubs",
		},
		card{
			value: 7,
			suit:  "Diamonds",
		},
	}

	odds = straightFlushOdds(commonCards, myHand)

	if odds != 0.0 {
		t.Errorf("4th case: We expected a odd to lose of 0.0 but got %v", odds)
	}

	commonCards = []card{
		card{
			value: 8,
			suit:  "Hearts",
		},
		card{
			value: 9,
			suit:  "Hearts",
		},
		card{
			value: 2,
			suit:  "Hearts",
		},
		card{
			value: 12,
			suit:  "Hearts",
		},
		card{
			value: 3,
			suit:  "Hearts",
		},
	}

	odds = straightFlushOdds(commonCards, myHand)

	if odds != 2.0/1980.0 {
		t.Errorf("5th case: We expected a odd to lose of 0.0 but got %v", odds)
	}

}

func TestFourOfAKindOdds(t *testing.T) {
	myHand := []card{card{value: 13, suit: "Hearts"}, card{value: 5, suit: "Diamonds"}}

	commonCards := []card{
		card{
			value: 4,
			suit:  "Spades",
		},
		card{
			value: 4,
			suit:  "Diamonds",
		},
		card{
			value: 10,
			suit:  "Diamonds",
		},
		card{
			value: 4,
			suit:  "Clubs",
		},
		card{
			value: 4,
			suit:  "Hearts",
		},
	}

	groupByValues := map[int]int{}
	sortCards(commonCards)
	for _, card := range commonCards {
		if groupByValues[card.value] == 0 {
			groupByValues[card.value] = 1
		} else {
			groupByValues[card.value]++
		}
	}

	hs := calculateFiveBestCards(append(commonCards, myHand...))

	odds := fourOfAKindOdds(groupByValues, hs, myHand, 45)

	if odds != 0.0 {
		t.Errorf("1st case: We expected a odd to lose of 0.0 but got %v", odds)
	}

	commonCards = []card{
		card{
			value: 4,
			suit:  "Spades",
		},
		card{
			value: 4,
			suit:  "Diamonds",
		},
		card{
			value: 10,
			suit:  "Diamonds",
		},
		card{
			value: 5,
			suit:  "Clubs",
		},
		card{
			value: 4,
			suit:  "Hearts",
		},
	}

	groupByValues = map[int]int{}
	sortCards(commonCards)
	for _, card := range commonCards {
		if groupByValues[card.value] == 0 {
			groupByValues[card.value] = 1
		} else {
			groupByValues[card.value]++
		}
	}

	hs = calculateFiveBestCards(append(commonCards, myHand...))

	odds = fourOfAKindOdds(groupByValues, hs, myHand, 45)

	if odds != 2.0/45.0 {
		t.Errorf("2nd case: We expected a odd to lose of 0.0444 but got %v", odds)
	}

	commonCards = []card{
		card{
			value: 4,
			suit:  "Spades",
		},
		card{
			value: 4,
			suit:  "Diamonds",
		},
		card{
			value: 10,
			suit:  "Diamonds",
		},
		card{
			value: 10,
			suit:  "Clubs",
		},
	}

	groupByValues = map[int]int{}
	sortCards(commonCards)
	for _, card := range commonCards {
		if groupByValues[card.value] == 0 {
			groupByValues[card.value] = 1
		} else {
			groupByValues[card.value]++
		}
	}

	hs = calculateFiveBestCards(append(commonCards, myHand...))

	odds = fourOfAKindOdds(groupByValues, hs, myHand, 46)

	if odds != 4.0/(2070.0) {
		t.Errorf("3rd case: We expected a odd to lose of 0.0019 but got %v", odds)
	}

	commonCards = []card{
		card{
			value: 4,
			suit:  "Spades",
		},
		card{
			value: 4,
			suit:  "Diamonds",
		},
		card{
			value: 10,
			suit:  "Diamonds",
		},
		card{
			value: 10,
			suit:  "Clubs",
		},
		card{
			value: 10,
			suit:  "Spades",
		},
	}

	groupByValues = map[int]int{}
	sortCards(commonCards)
	for _, card := range commonCards {
		if groupByValues[card.value] == 0 {
			groupByValues[card.value] = 1
		} else {
			groupByValues[card.value]++
		}
	}

	hs = calculateFiveBestCards(append(commonCards, myHand...))

	odds = fourOfAKindOdds(groupByValues, hs, myHand, 45)

	if odds != 2.0/(1980.0)+2.0/45.0 {
		t.Errorf("4th case: We expected a odd to lose of 0.0030 but got %v", odds)
	}

	myHand = []card{card{value: 10, suit: "Hearts"}, card{value: 5, suit: "Diamonds"}}

	hs = calculateFiveBestCards(append(commonCards, myHand...))

	odds = fourOfAKindOdds(groupByValues, hs, myHand, 45)

	if odds != 0.0 {
		t.Errorf("5th case: We expected a odd to lose of 0.0 but got %v", odds)
	}

	commonCards = []card{
		card{
			value: 14,
			suit:  "Spades",
		},
		card{
			value: 14,
			suit:  "Diamonds",
		},
		card{
			value: 10,
			suit:  "Diamonds",
		},
		card{
			value: 10,
			suit:  "Clubs",
		},
		card{
			value: 10,
			suit:  "Spades",
		},
	}

	groupByValues = map[int]int{}
	sortCards(commonCards)
	for _, card := range commonCards {
		if groupByValues[card.value] == 0 {
			groupByValues[card.value] = 1
		} else {
			groupByValues[card.value]++
		}
	}

	hs = calculateFiveBestCards(append(commonCards, myHand...))

	odds = fourOfAKindOdds(groupByValues, hs, myHand, 45)

	if odds != 2.0/1980.0 {
		t.Errorf("6th case: We expected a odd to lose of 0.00010 but got %v", odds)
	}
}

func TestFullHouseOdds(t *testing.T) {
	myHand := []card{card{value: 10, suit: "Hearts"}, card{value: 12, suit: "Diamonds"}}

	commonCards := []card{
		card{
			value: 4,
			suit:  "Spades",
		},
		card{
			value: 4,
			suit:  "Diamonds",
		},
		card{
			value: 10,
			suit:  "Diamonds",
		},
		card{
			value: 4,
			suit:  "Clubs",
		},
		card{
			value: 5,
			suit:  "Hearts",
		},
	}

	groupByValues := map[int]int{}
	sortCards(commonCards)
	for _, card := range commonCards {
		if groupByValues[card.value] == 0 {
			groupByValues[card.value] = 1
		} else {
			groupByValues[card.value]++
		}
	}

	hs := calculateFiveBestCards(append(commonCards, myHand...))

	odds := fullHouseOdds(groupByValues, hs, myHand, 45)
	// winning hands :
	// 55 p=6/(x*(x-1))
	// 1010 p=2/(x*(x-1))
	// 1212 p=6/(x*(x-1))
	// 1111 1313 1414 p=12/(x*(x-1))
	if odds != 50.0/1980.0 {
		t.Errorf("1st case: We expected a odd to lose of 0.02525 but got %v", odds)
	}

	myHand = []card{card{value: 5, suit: "Spades"}, card{value: 6, suit: "Diamonds"}}

	commonCards = []card{
		card{
			value: 4,
			suit:  "Spades",
		},
		card{
			value: 4,
			suit:  "Diamonds",
		},
		card{
			value: 10,
			suit:  "Diamonds",
		},
		card{
			value: 4,
			suit:  "Clubs",
		},
		card{
			value: 5,
			suit:  "Hearts",
		},
	}

	groupByValues = map[int]int{}
	sortCards(commonCards)
	for _, card := range commonCards {
		if groupByValues[card.value] == 0 {
			groupByValues[card.value] = 1
		} else {
			groupByValues[card.value]++
		}
	}

	hs = calculateFiveBestCards(append(commonCards, myHand...))

	odds = fullHouseOdds(groupByValues, hs, myHand, 45)
	// winning hands :
	// 55 p=2/(x*(x-1))
	// 66 p=6/(x*(x-1))
	// 77 88 99 1111 1212 1313 1414 p=12/(x*(x-1))
	// 10x p=3/x
	if math.Abs(odds-((92.0/(45.0*44.0))+3.0/45.0)) > 0.00001 {
		t.Errorf("2nd case: We expected a odd to lose of 0.1131 but got %v", odds)
	}

	myHand = []card{card{value: 5, suit: "Spades"}, card{value: 5, suit: "Diamonds"}}

	commonCards = []card{
		card{
			value: 4,
			suit:  "Spades",
		},
		card{
			value: 4,
			suit:  "Diamonds",
		},
		card{
			value: 10,
			suit:  "Diamonds",
		},
		card{
			value: 4,
			suit:  "Clubs",
		},
		card{
			value: 5,
			suit:  "Hearts",
		},
	}

	groupByValues = map[int]int{}
	sortCards(commonCards)
	for _, card := range commonCards {
		if groupByValues[card.value] == 0 {
			groupByValues[card.value] = 1
		} else {
			groupByValues[card.value]++
		}
	}

	hs = calculateFiveBestCards(append(commonCards, myHand...))

	odds = fullHouseOdds(groupByValues, hs, myHand, 45)
	// winning hands :
	// 1010 p=6/(x*(x-1))
	if odds == 6.0/1980.0 {
		t.Errorf("3rd case: We expected a odd to lose of 0.0030 but got %v", odds)
	}

	myHand = []card{card{value: 5, suit: "Spades"}, card{value: 5, suit: "Diamonds"}}

	commonCards = []card{
		card{
			value: 4,
			suit:  "Spades",
		},
		card{
			value: 4,
			suit:  "Diamonds",
		},
		card{
			value: 10,
			suit:  "Diamonds",
		},
		card{
			value: 4,
			suit:  "Clubs",
		},
	}

	groupByValues = map[int]int{}
	sortCards(commonCards)
	for _, card := range commonCards {
		if groupByValues[card.value] == 0 {
			groupByValues[card.value] = 1
		} else {
			groupByValues[card.value]++
		}
	}

	hs = calculateFiveBestCards(append(commonCards, myHand...))

	odds = fullHouseOdds(groupByValues, hs, myHand, 46)
	// winning hands :
	// 66 77 88 99 1111 1212 1313 1414 p=12/(x*(x-1))
	// 10x p=3/x
	if math.Abs(odds-((96.0/(45.0*46.0))+3.0/46.0)) > 0.00001 {
		t.Errorf("4th case: We expected a odd to lose of 0.11067 but got %v", odds)
	}

	myHand = []card{card{value: 7, suit: "Spades"}, card{value: 8, suit: "Diamonds"}}

	commonCards = []card{
		card{
			value: 4,
			suit:  "Spades",
		},
		card{
			value: 4,
			suit:  "Diamonds",
		},
		card{
			value: 6,
			suit:  "Diamonds",
		},
		card{
			value: 4,
			suit:  "Clubs",
		},
		card{
			value: 6,
			suit:  "Hearts",
		},
	}

	groupByValues = map[int]int{}
	sortCards(commonCards)
	for _, card := range commonCards {
		if groupByValues[card.value] == 0 {
			groupByValues[card.value] = 1
		} else {
			groupByValues[card.value]++
		}
	}

	hs = calculateFiveBestCards(append(commonCards, myHand...))

	odds = fullHouseOdds(groupByValues, hs, myHand, 45)
	// winning hands :
	// 77 88 p=6/(x*(x-1))
	// 99 1010 1111 1212 1313 1414 p=12/(x*(x-1))
	// 6x p=2/x
	if math.Abs(odds-((84.0/(45.0*44.0))+2.0/45.0)) > 0.00001 {
		t.Errorf("5th case: We expected a odd to lose of 0.086 but got %v", odds)
	}

	myHand = []card{card{value: 12, suit: "Spades"}, card{value: 11, suit: "Diamonds"}}

	commonCards = []card{
		card{
			value: 10,
			suit:  "Spades",
		},
		card{
			value: 10,
			suit:  "Spades",
		},
		card{
			value: 10,
			suit:  "Diamonds",
		},
	}

	groupByValues = map[int]int{}
	sortCards(commonCards)
	for _, card := range commonCards {
		if groupByValues[card.value] == 0 {
			groupByValues[card.value] = 1
		} else {
			groupByValues[card.value]++
		}
	}

	hs = calculateFiveBestCards(append(commonCards, myHand...))

	odds = fullHouseOdds(groupByValues, hs, myHand, 47)
	// winning hands :
	// 22 33 44 55 66 77 88 99 1313 1414 p=12/(x*(x-1))
	// 1111 1212 p=6/(x*(x-1))
	if math.Abs(odds-(132.0/2162.0)) > 0.00001 {
		t.Errorf("6th case: We expected a odd to lose of 0.061 but got %v", odds)
	}

	myHand = []card{card{value: 12, suit: "Spades"}, card{value: 11, suit: "Diamonds"}}

	commonCards = []card{
		card{
			value: 4,
			suit:  "Spades",
		},
		card{
			value: 10,
			suit:  "Spades",
		},
		card{
			value: 10,
			suit:  "Diamonds",
		},
		card{
			value: 4,
			suit:  "Clubs",
		},
	}

	groupByValues = map[int]int{}
	sortCards(commonCards)
	for _, card := range commonCards {
		if groupByValues[card.value] == 0 {
			groupByValues[card.value] = 1
		} else {
			groupByValues[card.value]++
		}
	}

	hs = calculateFiveBestCards(append(commonCards, myHand...))

	odds = fullHouseOdds(groupByValues, hs, myHand, 46)
	// winning hands :
	// 42 43 45 46 47 48 49 413 414 102 103 105 106 107 108 109 1013 1014 p=16/(x(x-1))
	// 411 412 1011 1012 p=12/(x(x-1))
	// 104  p=8/(x(x-1))
	if math.Abs(odds-((16.0*18.0)/2070.0+48.0/2070.0+8.0/2070.0)) > 0.00001 {
		t.Errorf("7th case: We expected a odd to lose of 0.0850 but got %v", odds)
	}

	myHand = []card{card{value: 4, suit: "Hearts"}, card{value: 11, suit: "Diamonds"}}

	hs = calculateFiveBestCards(append(commonCards, myHand...))

	odds = fullHouseOdds(groupByValues, hs, myHand, 46)
	// winning hands :
	// 102 103 105 106 107 108 109 1012 1013 1014 p=16/(x(x-1))
	// 104 p=4/(x(x-1))
	// 1011 p=12/(x(x-1))
	if math.Abs(odds-((16.0*10.0)/2070.0+4.0/2070.0+12.0/2070.0)) > 0.00001 {
		t.Errorf("8th case: We expected a odd to lose of 0.044 but got %v", odds)
	}

	commonCards = []card{
		card{
			value: 4,
			suit:  "Spades",
		},
		card{
			value: 10,
			suit:  "Spades",
		},
		card{
			value: 10,
			suit:  "Diamonds",
		},
		card{
			value: 4,
			suit:  "Clubs",
		},
		card{
			value: 5,
			suit:  "Spades",
		},
	}

	groupByValues = map[int]int{}
	sortCards(commonCards)
	for _, card := range commonCards {
		if groupByValues[card.value] == 0 {
			groupByValues[card.value] = 1
		} else {
			groupByValues[card.value]++
		}
	}

	hs = calculateFiveBestCards(append(commonCards, myHand...))

	odds = fullHouseOdds(groupByValues, hs, myHand, 45)
	// winning hands :
	// 55  p=6/(x(x-1))
	// 102 103 106 107 108 109 1012 1013 1014 p=16/(x(x-1))
	// 104 p=4/(x(x-1))
	// 1011 105 p=12/(x(x-1))
	if math.Abs(odds-(6.0/1980.0+(9.0*16.0)/1980.0+28.0/1980.0)) > 0.00001 {
		t.Errorf("9th case: We expected a odd to lose of 0.08383 but got %v", odds)
	}

	myHand = []card{card{value: 10, suit: "Hearts"}, card{value: 11, suit: "Diamonds"}}

	hs = calculateFiveBestCards(append(commonCards, myHand...))

	odds = fullHouseOdds(groupByValues, hs, myHand, 45)
	// winning hands :
	// 105 p=6/(x(x-1))
	if math.Abs(odds-6.0/1980.0) > 0.00001 {
		t.Errorf("10th case: We expected a odd to lose of 0.00303 but got %v", odds)
	}

	myHand = []card{card{value: 14, suit: "Spades"}, card{value: 11, suit: "Diamonds"}}

	commonCards = []card{
		card{
			value: 4,
			suit:  "Spades",
		},
		card{
			value: 10,
			suit:  "Spades",
		},
		card{
			value: 10,
			suit:  "Diamonds",
		},
		card{
			value: 14,
			suit:  "Clubs",
		},
		card{
			value: 13,
			suit:  "Clubs",
		},
	}

	groupByValues = map[int]int{}
	sortCards(commonCards)
	for _, card := range commonCards {
		if groupByValues[card.value] == 0 {
			groupByValues[card.value] = 1
		} else {
			groupByValues[card.value]++
		}
	}

	hs = calculateFiveBestCards(append(commonCards, myHand...))

	odds = fullHouseOdds(groupByValues, hs, myHand, 45)
	// winning hands :
	// 104 1013 p=12/(x(x-1))
	// 1014 p=8/(x(x-1))
	// 1313 44 p=6/(x(x-1))
	// 1414  p=2/(x(x-1))
	if math.Abs(odds-46.0/1980.0) > 0.00001 {
		t.Errorf("11th case: We expected a odd to lose of 0.0232 but got %v", odds)
	}

	myHand = []card{card{value: 10, suit: "Clubs"}, card{value: 11, suit: "Diamonds"}}

	hs = calculateFiveBestCards(append(commonCards, myHand...))

	odds = fullHouseOdds(groupByValues, hs, myHand, 45)
	// winning hands :
	// 104 1014 1013 p=6/(x(x-1))
	// 44 1414 1313 p=6/(x(x-1))

	if math.Abs(odds-36.0/1980.0) > 0.00001 {
		t.Errorf("12th case: We expected a odd to lose of 0.01818 but got %v", odds)
	}

	myHand = []card{card{value: 10, suit: "Clubs"}, card{value: 13, suit: "Diamonds"}}

	hs = calculateFiveBestCards(append(commonCards, myHand...))

	odds = fullHouseOdds(groupByValues, hs, myHand, 45)
	// winning hands :
	// 1014 p=6/(x(x-1))
	// 1313 p=2/(x(x-1))
	// 1414 p=6/(x(x-1))
	if math.Abs(odds-14.0/1980.0) > 0.00001 {
		t.Errorf("13th case: We expected a odd to lose of 0.0070 but got %v", odds)
	}

	myHand = []card{card{value: 13, suit: "Spades"}, card{value: 13, suit: "Diamonds"}}

	hs = calculateFiveBestCards(append(commonCards, myHand...))

	odds = fullHouseOdds(groupByValues, hs, myHand, 45)
	// winning hands :
	// 1414 p=6/(x(x-1))
	if math.Abs(odds-6.0/1980.0) > 0.00001 {
		t.Errorf("14th case: We expected a odd to lose of 0.0030 but got %v", odds)
	}
	commonCards = []card{
		card{
			value: 10,
			suit:  "Spades",
		},
		card{
			value: 10,
			suit:  "Diamonds",
		},
		card{
			value: 14,
			suit:  "Clubs",
		},
		card{
			value: 13,
			suit:  "Clubs",
		},
	}

	groupByValues = map[int]int{}
	sortCards(commonCards)
	for _, card := range commonCards {
		if groupByValues[card.value] == 0 {
			groupByValues[card.value] = 1
		} else {
			groupByValues[card.value]++
		}
	}

	hs = calculateFiveBestCards(append(commonCards, myHand...))

	odds = fullHouseOdds(groupByValues, hs, myHand, 46)
	// winning hands :
	// 1414 p=6/(x(x-1))
	if math.Abs(odds-6.0/2070.0) > 0.00001 {
		t.Errorf("14th case: We expected a odd to lose of 0.00289 but got %v", odds)
	}
}

func TestFlushOdds(t *testing.T) {
	myHand := []card{card{value: 13, suit: "Hearts"}, card{value: 5, suit: "Diamonds"}}

	commonCards := []card{
		card{
			value: 8,
			suit:  "Hearts",
		},
		card{
			value: 9,
			suit:  "Hearts",
		},
		card{
			value: 10,
			suit:  "Diamonds",
		},
	}

	odds := flushOdds(commonCards, myHand, 47)

	if odds != 0.0 {
		t.Errorf("1st case: We expected an odd of 0 but got %v", odds)
	}

	commonCards = []card{
		card{
			value: 8,
			suit:  "Hearts",
		},
		card{
			value: 9,
			suit:  "Hearts",
		},
		card{
			value: 10,
			suit:  "Hearts",
		},
	}

	odds = flushOdds(commonCards, myHand, 47)
	// winning hands: 32 42 43 52 53 54...
	if odds != 72.0/2162.0 {
		t.Errorf("2nd case: We expected an odd of 0.0333 but got %v", odds)
	}

	myHand = []card{card{value: 4, suit: "Hearts"}, card{value: 14, suit: "Hearts"}}

	odds = flushOdds(commonCards, myHand, 47)
	if odds != 0.0 {
		t.Errorf("3rd case: We expected an odd of 0 but got %v", odds)
	}

	myHand = []card{card{value: 4, suit: "Hearts"}, card{value: 7, suit: "Hearts"}}

	odds = flushOdds(commonCards, myHand, 47)
	// winning hands: 32 42 43 52 53 54...
	if odds != 44.0/2162.0 {
		t.Errorf("4th case: We expected an odd of 0.0203 but got %v", odds)
	}

	commonCards = []card{
		card{
			value: 8,
			suit:  "Hearts",
		},
		card{
			value: 9,
			suit:  "Hearts",
		},
		card{
			value: 5,
			suit:  "Hearts",
		},
		card{
			value: 12,
			suit:  "Hearts",
		},
		card{
			value: 10,
			suit:  "Spades",
		},
	}

	myHand = []card{card{value: 11, suit: "Hearts"}, card{value: 7, suit: "Diamonds"}}

	odds = flushOdds(commonCards, myHand, 45)
	// winning hands: 32 42 43 52 53 54...
	if math.Abs(odds-4.0/45.0+2/1980.0) > 0.00001 {
		t.Errorf("5th case: We expected an odd of 0.0878 but got %v", odds)
	}

	commonCards = []card{
		card{
			value: 8,
			suit:  "Hearts",
		},
		card{
			value: 9,
			suit:  "Hearts",
		},
		card{
			value: 5,
			suit:  "Hearts",
		},
		card{
			value: 12,
			suit:  "Hearts",
		},
	}

	myHand = []card{card{value: 11, suit: "Spades"}, card{value: 7, suit: "Diamonds"}}

	odds = flushOdds(commonCards, myHand, 46)
	if math.Abs(odds-18.0/46.0+72.0/2070.0) > 0.00001 {
		t.Errorf("6th case: We expected an odd of 0.356 but got %v", odds)
	}
}
