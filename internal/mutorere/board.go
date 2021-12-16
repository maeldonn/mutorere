package mutorere

import (
	"errors"
	"fmt"
	"os"
	"strings"
)

// Print the board in ascii art
func PrintBoard(b [9]string) {
	Clear()

	buf, err := os.ReadFile("assets/board.txt")

	if err != nil {
		printBasicBoard(b)
		return
	}

	board := string(buf)

	// Replace indicator chars in string
	board = strings.Replace(board, "0", b[0], 1)
	board = strings.Replace(board, "1", b[1], 1)
	board = strings.Replace(board, "2", b[2], 1)
	board = strings.Replace(board, "3", b[3], 1)
	board = strings.Replace(board, "4", b[4], 1)
	board = strings.Replace(board, "5", b[5], 1)
	board = strings.Replace(board, "6", b[6], 1)
	board = strings.Replace(board, "7", b[7], 1)
	board = strings.Replace(board, "8", b[8], 1)

	fmt.Println(board)
	fmt.Println()
}

// Print the board in ascii art
func printBasicBoard(b [9]string) {
	fmt.Println(b[1], "|", b[2], "|", b[3])
	fmt.Println(b[0], "|", b[8], "|", b[4])
	fmt.Println(b[7], "|", b[6], "|", b[5])
	fmt.Println()
}

// Print the board in tic tac toe mode
func PrintHelp() {
	Clear()

	buf, err := os.ReadFile("assets/board.txt")

	if err != nil {
		printBasicHelp()
		return
	}

	fmt.Println(string(buf))
	fmt.Println()
}

// Print the help in tic tac toe mode
func printBasicHelp() {
	fmt.Println("1 | 2 | 3")
	fmt.Println("0 | 8 | 4")
	fmt.Println("7 | 6 | 5")
	fmt.Println()
}

// Reset the board
func ResetBoard(b *[9]string) {
	*b = [9]string{
		"O", "O", "O",
		"O", "X", "X",
		"X", "X", " ",
	}
}

// Get the position of the empty square
func GetEmptyPostionIndex(board [9]string) (int, error) {
	for i, v := range board {
		if v == " " {
			return i, nil
		}
	}
	return 0, errors.New("there is no empty square")
}

// Convert the actual board to an array
func BoardToState(board [9]string) (state string) {
	for _, v := range board {
		state = state + v
	}
	return state
}

// Initialize all action-states
func intializeQ(board [9]string) (stateActions []Q) {
	states := Permutations(board[:])
	for _, s := range states {
		var stateAction Q
		var state [9]string
		copy(state[:], s[:9])
		stateAction.state = BoardToState(state)
		stateAction.actions = PermutationsActions()
		stateActions = append(stateActions, stateAction)
	}
	return stateActions
}
