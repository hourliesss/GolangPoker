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
	if len(sharedCards) != 0 {
		return 1.0 - instantOddsToLose(hand, sharedCards)
	}
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

	return 0
}

// calculates the odds that ont opponent has a better hand than ours
func instantOddsToLose(handInput []card, sharedCards []card) float64 {
	hand := copyDeck(handInput)
	sc := copyDeck(sharedCards)
	hs := calculateFiveBestCards(append(sc, hand...))
	remainingCardsNb := float64(52 - 2 - len(sc))
	if hs.score == 900 {
		return 0
	}
	losingOdds := 0.0
	if hs.score < 900 {
		losingOdds += straightFlushOdds(sc, hand)
	}

	groupByValues := map[int]int{}
	sortCards(sc)
	for _, card := range sc {
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

		if len(groupByValues) != len(sc) {
			losingOdds += fullHouseOdds(groupByValues, hs, hand, 52-2-len(sc))
		}
	}
	if hs.score < 600 {
		losingOdds += fullHouseOdds(groupByValues, hs, hand, 52-2-len(sc))

	}

	if hs.score < 500 {
		losingOdds += flushOdds(sc, hand, remainingCardsNb)
	}
	if hs.score < 400 {
		//losingOdds += straightOdds
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
