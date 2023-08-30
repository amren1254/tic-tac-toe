package main

import (
	"errors"
	"fmt"
)

type Board [3][3]Player

func createNewBoard() Board {
	return Board{}
}

func (b *Board) updateBoard(input Player, active Player) error {
	keymap := [9][2]int{
		{0, 0}, {0, 1}, {0, 2},
		{1, 0}, {1, 1}, {1, 2},
		{2, 0}, {2, 1}, {2, 2},
	}
	mappedSquare := keymap[input-1]
	square := b[mappedSquare[0]][mappedSquare[1]]
	if square != 0 {
		return errors.New("this square on the board is already chosen, please try again")
	}
	b[mappedSquare[0]][mappedSquare[1]] = active
	return nil
}

func (b Board) displayCurrentBoard() error {
	for _, row := range b {
		for _, val := range row {
			fmt.Print(val, "\t")
		}
		fmt.Println()
	}
	return nil
}
