package main

import (
	"fmt"
	"math"
	"time"
)

func main() {
	poker()
	//botMode()
}
func poker() {
	g := launchGame()

	coeff := coefficients{
		check:         0.65,
		raise:         0.3,
		call:          0.3,
		normalisation: 3,
		fold1:         0.5,
		fold2:         1.5,
		allIn:         0.95,
	}

	for g.isGameOn() {

		g.newRound(false)
		showCards(g.rounds[len(g.rounds)-1].players)
		g.betBlinds(false)

		numberOfPlayers := g.getNumberOfPlayersAlive()
		turnToPlay := (g.smallBlindTurn + 2) % numberOfPlayers
		for !g.rounds[len(g.rounds)-1].isBetTurnOver(turnToPlay) {
			play(&(g.rounds[len(g.rounds)-1]), turnToPlay, g.bigBlindValue, float64(g.totalStack), coeff, false)
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
				play(&(g.rounds[len(g.rounds)-1]), turnToPlay, g.bigBlindValue, float64(g.totalStack), coeff, false)
				turnToPlay = (turnToPlay + 1) % numberOfPlayers
			}
		}
		g.endRound(false)
		time.Sleep(time.Second * 3)

	}
	fmt.Println("\tCongratulations", g.getWinner(), "\n\n\tyou won!")
}

// func main() {
//	initialProbabilty()
//}

func botMode() {

	andrePoints := 0
	gomezPoints := 0

	coeffAndre := coefficients{
		check:         0.4,
		raise:         0.2,
		call:          0.1,
		normalisation: 1,
		fold1:         0.3,
		fold2:         1.0,
		allIn:         0.85,
	}

	coeffGomez := coefficients{
		check:         0.6, //  between 0.3 and 0.7
		raise:         1.0, // between 0 and 1
		call:          0.7, // between 0 and 1
		normalisation: 1,   // not used
		fold1:         0.8, // between 0.2 and 0.8
		fold2:         5.0, // between 0.1 and 10
		allIn:         1,   // between 0.7 and 1
	}
	first := true
	turn := 0
	difference := 0
	lastDifference := 0
	initialValue := float64(0.0)
	var updateDifference bool
	totalRounds := 0
	for math.Abs(coeffGomez.fold1-coeffAndre.fold1) > 0.1 || math.Abs(coeffGomez.fold2-coeffAndre.fold2) > 0.1 || math.Abs(coeffGomez.call-coeffAndre.call) > 0.1 || math.Abs(coeffGomez.raise-coeffAndre.raise) > 0.1 || math.Abs(coeffGomez.allIn-coeffAndre.allIn) > 0.1 || math.Abs(coeffGomez.check-coeffAndre.check) > 0.1 {

		difference = andrePoints - gomezPoints

		updateDifference = true
		if !first {
			if difference > 0 {
				if ((difference > lastDifference-20 && lastDifference > 0) || initialValue == 0) || totalRounds/(andrePoints+gomezPoints) < 20 {
					if turn == 0 {
						if initialValue == 0 {
							initialValue = coeffGomez.raise
							coeffGomez.raise = (initialValue + coeffAndre.raise) / 2.0
						} else {
							if initialValue > coeffGomez.raise {
								updateDifference = false
								coeffGomez.raise = 2*initialValue - coeffGomez.raise
							} else {
								turn++
								updateDifference = false
								coeffGomez.raise = initialValue
								initialValue = 0
							}
						}
					}
					if turn == 1 {
						if initialValue == 0 {
							initialValue = coeffGomez.call
							coeffGomez.call = (initialValue + coeffAndre.call) / 2.0
						} else {
							if initialValue > coeffGomez.call {
								updateDifference = false
								coeffGomez.call = 2*initialValue - coeffGomez.call
							} else {
								turn++
								updateDifference = false
								coeffGomez.call = initialValue
								initialValue = 0
							}
						}
					}
					if turn == 2 {
						if initialValue == 0 {
							initialValue = coeffGomez.fold1
							coeffGomez.fold1 = (initialValue + coeffAndre.fold1) / 2.0
						} else {
							if initialValue > coeffGomez.fold1 {
								updateDifference = false
								coeffGomez.fold1 = 2*initialValue - coeffGomez.fold1
							} else {
								turn++
								updateDifference = false
								coeffGomez.fold1 = initialValue
								initialValue = 0
							}
						}
					}
					if turn == 3 {
						if initialValue == 0 {
							initialValue = coeffGomez.fold2
							coeffGomez.fold2 = (initialValue + coeffAndre.fold2) / 2.0
						} else {
							if initialValue > coeffGomez.fold2 {
								updateDifference = false
								coeffGomez.fold2 = 2*initialValue - coeffGomez.fold2
							} else {
								turn++
								updateDifference = false
								coeffGomez.fold2 = initialValue
								initialValue = 0
							}
						}
					}
					if turn == 4 {
						if initialValue == 0 {
							initialValue = coeffGomez.allIn
							coeffGomez.allIn = (initialValue + coeffAndre.allIn) / 2.0
						} else {
							if initialValue > coeffGomez.allIn {
								updateDifference = false
								coeffGomez.allIn = 2*initialValue - coeffGomez.allIn
							} else {
								turn++
								updateDifference = false
								coeffGomez.allIn = initialValue
								initialValue = 0
							}
						}
					}
					if turn == 5 {
						if initialValue == 0 {
							initialValue = coeffGomez.check
							coeffGomez.check = (initialValue + coeffAndre.check) / 2.0
						} else {
							if initialValue > coeffGomez.check {
								updateDifference = false
								coeffGomez.check = 2*initialValue - coeffGomez.check
							} else {
								turn = 0
								updateDifference = false
								coeffGomez.check = initialValue
								initialValue = 0
							}
						}
					}
				} else {
					initialValue = 0
					if turn == 0 {
						fmt.Println("RAISE STABLE")
						turn++
					} else {
						if turn == 1 {
							fmt.Println("CALL STABLE")
							turn++
						} else {
							if turn == 2 {
								fmt.Println("FOLD1 STABLE")
								turn++
							} else {
								if turn == 3 {
									fmt.Println("FOLD2 STABLE")
									turn++
								} else {
									if turn == 4 {
										fmt.Println("ALLIN STABLE")
										turn++
									} else {
										fmt.Println("CHECK STABLE")
										turn = 0
									}
								}
							}
						}
					}
				}
			} else {
				if difference < 0 {
					if ((difference < lastDifference-8 && lastDifference < 0) || initialValue == 0) || totalRounds/(andrePoints+gomezPoints) < 20 {
						if turn == 0 {
							if initialValue == 0 {
								initialValue = coeffAndre.raise
								coeffAndre.raise = (initialValue + coeffGomez.raise) / 2.0
							} else {
								if initialValue < coeffAndre.raise {
									coeffAndre.raise = 2*initialValue - coeffAndre.raise
									updateDifference = false
								} else {
									turn++
									coeffAndre.raise = initialValue
									updateDifference = false
									initialValue = 0
								}
							}
						}
						if turn == 1 {
							if initialValue == 0 {
								initialValue = coeffAndre.call
								coeffAndre.call = (initialValue + coeffGomez.call) / 2.0
							} else {
								if initialValue < coeffAndre.call {
									coeffAndre.call = 2*initialValue - coeffAndre.call
									updateDifference = false
								} else {
									turn++
									coeffAndre.call = initialValue
									updateDifference = false
									initialValue = 0
								}
							}
						}
						if turn == 2 {
							if initialValue == 0 {
								initialValue = coeffAndre.fold1
								coeffAndre.fold1 = (initialValue + coeffGomez.fold1) / 2.0
							} else {
								if initialValue < coeffAndre.fold1 {
									coeffAndre.fold1 = 2*initialValue - coeffAndre.fold1
									updateDifference = false
								} else {
									turn++
									coeffAndre.fold1 = initialValue
									updateDifference = false
									initialValue = 0
								}
							}
						}
						if turn == 3 {
							if initialValue == 0 {
								initialValue = coeffAndre.fold2
								coeffAndre.fold2 = (initialValue + coeffGomez.fold2) / 2.0
							} else {
								if initialValue < coeffAndre.fold2 {
									coeffAndre.fold2 = 2*initialValue - coeffAndre.fold2
									updateDifference = false
								} else {
									turn++
									coeffAndre.fold2 = initialValue
									initialValue = 0
									updateDifference = false
								}
							}
						}
						if turn == 4 {
							if initialValue == 0 {
								initialValue = coeffAndre.allIn
								coeffAndre.allIn = (initialValue + coeffGomez.allIn) / 2.0
							} else {
								if initialValue < coeffAndre.allIn {
									coeffAndre.allIn = 2*initialValue - coeffAndre.allIn
									updateDifference = false
								} else {
									turn++
									coeffAndre.allIn = initialValue
									initialValue = 0
									updateDifference = false
								}
							}
						}
						if turn == 5 {
							if initialValue == 0 {
								initialValue = coeffAndre.check
								coeffAndre.check = (initialValue + coeffGomez.check) / 2.0
							} else {
								if initialValue < coeffAndre.check {
									coeffAndre.check = 2*initialValue - coeffAndre.check
									updateDifference = false
								} else {
									turn = 0
									coeffAndre.check = initialValue
									initialValue = 0
									updateDifference = false
								}
							}
						}
					} else {
						initialValue = 0
						if turn == 0 {
							fmt.Println("RAISE STABLE")
							turn++
						} else {
							if turn == 1 {
								fmt.Println("CALL STABLE")
								turn++
							} else {
								if turn == 2 {
									fmt.Println("FOLD1 STABLE")
									turn++
								} else {
									if turn == 3 {
										fmt.Println("FOLD2 STABLE")
										turn++
									} else {
										if turn == 4 {
											fmt.Println("ALLIN STABLE")
											turn++
										} else {
											fmt.Println("CHECK STABLE")
											turn = 0
										}
									}
								}
							}
						}
					}
				}

			}

			/*if andrePoints > gomezPoints {
				coeffGomez.raise = (coeffAndre.raise + coeffGomez.raise) / 2.0
				coeffGomez.call = (coeffAndre.call + coeffGomez.call) / 2.0
				coeffGomez.fold1 = (coeffAndre.fold1 + coeffGomez.fold1) / 2.0
				coeffGomez.fold2 = (coeffAndre.fold2 + coeffGomez.fold2) / 2.0
			} else {
				coeffAndre.raise = (coeffAndre.raise + coeffGomez.raise) / 2.0
				coeffAndre.call = (coeffAndre.call + coeffGomez.call) / 2.0
				coeffAndre.fold1 = (coeffAndre.fold1 + coeffGomez.fold1) / 2.0
				coeffAndre.fold2 = (coeffAndre.fold2 + coeffGomez.fold2) / 2.0
			}*/
		}
		andrePoints = 0
		gomezPoints = 0
		totalRounds = 0
		if updateDifference {
			lastDifference = difference
		}
		first = false

		fmt.Println("ANDRE VS GOMEZ, first in 300")
		fmt.Printf("%+v\n", coeffAndre)
		fmt.Printf("%+v\n", coeffGomez)
		for andrePoints < 300 && gomezPoints < 300 {
			//	fmt.Println("ANDRE", andrePoints, "-", gomezPoints, "GOMEZ")

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

			isBot := true
			g := newGame(players)

			for g.isGameOn() {

				g.newRound(isBot)
				if !isBot {
					showCards(g.rounds[len(g.rounds)-1].players)
				}
				g.betBlinds(isBot)

				numberOfPlayers := g.getNumberOfPlayersAlive()
				turnToPlay := (g.smallBlindTurn + 2) % numberOfPlayers
				for !g.rounds[len(g.rounds)-1].isBetTurnOver(turnToPlay) {
					if turnToPlay == 0 {
						play(&(g.rounds[len(g.rounds)-1]), turnToPlay, g.bigBlindValue, float64(g.totalStack), coeffAndre, isBot)
					} else {
						play(&(g.rounds[len(g.rounds)-1]), turnToPlay, g.bigBlindValue, float64(g.totalStack), coeffGomez, isBot)
					}

					turnToPlay = (turnToPlay + 1) % numberOfPlayers
				}
				for g.rounds[len(g.rounds)-1].playersAlive > 1 && len(g.rounds[len(g.rounds)-1].sharedCards) < 5 {
					switch len(g.rounds[len(g.rounds)-1].sharedCards) {
					case 0:
						g.rounds[len(g.rounds)-1].sharedCards = deal(&g.rounds[len(g.rounds)-1].deck, 3)
						//printDeck(g.rounds[len(g.rounds)-1].sharedCards)
					case 3:
						g.rounds[len(g.rounds)-1].sharedCards = append(g.rounds[len(g.rounds)-1].sharedCards, deal(&g.rounds[len(g.rounds)-1].deck, 1)...)
					//	printDeck(g.rounds[len(g.rounds)-1].sharedCards)
					case 4:
						g.rounds[len(g.rounds)-1].sharedCards = append(g.rounds[len(g.rounds)-1].sharedCards, deal(&g.rounds[len(g.rounds)-1].deck, 1)...)
						//	printDeck(g.rounds[len(g.rounds)-1].sharedCards)
					}

					for i := range g.rounds[len(g.rounds)-1].players {
						g.rounds[len(g.rounds)-1].players[i].hasSpoken = false
					}
					turnToPlay := g.smallBlindTurn
					for !g.rounds[len(g.rounds)-1].isBetTurnOver(turnToPlay) {
						if turnToPlay == 0 {
							play(&(g.rounds[len(g.rounds)-1]), turnToPlay, g.bigBlindValue, float64(g.totalStack), coeffAndre, isBot)
						} else {
							play(&(g.rounds[len(g.rounds)-1]), turnToPlay, g.bigBlindValue, float64(g.totalStack), coeffGomez, isBot)
						}
						turnToPlay = (turnToPlay + 1) % numberOfPlayers
					}
				}
				g.endRound(isBot)

			}
			if winner := g.getWinnerIndex(); winner == 0 {
				andrePoints++
			} else {
				gomezPoints++
			}
			totalRounds += len(g.rounds)
		}

		fmt.Println("ANDRE", andrePoints, "-", gomezPoints, "GOMEZ")
		fmt.Println("average rounds:", (totalRounds)/(andrePoints+gomezPoints))

	}

}
