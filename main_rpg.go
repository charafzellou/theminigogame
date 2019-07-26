package main

/*
	Files contents:
		structs.go		-->		contains all the game's structs and classes
		menus.go		-->		contains all the menus and display functions
		mechanics.go	-->		contains all the game mechanics functions
		demo.go			-->		contains the demo function to try out all the game mechanics
		README.md		-->		a nice and cool README file to understand the project
*/

// Importing necessary dependencies
import (
	"math/rand"
	"time"
)

// Setting up the needed variables
var (
	classesMap  = make(map[string]assetClass)
	attacksMap  = make(map[string]eventAttack)
	combosMap   = make(map[string]eventCombo)
	pvpBlueTeam = make(map[uint]*assetPlayer)
	pvpRedTeam  = make(map[uint]*assetPlayer)
	randSeed    = rand.NewSource(time.Now().UnixNano())
	random      = rand.New(randSeed)
)

// Initializing all the important assets
func init() {
	initClasses()
	initAttacks()
	initCombos()
	initPlayers()
}

// start1vs1 : This function launches 1VS1 Combat
func start1vs1() {
	setPvpParams(1, 1)
	startPvpFight(1, 1)
}

// start2vs2 : This function launches 2VS2 Combat
func start2vs2() {
	setPvpParams(2, 2)
	startPvpFight(2, 2)
}

// Launching Main Homepage
func main() {
	homePage()
}
