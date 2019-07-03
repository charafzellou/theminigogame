package main

// Importing necessary dependacies
import (
	"fmt"
)

// Initializing assets and events Structures
type assetPlayer struct {
	name  string
	class assetClass
}
type assetClass struct {
	name     string
	health   uint
	strength uint
}
type eventAttack struct {
	name   string
	class  assetClass
	damage uint
}
type eventCombo struct {
	name        string
	class       assetClass
	attackOne   eventAttack
	attackTwo   eventAttack
	damageBonus uint
}

// Initializing in-game assets
func initClasses() {
}
func initAttacks() {
}
func initCombos() {
}
func initPlayers() {
}

// Initializing in-game functions
func hit(attacker assetPlayer, target assetPlayer) {
	fmt.Println(target.name, "had", target.class.health , "HP and took" , attacker.class.strength, "damage from", attacker.name)
	target.class.health -= attacker.class.strength
	fmt.Println(target.name, "has currently", target.class.health, "HP!")
}
func getHit(target assetPlayer, amount uint){
	target.class.health -= amount
}
func heal(healer assetPlayer, target assetPlayer) {
	target.class.health += healer.class.strength
}
func getHealed(target assetPlayer, amount uint){
	target.class.health += amount
}

// Initializing Main Storyline
func main() {
	paladin := assetClass{"Paladin", 100, 20}
	archer := assetClass{"Archer", 75, 25}
	ninja := assetClass{"Ninja", 80, 35}

	paladinJuan := assetPlayer{"Juanitus", paladin}
	archerJuan := assetPlayer{"Juanito", archer}
	ninjaJuan := assetPlayer{"Juan", ninja}

	/*fmt.Println(paladinJuan.class.health)
	fmt.Println(archerJuan.class.health)
	fmt.Println(ninjaJuan.class.health)*/

	hit(paladinJuan, archerJuan)
	hit(ninjaJuan, paladinJuan)
	hit(archerJuan, ninjaJuan)

	/*fmt.Println(paladinJuan.class.health)
	fmt.Println(archerJuan.class.health)
	fmt.Println(ninjaJuan.class.health)*/
}