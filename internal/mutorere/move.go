package mutorere

type Move struct {
	initialPosition int
	finalPosition   int
}

// Get all valides moves for the current position
func GetValidesMoves(game *Game) []Move {

	var moves []Move

	emptyIndex, err := GetEmptyPostionIndex(game.board)

	if err != nil {
		panic(err)
	}

	if emptyIndex == 8 {

		for i, v := range game.board {
			if v == game.actualPlayer.mark && canMoveToPutahi(game, i) {
				moves = append(moves, Move{i, emptyIndex})
			}
		}
	} else {

		leftKewai := emptyIndex - 1
		rightKewai := emptyIndex + 1
		center := 8

		if leftKewai == -1 {
			leftKewai = 7
		}

		if rightKewai == 8 {
			rightKewai = 0
		}

		if game.board[leftKewai] == game.actualPlayer.mark {
			moves = append(moves, Move{leftKewai, emptyIndex})
		}

		if game.board[rightKewai] == game.actualPlayer.mark {
			moves = append(moves, Move{rightKewai, emptyIndex})
		}

		if game.board[center] == game.actualPlayer.mark {
			moves = append(moves, Move{center, emptyIndex})
		}
	}

	return moves
}

// Return if the move is in the valide moves
func IsValideMove(moves []Move, move Move) bool {

	for _, v := range moves {
		if v.initialPosition == move.initialPosition && v.finalPosition == move.finalPosition {
			return true
		}
	}
	return false
}

// Return if the kewai can move to putahi
func canMoveToPutahi(game *Game, initialPosition int) bool {

	leftKewai := initialPosition - 1
	rightKewai := initialPosition + 1

	if leftKewai == -1 {
		leftKewai = 7
	}

	if rightKewai == 8 {
		rightKewai = 0
	}

	if game.board[leftKewai] != game.actualPlayer.mark || game.board[rightKewai] != game.actualPlayer.mark {
		return true
	}
	return false
}
