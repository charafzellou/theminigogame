package main

// Importing necessary dependacies
import (
	"fmt"
	"math/rand"
)

// Initializing assets and events Structures
type assetPlayer struct {
	name		 string
	class		 assetClass
	inventory	 uint
	lastAttack	 eventAttack
}
type assetClass struct {
	name     	string
	health		uint
	healing 	uint
	strength	uint
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
type objectConsumable struct{
	name		string
	amountBonus	uint
}

// Initializing in-game assets
const (
)
func initClasses() {
}
func initAttacks() {
}
func initCombos() {
}
func initPlayers() {
}

// Initializing in-game functions
func calculateDamage(attacker assetPlayer, attack eventAttack) uint{
	return (uint)(attack.damage + (uint)(rand.Intn((int)(attacker.class.strength))))
}
func calculateHealing(healer assetPlayer) uint{
	return (uint)(1 + rand.Intn((int)(healer.class.healing)))
}
func hit(attacker assetPlayer, attack eventAttack, target *assetPlayer) {
	doneDamage := calculateDamage(attacker, attack)
	fmt.Println(target.name, "took", doneDamage, "damage from", attacker.name)
	target.class.health -= doneDamage
	fmt.Println(target.name, "has currently", target.class.health, "HP!")
}
func getHit(target *assetPlayer, amount uint){
	fmt.Println(target.name, "suffered", amount, "damage")
	target.class.health -= amount
	fmt.Println(target.name, "dropped to", target.class.health, "HP!")
}
func heal(healer assetPlayer, target *assetPlayer) {
	target.class.health += healer.class.strength
	fmt.Println(healer.name, "healed", target.name, ", who has", target.class.health, "HP!")
}
func getHealed(target *assetPlayer, amount uint){
	target.class.health += amount
	fmt.Println(target.name, "got healed and now has", target.class.health, "HP!")
}

// Initializing Main Storyline
func main() {
	paladin := assetClass{"Paladin", 300, 20, 10}
	archer := assetClass{"Archer", 245, 5, 25}
	ninja := assetClass{"Ninja", 285, 10, 20}

	attackPaladin := eventAttack{"attackPaladin", paladin, 4}
	attackArcher := eventAttack{"attackArcher", archer, 4}
	attackNinja := eventAttack{"attackNinja", ninja, 4}

	paladinJuan := assetPlayer{"Juanitus", paladin, 0, attackPaladin}
	archerJuan := assetPlayer{"Juanito", archer, 0, attackArcher}
	ninjaJuan := assetPlayer{"Juan", ninja, 0, attackNinja}

	hit(paladinJuan, attackPaladin, &archerJuan)
	hit(paladinJuan, attackPaladin, &archerJuan)
	hit(paladinJuan, attackPaladin, &archerJuan)
	hit(paladinJuan, attackPaladin, &archerJuan)
	hit(paladinJuan, attackPaladin, &archerJuan)
	hit(ninjaJuan, attackArcher, &paladinJuan)
	hit(archerJuan, attackNinja, &ninjaJuan)
}