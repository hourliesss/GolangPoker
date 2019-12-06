package main

import (
	"fmt"
	"time"
)

func main() {
	//poker()
	botMode()
}
func poker() {
	g := launchGame()

	coeff := coefficients{
		check:         0.5,
		raise:         0.3,
		call:          0.2,
		normalisation: 3,
		fold1:         0.2,
		fold2:         0,
		allIn:         0.2,
	}

	for g.isGameOn() {

		g.newRound()
		showCards(g.rounds[len(g.rounds)-1].players)
		g.betBlinds()

		numberOfPlayers := g.getNumberOfPlayersAlive()
		turnToPlay := (g.smallBlindTurn + 2) % numberOfPlayers
		for !g.rounds[len(g.rounds)-1].isBetTurnOver(turnToPlay) {
			play(&(g.rounds[len(g.rounds)-1]), turnToPlay, g.bigBlindValue, float64(g.totalStack), coeff)
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
				play(&(g.rounds[len(g.rounds)-1]), turnToPlay, g.bigBlindValue, float64(g.totalStack), coeff)
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

func botMode() {

	fmt.Println("ANDRE VS GOMEZ, first in 100")
	andrePoints := 0
	gomezPoints := 0
	totalRounds := 0
	for andrePoints < 1 && gomezPoints < 1 {
		fmt.Println("ANDRE", andrePoints, "-", gomezPoints, "GOMEZ")

		players := []person{
			person{
				name:  "Andre",
				stack: 1000,
				isBot: true,
			},
			person{
				name:  "Gomez",
				stack: 1000,
				isBot: true,
			},
		}

		coeffAndre := coefficients{
			check:         0.5,
			raise:         0.3,
			call:          0.2,
			normalisation: 1,
			fold1:         0.2,
			fold2:         0,
			allIn:         0.2,
		}

		coeffGomez := coefficients{
			check:         0.5,
			raise:         0.3,
			call:          0.2,
			normalisation: 1,
			fold1:         0.2,
			fold2:         0,
			allIn:         0.2,
		}
		g := newGame(players)

		for g.isGameOn() {

			g.newRound()
			showCards(g.rounds[len(g.rounds)-1].players)
			g.betBlinds()

			numberOfPlayers := g.getNumberOfPlayersAlive()
			turnToPlay := (g.smallBlindTurn + 2) % numberOfPlayers
			for !g.rounds[len(g.rounds)-1].isBetTurnOver(turnToPlay) {
				if turnToPlay == 0 {
					play(&(g.rounds[len(g.rounds)-1]), turnToPlay, g.bigBlindValue, float64(g.totalStack), coeffAndre)
				} else {
					play(&(g.rounds[len(g.rounds)-1]), turnToPlay, g.bigBlindValue, float64(g.totalStack), coeffGomez)
				}

				turnToPlay = (turnToPlay + 1) % numberOfPlayers
			}
			for g.rounds[len(g.rounds)-1].playersAlive > 1 && len(g.rounds[len(g.rounds)-1].sharedCards) < 5 {
				switch len(g.rounds[len(g.rounds)-1].sharedCards) {
				case 0:
					g.rounds[len(g.rounds)-1].sharedCards = deal(&g.rounds[len(g.rounds)-1].deck, 3)
					fmt.Println(g.rounds[len(g.rounds)-1].sharedCards[0].toString())
					fmt.Println(g.rounds[len(g.rounds)-1].sharedCards[1].toString())
					fmt.Println(g.rounds[len(g.rounds)-1].sharedCards[2].toString())
				case 3:
					g.rounds[len(g.rounds)-1].sharedCards = append(g.rounds[len(g.rounds)-1].sharedCards, deal(&g.rounds[len(g.rounds)-1].deck, 1)...)
					printDeck(g.rounds[len(g.rounds)-1].sharedCards)
				case 4:
					g.rounds[len(g.rounds)-1].sharedCards = append(g.rounds[len(g.rounds)-1].sharedCards, deal(&g.rounds[len(g.rounds)-1].deck, 1)...)
					printDeck(g.rounds[len(g.rounds)-1].sharedCards)
				}

				for i := range g.rounds[len(g.rounds)-1].players {
					g.rounds[len(g.rounds)-1].players[i].hasSpoken = false
				}
				turnToPlay := g.smallBlindTurn
				for !g.rounds[len(g.rounds)-1].isBetTurnOver(turnToPlay) {
					if turnToPlay == 0 {
						play(&(g.rounds[len(g.rounds)-1]), turnToPlay, g.bigBlindValue, float64(g.totalStack), coeffAndre)
					} else {
						play(&(g.rounds[len(g.rounds)-1]), turnToPlay, g.bigBlindValue, float64(g.totalStack), coeffGomez)
					}
					turnToPlay = (turnToPlay + 1) % numberOfPlayers
				}
			}
			g.endRound()

		}
		if winner := g.getWinnerIndex(); winner == 0 {
			andrePoints++
		} else {
			gomezPoints++
		}
		totalRounds += len(g.rounds)
	}

	fmt.Println("ANDRE", andrePoints, "-", gomezPoints, "GOMEZ")
	fmt.Println("average rounds:", float64(totalRounds)/float64(andrePoints+gomezPoints))

}
