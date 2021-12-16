package mutorere

import (
	"github.com/inancgumus/screen"
	"strconv"
)

// Clear screen
func Clear() {
	screen.Clear()
	screen.MoveTopLeft()
}

// Return the index and the max value of an array
func ArrayMaxIndex(a []Action) (max float64, index int) {
	max = a[0].value
	for i, v := range a {
		if max < v.value {
			max = v.value
			index = i
		}
	}
	return max, index
}

// Return the index of the state in Q
func FindState(Q []Q, state string) int {
	for i, v := range Q {
		if v.state == state {
			return i
		}
	}
	return -1
}

// Return the index of an actions in actions
func FindAction(actions []Action, action Move) int {
	for i, v := range actions {
		if v.action == action {
			return i
		}
	}
	return -1
}

// Convert a string to an integer
func ConvertStringToInteger(str string) int {
	if i, err := strconv.Atoi(str); err == nil {
		return i
	}
	return -1
}

// Generate permutations
// TODO: Remove all duplicates
func Permutations(arr []string) [][]string {
	var helper func([]string, int)
	res := [][]string{}

	helper = func(arr []string, n int) {
		if n == 1 {
			tmp := make([]string, len(arr))
			copy(tmp, arr)
			res = append(res, tmp)
		} else {
			for i := 0; i < n; i++ {
				helper(arr, n-1)
				if n%2 == 1 {
					tmp := arr[i]
					arr[i] = arr[n-1]
					arr[n-1] = tmp
				} else {
					tmp := arr[0]
					arr[0] = arr[n-1]
					arr[n-1] = tmp
				}
			}
		}
	}
	helper(arr, len(arr))
	return res
}

// Generate permutations of different moves
func permutationsMoves() []Move {
	var moves []Move
	for i := 0; i < 9; i++ {
		for j := 0; j < 9; j++ {
			moves = append(moves, Move{i, j})
		}
	}
	return moves
}

// Generate permutations of different actions
func PermutationsActions() []Action {
	var actions []Action
	moves := permutationsMoves()
	for _, move := range moves {
		var action Action
		action.value = 0
		action.action = move
		actions = append(actions, action)
	}
	return actions
}
