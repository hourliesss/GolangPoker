package main

import (
	"testing"
)

func TestSortCards(t *testing.T) {
	s := []int{8, 6, 2, 14, 13, 12, 13, 2, 9, 5, 4, 6, 6, 13, 14}

	orderSliceOfInt(s)

	isOrdered := true
	for i := range s {
		if i < len(s)-1 && s[i] > s[i+1] {
			isOrdered = false
			break
		}

	}
	if !isOrdered {
		t.Errorf("Expected a sorted slice but got %v", s)
	}

}

func TestTwoPairs(t *testing.T) {

	t.Logf("Testing Two Pairs...")
	hand := deck{
		createCard(8, "Spades"),
		createCard(14, "Clubs"),
		createCard(13, "Clubs"),
		createCard(13, "Hearts"),
		createCard(14, "Hearts"),
		createCard(8, "Diamonds"),
		createCard(11, "Diamonds"),
	}

	score := calculateFiveBestCards(hand)

	if score.score != 200 {
		t.Errorf("Expected two pairs of 14 and 13 but got a %v", score.toString())
	}

}

func TestStraightFlushWithAce(t *testing.T) {

	t.Logf("Testing Straight Flush without Ace...")

	hand := deck{
		createCard(5, "Hearts"),
		createCard(4, "Hearts"),
		createCard(5, "Clubs"),
		createCard(3, "Hearts"),
		createCard(14, "Hearts"),
		createCard(2, "Hearts"),
		createCard(4, "Clubs"),
	}

	score := calculateFiveBestCards(hand)

	if score.score != 800 || score.card1 != 5 {
		t.Errorf("Expected a straight flush with high card 5  but got %v", score.toString())
	}

}

func TestStraightFlushWithoutAce(t *testing.T) {

	t.Logf("Testing Straight Flush without Ace...")

	hand := deck{
		createCard(5, "Hearts"),
		createCard(4, "Hearts"),
		createCard(6, "Hearts"),
		createCard(3, "Hearts"),
		createCard(14, "Diamonds"),
		createCard(2, "Hearts"),
		createCard(4, "Clubs"),
	}

	score := calculateFiveBestCards(hand)

	if score.score != 800 || score.card1 != 6 {
		t.Errorf("Expected a straight flush with high card 5  but got %v", score.toString())
	}

}

func TestRoyalFlush(t *testing.T) {

	t.Logf("Testing Royal Flush...")

	hand := deck{
		createCard(10, "Hearts"),
		createCard(9, "Hearts"),
		createCard(8, "Hearts"),
		createCard(11, "Hearts"),
		createCard(14, "Hearts"),
		createCard(12, "Hearts"),
		createCard(13, "Hearts"),
	}

	score := calculateFiveBestCards(hand)

	if score.score != 900 {
		t.Errorf("Expected a royal flush but got %v", score.toString())
	}

}

func TestFourOfAKind(t *testing.T) {

	t.Logf("Testing 4 of a kind...")

	hand := deck{
		createCard(9, "Clubs"),
		createCard(10, "Hearts"),
		createCard(10, "Diamonds"),
		createCard(10, "Spades"),
		createCard(9, "Spades"),
		createCard(9, "Hearts"),
		createCard(10, "Cubes"),
	}

	score := calculateFiveBestCards(hand)

	if score.score != 700 || score.card1 != 10 || score.remainingCards[0] != 9 {
		t.Errorf("Expected a 4 of a kind of 10 with a 9 but got a %v of %v and a %v", score.toString(), score.card1, score.remainingCards[0])
	}

}

func TestFullHouse(t *testing.T) {

	t.Logf("Testing Full House...")

	hand := deck{
		createCard(9, "Clubs"),
		createCard(10, "Hearts"),
		createCard(10, "Diamonds"),
		createCard(10, "Spades"),
		createCard(9, "Spades"),
		createCard(9, "Hearts"),
		createCard(6, "Hearts"),
	}

	score := calculateFiveBestCards(hand)

	if score.score != 600 || score.card1 != 10 || score.card2 != 9 {
		t.Errorf("Expected a full house but got %v of the %v by the %v", score.toString(), score.card1, score.card2)
	}
}

func TestFlush(t *testing.T) {

	t.Logf("Testing flush...")

	hand := deck{
		createCard(5, "Clubs"),
		createCard(11, "Clubs"),
		createCard(2, "Clubs"),
		createCard(3, "Hearts"),
		createCard(3, "Spades"),
		createCard(3, "Clubs"),
		createCard(4, "Clubs"),
	}

	score := calculateFiveBestCards(hand)

	if score.score != 500 {
		t.Errorf("Expected a flush but got %v", score.toString())
	}

}

func TestStraight(t *testing.T) {

	t.Logf("Testing Straight...")

	hand := deck{
		createCard(5, "Clubs"),
		createCard(3, "Clubs"),
		createCard(2, "Clubs"),
		createCard(14, "Hearts"),
		createCard(3, "Spades"),
		createCard(3, "Diamonds"),
		createCard(4, "Clubs"),
	}

	score := calculateFiveBestCards(hand)

	if score.score != 400 || score.card1 != 5 {
		t.Errorf("Expected a straight with high card 5 but got %v with high card %v", score.toString(), score.card1)
	}

}

func TestThreeOfAKind(t *testing.T) {

	t.Logf("Testing 3 of a kind...")

	hand := deck{
		createCard(8, "Clubs"),
		createCard(9, "Clubs"),
		createCard(12, "Clubs"),
		createCard(13, "Hearts"),
		createCard(9, "Spades"),
		createCard(9, "Diamonds"),
		createCard(3, "Clubs"),
	}

	score1 := calculateFiveBestCards(hand)

	if score1.score != 300 || score1.card1 != 9 || score1.remainingCards[0] != 12 || score1.remainingCards[1] != 13 {
		t.Errorf("Expected a three 9s with a 13 and a 14 but got %v of %v with %v and %v", score1.toString(), score1.card1, score1.remainingCards[0], score1.remainingCards[1])
	}

}

func TestPair(t *testing.T) {

	t.Logf("Testing a pair...")

	pairHand := deck{
		createCard(8, "Clubs"),
		createCard(5, "Clubs"),
		createCard(6, "Clubs"),
		createCard(13, "Hearts"),
		createCard(12, "Spades"),
		createCard(9, "Diamonds"),
		createCard(13, "Clubs"),
	}
	pairScore := calculateFiveBestCards(pairHand)
	if pairScore.score != 100 && pairScore.card1 != 13 || pairScore.remainingCards[0] != 8 || pairScore.remainingCards[1] != 9 || pairScore.remainingCards[2] != 12 {
		t.Errorf("Expected a pair of 13 with 8 9 12 but got %v of %v with %v %v %v", pairScore.toString(), pairScore.card1, pairScore.remainingCards[0], pairScore.remainingCards[1], pairScore.remainingCards[2])
	}

}

func TestHighCard(t *testing.T) {

	t.Logf("Testing high card...")

	hand := deck{
		createCard(8, "Clubs"),
		createCard(5, "Clubs"),
		createCard(6, "Clubs"),
		createCard(13, "Hearts"),
		createCard(12, "Spades"),
		createCard(9, "Diamonds"),
		createCard(14, "Clubs"),
	}

	score := calculateFiveBestCards(hand)

	if score.score != 0 && score.card1 != 14 || score.remainingCards[0] != 8 || score.remainingCards[1] != 9 || score.remainingCards[2] != 12 || score.remainingCards[3] != 13 {
		t.Errorf("Expected a high card 14 with 8 9 12 13 but got %v of %v with %v %v %v %v", score.toString(), score.card1, score.remainingCards[0], score.remainingCards[1], score.remainingCards[2], score.remainingCards[3])
	}

}
