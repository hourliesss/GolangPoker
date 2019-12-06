package main

import (
	"fmt"
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

func TestStraightOdds(t *testing.T) {
	myHand := []card{card{value: 7, suit: "Hearts"}, card{value: 7, suit: "Diamonds"}}

	commonCards := []card{
		card{
			value: 4,
			suit:  "Spades",
		},
		card{
			value: 5,
			suit:  "Diamonds",
		},
		card{
			value: 10,
			suit:  "Diamonds",
		},
		card{
			value: 8,
			suit:  "Clubs",
		},
	}

	hs := calculateFiveBestCards(append(commonCards, myHand...))

	odds := straightOdds(commonCards, hs, myHand)
	// winning hands :
	// 7 and 8
	if odds != 16.0/2070.0 {
		t.Errorf("1st case: We expected a odd to lose of 0.007 but got %v", odds)
	}

	commonCards = []card{
		card{
			value: 4,
			suit:  "Spades",
		},
		card{
			value: 5,
			suit:  "Diamonds",
		},
		card{
			value: 11,
			suit:  "Diamonds",
		},
		card{
			value: 8,
			suit:  "Clubs",
		},
		card{
			value: 6,
			suit:  "Clubs",
		},
	}

	hs = calculateFiveBestCards(append(commonCards, myHand...))
	odds = straightOdds(commonCards, hs, myHand)
	// winning hands :
	// 7 9 p=4/x
	if odds != 16.0/1980.0 {
		t.Errorf("2nd case: We expected a odd to lose of 0.088 but got %v", odds)
	}

	myHand = []card{card{value: 2, suit: "Hearts"}, card{value: 10, suit: "Clubs"}}

	hs = calculateFiveBestCards(append(commonCards, myHand...))

	odds = straightOdds(commonCards, hs, myHand)
	// winning hands :
	// 7x p=4/x + (x-4)/x * 4/(x-1)
	// 3 7 included in the 7x
	// 2 3 p= 24.0/1980
	if math.Abs(odds-(4.0/45.0+(41.0/45.0)*(4.0/44.000)+24.0/1980.0)) > 0.0001 {
		t.Errorf("3rd case: We expected a odd to lose of 0.1777 but got %v", odds)
	}

	myHand = []card{card{value: 7, suit: "Hearts"}, card{value: 10, suit: "Hearts"}}

	hs = calculateFiveBestCards(append(commonCards, myHand...))

	odds = straightOdds(commonCards, hs, myHand)
	// winning hands :
	// 79 p=24/x
	if odds != 24.0/1980.0 {
		t.Errorf("4th case: We expected a odd to lose of 0 but got %v", odds)
	}

	commonCards = []card{
		card{
			value: 4,
			suit:  "Spades",
		},
		card{
			value: 5,
			suit:  "Diamonds",
		},
		card{
			value: 10,
			suit:  "Diamonds",
		},
		card{
			value: 8,
			suit:  "Clubs",
		},
		card{
			value: 6,
			suit:  "Clubs",
		},
	}

	hs = calculateFiveBestCards(append(commonCards, myHand...))

	odds = straightOdds(commonCards, hs, myHand)
	// winning hands :
	// 7 and 9 p=24/1980
	if odds != 24.0/1980.0 {
		t.Errorf("5th case: We expected a odd to lose of 0.012 but got %v", odds)
	}

	commonCards = []card{
		card{
			value: 6,
			suit:  "Spades",
		},
		card{
			value: 5,
			suit:  "Diamonds",
		},
		card{
			value: 7,
			suit:  "Diamonds",
		},
		card{
			value: 9,
			suit:  "Clubs",
		},
		card{
			value: 10,
			suit:  "Clubs",
		},
	}

	myHand = []card{card{value: 3, suit: "Hearts"}, card{value: 4, suit: "Hearts"}}

	hs = calculateFiveBestCards(append(commonCards, myHand...))

	odds = straightOdds(commonCards, hs, myHand)
	// winning hands :
	// 8x
	// 8 12 uncluded in 8x
	if odds != 4.0/45.0+(164.0/1980.0) {
		t.Errorf("6th case: We expected a odd to lose of 0.012 but got %v", odds)
	}
}

func TestThreeInARowOdds(t *testing.T) {

	myHand := []card{card{value: 5, suit: "Hearts"}, card{value: 12, suit: "Diamonds"}}

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

	odds := threeInARowOdds(groupByValues, hs, myHand, 46)
	// winning hands :
	// 4x 39 winning combi
	if odds != 78.0/1035.0 {
		t.Errorf("1st case: We expected a odd to lose of 0.073 but got %v", odds)
	}

	commonCards = []card{
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

	odds = threeInARowOdds(groupByValues, hs, myHand, 47)
	// winning hands :
	// 44 1010 p=3/1084
	// 55 p=1/1084
	if odds-7.0/1081.0 > 0.00001 {
		t.Errorf("2nd case: We expected a odd to lose of 0.006 but got %v", odds)
	}

	myHand = []card{card{value: 5, suit: "Clubs"}, card{value: 5, suit: "Diamonds"}}

	commonCards = []card{
		card{
			value: 4,
			suit:  "Spades",
		},
		card{
			value: 10,
			suit:  "Diamonds",
		},
		card{
			value: 5,
			suit:  "Hearts",
		},
		card{
			value: 13,
			suit:  "Diamonds",
		},
		card{
			value: 12,
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

	odds = threeInARowOdds(groupByValues, hs, myHand, 45)
	// winning hands :
	// 1313 1212 1010 p=3/990
	if odds != 9.0/990.0 {
		t.Errorf("3rd case: We expected a odd to lose of 0.009 but got %v", odds)
	}

	myHand = []card{card{value: 10, suit: "Clubs"}, card{value: 3, suit: "Diamonds"}}

	commonCards = []card{
		card{
			value: 4,
			suit:  "Spades",
		},
		card{
			value: 10,
			suit:  "Diamonds",
		},
		card{
			value: 5,
			suit:  "Hearts",
		},
		card{
			value: 5,
			suit:  "Diamonds",
		},
		card{
			value: 12,
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

	odds = threeInARowOdds(groupByValues, hs, myHand, 45)
	// winning hands :
	// 5x p=70/990
	if odds != 70.0/990.0 {
		t.Errorf("4th case: We expected a odd to lose of 0.0767 but got %v", odds)
	}
}

func TestTwoPairsOdds(t *testing.T) {

	myHand := []card{card{value: 11, suit: "Clubs"}, card{value: 2, suit: "Clubs"}}

	commonCards := []card{
		card{
			value: 3,
			suit:  "Spades",
		},
		card{
			value: 12,
			suit:  "Diamonds",
		},
		card{
			value: 10,
			suit:  "Diamonds",
		},
		card{
			value: 13,
			suit:  "Hearts",
		},
		card{
			value: 7,
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

	odds := twoPairsOdds(groupByValues, hs, myHand, 45)
	// winning hands :
	// 73 103 107 Q3 Q7 Q10 K3 K7 K10 KQ p=18.0/1980.0
	if odds != 180.0/1980.0 {
		t.Errorf("1st case: We expected a odd to lose of 0.0909 but got %v", odds)
	}

	myHand = []card{card{value: 10, suit: "Clubs"}, card{value: 2, suit: "Clubs"}}

	hs = calculateFiveBestCards(append(commonCards, myHand...))

	odds = twoPairsOdds(groupByValues, hs, myHand, 45)
	// winning hands :
	// 73  Q3 Q7 K3 K7 KQ p=18.0/1980.0
	// 103 107 Q10 K10 p=12.0/1980.0
	if odds-156.0/1980.0 > 0.000001 {
		t.Errorf("2nd case: We expected a odd to lose of 0.07878 but got %v", odds)
	}

	myHand = []card{card{value: 12, suit: "Clubs"}, card{value: 7, suit: "Clubs"}}

	hs = calculateFiveBestCards(append(commonCards, myHand...))

	odds = twoPairsOdds(groupByValues, hs, myHand, 45)
	// winning hands :
	// Q10 p=12.0/1980.0
	// K3 K10 p=18.0/1980.0
	// K7 KQ p= 12.0/1980.0
	if odds-72.0/1980.0 > 0.000001 {
		t.Errorf("3rd case: We expected a odd to lose of 0.036363 but got %v", odds)
	}

	commonCards = []card{
		card{
			value: 3,
			suit:  "Spades",
		},
		card{
			value: 10,
			suit:  "Clubs",
		},
		card{
			value: 10,
			suit:  "Diamonds",
		},
		card{
			value: 13,
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

	odds = twoPairsOdds(groupByValues, hs, myHand, 46)
	// winning hands :
	// 22 44 55 66 88 99 JJ AA p=12.0/2070.0
	// 77 QQ p=6.0/2070.0
	// 3x (x not in 10 K 3) p=2* (3.0/46 * 45-7 / 45)
	// Kx (x not in 10 K)  p=2* (3.0/46 * 45-4 / 45)
	if math.Abs(odds-(108.0/2070.0+(6.0/46.0)*((38.0/45.0)+(41.0/45.0)))) > 0.000001 {
		t.Errorf("4th case: We expected a odd to lose of 0.28695 but got %v", odds)
	}

	myHand = []card{card{value: 3, suit: "Clubs"}, card{value: 7, suit: "Clubs"}}

	hs = calculateFiveBestCards(append(commonCards, myHand...))

	odds = twoPairsOdds(groupByValues, hs, myHand, 46)
	// winning hands :
	// 44 55 66 88 99 JJ QQ AA p=12.0/2070.0
	// 77 p=6.0/2070.0
	// 3A p=16.0/2070.0
	// Kx (x not in 10 K)  p=2* (3.0/46 * 45-4 / 45)
	if math.Abs(odds-(118.0/2070.0+(6.0/46.0)*(41.0/45.0))) > 0.000001 {
		t.Errorf("5th case: We expected a odd to lose of 0.17584 but got %v", odds)
	}

	myHand = []card{card{value: 11, suit: "Clubs"}, card{value: 7, suit: "Clubs"}}

	commonCards = []card{
		card{
			value: 5,
			suit:  "Spades",
		},
		card{
			value: 10,
			suit:  "Clubs",
		},
		card{
			value: 10,
			suit:  "Diamonds",
		},
		card{
			value: 12,
			suit:  "Hearts",
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

	odds = twoPairsOdds(groupByValues, hs, myHand, 45)
	// winning hands :
	// 66 88 99 1313 1414 p=12.0/1980.0
	// 77 1111 p=6.0/1980.0
	// 12 x (x not in 5 10 12)  p=2* (3.0/45 * 44-6 / 44)
	// 13x  (x not in 10 5 12 13 14) p=2* (4.0/45 * 44-14 / 44)
	// 14x (x not in 10 5 12 14) p=2* (4.0/45 * 44-10 / 44)
	if math.Abs(odds-(72.0/1980.0+(6.0/45.0)*(38.0/44.0)+(8.0/45.0)*(64.0/44.0))) > 0.000001 {
		t.Errorf("6th case: We expected a odd to lose of 0.410101010101010 but got %v", odds)
	}

	myHand = []card{card{value: 10, suit: "Clubs"}, card{value: 12, suit: "Clubs"}}

	commonCards = []card{
		card{
			value: 6,
			suit:  "Spades",
		},
		card{
			value: 7,
			suit:  "Clubs",
		},
		card{
			value: 6,
			suit:  "Diamonds",
		},
		card{
			value: 11,
			suit:  "Hearts",
		},
		card{
			value: 10,
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

	odds = twoPairsOdds(groupByValues, hs, myHand, 45)
	// winning hands :
	// 1313 1414 p=12.0/1980.0
	// 1212 p=6.0/1980.0
	// 1013 1014 p=16/1980.0
	// 107 p=12/1980.0
	// 11x  (x not in 6 11) p=2* (3.0/45 * 44-4 / 44)
	if math.Abs(odds-(74.0/1980.0+(6.0/45.0)*(40.0/44.0))) > 0.000001 {
		t.Errorf("7th case: We expected a odd to lose of 0.1585858585 but got %v", odds)
	}

}

func TestPairOdds(t *testing.T) {

	myHand := []card{card{value: 11, suit: "Clubs"}, card{value: 4, suit: "Clubs"}}

	commonCards := []card{
		card{
			value: 3,
			suit:  "Spades",
		},
		card{
			value: 12,
			suit:  "Diamonds",
		},
		card{
			value: 10,
			suit:  "Diamonds",
		},
		card{
			value: 13,
			suit:  "Hearts",
		},
		card{
			value: 7,
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

	odds := pairOdds(groupByValues, hs, myHand, 45)
	// winning hands :
	// 22 55 66 88 99 1414 ===> 6 combin
	// 1111 44 ===> 3 combin
	// 32 35 36 38 39 314  ===> 12 combin
	// 72 75 76 78 79 714  ===> 12 combin
	// 102 105 106 108 109 1014  ===> 12
	// idem 12  ===> 12 combin
	// idem 13  ===> 12 combin
	// 34 311 74 711 104 1011 124 1211 134 1311 ===> 9 combin
	if math.Abs(odds-(6.0*6.0+3.0*2.0+12.0*30.0+9.0*10.0)/990.0) > 0.0001 {
		t.Errorf("1st case: We expected a odd to lose of 0.496969696969 but got %v", odds)
	}

	myHand = []card{card{value: 10, suit: "Clubs"}, card{value: 4, suit: "Clubs"}}

	commonCards = []card{
		card{
			value: 12,
			suit:  "Diamonds",
		},
		card{
			value: 10,
			suit:  "Diamonds",
		},
		card{
			value: 13,
			suit:  "Hearts",
		},
		card{
			value: 7,
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

	odds = pairOdds(groupByValues, hs, myHand, 46)
	// winning hands :
	// 1414 1111 ===> 6 combin
	// 108 109 1011 1014 ===> 8 combin
	// 122 123 125 126 128 129 1211 1214 ===> 12 combin
	// 132 133 135 136 138 139 1311 1314 ===> 12 combin
	// 124 134 ===> 9 combin
	if math.Abs(odds-(6.0*2.0+8.0*4.0+12.0*16.0+9.0*2.0)/1035.0) > 0.0001 {
		t.Errorf("2nd case: We expected a odd to lose of 0.24541 but got %v", odds)
	}

	myHand = []card{card{value: 12, suit: "Clubs"}, card{value: 4, suit: "Clubs"}}

	commonCards = []card{
		card{
			value: 3,
			suit:  "Spades",
		},
		card{
			value: 10,
			suit:  "Diamonds",
		},
		card{
			value: 10,
			suit:  "Diamonds",
		},
		card{
			value: 13,
			suit:  "Hearts",
		},
		card{
			value: 7,
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

	odds = pairOdds(groupByValues, hs, myHand, 45)
	// winning hands :
	// 128 129 1211 ===> 12 combin
	// 142 145 146 148 149 1411 ===> 16 combin
	// 142 1412 ==> 12 combin
	if math.Abs(odds-(12.0*3.0+16.0*6.0+12.0*2.0)/990.0) > 0.0001 {
		t.Errorf("3rd case: We expected a odd to lose of 0.15757 but got %v", odds)
	}

	myHand = []card{card{value: 3, suit: "Clubs"}, card{value: 5, suit: "Clubs"}}

	commonCards = []card{
		card{
			value: 12,
			suit:  "Spades",
		},
		card{
			value: 10,
			suit:  "Diamonds",
		},
		card{
			value: 10,
			suit:  "Diamonds",
		},
		card{
			value: 13,
			suit:  "Hearts",
		},
		card{
			value: 9,
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

	odds = pairOdds(groupByValues, hs, myHand, 45)
	// winning hands :
	// 112 114 116 117 118 142 144 146 147 148 1411 ===> 16 combin
	// 113 115 143 145===> 12 combin
	if math.Abs(odds-(12.0*4.0+16.0*11.0)/990.0) > 0.0001 {
		t.Errorf("3rd case: We expected a odd to lose of 0.1151515 but got %v", odds)
	}
}

func TestHighCardOdds(t *testing.T) {

	myHand := []card{card{value: 11, suit: "Clubs"}, card{value: 4, suit: "Clubs"}}

	commonCards := []card{
		card{
			value: 3,
			suit:  "Spades",
		},
		card{
			value: 12,
			suit:  "Diamonds",
		},
		card{
			value: 10,
			suit:  "Diamonds",
		},
		card{
			value: 9,
			suit:  "Hearts",
		},
		card{
			value: 7,
			suit:  "Hearts",
		},
	}

	odds := highCardOdds(commonCards, myHand, 45)
	// winning hands :
	// 118 ===> 12 combin
	// 132 135 136 138 142 145 146 148 1413 ===> 16 combin
	// 134 1311 144 1411 ===> 12 combin

	if math.Abs(odds-(12.0+9.0*16.0+4.0*12.0)/990.0) > 0.0001 {
		t.Errorf("1st case: We expected a odd to lose of 0.218 but got %v", odds)
	}

	myHand = []card{card{value: 3, suit: "Clubs"}, card{value: 4, suit: "Clubs"}}

	commonCards = []card{
		card{
			value: 8,
			suit:  "Spades",
		},
		card{
			value: 12,
			suit:  "Diamonds",
		},
		card{
			value: 10,
			suit:  "Diamonds",
		},
		card{
			value: 6,
			suit:  "Hearts",
		},
		card{
			value: 7,
			suit:  "Hearts",
		},
	}

	odds = highCardOdds(commonCards, myHand, 45)
	// winning hands :
	// 92 95 112 115 119 132 135 139 1311 142 145 149 1411 1413 ===> 16 combin
	// 93 94 113 114 133 134 143 144===> 12 combin

	if math.Abs(odds-(14.0*16.0+8.0*12.0)/990.0) > 0.0001 {
		t.Errorf("2nd case: We expected a odd to lose of 0.218 but got %v", odds)
	}

	myHand = []card{card{value: 14, suit: "Clubs"}, card{value: 11, suit: "Clubs"}}

	odds = highCardOdds(commonCards, myHand, 45)
	// winning hands :
	// 1413 ===> 12 combin

	if math.Abs(odds-(12.0)/990.0) > 0.0001 {
		t.Errorf("3rd case: We expected a odd to lose of 0.218 but got %v", odds)
	}
}

func TestInstantOddsToLose(t *testing.T) {

	myHand := []card{card{value: 11, suit: "Clubs"}, card{value: 4, suit: "Clubs"}}

	commonCards := []card{
		card{
			value: 3,
			suit:  "Spades",
		},
		card{
			value: 12,
			suit:  "Diamonds",
		},
		card{
			value: 10,
			suit:  "Diamonds",
		},
		card{
			value: 9,
			suit:  "Hearts",
		},
		card{
			value: 7,
			suit:  "Hearts",
		},
	}

	odds1 := instantOddsToLose(myHand, commonCards)

	fmt.Println("the odds to lose with nothing are:", odds1)

	myHand = []card{card{value: 10, suit: "Clubs"}, card{value: 4, suit: "Clubs"}}

	commonCards = []card{
		card{
			value: 3,
			suit:  "Spades",
		},
		card{
			value: 12,
			suit:  "Diamonds",
		},
		card{
			value: 10,
			suit:  "Diamonds",
		},
		card{
			value: 9,
			suit:  "Hearts",
		},
		card{
			value: 7,
			suit:  "Hearts",
		},
	}

	odds2 := instantOddsToLose(myHand, commonCards)

	fmt.Println("the odds to lose with the second best pair are:", odds2)

	commonCards = []card{
		card{
			value: 3,
			suit:  "Spades",
		},
		card{
			value: 12,
			suit:  "Diamonds",
		},
		card{
			value: 10,
			suit:  "Diamonds",
		},
		card{
			value: 9,
			suit:  "Hearts",
		},
		card{
			value: 7,
			suit:  "Diamonds",
		},
	}

	odds3 := instantOddsToLose(myHand, commonCards)

	fmt.Println("the odds to lose with the second best pair and 3 diamonds without having one are:", odds3)

	myHand = []card{card{value: 10, suit: "Clubs"}, card{value: 4, suit: "Diamonds"}}

	odds4 := instantOddsToLose(myHand, commonCards)

	fmt.Println("the odds to lose with the second best pair and 3 diamonds with one are:", odds4)

	myHand = []card{card{value: 10, suit: "Clubs"}, card{value: 9, suit: "Diamonds"}}

	odds5 := instantOddsToLose(myHand, commonCards)

	fmt.Println("the odds to lose with two pairs and 3 diamonds with one are:", odds5)

	myHand = []card{card{value: 9, suit: "Clubs"}, card{value: 9, suit: "Diamonds"}}

	odds6 := instantOddsToLose(myHand, commonCards)

	fmt.Println("the odds to lose with a three in a row and 3 diamonds with one are:", odds6)

	myHand = []card{card{value: 6, suit: "Clubs"}, card{value: 8, suit: "Diamonds"}}

	odds7 := instantOddsToLose(myHand, commonCards)

	fmt.Println("the odds to lose with a bad straight and 3 diamonds with one are:", odds7)

	myHand = []card{card{value: 11, suit: "Clubs"}, card{value: 13, suit: "Diamonds"}}

	odds8 := instantOddsToLose(myHand, commonCards)

	fmt.Println("the odds to lose with a almost best straight and 3 diamonds with one are:", odds8)

	myHand = []card{card{value: 3, suit: "Diamonds"}, card{value: 4, suit: "Diamonds"}}

	odds9 := instantOddsToLose(myHand, commonCards)

	fmt.Println("the odds to lose with a low flush with 2 cards one are:", odds9)

	myHand = []card{card{value: 3, suit: "Diamonds"}, card{value: 11, suit: "Diamonds"}}

	odds10 := instantOddsToLose(myHand, commonCards)

	fmt.Println("the odds to lose with a good flush with 2 cards one are:", odds10)

	myHand = []card{card{value: 3, suit: "Diamonds"}, card{value: 14, suit: "Diamonds"}}

	odds11 := instantOddsToLose(myHand, commonCards)

	fmt.Println("the odds to lose with a the best flush and the best hand are:", odds11)

	commonCards = []card{
		card{
			value: 9,
			suit:  "Spades",
		},
		card{
			value: 12,
			suit:  "Diamonds",
		},
		card{
			value: 10,
			suit:  "Diamonds",
		},
		card{
			value: 7,
			suit:  "Hearts",
		},
		card{
			value: 7,
			suit:  "Diamonds",
		},
	}

	myHand = []card{card{value: 3, suit: "Diamonds"}, card{value: 11, suit: "Diamonds"}}

	odds12 := instantOddsToLose(myHand, commonCards)

	fmt.Println("the odds to lose with a the good flush and one pair in the commons:", odds12)

	commonCards = []card{
		card{
			value: 10,
			suit:  "Spades",
		},
		card{
			value: 12,
			suit:  "Diamonds",
		},
		card{
			value: 10,
			suit:  "Diamonds",
		},
		card{
			value: 7,
			suit:  "Hearts",
		},
		card{
			value: 7,
			suit:  "Diamonds",
		},
	}

	odds13 := instantOddsToLose(myHand, commonCards)

	fmt.Println("the odds to lose with a the good flush and 2 pairs in the commons:", odds13)

	commonCards = []card{
		card{
			value: 10,
			suit:  "Spades",
		},
		card{
			value: 12,
			suit:  "Diamonds",
		},
		card{
			value: 10,
			suit:  "Diamonds",
		},
		card{
			value: 10,
			suit:  "Hearts",
		},
		card{
			value: 7,
			suit:  "Diamonds",
		},
	}

	odds14 := instantOddsToLose(myHand, commonCards)

	fmt.Println("the odds to lose with a the good flush and a three in a row in the commons:", odds14)

	myHand = []card{card{value: 7, suit: "Spades"}, card{value: 11, suit: "Diamonds"}}

	odds15 := instantOddsToLose(myHand, commonCards)

	fmt.Println("the odds to lose with a the low full and a three in a row in the commons:", odds15)

	myHand = []card{card{value: 14, suit: "Spades"}, card{value: 14, suit: "Diamonds"}}

	odds16 := instantOddsToLose(myHand, commonCards)

	fmt.Println("the odds to lose with a the one of the best full and a three in a row in the commons:", odds16)

	myHand = []card{card{value: 12, suit: "Spades"}, card{value: 12, suit: "Hearts"}}

	odds17 := instantOddsToLose(myHand, commonCards)

	fmt.Println("the odds to lose with the best full and a three in a row in the commons:", odds17)

	myHand = []card{card{value: 10, suit: "Clubs"}, card{value: 2, suit: "Hearts"}}

	odds18 := instantOddsToLose(myHand, commonCards)

	fmt.Println("the odds to lose with a four in a row and the best hand are:", odds18)

	commonCards = []card{
		card{
			value: 10,
			suit:  "Spades",
		},
		card{
			value: 12,
			suit:  "Diamonds",
		},
		card{
			value: 10,
			suit:  "Diamonds",
		},
		card{
			value: 10,
			suit:  "Hearts",
		},
		card{
			value: 9,
			suit:  "Diamonds",
		},
	}

	odds19 := instantOddsToLose(myHand, commonCards)

	fmt.Println("the odds to lose with a four in a row and 2 possibilities of straight flush are:", odds19)

	commonCards = []card{
		card{
			value: 12,
			suit:  "Spades",
		},
		card{
			value: 8,
			suit:  "Diamonds",
		},
		card{
			value: 8,
			suit:  "Spades",
		},
		card{
			value: 12,
			suit:  "Hearts",
		},
	}

	myHand = []card{card{value: 3, suit: "Diamonds"}, card{value: 11, suit: "clubs"}}

	odds20 := instantOddsToLose(myHand, commonCards)

	fmt.Println("the odds to lose with 2 pairs in the commons:", odds20)

	commonCards = []card{
		card{
			value: 12,
			suit:  "Spades",
		},
		card{
			value: 8,
			suit:  "Diamonds",
		},
		card{
			value: 8,
			suit:  "Spades",
		},
		card{
			value: 12,
			suit:  "Hearts",
		},
		card{
			value: -8,
			suit:  "Fake",
		},
	}

	hs := calculateFiveBestCards(commonCards)
	fmt.Println(hs.toString())

}
