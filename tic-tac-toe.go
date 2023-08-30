package main

import (
	"fmt"
)

func main() {
	fmt.Println("Let's start TIC-TAC-TOE")
	board := createNewBoard()
	var active Player
	movesCount := 0
	activePlayer := active.Switch()
	for movesCount < 9 {
		fmt.Println("Active : ", activePlayer.String())
		board.displayCurrentBoard()
		input := getPlayerInput()
		err := board.updateBoard(input, activePlayer)
		if err != nil {
			fmt.Println("Warning :", err.Error())
		}
		movesCount += 1
		activePlayer = activePlayer.Switch()
		//check for winner
		hasMatch, winner := board.hasAnyHorizontalMatch(activePlayer)
		if hasMatch {
			fmt.Printf("%v wins! (horizontal row)", winner)
			return
		}

		hasMatch, winner = board.hasAnyVerticalMatch(activePlayer)
		if hasMatch {
			fmt.Printf("%v wins! (vertical row)", winner)
			return
		}

		hasMatch, winner = board.hasAnyDiagonalMatch(activePlayer)
		if hasMatch {
			fmt.Printf(" wins! (diagonal), %d", winner)
			return
		}
	}
}
