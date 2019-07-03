package main

import (
	"fmt"
	"math/rand"
)

//Init is called by main to instantiate the playground
func Init() {
	//1. Initialize the size of tic tac toe array
	fmt.Println("What size of the grid do you want?")
	fmt.Scanf("%d", &gridSize)
	//Below int is used to fill up the slots in the array
	init := 1
	for i := 0; i < gridSize; i++ {
		var row []string
		design.Positions = append(design.Positions, row)
		for j := 0; j < gridSize; j++ {
			design.Positions[i] = append(design.Positions[i], emptySlot)
			//Filling the free slots array with the values
			freeSlots = append(freeSlots, init)
			init++
		}

	}
	fmt.Println("1 - Multi Player")
	fmt.Println("2 - Single Player")
	fmt.Println("3 - Watch the computer play")
	fmt.Scanf("%d", &playType)
	//2. Initialize the players
	player1 = Player{Notation: 1, Symbol: "X"}
	player2 = Player{Notation: 2, Symbol: "0"}
	currentPlayer = player1
}

//StartPlay will be called after Init
func StartPlay() {
	// 1. Toggle between the players to get input- select the player
	// 2. Take the input from the player
	// 3. Check whether the input is valid( and not occupied so far)
	// 4. Update the array
	// 5. Check for the winner
	// 6. If player wins, declare and exit, else, goto 1
	i := 0
	for {
		ShowGrid()
		fmt.Println("Player ", currentPlayer.Notation, " : Enter your spot")
		switch playType {
		case 1:
			fmt.Scanf("%d", &i)
		case 2:
			if currentPlayer == player1 {
				fmt.Scanf("%d", &i)
			} else {
				i = freeSlots[rand.Intn(len(freeSlots))]
			}
		case 3:
			i = freeSlots[rand.Intn(len(freeSlots))]
		}
		if CheckValidMove(i) {
			if CheckIfSlotOccupied(i) {
				fmt.Println("Already occupied, choose another slot")
			} else {
				UpdateGrid(i)
				resultForWinner, playerWinner := CheckForWinner()
				if resultForWinner {
					fmt.Println("We have found a winner here, which is :", playerWinner.Notation)
					ShowGrid()
					break
				}
				if CheckForDraw() {
					fmt.Println("The match is drawn")
					ShowGrid()
					break
				}
				TogglePlayer()
			}
		} else {
			fmt.Println("Invalid input, try again")
		}
	}

}
func UpdateGrid(input int) {
	row, column := GetRowColumn(input)
	for i, v := range freeSlots {
		if v == input {
			freeSlots = append(freeSlots[:i], freeSlots[i+1:]...)
			break
		}
	}
	design.Positions[row][column] = currentPlayer.Symbol
}
