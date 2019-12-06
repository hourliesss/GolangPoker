package main

import (
	"fmt"
	"io/ioutil"
	"os"
	"strconv"
	"strings"
)

// calculates the odds that we hav a better hand than the opponent -- if we are before the flop, the odds are taken from the initialOdds2players file
// the odds inside are statistics of wining with each hand calculated in 100,000,000 rounds
func calculateOdds(hand []card, sharedCards []card) float64 {
	switch len(sharedCards) {
	case 0:
		bytes, err := ioutil.ReadFile("initialOdds2Players.csv")
		if err != nil {
			fmt.Println("Error while reading initialOdds2Players.csv:", err)
			os.Exit(1)
		}

		lines := strings.Split(string(bytes), "\r\n")
		for _, line := range lines[1:] {
			data := strings.Split(line, ",")
			if ((strconv.Itoa(hand[0].value) == data[0] && strconv.Itoa(hand[1].value) == data[1]) ||
				(strconv.Itoa(hand[1].value) == data[0] && strconv.Itoa(hand[0].value) == data[1])) &&
				((hand[0].suit == hand[1].suit && data[2] == "1") || (hand[0].suit != hand[1].suit && data[2] == "0")) {
				f, err := strconv.ParseFloat(data[5], 64)
				if err == nil {
					return f
				}
			}
		}

	case 5:
		return 1.0 - instantOddsToLose(hand, sharedCards)
	case 3:
		totalOdds := 0.0
		combi := 0.0
		suits := []string{"Spades", "Diamonds", "Hearts", "Clubs"}
		for i := 2; i < 15; i++ {
			for index, j := range suits {
				c1 := card{suit: j, value: i}
				if !deckContains(hand, c1) && !deckContains(sharedCards, c1) {
					if index < 3 {
						for _, l := range suits[index+1:] {
							c2 := card{suit: l, value: i}
							if !deckContains(hand, c2) && !deckContains(sharedCards, c2) {
								totalOdds += instantOddsToLose(hand, append(append(sharedCards, c1), c2))
								combi += 1.0
							}
						}
					}
					for k := i + 1; k < 15; k++ {
						for _, l := range suits {
							c2 := card{suit: l, value: k}
							if !deckContains(hand, c2) && !deckContains(sharedCards, c2) {
								totalOdds += instantOddsToLose(hand, append(append(sharedCards, c1), c2))
								combi += 1.0
							}
						}
					}
				}
			}
		}
		return totalOdds / combi
	case 4:
		totalOdds := 0.0
		combi := 0.0
		suits := []string{"Spades", "Diamonds", "Hearts", "Clubs"}
		for i := 2; i < 15; i++ {
			for _, j := range suits {
				c1 := card{suit: j, value: i}
				if !deckContains(hand, c1) && !deckContains(sharedCards, c1) {
					totalOdds += instantOddsToLose(hand, append(sharedCards, c1))
					combi += 1.0
				}
			}
		}
		return totalOdds / combi
	}
	return 0

}

// calculates the odds that ont opponent has a better hand than ours
func instantOddsToLose(handInput []card, sharedCards []card) float64 {
	hand := copyDeck(handInput)
	cc := copyDeck(sharedCards)
	hs := calculateFiveBestCards(append(cc, hand...))

	remainingCardsNb := float64(52 - 2 - len(cc))
	if hs.score == 900 {
		return 0
	}
	losingOdds := 0.0
	if hs.score < 900 {
		losingOdds += straightFlushOdds(cc, hand)
	}

	groupByValues := map[int]int{}
	sortCards(cc)
	for _, card := range cc {
		if groupByValues[card.value] == 0 {
			groupByValues[card.value] = 1
		} else {
			groupByValues[card.value]++
		}
	}

	if hs.score < 800 {
		losingOdds += fourOfAKindOdds(groupByValues, hs, hand, remainingCardsNb)
	}
	if hs.score < 700 {

		if len(groupByValues) != len(cc) {
			losingOdds += fullHouseOdds(groupByValues, hs, hand, 52-2-len(cc))
		}
	}
	if hs.score < 600 {
		losingOdds += flushOdds(cc, hand, remainingCardsNb)

	}

	if hs.score < 500 {
		losingOdds += straightOdds(cc, hs, hand)
	}
	if hs.score < 400 {
		losingOdds += threeInARowOdds(groupByValues, hs, hand, remainingCardsNb)
	}
	if hs.score < 300 {
		losingOdds += twoPairsOdds(groupByValues, hs, hand, remainingCardsNb)
	}
	if hs.score < 200 {
		losingOdds += pairOdds(groupByValues, hs, hand, 52-2-len(cc))
	}
	if hs.score < 100 {
		losingOdds += highCardOdds(cc, hand, remainingCardsNb)
	}

	return losingOdds
}

// returns the odds that the opponent has a flush. If we have a flush, it only returns the odds that the opponent has a better flush
func flushOdds(cc []card, hand []card, remainingCardsNb float64) float64 {
	groupBySuits := map[string][]int{}
	odds := 0.0
	for _, card := range cc {
		if len(groupBySuits[card.suit]) == 0 {
			groupBySuits[card.suit] = []int{card.value}
		} else {
			groupBySuits[card.suit] = append(groupBySuits[card.suit], card.value)
		}
	}
	for key, ss := range groupBySuits {
		if len(ss) > 2 {
			commonCards := len(ss)
			if hand[0].suit == key {
				ss = append(ss, hand[0].value)
			}
			if hand[1].suit == key {
				ss = append(ss, hand[1].value)
			}
			valueToDefeat := 2
			if len(ss) > 4 {
				orderSliceOfIntDesc(ss)
				valueToDefeat = ss[len(ss)-1]
				for _, value := range ss[:5] {
					if (hand[0].value == value && hand[0].suit == key) || (hand[1].value == value && hand[1].suit == key) {
						valueToDefeat = value
						break
					}
				}
			}
			if commonCards == 3 && valueToDefeat < 14 {
				for i := valueToDefeat + 1; i < 15; i++ {
					if !contains(ss, i) {
						for j := 2; j < i; j++ {
							if !contains(ss, j) {
								odds += 2.0 / (remainingCardsNb * (remainingCardsNb - 1))
							}
						}
					}

				}
			} else {
				winningCards := 0
				for i := valueToDefeat; i < 15; i++ {
					if !contains(ss, i) {
						winningCards++
					}
				}
				for i := 0; i < winningCards; i++ {
					odds += 2.0 * (remainingCardsNb - float64(winningCards-i)) / (remainingCardsNb * (remainingCardsNb - 1))
				}
			}

		}
	}

	return odds

}

// returns the odds that the opponent has a straight flush. If we have a straight flush, it only returns the odds that the opponent has a better straight flush
func straightFlushOdds(sc []card, h []card) float64 {
	groupOfSameColours := map[string][]int{}
	odds := 0.0
	sortCards(sc)

	for _, card := range sc {
		groupOfSameColours[card.suit] = append(groupOfSameColours[card.suit], card.value)
	}

	for suit, slice := range groupOfSameColours {
		remainingCardsNb := float64(52 - 2 - len(sc))
		if len(slice) > 2 {
			for i := 2; i < 11; i++ {
				holes := []int{}
				matches := []int{}
				j := i
				for len(holes) < 3 && j < i+5 {
					if !contains(slice, j) {
						if (h[0].suit == suit && h[0].value == j) || (h[1].suit == suit && h[1].value == j) {
							matches = append(holes, j)
						} else {
							holes = append(holes, j)
						}
					}
					j++
				}
				if len(holes) == 0 && len(matches) == 2 {
					odds = 0.0
				} else {
					if len(matches) == 0 {
						if len(holes) == 2 {
							odds += 2.0 / (remainingCardsNb * (remainingCardsNb - 1.0))
						}
						if len(holes) == 1 {
							odds += 2.0 / remainingCardsNb
						}
					}
				}

			}
		}
	}

	return odds
}

// returns the odds that the opponent has a four of a kind. If we have a four of a kind, it only returns the odds that the opponent has a better one
func fourOfAKindOdds(groupByValues map[int]int, hs handScore, hand []card, remainingCardsNb float64) float64 {
	odds := 0.0
	for value, number := range groupByValues {
		if hs.score < 700 || hs.card1 < value {
			if number == 4 {
				return 0
			}
			if number == 3 {
				odds += 2.0 / remainingCardsNb
			}
			if number == 2 {
				if hand[0].value != value && hand[1].value != value {
					odds += 2.0 / (remainingCardsNb * (remainingCardsNb - 1.0))
				}
			}
		}

	}
	return odds
}

// returns the odds that the opponent has a full house. If we have a full house, it only returns the odds that the opponent has a better one
func fullHouseOdds(groupByValues map[int]int, hs handScore, hand []card, remainingCardsNb int) float64 {
	odds := 0.0
	fullValueCard1 := 0
	fullValueCard2 := 0
	if hs.score == 600 {
		fullValueCard1 = hs.card1
		fullValueCard2 = hs.card2
	}
	// the community cards are the full
	if 50-remainingCardsNb == len(groupByValues)+3 && remainingCardsNb == 45 {
		threeInARowValue := 0
		pairValue := 0
		for key, value := range groupByValues {
			if value == 3 {
				threeInARowValue = key
			}
			if value == 2 {
				pairValue = key
			}
		}
		if fullValueCard1 == threeInARowValue {
			// in this other case we have the card of the pair, we are unbeattable, only draw possible
			if pairValue > threeInARowValue {
				// if the pair is greater than the three in a row and the player has this card he wins
				odds += 2.0 / float64(remainingCardsNb)
			}
			for i := fullValueCard2 + 1; i < 15; i++ {
				// a pair greater than our best pair can defeat us
				if i != threeInARowValue {
					if hand[0].value != i && hand[1].value != i {
						// 4 cards remaining if we dont have one
						odds += 12.0 / (float64(remainingCardsNb) * (float64(remainingCardsNb) - 1.0))
					} else {
						odds += 6.0 / (float64(remainingCardsNb) * (float64(remainingCardsNb) - 1.0))
					}
				}
			}
		}
	}
	// there is a three in a row in the community cards or 2 pairs
	if 50-remainingCardsNb == len(groupByValues)+2 {
		isThreeInARow := false
		for _, value := range groupByValues {
			if value == 3 {
				isThreeInARow = true
			}
		}
		// three in a row
		if isThreeInARow {
			// if the player has a full house with a different major than the comm cards
			if fullValueCard1 != 0 && groupByValues[fullValueCard1] != 3 {
				for key := range groupByValues {
					if key > fullValueCard1 {
						odds += 2.0 / (float64(remainingCardsNb) * (float64(remainingCardsNb) - 1.0))
					}
				}
			} else {
				for i := maxInt(fullValueCard2+1, 2); i < 15; i++ {
					if groupByValues[i] != 3 {
						// if the player has a full house we are not interested by cards of value lower than his pair; if not fullValueCard2 equals 0
						if groupByValues[i] == 1 {
							// we dont have this card so 3 remains
							odds += 3.0 / float64(remainingCardsNb)
						}
						if groupByValues[i] == 0 {
							// no card of value i is in the common cards
							if hand[0].value != i && hand[1].value != i {
								// 4 cards left in the deck
								odds += 12.0 / (float64(remainingCardsNb) * (float64(remainingCardsNb) - 1.0))
							} else {
								// 3 cards left in the deck - we cant have two, absurd because of i > fullValueCard2
								odds += 6.0 / (float64(remainingCardsNb) * (float64(remainingCardsNb) - 1.0))
							}
						}
					}
				}
				// if the player has a pair in hand that creates a major with a card of the common cards
				for key, value := range groupByValues {
					// if the value of the card is greater than fullValueCard2, it has already has been counted in if groupByValues[i] == 1 above

					if value != 3 && fullValueCard2 >= key && fullValueCard1 < key {
						if hand[0].value != key && hand[1].value != key {
							// 3 cards left in the deck
							odds += 6.0 / (float64(remainingCardsNb) * (float64(remainingCardsNb) - 1.0))
						} else {
							// 2 cards left in the deck
							odds += 2.0 / (float64(remainingCardsNb) * (float64(remainingCardsNb) - 1.0))
						}
					}
				}
			}
		} else {
			// 2 pairs here
			for key, value := range groupByValues {
				// if the key is a common pair
				if value == 2 {
					// if we have a full with this card major
					if fullValueCard1 == key {
						// one card left we can only beat us with a better full using the common cards
						for key2, value2 := range groupByValues {
							// if it's not the same card && this card is better than the opponent second card if he has a full && this card doesnt bring a better full with this card in major
							if key2 != key && key2 > fullValueCard2 && (key2 < key || value2 == 1) {
								odds += 6.0 / (float64(remainingCardsNb) * (float64(remainingCardsNb) - 1.0))
							}
						}
					}
					// the opponent does not have a full with a card that high, one is enough to have the best hand, 2 remaining in the deck
					if fullValueCard1 < key {
						// to have a full house with this key major, the opponent needs this card with another of the common
						for i := 2; i < 15; i++ {
							if i != key {
								if groupByValues[i] == 0 {
									if hand[0].value == i && hand[1].value == i {
										odds += 8.0 / (float64(remainingCardsNb) * (float64(remainingCardsNb) - 1.0))
									} else {
										if hand[0].value == i || hand[1].value == i {
											odds += 12.0 / (float64(remainingCardsNb) * (float64(remainingCardsNb) - 1.0))
										} else {
											odds += 16.0 / (float64(remainingCardsNb) * (float64(remainingCardsNb) - 1.0))
										}
									}
								}
								if groupByValues[i] == 2 && i < key {
									if hand[0].value == i || hand[1].value == i {
										odds += 4.0 / (float64(remainingCardsNb) * (float64(remainingCardsNb) - 1.0))
									} else {
										odds += 8.0 / (float64(remainingCardsNb) * (float64(remainingCardsNb) - 1.0))
									}
								}
								if groupByValues[i] == 1 {
									if hand[0].value == i && hand[1].value == i {
										odds += 4.0 / (float64(remainingCardsNb) * (float64(remainingCardsNb) - 1.0))
									} else {
										if hand[0].value == i || hand[1].value == i {
											odds += 8.0 / (float64(remainingCardsNb) * (float64(remainingCardsNb) - 1.0))
										} else {
											odds += 12.0 / (float64(remainingCardsNb) * (float64(remainingCardsNb) - 1.0))
										}
									}
								}
							}
						}

					}
				}
				// if the key is not a common pair && the opponent doesn't have a full house that strong
				if value == 1 && fullValueCard1 < key {
					if hand[0].value == key || hand[1].value == key {
						odds += 2.0 / (float64(remainingCardsNb) * (float64(remainingCardsNb) - 1.0))
					} else {
						odds += 6.0 / (float64(remainingCardsNb) * (float64(remainingCardsNb) - 1.0))
					}
				}
			}
		}

	}

	// 1 pair in the common cards
	if 50-remainingCardsNb == len(groupByValues)+1 {
		for key, value := range groupByValues {
			// if it is a common pair
			if value == 2 {
				// if the opponent has a full with this card major
				if fullValueCard1 == key {
					// one card left we can only beat him with a better full using the common cards
					for key2 := range groupByValues {
						// if it's not the same card && this card is better than the opponent second card
						if key2 != key && key2 > fullValueCard2 {
							odds += 6.0 / (float64(remainingCardsNb) * (float64(remainingCardsNb) - 1.0))
						}
					}
				}
				// if the opponent has not a full that high but still can have a three of a kind with this card
				if fullValueCard1 < key {
					keyCardsRemaining := 2.0
					if hand[0].value == key || hand[1].value == key {
						keyCardsRemaining -= 1.0
					}
					for key2 := range groupByValues {
						// if it's not the same card
						if key2 != key {
							// the opponent has a pair in hand
							key2CardsRemaining := 3.0
							if hand[0].value == key2 && hand[1].value == key2 {
								key2CardsRemaining -= 2.0

							} else {
								// the opponent has one card in hand
								if hand[0].value == key2 || hand[1].value == key2 {
									key2CardsRemaining -= 1.0
								}
							}
							odds += (2 * keyCardsRemaining * key2CardsRemaining) / (float64(remainingCardsNb) * (float64(remainingCardsNb) - 1.0))
						}

					}

				}
			}
			if value == 1 {
				if fullValueCard1 < key {
					if hand[0].value == key || hand[1].value == key {
						odds += 2.0 / (float64(remainingCardsNb) * (float64(remainingCardsNb) - 1.0))
					} else {
						odds += 6.0 / (float64(remainingCardsNb) * (float64(remainingCardsNb) - 1.0))
					}
				}
			}
		}
	}

	return odds
}

// returns the odds that the opponent has a straight. If we have a straight, it only returns the odds that the opponent has a better one
func straightOdds(commonCards []card, hs handScore, hand []card) float64 {
	simpleCards := []int{}
	odds := 0.0
	remainingCardsNb := float64(50 - len(commonCards))
	for _, card := range commonCards {
		if !contains(simpleCards, card.value) {
			simpleCards = append(simpleCards, card.value)
		}
	}
	orderSliceOfInt(simpleCards)
	cardToDefeat := 0
	if hs.score == 400 {
		cardToDefeat = hs.card1
	}
	if cardToDefeat < 14 {
		alreadyUsedSlice := []int{}
		for i := maxInt(2, cardToDefeat-3); i < 11; i++ {
			holes := []float64{}
			j := i
			var alreadyUsedValue int
			for len(holes) < 3 && j < i+5 {
				if !contains(simpleCards, j) {
					if contains(alreadyUsedSlice, j) {
						break
					}
					alreadyUsedValue = j
					if hand[0].value == j && hand[1].value == j {
						// we store for each missing card the number of cards remaining in the deck
						holes = append(holes, 2.0)
					} else {
						if hand[0].value == j || hand[1].value == j {
							holes = append(holes, 3.0)
						} else {
							holes = append(holes, 4.0)
						}
					}
				}
				j++
			}
			if len(holes) == 2 {
				if !contains(simpleCards, i+5) {
					// if the next card is also in the common cards, we will have a straight at the next value of i with only one card.
					// this straight will include this one, we dont count it twice1
					odds += (2.0 * holes[0] * holes[1]) / (remainingCardsNb * (remainingCardsNb - 1.0))
				}

			}
			if len(holes) == 1 {
				alreadyUsedSlice = append(alreadyUsedSlice, alreadyUsedValue)
				// if holes[0] == 4 so 4 cards remaining:
				// 4/x to have it at the first card + x-4/x to not have it at the first card and then 4/x-1 to have it at the second
				odds += (holes[0] / remainingCardsNb) * (1.0 + ((remainingCardsNb - holes[0]) / (remainingCardsNb - 1.0)))
			}
			// if len == 0 we have the straight -> impossible as we began from the biggest card of the straight (if we have one) minus 3 so the last card would be greater it's absurd

		}
	}

	return odds

}

// returns the odds that the opponent has the best hand with a 3 in a row.
func threeInARowOdds(group map[int]int, hs handScore, hand []card, remainingCardsNb float64) float64 {
	cardToDefeat := 0
	odds := 0.0
	if hs.score == 300 {
		cardToDefeat = hs.card1
	}
	commonCardsNb := 50.0 - remainingCardsNb

	for key, value := range group {
		if key > cardToDefeat {
			if value == 2 {
				// if not there is two pairs in the common cards we already calculate the full house odds
				if commonCardsNb < float64(len(group)+2) {
					// one card is enough to have a 3 in a row and no need to look in the hand, this value could not be in it (absurd)
					// only one because we dont want the four in a row
					// we dont want to catch another card of the common (full)
					cardsToAvoid := 2.0 // the pair
					for key2 := range group {
						if key2 != key {
							if hand[0].value == key2 || hand[1].value == key2 {
								cardsToAvoid += 2.0
							} else {
								cardsToAvoid += 3.0
							}
						}
					}
					odds += (2.0 * 2.0 * (remainingCardsNb - cardsToAvoid)) / (remainingCardsNb * (remainingCardsNb - 1.0))
				}
			}
			if value == 1 {
				// we dont want any pair in the common cards
				if commonCardsNb == float64(len(group)) {
					if hand[0].value == key || hand[1].value == key {
						odds += 2.0 / (remainingCardsNb * (remainingCardsNb - 1.0))
					} else {
						odds += 6.0 / (remainingCardsNb * (remainingCardsNb - 1.0))
					}
				}
			}
		}
	}
	return odds
}

// returns the odds that the opponent has the best hand with 2 pairs.
func twoPairsOdds(group map[int]int, hs handScore, hand []card, remainingCardsNb float64) float64 {
	cardToDefeat1 := 0
	cardToDefeat2 := 0
	odds := 0.0
	if hs.score == 200 {
		cardToDefeat1 = hs.card1
		cardToDefeat2 = hs.card2
	}
	commonCardsNb := 50.0 - remainingCardsNb

	for i := maxInt(cardToDefeat2, 2); i < 15; i++ {
		if group[i] == 1 {
			if cardToDefeat1 < i {
				if float64(len(group)) != commonCardsNb {
					// 1 card of this value is enough to win
					// we need to count all the second cards that give us a better hand to not count this twice or more
					cardsToAvoid := 3.0
					for j := 2; j < 15; j++ {
						if i != j {
							if group[j] == 2 {
								cardsToAvoid += 2.0
							}
							if group[j] == 1 && j > i {
								// we dont want to get this better card we will deal with it later
								cardsToAvoid += 3.0
							}
						}
					}
					odds += 2.0 * (3.0 / remainingCardsNb) * ((remainingCardsNb - cardsToAvoid) / (remainingCardsNb - 1.0))
				} else {
					// we need a second pair
					cardsToGet := 0.0
					for key := range group {
						if key < i {
							if hand[0].value == key || hand[1].value == key || hand[0].value == i || hand[1].value == i {
								cardsToGet += 2.0
							} else {
								cardsToGet += 3.0
							}
						}

					}
					odds += 2.0 * (3.0 / remainingCardsNb) * (cardsToGet / (remainingCardsNb - 1.0))
				}

			}
			if cardToDefeat1 == i {
				// we need this card and the second pair better or equal than the opponenent
				cardsToGet := 0.0
				if group[cardToDefeat2] == 2 {
					// if the second pair is on the common cards

					for j := 2; j < 15; j++ {
						if j != i {
							if group[j] == 1 && j > cardToDefeat2 && j < cardToDefeat1 {
								cardsToGet += 3.0
							}
							if group[j] == 0 && j > hs.remainingCards[0] {
								cardsToGet += 4.0
							}
						}
					}
				} else {
					// in this situation the opponent has 2 pairs with two of the common cards in hand
					// we need a better pair to win
					cardsToGet := 0.0
					for key, value := range group {
						if value == 1 && key > cardToDefeat2 && key < cardToDefeat1 {
							cardsToGet += 3.0
						}
					}
				}
				odds += 2.0 * (2.0 / remainingCardsNb) * (cardsToGet / (remainingCardsNb - 1.0))
			}
			if cardToDefeat2 == i {
				// we need this card and the second pair better or equal than the opponenent remaining card
				cardsToGet := 0.0
				if group[cardToDefeat1] == 2 {
					// if the first pair is on the common cards

					for j := hs.remainingCards[0] + 1; j < 15; j++ {
						if j != i && group[j] == 0 {
							cardsToGet += 4.0
						}
					}
				}
				odds += 2.0 * (2.0 / remainingCardsNb) * (cardsToGet / (remainingCardsNb - 1.0))
			}
		}
		// if this value is not is the common cards, we cant defeat the opponent with this pair if he already got it and we cant have 2 pairs with it if there is no other pair in the commons
		if group[i] == 0 && cardToDefeat2 != i && cardToDefeat1 != i && float64(len(group)) != commonCardsNb {
			if group[cardToDefeat1] == 2 || cardToDefeat1 < i {
				// a pair of this value wins
				if hand[0].value == i || hand[1].value == i {
					// 3 cards remaining
					odds += 6.0 / (remainingCardsNb * (remainingCardsNb - 1.0))
				} else {
					// 4 cards remaining
					odds += 12.0 / (remainingCardsNb * (remainingCardsNb - 1.0))
				}
			}

			if group[cardToDefeat1] == 2 && group[cardToDefeat2] == 2 && i > hs.remainingCards[0] {
				//no need a pair, a high card is enough
				cardsToAvoid := 4.0
				for j := 2; j < 15; j++ {
					if j != i {
						if group[j] == 1 && j > cardToDefeat2 {
							cardsToAvoid += 3.0
						}
						if group[j] == 2 {
							cardsToAvoid += 2.0
						}
						if group[j] == 0 && j > i {
							cardsToAvoid += 4.0
						}
					}
				}
				odds += 2.0 * (4.0 / remainingCardsNb) * ((remainingCardsNb - cardsToAvoid) / (remainingCardsNb - 1.0))
			}
		}
	}

	return odds
}

// returns the odds that the opponent has the best hand with 1 pair.
func pairOdds(group map[int]int, hs handScore, hand []card, remainingCardsNb int) float64 {
	pairToDefeat := 0
	var highCard1ToDefeat int
	var highCard2ToDefeat int
	odds := 0.0
	if hs.score == 100 {
		highCard1ToDefeat = hs.remainingCards[0]
		highCard2ToDefeat = 2
		pairToDefeat = hs.card1
		if group[hs.remainingCards[2]] == 0 {
			highCard1ToDefeat = hs.remainingCards[2]
			if group[hs.remainingCards[1]] == 0 {
				highCard2ToDefeat = hs.remainingCards[1]
			} else {
				highCard2ToDefeat = hs.remainingCards[0]
			}
		} else {
			if group[hs.remainingCards[1]] == 0 {
				highCard1ToDefeat = hs.remainingCards[1]
				highCard2ToDefeat = hs.remainingCards[0]
			}
		}

	}
	commonCardsNb := 50 - remainingCardsNb
	// if not there is a pair inside we are looking for high cards
	if len(group) == commonCardsNb {
		// no pair in the commons
		for i := maxInt(pairToDefeat, 2); i < 15; i++ {
			if group[i] == 1 {
				if pairToDefeat == i {
					cardsToGet := 0.0
					for j := highCard1ToDefeat + 1; j < 15; j++ {
						if group[j] == 0 {
							cardsToGet += 4.0
						}
					}
					odds += (2.0 * (2.0 * cardsToGet)) / (float64(remainingCardsNb * (remainingCardsNb - 1)))
				} else {
					cardsToAvoid := 3
					for key := range group {
						if key != i {
							if hand[0].value == key || hand[1].value == key {
								cardsToAvoid += 2
							} else {
								cardsToAvoid += 3
							}
						}
					}
					odds += 2.0 * (3.0 * float64(remainingCardsNb-cardsToAvoid) / (float64(remainingCardsNb * (remainingCardsNb - 1))))
				}
			}
			if group[i] == 0 && i > pairToDefeat {
				if hand[0].value == i || hand[1].value == i {
					odds += 6.0 / (float64(remainingCardsNb * (remainingCardsNb - 1)))
				} else {
					odds += 12.0 / (float64(remainingCardsNb * (remainingCardsNb - 1)))
				}
			}
		}
	} else {
		// there is a pair in the commons
		// the opponent has this pair because there he cant have 2 pairs here
		// if there is only one card to defeat we just need one card better
		for j := highCard1ToDefeat; j < 15; j++ {
			if j > highCard1ToDefeat {
				if group[j] == 0 {
					cardsToGet := 0.0
					for k := 2; k < j; k++ {
						if group[k] == 0 {
							if hand[0].value == k || hand[1].value == k {
								cardsToGet += 3.0
							} else {
								cardsToGet += 4.0
							}
						}
					}
					odds += (2.0 * 4.0 * cardsToGet) / (float64(remainingCardsNb * (remainingCardsNb - 1)))
				}
			} else {
				// j is the card to defeat
				if group[j] == 0 {
					// if j is in the commons or its the pair you dont want this card
					cardsToGet := 0.0
					for k := highCard2ToDefeat + 1; k < j; k++ {
						if group[k] == 0 {
							cardsToGet += 4.0
						}
					}
					odds += (2.0 * 3.0 * cardsToGet) / (float64(remainingCardsNb * (remainingCardsNb - 1)))
				}
			}
		}
	}
	return odds
}

// returns the odds that the opponent has the best hand without even a pair.
func highCardOdds(cc []card, hand []card, remainingCardsNb float64) float64 {
	var highCard1ToDefeat int
	var highCard2ToDefeat int
	odds := 0.0

	allCardsKnown := []int{}
	for _, c := range append(cc, hand...) {
		allCardsKnown = append(allCardsKnown, c.value)
	}
	orderSliceOfIntDesc(allCardsKnown)

	if index := indexOf(allCardsKnown, maxInt(hand[0].value, hand[1].value)); index < 3 {
		highCard1ToDefeat = allCardsKnown[index]
	} else {
		highCard1ToDefeat = allCardsKnown[3]
	}
	if index := indexOf(allCardsKnown, minInt(hand[0].value, hand[1].value)); index < 4 {
		highCard2ToDefeat = allCardsKnown[index]
	} else {
		highCard2ToDefeat = allCardsKnown[4]
	}

	// if we have highCard1 and it's not in the commons
	if maxInt(hand[0].value, hand[1].value) == highCard1ToDefeat {
		for j := highCard2ToDefeat + 1; j < highCard1ToDefeat; j++ {
			// we dont want a pair and j cant be in the hand so allcardsknown works as the cc
			if !contains(allCardsKnown, j) {
				odds += (2.0 * 3.0 * 4.0) / (remainingCardsNb * (remainingCardsNb - 1.0))
			}

		}
	}
	for i := highCard1ToDefeat + 1; i < 15; i++ {
		if !contains(allCardsKnown, i) {
			cardsToGet := 0.0
			for j := 2; j < i; j++ {
				if hand[0].value == j || hand[1].value == j {
					cardsToGet += 3.0
				} else {
					if !contains(allCardsKnown, j) {
						cardsToGet += 4.0
					}
				}
			}
			odds += (2.0 * 4.0 * cardsToGet) / (remainingCardsNb * (remainingCardsNb - 1.0))
		}
	}

	return odds
}
