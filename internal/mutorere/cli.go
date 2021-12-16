package mutorere

import (
	"errors"
	"fmt"
	"os"
	"strings"

	"github.com/manifoldco/promptui"
)

// Interface to select the application mode to launch
func ModeSelection() {

	Clear()

	prompt := promptui.Select{
		Label: "Select mode",
		Items: []string{"Train AI", "Challenge AI", "Challenge HumanAgent", "Rules", "Quit"},
	}

	_, mode, err := prompt.Run()

	if err != nil {
		fmt.Printf("Prompt failed %v\n", err)
		os.Exit(1)
	}

	switch mode {
	case "Train AI":
		trainArtificialIntelligence()
	case "Challenge AI":
		challengeArtificialIntelligence()
	case "Challenge HumanAgent":
		challengeHuman()
	case "Rules":
		showRules()
	case "Quit":
		quit()
	default:
		os.Exit(1)
	}
}

// Mode to train the sarsa agent and plot statistics
func trainArtificialIntelligence() {

	validate := func(input string) error {
		if len(input) < 2 {
			return errors.New("you should specify a larger number for best results")
		}
		if ConvertStringToInteger(input) == -1 {
			return errors.New("you should specify a number")
		}
		return nil
	}

	prompt := promptui.Prompt{
		Label:    "Number of games",
		Validate: validate,
		Default:  "150",
	}

	result, err := prompt.Run()

	if err != nil {
		fmt.Printf("Prompt failed %v\n", err)
		return
	}

	// Create the game
	var game Game
	game.CreateNewGame(AiVersusRandom)

	total := ConvertStringToInteger(result)
	var winrate []int
	var epsilons []float64
	var turns []int
	var badMoves []int
	var rewards []int

	// Launch "total" games
	for i := 0; i < total; i++ {
		fmt.Println("Game ", i+1, " / ", total)
		game.Start()
		winrate = append(winrate, game.GetWinner())
		turns = append(turns, game.turn)
		epsilons = append(epsilons, game.firstPlayer.epsilon)
		badMoves = append(badMoves, game.firstPlayer.badMoves)
		rewards = append(rewards, game.firstPlayer.rewards)
	}

	PlotEpsilonEvolution(epsilons)
	PlotRewards(rewards)
	PlotBadMoves(badMoves)
}

// Mode to play against a trained artificial intelligence
func challengeArtificialIntelligence() {

	// Create the training
	var training Game
	training.CreateNewGame(AiVersusRandom)

	// Train the algorithm in 150 games
	for i := 0; i < 150; i++ {
		fmt.Println("Training ", i+1, " / ", 150)
		training.Start()
	}

	// Play against human
	var game Game
	game.CreateNewGame(HumanVersusAi)

	// Setting to AI what first AI have learned
	game.firstPlayer.epsilonDecay = training.firstPlayer.epsilonDecay
	game.firstPlayer.alpha = training.firstPlayer.alpha
	game.firstPlayer.gamma = training.firstPlayer.gamma
	game.firstPlayer.Q = training.firstPlayer.Q

	// Force the sarsa agent to be in exploitation mode
	game.firstPlayer.epsilon = 0.05

	// Play the game
	game.Start()
}

// Mode to play against another Human
func challengeHuman() {
	var game Game
	game.CreateNewGame(HumanVersusHuman)
	game.Start()
}

// Display the rules of Mu Torere
func showRules() {

	fmt.Println("RULES")
	fmt.Println("")
	fmt.Println("Le Mu torere est un jeu traditionnel des Maoris de la côte est de l'Ile du Nord.")
	fmt.Println("")
	fmt.Println("C'est un jeu opposant deux joueurs, représentés par des pions noirs et blancs autour d'un plateau")
	fmt.Println("en forme d'étoile à 8 branches. Le but est de bloquer les pions de son adversaire.")
	fmt.Println("Chaque joueur déplace alternativement l'un de ses pions vers une intersection vide adjacente ;")
	fmt.Println("il n'est possible de déplacer au centre qu'un pion qui est adjacent à un pion de l'adversaire.")
	fmt.Println("Le vainqueur est celui qui prive son adversaire de coups légaux.")

	quit()
}

// Interface to get help with movements
func Help() bool {
	prompt := promptui.Prompt{
		Label:     "Show help",
		IsConfirm: true,
	}

	confirm, _ := prompt.Run()

	switch strings.ToLower(confirm) {
	case "y", "yes":
		return true
	case "n", "no", "":
		return false
	default:
		return false
	}
}

// Quit the application
func quit() {

	prompt := promptui.Prompt{
		Label:     "Exit program",
		IsConfirm: true,
	}

	confirm, _ := prompt.Run()

	switch strings.ToLower(confirm) {
	case "y", "yes":
		os.Exit(0)
	case "n", "no", "":
		ModeSelection()
	default:
		quit()
	}
}
