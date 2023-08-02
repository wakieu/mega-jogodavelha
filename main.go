package main

import (
	"fmt"
	"strconv"
)

func welcome() {
	fmt.Println("#####################")
	fmt.Print("Welcome!\n\n")
}

func menu() {	
	fmt.Println("Select your playmode:")
	fmt.Println("1 - Player vs Bot")
	fmt.Println("2 - Local Multiplayer")
	fmt.Print(">> ")
}

func showBoard(game *Board) {
	for i := 0; i < 20; i++ {
		fmt.Println()
	}

	fmt.Println(game.bigBoard)
}

func playVsBot(game *Board) {
	// var win bool = false
	// turns := 0
	// for turns < 81 {

	// }

	fmt.Println("Not implemented yet :(")
}

func startGame(gamemode string) {
	// game := NewBoard()

	// if (gamemode == "1") {playVsBot(&game)}
	// if (gamemode == "2") {playLocal(&game)}
}

var letters = map[string]int{
	"a" : 1, "A" : 1,
	"b" : 2, "B" : 2,
	"c" : 3, "C" : 3,
	"d" : 4, "D" : 4,
	"e" : 5, "E" : 5,
	"f" : 6, "F" : 6,
	"g" : 7, "G" : 7,
	"h" : 8, "H" : 8,
	"i" : 9, "I" : 9,
}

func askCoord() (int, int){
	var s string
	fmt.Println("Input coord as \"a1\" or \"A1\"")
	fmt.Print(">> ")
	for {
		fmt.Scan(&s)
		x := letters[s[:1]]
		y1 := s[1:]
		y, _ := strconv.Atoi(y1)

		if y > 9 { y = 0 }

		if x != 0 && y != 0 {
			return x, y
		} else {
			fmt.Println("Please input valid coords")
			fmt.Print(">> ")
		}
	}
}

func startLocalGame() {
	b := NewBoard()
	b.render()
	for i := 0; i < 81; i++ {
		for {
			x, y := askCoord()
			err := b.validate(x, y)
			if err != nil {
				fmt.Println(err)
				continue
			}
			b.play(x, y)
			b.render()
			break
		}
		win, winner := b.checkWin()
		if win {
			fmt.Println()
			fmt.Println(" Winner is", string(winner), ", CONGRATULATIONS!")
			return
		}
	}

	fmt.Println("=======\n  TIE  \n=======")
}

func main() {
	startLocalGame()
}