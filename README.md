# GolangPoker
Texas Hold'em poker in Golang

This is my first project in Golang

You can play multiplayer on a single computer or play against a bot
If you want to play against a bot, you need to play 1vs1 (more players not yet implemented, the IA will fold most of the time)
Before play, comment or uncomment the main function in main.go
func main() {
	poker() // this is the mode to play mutiplayer or single player 
	//botMode() // this is the bot mode to see andre playing vs gomez, this mode has been created to improve the IA
}


To run the project:

--> go run main.go cards.go game.go odds.go player.go math_functions.go bot.go