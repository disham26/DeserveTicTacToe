package main

import "fmt"

// 1 -> 0,0
// 2 -> 0,1
// 3 -> 0,2
// 4 -> 1,0
// 5 -> 1,1
// 6 -> 1,2
// 7 -> 2,0
// 8 -> 2,1
// 9 -> 2,2
func CheckIfSlotOccupied(input int) bool {
	row, column := GetRowColumn(input)
	return design.Positions[row][column] != emptySlot

}
func GetRowColumn(input int) (int, int) {
	return (input - 1) / gridSize, (input - 1) % gridSize
}
func CheckValidMove(input int) bool {
	fmt.Println(input)
	return 0 < input && input < (gridSize*gridSize)+1
}
func TogglePlayer() {
	//Will make this better
	if currentPlayer == player1 {
		currentPlayer = player2
	} else {
		currentPlayer = player1
	}
}
func ShowGrid() {
	for i := 0; i < gridSize; i++ {
		for j := 0; j < gridSize; j++ {
			fmt.Print(design.Positions[i][j] + " ")
		}
		fmt.Println()

	}
}

func CheckForWinner() (bool, Player) {
	//1. Check for a row winner
	var result bool
	var player Player
	result, player = CheckForRowWinner()
	if result {
		return true, player
	} else {
		//2. Check for a column winner
		result, player = CheckForColumnWinner()
		if result {
			return true, player
		} else {
			//3. Check for a diagonal winner
			result, player = CheckForDiagonalWinner()
			if result {
				return true, player
			}
		}
	}
	return false, player

}

func CheckForRowWinner() (bool, Player) {
	var player Player

	for i := 0; i < gridSize; i++ {
		found := true
		if design.Positions[i][0] != emptySlot {
			for j := 1; j < gridSize; j++ {
				if design.Positions[i][j] != design.Positions[i][0] {
					found = false
				}
			}
			if found {
				if design.Positions[i][0] == player1.Symbol {
					return true, player1
				} else {
					return true, player2
				}
			}

		}

	}
	return false, player
}
func CheckForColumnWinner() (bool, Player) {
	var player Player
	for i := 0; i < gridSize; i++ {
		found := true
		if design.Positions[0][i] != emptySlot {
			for j := 1; j < gridSize; j++ {
				if design.Positions[j][i] != design.Positions[0][i] {
					found = false
				}
			}
			if found {
				if design.Positions[0][i] == player1.Symbol {
					return true, player1
				} else {
					return true, player2
				}
			}

		}

	}
	return false, player
}

func CheckForDiagonalWinner() (bool, Player) {
	fmt.Println("checking diagonal")
	var player Player
	//first diagonal
	if design.Positions[0][0] != emptySlot {
		found := true
		for i := 1; i < gridSize; i++ {
			if design.Positions[i][i] != design.Positions[i-1][i-1] {
				found = false
			}
		}
		if found {
			if design.Positions[0][0] == player1.Symbol {
				return true, player1
			} else {
				return true, player2
			}
		}
	} else if design.Positions[0][gridSize-1] != emptySlot {
		found := true
		for i := 1; i < gridSize; i++ {
			if design.Positions[i][gridSize-i-1] != design.Positions[i-1][gridSize-i] {
				found = false
			}
		}
		if found {
			if design.Positions[0][gridSize-1] == player1.Symbol {
				return true, player1
			} else {
				return true, player2
			}
		}

	}
	return false, player
}

func CheckForDraw() bool {
	for i := 0; i < gridSize; i++ {
		for j := 0; j < gridSize; j++ {
			if design.Positions[i][j] == emptySlot {
				return false
			}
		}
	}
	return true
}
