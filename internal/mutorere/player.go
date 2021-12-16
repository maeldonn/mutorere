package mutorere

import (
	"errors"
	"fmt"
	"math/rand"
)

type Entity int32

const (
	HumanAgent  Entity = 0
	RandomAgent Entity = 1
	SarsaAgent  Entity = 2
)

const (
	IllegalMove int = -2
	LegalMove   int = 100
)

type Player struct {
	mark         string
	entity       Entity
	game         *Game
	epsilon      float64
	epsilonDecay float64
	alpha        float64 // Learning rate
	gamma        float64 // Decay rate
	Q            []Q
	rewards      int
	badMoves     int
}

type Q struct {
	state   string
	actions []Action
}

type Action struct {
	value  float64
	action Move
}

// Ask a movement to the player
func (player *Player) askforplay(moves []Move) {
	var move Move
	var kewai int
	var newPosition int
	valid := false

	for !valid {

		// Display the map
		printBasicBoard(player.game.board)

		// Display help
		if Help() {
			printBasicHelp()
		}

		fmt.Printf("It's %s turn to play. ", player.mark)

		fmt.Println("Enter Kewai to move:")
		fmt.Scan(&kewai)

		fmt.Println("Enter a new position for kewai:")
		fmt.Scan(&newPosition)

		move = Move{kewai, newPosition}

		// if not valid print an error and ask again for a move
		if IsValideMove(moves, move) {
			valid = true
		} else {
			fmt.Println(errors.New("try another move"))
		}
	}

	// Update the board
	player.game.UpdateBoard(move)
}

// Play a random legal move
func (player *Player) playRandomMove(moves []Move) {
	player.game.UpdateBoard(moves[rand.Intn(len(moves))])
}

// Make agent learn rules with SARSA algorithm
func (player *Player) learn(moves []Move) {

	valid := false
	var state string
	var newState string
	var action Move
	var newAction Move

	// Decrease epsilon
	player.decreaseEpsilon(player.epsilon > 0.1, player.epsilonDecay)

	// Determine actual state and action
	state = BoardToState(player.game.board)
	action = player.epsilonGreedy(state)

	// while the determined action is not valid
	for !valid {
		// If it's a valid move
		if IsValideMove(moves, action) {
			player.game.UpdateBoard(action)
			valid = true
		}

		// Determine next state and action
		newState = BoardToState(player.game.board)
		newAction := player.epsilonGreedy(newState)

		// If it's not a valid move increase epsilon and update Q
		if !valid {
			player.badMoves++
			player.decreaseEpsilon(player.epsilon < 0.9, float64(IllegalMove)*player.epsilonDecay)
			player.update(state, newState, action, newAction, IllegalMove)
		}

		state = newState
		action = newAction
	}

	// If it's a valid move decrease epsilon and update Q
	player.decreaseEpsilon(player.epsilon > 0.01, float64(LegalMove)*player.epsilonDecay)
	player.update(state, newState, action, newAction, LegalMove)
}

// Test if we are going to do exploitation or exploration
func (player *Player) epsilonGreedy(state string) Move {
	if rand.Float64() < player.epsilon {
		return Move{rand.Intn(8), rand.Intn(8)}
	} else {
		return player.greedyStep(state)
	}
}

// Returns the index corresponding to the maximum of the actions-state value
func (player *Player) greedyStep(state string) Move {
	i := FindState(player.Q, state)
	_, j := ArrayMaxIndex(player.Q[i].actions)
	return player.Q[i].actions[j].action
}

// SARSA algorithm
func (player *Player) update(state string, newState string, action Move, newAction Move, reward int) {
	// Find all index in Q
	stateIndex := FindState(player.Q, state)
	actionIndex := FindAction(player.Q[stateIndex].actions, action)
	newStateIndex := FindState(player.Q, newState)
	newActionIndex := FindAction(player.Q[newStateIndex].actions, newAction)

	// Update total rewards
	player.rewards += reward

	// Update Q
	player.Q[stateIndex].actions[actionIndex].value = player.Q[stateIndex].actions[actionIndex].value + player.alpha*(float64(reward)+player.gamma*player.Q[newStateIndex].actions[newActionIndex].value-player.Q[stateIndex].actions[actionIndex].value)
}

// Play a move
func (player *Player) Play(moves []Move) {
	switch player.entity {
	case HumanAgent:
		player.askforplay(moves)
	case RandomAgent:
		player.playRandomMove(moves)
	case SarsaAgent:
		player.learn(moves)
	default:
		player.playRandomMove(moves)
	}
}

// Decrease epsilon if condition is true
func (player *Player) decreaseEpsilon(condition bool, value float64) {
	if condition {
		player.epsilon -= value
	}
}
