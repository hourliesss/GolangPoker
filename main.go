package main

import (
	"fmt"
	"time"
)

func main() {
	poker()
}
func poker() {
	g := launchGame()

	for g.isGameOn() {

		g.newRound()
		showCards(g.rounds[len(g.rounds)-1].players)
		g.betBlinds()

		numberOfPlayers := g.getNumberOfPlayersAlive()
		turnToPlay := (g.smallBlindTurn + 2) % numberOfPlayers
		for !g.rounds[len(g.rounds)-1].isBetTurnOver(turnToPlay) {
			play(&(g.rounds[len(g.rounds)-1]), turnToPlay, g.bigBlindValue)
			turnToPlay = (turnToPlay + 1) % numberOfPlayers
		}
		for g.rounds[len(g.rounds)-1].playersAlive > 1 && len(g.rounds[len(g.rounds)-1].sharedCards) < 5 {
			switch len(g.rounds[len(g.rounds)-1].sharedCards) {
			case 0:
				fmt.Print("\n\tHere comes the flop...\n\n")
				g.rounds[len(g.rounds)-1].sharedCards = deal(&g.rounds[len(g.rounds)-1].deck, 3)
				time.Sleep(time.Second * 1)
				fmt.Println(g.rounds[len(g.rounds)-1].sharedCards[0].toString())
				time.Sleep(time.Second * 1)
				fmt.Println(g.rounds[len(g.rounds)-1].sharedCards[1].toString())
				time.Sleep(time.Second * 1)
				fmt.Println(g.rounds[len(g.rounds)-1].sharedCards[2].toString())
				time.Sleep(time.Second * 1)
			case 3:
				fmt.Print("\n\tHere comes the turn...\n\n")
				g.rounds[len(g.rounds)-1].sharedCards = append(g.rounds[len(g.rounds)-1].sharedCards, deal(&g.rounds[len(g.rounds)-1].deck, 1)...)
				time.Sleep(time.Second * 1)
				fmt.Println(g.rounds[len(g.rounds)-1].sharedCards[3].toString())
				time.Sleep(time.Second * 1)
				fmt.Print("\n\tCARDS:\n")
				printDeck(g.rounds[len(g.rounds)-1].sharedCards)
			case 4:
				fmt.Print("\n\tHere comes the river...\n\n")
				g.rounds[len(g.rounds)-1].sharedCards = append(g.rounds[len(g.rounds)-1].sharedCards, deal(&g.rounds[len(g.rounds)-1].deck, 1)...)
				time.Sleep(time.Second * 1)
				fmt.Println(g.rounds[len(g.rounds)-1].sharedCards[4].toString())
				time.Sleep(time.Second * 1)
				fmt.Print("\n\tCARDS:\n")
				printDeck(g.rounds[len(g.rounds)-1].sharedCards)
			}

			for i := range g.rounds[len(g.rounds)-1].players {
				g.rounds[len(g.rounds)-1].players[i].hasSpoken = false
			}
			turnToPlay := g.smallBlindTurn
			for !g.rounds[len(g.rounds)-1].isBetTurnOver(turnToPlay) {
				play(&(g.rounds[len(g.rounds)-1]), turnToPlay, g.bigBlindValue)
				turnToPlay = (turnToPlay + 1) % numberOfPlayers
			}
		}
		g.endRound()
		time.Sleep(time.Second * 3)

	}
	fmt.Println("\tCongratulations", g.getWinner(), "\n\n\tyou won!")
}

// func main() {
//	initialProbabilty()
//}
