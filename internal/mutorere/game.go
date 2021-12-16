package mutorere

import (
	"fmt"
)

type State int32
type GameType int32

const (
	InProgress State = 0
	Finished   State = 1
)

const (
	HumanVersusHuman GameType = 0
	AiVersusRandom   GameType = 1
	HumanVersusAi    GameType = 2
)

type Game struct {
	board        [9]string
	firstPlayer  Player
	secondPlayer Player
	actualPlayer *Player
	state        State
	turn         int
}

// Start a game
func (game *Game) Start() {

	// Check if players are defined before starting the game
	if game.firstPlayer.mark == "" || game.secondPlayer.mark == "" {
		panic("Players are not defined")
	}

	// ResetBoard parmeters before starting the actual game
	game.Reset()

	// Play moves while game is in progress
	for game.state == InProgress {
		game.Play()
	}

	// Print the result of the game if there is a human player
	if game.isHumanPlayer() {
		printBasicBoard(game.board)
		fmt.Printf("Congrats, %s is the winner ! The game lasted %d turns.\n", game.actualPlayer.mark, game.turn)
	}
}

// Play a movement
func (game *Game) Play() {

	// Recover all valid moves
	moves := GetValidesMoves(game)

	// If moves is empty then the game is finished
	if len(moves) <= 0 {
		game.state = Finished
		game.SwitchPlayers()
		return
	}

	// Play a move and switch player turn
	game.actualPlayer.Play(moves)
	game.SwitchPlayers()
	game.turn++
}

// Change turn
func (game *Game) SwitchPlayers() {
	if game.actualPlayer.mark == game.firstPlayer.mark {
		game.actualPlayer = &game.secondPlayer
	} else {
		game.actualPlayer = &game.firstPlayer
	}
}

// Creation of a new game according to the game type
func (game *Game) CreateNewGame(gameType GameType) {
	game.Reset()
	game.setPlayers(gameType)
}

// Set the player according to the game type
func (game *Game) setPlayers(gameType GameType) {

	game.firstPlayer.mark = "X"
	game.secondPlayer.mark = "O"

	switch gameType {
	case HumanVersusHuman:
		// First Player
		game.firstPlayer.entity = HumanAgent
		game.firstPlayer.game = game

		// Second Player
		game.secondPlayer.entity = HumanAgent
		game.secondPlayer.game = game
	case AiVersusRandom:
		// First Player
		game.firstPlayer.entity = SarsaAgent
		game.firstPlayer.game = game
		game.firstPlayer.epsilon = 0.7
		game.firstPlayer.epsilonDecay = 0.0001
		game.firstPlayer.alpha = 0.35
		game.firstPlayer.gamma = 0.9
		game.firstPlayer.Q = intializeQ(game.board)

		// Second Player
		game.secondPlayer.entity = RandomAgent
		game.secondPlayer.game = game
	case HumanVersusAi:
		// First Player
		game.firstPlayer.entity = SarsaAgent
		game.firstPlayer.game = game

		// Second Player
		game.secondPlayer.entity = HumanAgent
		game.secondPlayer.game = game
	default:
		panic("No game type provided")
	}

	game.actualPlayer = &game.firstPlayer
}

// Reset all parameters linked to one game
func (game *Game) Reset() {
	ResetBoard(&game.board)
	game.turn = 0
	game.state = InProgress
}

// Get the winner of the actual game
func (game *Game) GetWinner() int {
	if game.state == Finished {
		if game.actualPlayer.mark == game.firstPlayer.mark {
			return 1
		} else {
			return -1
		}
	}
	return 0
}

// Return if there is a human player in the game
func (game *Game) isHumanPlayer() bool {
	return game.firstPlayer.entity == HumanAgent || game.secondPlayer.entity == HumanAgent
}

// Update the board
func (game *Game) UpdateBoard(move Move) {
	game.board[move.initialPosition] = " "
	game.board[move.finalPosition] = game.actualPlayer.mark
}
