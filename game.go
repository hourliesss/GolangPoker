package main

import (
	"bufio"
	"fmt"
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
}

func launchGame() game {

	sp := []person{}

	printSimPoker()

	for {
		fmt.Print("-> How many real players? (2 to 9)\n")

		text := readFromTerminal()

		number, err := strconv.Atoi(text)
		if err == nil {
			if number > 1 && number < 10 {
				for i := 0; i < number; i++ {
					p := person{stack: 1000}
					fmt.Print("-> Name of player ", i+1, "?\n")
					name := readFromTerminal()
					p.name = name
					sp = append(sp, p)
				}

				// AI IN PROCESS

				/*fmt.Print("-> How many virtual players? (" + strconv.Itoa(9-number) + " max)\n")
				text2 := readFromTerminal()

				numberOfBots, err2 := strconv.Atoi(text2)
				if err2 == nil || number < 0 || numberOfBots+number > 9 {
					for i := 0; i < number; i++ {
						p := person{stack: 1000, isBot: true}
						fmt.Print("-> Name of the virtual player ", i+1, "?\n")
						name := readFromTerminal()
						p.name = name
						sp = append(sp, p)
					}
				} else {
					fmt.Println("Unvalid number")
					os.Exit(1)
				}*/
			} else {
				fmt.Println("Sorry, Only 2 to 9 players can play")
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
	}

	return g
}

func (g *game) newRound() {
	d := newDeck()
	players := []player{}
	for _, person := range (*g).players {
		if person.stack > 0 {
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
	}
	(*g).rounds = append((*g).rounds, r)
	fmt.Println()
	fmt.Println("ROUND", len((*g).rounds))
	fmt.Println("\tPlayers:")
	for _, rp := range (*g).rounds[len((*g).rounds)-1].players {
		fmt.Println("\t\t", rp.name, "stack:", rp.initialStack)
	}

	fmt.Println()
}

func (g game) betBlinds() {
	g.rounds[len(g.rounds)-1].bet(g.smallBlindTurn, g.smallBlindValue)
	fmt.Println("\t", g.rounds[len(g.rounds)-1].players[g.smallBlindTurn].name, "is the small blind and bet", g.smallBlindValue)

	g.rounds[len(g.rounds)-1].bet(g.bigBlindTurn, g.bigBlindValue)
	fmt.Print("\t", g.rounds[len(g.rounds)-1].players[g.bigBlindTurn].name, " is the big blind and bet ", g.bigBlindValue, "\n\n")
}
func (pointerToRound *round) bet(playerIndex int, amount int) {
	if (*pointerToRound).players[playerIndex].initialStack-(*pointerToRound).players[playerIndex].roundBet <= amount {
		(*pointerToRound).pot += (*pointerToRound).players[playerIndex].initialStack - (*pointerToRound).players[playerIndex].roundBet
		(*pointerToRound).players[playerIndex].roundBet = (*pointerToRound).players[playerIndex].initialStack
		(*pointerToRound).maxBet = maxInt((*pointerToRound).players[playerIndex].initialStack, (*pointerToRound).maxBet)
		(*pointerToRound).players[playerIndex].isAllIn = true
		(*pointerToRound).playersAllIn++
		fmt.Println("\t", (*pointerToRound).players[playerIndex].name, " is all in!")
	} else {
		(*pointerToRound).players[playerIndex].roundBet += amount
		(*pointerToRound).pot += amount
		(*pointerToRound).maxBet = (*pointerToRound).players[playerIndex].roundBet
	}
}

func (g *game) endRound() {
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
					if playerInRound.won != 0 {
						for _, sp := range results {
							for pr := range sp {
								if sp[pr].playerIndex == playerIndex {
									if r.playersAlive > 1 {
										fmt.Println("\n\t"+playerInRound.name, "won", playerInRound.won, "with "+sp[pr].handScore.toString())
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

	// actualize the blind turns
	playersNumber := (*g).getNumberOfPlayersAlive()
	(*g).smallBlindTurn = ((*g).smallBlindTurn + 1) % playersNumber
	(*g).bigBlindTurn = ((*g).bigBlindTurn + 1) % playersNumber
	(*g).dealerTurn = ((*g).dealerTurn + 1) % playersNumber
	// actualize the blind values
	if len((*g).rounds)%10 == 0 {
		(*g).bigBlindValue *= 2
		(*g).smallBlindValue *= 2
		fmt.Println("the blinds increases!")
		fmt.Println("Small blind: ", (*g).smallBlindValue)
		fmt.Println("Big blind: ", (*g).bigBlindValue)
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

func play(r *round, playerIndex int, bigBlind int) {
	pp := &((*r).players[playerIndex])
	p := (*pp)
	if p.isAllIn {
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
						(*r).bet(playerIndex, (*r).maxBet-p.roundBet)
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
										(*r).bet(playerIndex, (*r).maxBet-p.roundBet+number)
									}
								}
							} else {
								(*r).bet(playerIndex, p.initialStack-p.roundBet)
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
		} // AI SECTION IN PROCESS //
		/*else {
			odds := calculateOdds(p.cards, r.sharedCards)
			fmt.Println("odds:", odds)
			exp := expoFunction(odds)
			fmt.Println("expo:", exp)
			nlr := normaleLawRandomization(1)
			fmt.Println("normalization:", nlr)
			decision := exp * nlr
			fmt.Println("decision:", decision)
			if (*r).maxBet == p.roundBet {
				if decision < 0.2 {
					fmt.Print("\t", p.name, ", checks\n\n")
				} else {
					raiseValue := maxInt(int(math.Ceil(decision*float64(p.initialStack-p.roundBet)*0.5)), bigBlind)
					(*r).bet(playerIndex, raiseValue)
					fmt.Print("\t", p.name, " raises ", raiseValue, "\n\n")
				}
			} else {
				toCallRatio := float64((*r).maxBet-p.roundBet) / float64(p.initialStack-p.roundBet)
				if decision < toCallRatio {
					(*pp).hasFolded = true
					(*r).playersAlive--
					fmt.Print("\t", p.name, ", folds\n\n")
				} else {
					if decision < toCallRatio+0.2 {
						(*r).bet(playerIndex, (*r).maxBet-p.roundBet)
						fmt.Print("\t", p.name, " calls\n\n")
					} else {
						raiseValue := maxInt(int(math.Ceil((decision-toCallRatio)*(float64(p.initialStack-p.roundBet)))), bigBlind)
						(*r).bet(playerIndex, (*r).maxBet-p.roundBet+raiseValue)
						fmt.Print("\t", p.name, " raises ", raiseValue, "\n\n")
					}
				}
			}
		}*/

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
