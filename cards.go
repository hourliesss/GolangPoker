package main

import (
	"fmt"
	"math/rand"
	"strings"
	"time"
)

type card struct {
	value int
	suit  string
}

type handValue struct {
	points   int
	highCard int
}
type handScore struct {
	score            int
	suit             string
	card1            int
	card2            int
	remainingCards   []int
	playerIndex      int
	isPrecedentEqual bool
}

type deck = []card

func newDeck() deck {
	suitValues := []string{"Spades", "Clubs", "Hearts", "Diamonds"}
	var d deck
	for i := 2; i < 15; i++ {
		for _, suit := range suitValues {
			d = append(d, createCard(i, suit))
		}
	}
	shuffle(d)
	return d
}

func createCard(value int, suit string) card {
	c := card{
		value: value,
		suit:  suit,
	}
	return c
}

func printDeck(d deck) {
	s := ""
	for _, card := range d {
		s += card.toString() + " "
	}
	fmt.Printf(s + "\n")
}

func (c card) toString() string {
	valuesMapping := map[int]string{
		2:  "2",
		3:  "3",
		4:  "4",
		5:  "5",
		6:  "6",
		7:  "7",
		8:  "8",
		9:  "9",
		10: "10",
		11: "J",
		12: "Q",
		13: "K",
		14: "A",
	}

	suitMapping := map[string]string{
		"Spades":   "♠",
		"Hearts":   "♥",
		"Diamonds": "♦",
		"Clubs":    "♣",
	}
	return "\t" + suitMapping[c.suit] + valuesMapping[c.value]

}

func (s handScore) toString() string {
	valuesMappingInt := map[int]string{
		2:  "Two",
		3:  "Three",
		4:  "Four",
		5:  "Five",
		6:  "Six",
		7:  "Seven",
		8:  "Eight",
		9:  "Nine",
		10: "Ten",
		11: "Jalet",
		12: "Queen",
		13: "King",
		14: "Ace",
	}
	valuesMapping := map[int]string{
		0:    "a High Card " + valuesMappingInt[s.card1],
		100:  "a Single Pair of " + valuesMappingInt[s.card1] + "s",
		200:  "Two Pairs of " + valuesMappingInt[s.card1] + "s" + " and " + valuesMappingInt[s.card2] + "s",
		300:  "a Three Of A Kind of " + valuesMappingInt[s.card1] + "s",
		400:  "a " + valuesMappingInt[s.card1] + "-high Straigth",
		500:  "a Flush of " + s.suit,
		600:  "a Full House, " + valuesMappingInt[s.card1] + "s" + " over " + valuesMappingInt[s.card2] + "s",
		700:  "a Four Of A Kind of " + valuesMappingInt[s.card1] + "s",
		800:  "a " + valuesMappingInt[s.card1] + "-high Straight Flush of " + s.suit,
		900:  "Royal Flush of " + s.suit,
		-100: "No Value",
	}

	return valuesMapping[s.score]

}

func showCards(sp []player) {
	realPlayers := 0
	for _, p := range sp {
		if !p.isBot {
			realPlayers++
		}
	}
	for _, p := range sp {
		if realPlayers > 1 {
			if p.isBot {
				//fmt.Print("\t", p.name, ", your cards are:\n\n")
				//fmt.Println(p.cards[0].toString(), " ", p.cards[1].toString())
			} else {
				c := ""
				for strings.ToLower(c) != "ok" {
					fmt.Print("\t", p.name, " press OK to see your cards ")
					c = strings.ToLower(readFromTerminal())
				}
				fmt.Print("\t", p.name, ", your cards are:\n\n")
				fmt.Println(p.cards[0].toString(), " ", p.cards[1].toString())
				c = ""
				for strings.ToLower(c) != "ok" {
					fmt.Print("\n\n\t", p.name, " press OK to hide your cards ")
					c = strings.ToLower(readFromTerminal())
				}

				printSimPoker()
			}
		} else {
			if realPlayers == 1 {
				if !p.isBot {
					fmt.Print("\t", p.name, ", your cards are:\n\n")
					fmt.Println(p.cards[0].toString(), " ", p.cards[1].toString())
				}
			} else {
				fmt.Print("\t", p.name, ",'s cards are:\n\n")
				fmt.Println(p.cards[0].toString(), " ", p.cards[1].toString())
			}

		}
	}
}

func printSimPoker() {
	fmt.Println("\n\n\n\t♠♠♠♠♠♠♠♠♠♠♠♠♠♠♠♠♠♠♠♠♠♠♠♠♠♠♠♠♠♠♠♠♠♠♠♠♠♠♠♠♠♠♠♠♠♠♠♠♠♠♠♠♠♠♠♠♠♠♠♠♠♠♠♠♠♠♠♠♠♠♠♠♠♠♠♠♠")
	fmt.Println("\n\n\t♥♥♥♥♥♥♥♥♥♥♥♥♥♥♥♥♥♥♥♥♥♥♥♥♥♥♥♥♥♥♥♥♥♥♥♥♥♥♥♥♥♥♥♥♥♥♥♥♥♥♥♥♥♥♥♥♥♥♥♥♥♥♥♥♥♥♥♥♥♥♥♥♥♥♥♥♥")
	fmt.Println("\n\n\t♣♣♣♣♣♣♣♣♣♣♣♣♣♣♣♣♣♣♣♣♣♣♣♣♣♣♣♣♣♣♣♣♣♣♣♣♣♣♣♣♣♣♣♣♣♣♣♣♣♣♣♣♣♣♣♣♣♣♣♣♣♣♣♣♣♣♣♣♣♣♣♣♣♣♣♣♣")
	fmt.Println("\n\n\t♦♦♦♦♦♦♦♦♦♦♦♦♦♦♦♦♦♦♦♦♦♦♦♦♦♦♦♦♦♦♦♦♦♦♦♦♦♦♦♦♦♦♦♦♦♦♦♦♦♦♦♦♦♦♦♦♦♦♦♦♦♦♦♦♦♦♦♦♦♦♦♦♦♦♦♦♦")

	fmt.Println("\n\n\n\t ==========   ==========    \\\\        //     ==========     \\\\            ||")
	fmt.Println("\t||                ||       ||\\\\      //||   ||        ||   ||\\\\           ||")
	fmt.Println("\t||                ||       || \\\\    // ||   ||        ||   || \\\\          ||")
	fmt.Println("\t||                ||       ||  \\\\  //  ||   ||        ||   ||  \\\\         ||")
	fmt.Println("\t||                ||       ||   \\\\//   ||   ||        ||   ||   \\\\        ||")
	fmt.Println("\t||                ||       ||    \\/    ||   ||        ||   ||    \\\\       ||")
	fmt.Println("\t||                ||       ||          ||   ||        ||   ||     \\\\      ||")
	fmt.Println("\t===========       ||       ||          ||   ||        ||   ||      \\\\     ||")
	fmt.Println("\t         ||       ||       ||          ||   ||        ||   ||       \\\\    ||")
	fmt.Println("\t         ||       ||       ||          ||   ||        ||   ||        \\\\   ||")
	fmt.Println("\t         ||       ||       ||          ||   ||        ||   ||         \\\\  ||")
	fmt.Println("\t         ||       ||       ||          ||   ||        ||   ||          \\\\ ||")
	fmt.Println("\t         ||       ||       ||          ||   ||        ||   ||           \\\\||")
	fmt.Print("\t==========    ==========   ||          ||    ==========    ||             ||\n\n\n")

	fmt.Println("\n\n\n\t||=======||    ||=======||   ||      //      ||=========    ||=======//")
	fmt.Println("\t||       ||    ||       ||   ||     //       ||             ||      //")
	fmt.Println("\t||       ||    ||       ||   ||    //        ||             ||     //")
	fmt.Println("\t||       ||    ||       ||   ||   //         ||             ||    //")
	fmt.Println("\t||       ||    ||       ||   ||  //          ||             ||   //")
	fmt.Println("\t||       ||    ||       ||   || //           ||             ||  //")
	fmt.Println("\t||       ||    ||       ||   ||//            ||             || //")
	fmt.Println("\t||=======||    ||       ||   ||\\\\            ||=========    || \\\\")
	fmt.Println("\t||             ||       ||   || \\\\           ||             ||  \\\\")
	fmt.Println("\t||             ||       ||   ||  \\\\          ||             ||   \\\\")
	fmt.Println("\t||             ||       ||   ||   \\\\         ||             ||    \\\\")
	fmt.Println("\t||             ||       ||   ||    \\\\        ||             ||     \\\\")
	fmt.Println("\t||             ||       ||   ||     \\\\       ||             ||      \\\\")
	fmt.Print("\t||             ||=======||   ||      \\\\      ||=========    ||       \\\\\n\n\n")

	fmt.Print("\n\n\n\t♠♠♠♠♠♠♠♠♠♠♠♠♠♠♠♠♠♠♠♠♠♠♠♠♠♠♠♠♠♠♠♠♠♠♠♠♠♠♠♠♠♠♠♠♠♠♠♠♠♠♠♠♠♠♠♠♠♠♠♠♠♠♠♠♠♠♠♠♠♠♠♠♠♠♠♠♠")
	fmt.Print("\n\n\t♥♥♥♥♥♥♥♥♥♥♥♥♥♥♥♥♥♥♥♥♥♥♥♥♥♥♥♥♥♥♥♥♥♥♥♥♥♥♥♥♥♥♥♥♥♥♥♥♥♥♥♥♥♥♥♥♥♥♥♥♥♥♥♥♥♥♥♥♥♥♥♥♥♥♥♥♥")
	fmt.Print("\n\n\t♣♣♣♣♣♣♣♣♣♣♣♣♣♣♣♣♣♣♣♣♣♣♣♣♣♣♣♣♣♣♣♣♣♣♣♣♣♣♣♣♣♣♣♣♣♣♣♣♣♣♣♣♣♣♣♣♣♣♣♣♣♣♣♣♣♣♣♣♣♣♣♣♣♣♣♣♣")
	fmt.Print("\n\n\t♦♦♦♦♦♦♦♦♦♦♦♦♦♦♦♦♦♦♦♦♦♦♦♦♦♦♦♦♦♦♦♦♦♦♦♦♦♦♦♦♦♦♦♦♦♦♦♦♦♦♦♦♦♦♦♦♦♦♦♦♦♦♦♦♦♦♦♦♦♦♦♦♦♦♦♦♦\n\n\n\n\n\n")
}

func deal(d *deck, handSize int) deck {
	hand := (*d)[:handSize]
	(*d) = (*d)[handSize:]
	return hand
}

func shuffle(d deck) {
	source := rand.NewSource(time.Now().UnixNano())
	r := rand.New(source)
	r.Shuffle(len(d), func(i, j int) {
		d[i], d[j] = d[j], d[i]
	})

}
func calculateFiveBestCards(cardsInput []card) handScore {
	cards := copyDeck(cardsInput)
	hS := handScore{score: -100}
	var flush []int
	var pairs []int
	var threeOfAKind []int
	var remainingCards []int
	groupOfSameColours := map[string]int{}
	groupByValues := map[int]int{}

	sortCards(cards)

	for _, card := range cards {
		if groupByValues[card.value] == 0 {
			groupByValues[card.value] = 1
		} else {
			groupByValues[card.value]++
		}
		if groupOfSameColours[card.suit] == 0 {
			groupOfSameColours[card.suit] = 1
		} else {
			groupOfSameColours[card.suit]++
		}
	}

	// find royal flush, straight flush or flush
	for key, value := range groupOfSameColours {
		if value > 4 {
			hS.suit = key
			sc := []int{}
			for _, card := range cards {
				if card.suit == key {
					sc = append(sc, card.value)
				}
			}
			if sc[len(sc)-5] == 10 {
				hS.score = 900
			} else {
				if highestCard := findStraight(sc); highestCard != 0 {
					hS.score = 800
					hS.card1 = highestCard
				} else {
					flush = sc[len(sc)-5:]
				}
			}
		}
	}

	// find four of a kind
	if hS.score == -100 {
		remains := []int{}
		for key, value := range groupByValues {
			if value == 4 {
				hS.score = 700
				hS.card1 = key
			} else {
				remains = append(remains, key)
			}
		}
		if hS.score == 700 {
			hS.remainingCards = remains[len(remains)-1:]
		}
	}

	// find full house - three of a kind - 1 pair - 2 pairs
	if hS.score == -100 {
		for key, value := range groupByValues {
			if value == 3 {
				threeOfAKind = append(threeOfAKind, key)
			} else {
				if value == 2 {
					pairs = append(pairs, key)
				} else {
					remainingCards = append(remainingCards, key)
					orderSliceOfInt(remainingCards)
				}
			}
		}
		orderSliceOfInt(threeOfAKind)
		if len(threeOfAKind) == 2 {
			pairs = append(pairs, threeOfAKind[0])
			threeOfAKind = threeOfAKind[1:]
		}

		orderSliceOfInt(pairs)

		if len(threeOfAKind) == 1 {
			if len(pairs) > 0 {
				hS.score = 600
				hS.card1 = threeOfAKind[0]
				hS.card2 = pairs[len(pairs)-1]
			}
		}

	}

	if hS.score == -100 && len(flush) > 0 {
		hS.score = 500
		hS.remainingCards = flush
		hS.card1 = flush[len(flush)-1]
	}

	// find straight including the straight Ace 2 3 4 5 (Ace=14)
	if hS.score == -100 {
		cardValues := []int{}
		for _, card := range cards {
			cardValues = append(cardValues, card.value)
		}
		if highestCard := findStraight(cardValues); highestCard != 0 {
			hS.score = 400
			hS.card1 = highestCard
		}

	}

	// remains three of a kind, single pair, two pairs, or high card
	if hS.score == -100 {
		if len(threeOfAKind) > 0 {
			hS.score = 300
			hS.card1 = threeOfAKind[0]
			hS.remainingCards = remainingCards[len(remainingCards)-2:]

		} else {
			if len(pairs) > 2 {
				remainingCards = append(remainingCards, pairs[(len(pairs)-3)])
				orderSliceOfInt(remainingCards)
				pairs = pairs[len(pairs)-2:]
			}
			orderSliceOfInt(remainingCards)
			switch len(pairs) {
			case 2:
				hS.score = 200
				hS.card1 = pairs[len(pairs)-1]
				hS.card2 = pairs[len(pairs)-2]
				hS.remainingCards = remainingCards[len(remainingCards)-1:]
			case 1:
				hS.score = 100
				hS.card1 = pairs[0]
				hS.remainingCards = remainingCards[len(remainingCards)-3:]
			case 0:
				hS.score = 0
				hS.card1 = remainingCards[len(remainingCards)-1]
				hS.remainingCards = remainingCards[len(remainingCards)-5 : len(remainingCards)-1]
			}

		}
	}

	return hS
}

func copyDeck(cards []card) []card {
	copy := []card{}
	for _, card := range cards {
		copy = append(copy, card)
	}

	return copy
}

func isBetterHand(h1 handScore, h2 handScore) int {
	if h1.score > h2.score {
		return -1
	}
	if h1.score < h2.score {
		return 1
	}
	if h1.card1 > h2.card1 {
		return -1
	}
	if h1.card1 < h2.card1 {
		return 1
	}
	if h1.card2 > h2.card2 {
		return -1
	}
	if h1.card2 < h2.card2 {
		return 1
	}
	for i := len(h1.remainingCards) - 1; i >= 0; i-- {
		if h1.remainingCards[i] > h2.remainingCards[i] {
			return -1
		}
		if h1.remainingCards[i] < h2.remainingCards[i] {
			return 1
		}
	}

	return 0
}

func sortCards(c []card) {
	for i, _ := range c {
		j := i
		for j >= 0 && j < len(c)-1 && c[j+1].value < c[j].value {
			c[j+1], c[j] = c[j], c[j+1]
			j--
		}

	}
}

func orderSliceOfInt(s []int) {
	for i, _ := range s {
		j := i
		for j >= 0 && j < len(s)-1 && s[j+1] < s[j] {
			s[j+1], s[j] = s[j], s[j+1]
			j--
		}

	}
}

func orderSliceOfIntDesc(s []int) {
	for i, _ := range s {
		j := i
		for j >= 0 && j < len(s)-1 && s[j+1] > s[j] {
			s[j+1], s[j] = s[j], s[j+1]
			j--
		}

	}
}

// the returned value is 0 if not found, or the highest card of the straight
func findStraight(c []int) int {

	best := []int{}
	curr := []int{c[0]}

	for i := range c {
		if i > 0 && c[i] == c[i-1]+1 {
			curr = append(curr, c[i])
		} else {
			if i > 0 && c[i] != c[i-1] {
				if len(curr) == 5 || (len(curr) == 4 && curr[0] == 2 && c[len(c)-1] == 14) {
					best = curr
				}
				curr = []int{c[i]}
			}
		}
	}
	if len(curr) >= 5 || (len(curr) == 4 && curr[0] == 2 && c[len(c)-1] == 14) {
		best = curr
	}
	if len(best) > 0 {
		return best[len(best)-1]
	}
	return 0
}

func maxInt(a int, b int) int {
	if a < b {
		return b
	}

	return a
}

func minInt(a int, b int) int {
	if a < b {
		return a
	}

	return b
}

func contains(s []int, e int) bool {
	for _, a := range s {
		if a == e {
			return true
		}
	}
	return false
}

func deckContains(deck []card, c card) bool {
	for _, a := range deck {
		if a.value == c.value && c.suit == a.suit {
			return true
		}
	}
	return false
}

func indexOf(s []int, e int) int {
	for i, a := range s {
		if a == e {
			return i
		}
	}
	return -1
}
