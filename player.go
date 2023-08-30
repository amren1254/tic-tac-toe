package main

import "fmt"

type Player int

const (
	Player0 Player = iota
	Player1
	Player2
)

func (p Player) String() string {
	switch p {
	case Player0:
		return " . "
	case Player1:
		return " X "
	case Player2:
		return " O "
	}
	return fmt.Sprintf("%d", int(p)) // this shouldn't happen
}

func (p Player) Switch() Player {
	if p == Player1 {
		return Player2
	}
	return Player1 // if Player0 || Player2, so we don't have to initialise
}

func getPlayerInput() Player {
	var choice Player
	fmt.Println("Please input your choice")
	fmt.Scanln(&choice)
	for choice > 9 && choice == 0 {
		fmt.Println("Please enter a number between 1-9.")
		choice = getPlayerInput()
	}
	return choice
}
