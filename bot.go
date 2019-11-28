package main

import (
	"fmt"
	"io/ioutil"
	"math/rand"
	"strconv"
	"time"
)

type initialProbabilityMap struct {
	card1  int
	card2  int
	suited bool
}

type initialProbabiltyResult struct {
	success int
	try     int
}

func initialProbabilty() error {

	probs := map[initialProbabilityMap]initialProbabiltyResult{}
	for i := 2; i < 15; i++ {
		for j := 2; j < i+1; j++ {
			initSuited := initialProbabilityMap{
				card1:  i,
				card2:  j,
				suited: true,
			}
			initNotSuited := initialProbabilityMap{
				card1:  i,
				card2:  j,
				suited: false,
			}
			probs[initSuited] = initialProbabiltyResult{}
			probs[initNotSuited] = initialProbabiltyResult{}
		}
	}

	for count := 0; count < 100000000; count++ {
		d := newDeckForTest(count)
		if count%1000000 == 0 {
			fmt.Println("Done:", count/1000000, "%")
		}
		hand1 := deal(&d, 2)
		if hand1[0].value < hand1[1].value {
			hand1[0], hand1[1] = hand1[1], hand1[0]
		}

		hand2 := deal(&d, 2)
		if hand2[0].value < hand2[1].value {
			hand2[0], hand2[1] = hand2[1], hand2[0]
		}
		sharedCards := deal(&d, 5)
		hs1 := calculateFiveBestCards(append(sharedCards, hand1...))
		hs2 := calculateFiveBestCards(append(sharedCards, hand2...))
		i := isBetterHand(hs1, hs2)

		if i != 0 {
			initialProbability1 := initialProbabilityMap{
				card1:  hand1[0].value,
				card2:  hand1[1].value,
				suited: hand1[0].suit == hand1[1].suit,
			}
			initialProbability2 := initialProbabilityMap{
				card1:  hand2[0].value,
				card2:  hand2[1].value,
				suited: hand2[0].suit == hand2[1].suit,
			}
			res1 := probs[initialProbability1]
			res1.try++
			res2 := probs[initialProbability2]
			res2.try++
			if i == -1 {
				res1.success++
			} else {
				res2.success++
			}
			probs[initialProbability1] = res1
			probs[initialProbability2] = res2
		}
	}
	resultString := "card1\tcard2\tsuited\tsuccess\ttry\n"
	for key, value := range probs {
		resultString += strconv.Itoa(key.card1) + "\t" + strconv.Itoa(key.card2) + "\t" + strconv.FormatBool(key.suited) + "\t" + strconv.Itoa(value.success) + "\t" + strconv.Itoa(value.try) + "\n"
	}

	return ioutil.WriteFile("probability.csv", []byte(resultString), 0666)
}


func newDeckForTest(i int) deck {
	suitValues := []string{"Spades", "Clubs", "Hearts", "Diamonds"}
	var d deck
	for i := 2; i < 15; i++ {
		for _, suit := range suitValues {
			d = append(d, createCard(i, suit))
		}
	}
	shuffleForTest(d, i)
	return d
}

func shuffleForTest(d deck, i int) {
	source := rand.NewSource(time.Now().UnixNano() + int64(i))
	r := rand.New(source)
	for i := range d {
		newPosition := r.Intn(len(d) - 1)
		d[i], d[newPosition] = d[newPosition], d[i]
	}

}
