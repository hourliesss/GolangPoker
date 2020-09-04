package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"strconv"
	"strings"
)

type round struct {
	sharedCards  deck
	players      []player
	deck         deck
	pot          int
	maxBet       int
	playersAlive int
	playersAllIn int
	limitBet     int
}

type roundResults = [][]playerResult

type result struct {
	position      int
	playerResults []playerResult
}

type playerResult struct {
	playerIndex int
	maxWin      int
	handScore   handScore
}

type sidePot struct {
	value         int
	playerIndexes []int
}

type game struct {
	players         []person
	rounds          []round
	smallBlindTurn  int
	bigBlindTurn    int
	dealerTurn      int
	smallBlindValue int
	bigBlindValue   int
	totalStack      int
}

func launchGame() game {

	sp := []person{}

	printSimPoker()

	for {
		fmt.Print("-> How many real players? (0 to 9)\n")

		text := readFromTerminal()

		number, err := strconv.Atoi(text)
		if err == nil {
			if number >= 0 && number < 10 {
				for i := 0; i < number; i++ {
					p := person{stack: 1000}
					fmt.Print("-> Name of player ", i+1, "?\n")
					name := readFromTerminal()
					p.name = name
					sp = append(sp, p)
				}

				fmt.Print("-> How many virtual players? (" + strconv.Itoa(9-number) + " max)\n")
				text2 := readFromTerminal()

				numberOfBots, err2 := strconv.Atoi(text2)
				if err2 == nil && numberOfBots >= 0 && numberOfBots+number < 10 {
					for i := 0; i < numberOfBots; i++ {
						p := person{stack: 1000, isBot: true}
						fmt.Print("-> Name of the virtual player ", i+1, "?\n")
						name := readFromTerminal()
						p.name = name
						sp = append(sp, p)
					}
				} else {
					fmt.Println("Unvalid number")
					os.Exit(1)
				}
			} else {
				fmt.Println("Sorry, Only 1 to 9 players can play")
				os.Exit(1)
			}
		} else {
			fmt.Println("Unvalid number")
			os.Exit(1)
		}

		return newGame(sp)

	}
}

func newGame(p []person) game {
	g := game{
		smallBlindTurn:  0,
		bigBlindTurn:    1,
		dealerTurn:      len(p) - 1,
		smallBlindValue: 10,
		bigBlindValue:   20,
		rounds:          []round{},
		players:         p,
		totalStack:      p[0].stack * len(p),
	}

	return g
}

func (g *game) newRound(botMode bool) {
	d := newDeck()
	players := []player{}
	biggestStack := 0
	secondStack := 0
	for _, person := range (*g).players {
		if person.stack > 0 {
			if biggestStack < person.stack {
				secondStack = biggestStack
				biggestStack = person.stack
			} else {
				if secondStack < person.stack {
					secondStack = person.stack
				}
			}
			c := deal(&d, 2)
			player := player{
				name:         person.name,
				initialStack: person.stack,
				cards:        c,
				roundBet:     0,
				hasSpoken:    false,
				hasFolded:    false,
				isAllIn:      false,
				isBot:        person.isBot,
			}

			players = append(players, player)

		}
	}

	r := round{
		deck:         d,
		players:      players,
		playersAlive: len(players),
		limitBet:     secondStack,
	}
	(*g).rounds = append((*g).rounds, r)
	if !botMode {
		fmt.Println()
		fmt.Println("ROUND", len((*g).rounds))
		fmt.Println("\tPlayers:")
		for _, rp := range (*g).rounds[len((*g).rounds)-1].players {
			fmt.Println("\t\t", rp.name, "stack:", rp.initialStack)
		}

		fmt.Println()
	}

}

func (g game) betBlinds(botMode bool) {
	g.rounds[len(g.rounds)-1].bet(g.smallBlindTurn, g.smallBlindValue, botMode)
	if !botMode {
		fmt.Println("\t", g.rounds[len(g.rounds)-1].players[g.smallBlindTurn].name, "is the small blind and bet", g.smallBlindValue)
	}

	if len(g.rounds[len(g.rounds)-1].players) == 2 && g.rounds[len(g.rounds)-1].playersAllIn == 1 {
		g.rounds[len(g.rounds)-1].bet(g.bigBlindTurn, g.rounds[len(g.rounds)-1].maxBet, botMode)
	} else {
		g.rounds[len(g.rounds)-1].bet(g.bigBlindTurn, g.bigBlindValue, botMode)
	}
	if !botMode {
		fmt.Print("\t", g.rounds[len(g.rounds)-1].players[g.bigBlindTurn].name, " is the big blind and bet ", g.bigBlindValue, "\n\n")
	}

}
func (pointerToRound *round) bet(playerIndex int, amount int, botMode bool) {
	bet := minInt(amount, (*pointerToRound).limitBet-(*pointerToRound).players[playerIndex].roundBet)
	if (*pointerToRound).players[playerIndex].initialStack-(*pointerToRound).players[playerIndex].roundBet <= bet {
		(*pointerToRound).pot += (*pointerToRound).players[playerIndex].initialStack - (*pointerToRound).players[playerIndex].roundBet
		(*pointerToRound).players[playerIndex].roundBet = (*pointerToRound).players[playerIndex].initialStack
		(*pointerToRound).maxBet = maxInt((*pointerToRound).players[playerIndex].initialStack, (*pointerToRound).maxBet)
		(*pointerToRound).players[playerIndex].isAllIn = true
		(*pointerToRound).playersAllIn++
		if !botMode {
			fmt.Println("\t", (*pointerToRound).players[playerIndex].name, " is all in!")
		}

	} else {
		(*pointerToRound).players[playerIndex].roundBet += bet
		(*pointerToRound).pot += bet
		(*pointerToRound).maxBet = (*pointerToRound).players[playerIndex].roundBet
	}
}

func (g *game) endRound(botMode bool) {
	r := (*g).rounds[len((*g).rounds)-1]
	// actualize the stacks

	results := r.getResults()
	alreadyWon := 0
	for r.pot > alreadyWon {
		for _, sp := range results {
			for pr := range sp {
				for _, roundPlayer := range r.players {
					sp[pr].maxWin += minInt(r.players[sp[pr].playerIndex].roundBet, roundPlayer.roundBet)
				}
				sp[pr].maxWin -= alreadyWon
			}
			orderResultsWithMaxWins(sp)
			r.players[sp[0].playerIndex].won = sp[0].maxWin / len(sp)
			alreadyWon += r.players[sp[0].playerIndex].won
			if len(sp) > 1 {
				for i := 1; i < len(sp); i++ {
					r.players[sp[i].playerIndex].won = r.players[sp[i-1].playerIndex].won + (sp[i].maxWin-sp[i-1].maxWin)/(len(sp)-i)
					alreadyWon += r.players[sp[i].playerIndex].won
				}
			}
		}
	}
	for i := range (*g).players {
		if (*g).players[i].stack > 0 {
			for playerIndex, playerInRound := range r.players {
				if (*g).players[i].name == playerInRound.name {
					(*g).players[i].stack -= playerInRound.roundBet
					(*g).players[i].stack += playerInRound.won
					if !botMode {
						for _, sp := range results {
							for pr := range sp {
								if sp[pr].playerIndex == playerIndex {
									if r.playersAlive > 1 {
										fmt.Print("\n\t"+playerInRound.name, "has ")
										printDeck(playerInRound.cards)
										if playerInRound.won != 0 {
											fmt.Println("\n\t"+playerInRound.name, "won", playerInRound.won, "with "+sp[pr].handScore.toString())
										}

									} else {
										fmt.Println("\n\t"+playerInRound.name, "won", playerInRound.won)
									}
									break
								}
							}
						}
					}

				}
			}

		}
	}

	/*for _, p := range r.players {
		if !p.hasFolded {
			fmt.Println(p.name, "has", p.cards[0].toString(), " ", p.cards[1].toString())
		}
	}*/
	// actualize the blind turns
	playersNumber := (*g).getNumberOfPlayersAlive()
	(*g).smallBlindTurn = ((*g).smallBlindTurn + 1) % playersNumber
	(*g).bigBlindTurn = ((*g).bigBlindTurn + 1) % playersNumber
	(*g).dealerTurn = ((*g).dealerTurn + 1) % playersNumber
	// actualize the blind values
	if len((*g).rounds)%10 == 0 {
		(*g).bigBlindValue *= 2
		(*g).smallBlindValue *= 2
		if !botMode {
			fmt.Println("the blinds increases!")
			fmt.Println("Small blind: ", (*g).smallBlindValue)
			fmt.Println("Big blind: ", (*g).bigBlindValue)
		}

	}

}

func orderResultsWithMaxWins(prs []playerResult) {
	for i := range prs {
		j := i
		for j >= 0 && j < len(prs)-1 && prs[j+1].maxWin < prs[j].maxWin {
			prs[j+1], prs[j] = prs[j], prs[j+1]
			j--
		}

	}
}

func (r round) getResults() roundResults {
	if r.playersAlive == 1 {
		for i := range r.players {
			if !r.players[i].hasFolded {
				winner := playerResult{
					playerIndex: i,
					maxWin:      0,
				}
				results := roundResults{}
				return append(results, []playerResult{winner})
			}
		}
	}
	sh := []handScore{}
	for i, p := range r.players {
		if !p.hasFolded {
			sh = append(sh, calculateFiveBestCards(append(r.sharedCards, p.cards...)))
			sh[len(sh)-1].playerIndex = i
		}
	}

	for i := range sh {
		j := i
		for j >= 0 && j < len(sh)-1 && sh[j+1].score > sh[j].score {
			sh[j+1], sh[j] = sh[j], sh[j+1]
			j--
		}
	}
	noswap := false
	for !noswap {
		noswap = true
		j := 0
		for j >= 0 && j < len(sh)-1 {
			x := isBetterHand(sh[j], sh[j+1])
			if x == 1 {
				sh[j+1], sh[j] = sh[j], sh[j+1]
				noswap = false
				sh[j+1].isPrecedentEqual = false
				sh[j].isPrecedentEqual = false
			}
			if x == 0 {
				sh[j+1].isPrecedentEqual = true
			}
			j++
		}
	}
	results := roundResults{}
	for _, h := range sh {
		playerR := playerResult{
			playerIndex: h.playerIndex,
			maxWin:      0,
			handScore:   h,
		}
		if !h.isPrecedentEqual {
			slicePlayerResult := []playerResult{playerR}
			results = append(results, slicePlayerResult)
		} else {
			results[len(results)-1] = append(results[len(results)-1], playerR)
		}
	}

	return results

}

func (r round) isBetTurnOver(playerIndex int) bool {
	p := r.players[playerIndex]
	if (!p.hasFolded && p.hasSpoken && (p.roundBet == r.maxBet || p.isAllIn)) || r.playersAlive == 1 || (r.playersAlive-r.playersAllIn <= 1 && (p.roundBet == r.maxBet || p.isAllIn)) {
		return true
	}

	return false
}

func play(r *round, playerIndex int, bigBlind int, totalStacks float64, coeff coefficients, botMode bool) {
	pp := &((*r).players[playerIndex])
	p := (*pp)
	if p.isAllIn && !botMode {
		fmt.Print("\t", p.name, " is all in\n")
	}
	if !p.isAllIn && !p.hasFolded {
		if !p.isBot {
			fmt.Print("\t-> ", p.name, " What do you want to do ? You have ", p.initialStack-p.roundBet, "\n")
			if (*r).maxBet == p.roundBet {
				fmt.Println("\t\tCheck (ch)")
			} else {
				fmt.Println("\t\tCall", (*r).maxBet-p.roundBet, "(ca)")
			}
			if p.initialStack > (*r).maxBet {
				fmt.Println("\t\tRaise (r)")
			}
			fmt.Print("\t\tFold (f) ")
			validInput := false
			for !validInput {
				validInput = true
				c := strings.ToLower(readFromTerminal())
				switch c {
				case "check", "ch":
					if (*r).maxBet != p.roundBet {
						fmt.Print("\t", p.name, ", checks\n\n")
						validInput = false
					}
				case "call", "ca":
					if (*r).maxBet != p.roundBet {
						fmt.Print("\t", p.name, " calls\n\n")
						(*r).bet(playerIndex, (*r).maxBet-p.roundBet, false)
					}
				case "fold", "f":
					(*pp).hasFolded = true
					(*r).playersAlive--
					fmt.Print("\t", p.name, " folds\n\n")
				case "raise", "r":
					if p.initialStack > (*r).maxBet {
						valueRaise := 0
						for valueRaise == 0 {
							if p.initialStack-(*r).maxBet > bigBlind {
								fmt.Print("\t---> How much do you want to raise? (", bigBlind, "-", p.initialStack-(*r).maxBet, ")")
								text := readFromTerminal()
								number, err := strconv.Atoi(text)
								if err == nil {
									if number >= bigBlind {
										valueRaise = number
										(*r).bet(playerIndex, (*r).maxBet-p.roundBet+number, false)
									}
								}
							} else {
								(*r).bet(playerIndex, p.initialStack-p.roundBet, false)
							}
						}
						fmt.Print("\t", p.name, " raises ", valueRaise, "\n\n")
					}
				default:
					validInput = false
				}
				if !validInput {
					fmt.Println("Your choice has not been understood")
				}
			}
		} else {
			var decision float64

			if (*pp).hasSpoken {
				decision = (*pp).decision
			} else {
				odds := calculateOdds(p.cards, r.sharedCards)
				//fmt.Println("odds:", odds)
				/*	var exp float64
					if len(r.sharedCards) == 0 {
						exp = expoFunction(odds, 3.0)
					} else {
						exp = expoFunction(odds, 2.0)
					}
					//		fmt.Println("expo:", exp)
					nlr := normaleLawRandomization(1, 0)*/
				//		fmt.Println("normalization:", nlr)
				//decision = exp * nlr
				//fmt.Println("decision before random:", odds)
				decision = odds * normaleLawRandomization(0.7, 0)
				(*pp).decision = decision
			}

			//fmt.Println("decision after random:", decision)
			if (*r).maxBet == p.roundBet {
				if decision < coeff.check {
					if !botMode {
						fmt.Print("\t", p.name, " checks\n\n")
					}
				} else {
					raiseValue := maxInt(int(math.Ceil(expoFunction(decision, 3.0)*(totalStacks/float64(len(r.players)))*coeff.raise*normaleLawRandomization(1, 0))), bigBlind)
					raiseValue -= raiseValue % (bigBlind / 2)
					(*r).bet(playerIndex, raiseValue, botMode)
					if !botMode {
						fmt.Print("\t", p.name, " raises ", raiseValue, "\n\n")
					}

				}
			} else {
				toCallRatio := float64((*r).maxBet-p.roundBet) / (totalStacks * math.Exp(float64(p.roundBet)/(totalStacks/float64(len(r.players)))))

				limit1 := coeff.fold1
				limit2 := raiseExpo(toCallRatio * coeff.fold2)
				fmt.Println("tocallratio:", limit2)
				if decision < limit1 || (decision < limit2) /* && (float64(p.initialStack-p.roundBet) > 0.05*totalStacks*/ {
					(*pp).hasFolded = true
					(*r).playersAlive--
					if !botMode {
						fmt.Print("\t", p.name, " folds\n\n")
					}

				} else {
					overFold := decision - limit2
					if overFold < coeff.call {
						(*r).bet(playerIndex, (*r).maxBet-p.roundBet, botMode)
						if !botMode {
							fmt.Print("\t", p.name, " calls\n\n")
						}

					} else {
						var raiseValue int
						if decision > coeff.allIn {
							raiseValue = p.initialStack - p.roundBet
						} else {
							raiseValue = maxInt(int(math.Ceil(expoFunction(overFold, 3.0)*(totalStacks/float64(len(r.players)))*coeff.raise*normaleLawRandomization(0.5, 0))), bigBlind)
							raiseValue -= raiseValue % (bigBlind / 2)
						}

						(*r).bet(playerIndex, (*r).maxBet-p.roundBet+raiseValue, botMode)
						if !botMode {
							fmt.Print("\t", p.name, " raises ", raiseValue, "\n\n")
						}

					}
				}
			}
		}
		(*pp).hasSpoken = true

	}

}

func readFromTerminal() string {
	reader := bufio.NewReader(os.Stdin)
	text, _ := reader.ReadString('\n')
	text = strings.Replace(text, "\n", "", -1)
	text = strings.Replace(text, "\r", "", -1)

	return text
}

func (g game) isGameOn() bool {
	n := 0
	for _, player := range g.players {
		if player.stack > 0 {
			n++
		}
		if n > 1 {
			return true
		}
	}
	return false
}

func (g game) getNumberOfPlayersAlive() int {
	n := 0
	for _, player := range g.players {
		if player.stack > 0 {
			n++
		}
	}
	return n
}

func (g game) getWinner() string {
	for _, player := range g.players {
		if player.stack > 0 {
			return player.name
		}
	}
	return ""
}

func (g game) getWinnerIndex() int {
	for i, player := range g.players {
		if player.stack > 0 {
			return i
		}
	}
	return 0
}
