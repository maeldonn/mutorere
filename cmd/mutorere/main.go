package main

import (
	"math/rand"
	"mutorere/internal/mutorere"
	"time"
)

// Entry point of our application
func main() {

	// Generate seed for package random
	rand.Seed(time.Now().UnixNano())

	// Launch interactive menu
	mutorere.ModeSelection()
}
