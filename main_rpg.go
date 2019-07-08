package main

// Importing necessary dependencies
import (
	"fmt"
	"math/rand"
)

// Initializing assets and events Structures
type assetPlayer struct {
	name		    string
	class		    assetClass
	inventory	    uint
	attackList      map[int]eventAttack
	specialAttack   eventCombo
	reloadTime      uint
	lastAttack	    eventAttack
}
type assetClass struct {
	name     	string
	health		uint
	healing 	uint
	strength	uint
}
/*
 * name   : attack name
 * class  : class that can use the attack
 * effect : 0 => healing 1 => damage
 * damage : amount of damage without class' strength
 */
type eventAttack struct {
	name   string
	class  string
	effect uint
	damage uint
}
type eventCombo struct {
	name        string
	class       string
	attackOne   eventAttack
	attackTwo   eventAttack
	damageBonus uint
	reloadTime  uint
}
type objectConsumable struct{
	name		string
	amountBonus	uint
}

// Initializing in-game assets
const (
)
var (
	classesMap = make(map[string]assetClass)
	attacksMap = make(map[string]eventAttack)
	combosMap  = make(map[string]eventCombo)
)
func init(){
	initClasses()
	initAttacks()
	initCombos()
}
func initClasses() {
	classesMap["Paladin"] = assetClass{"Paladin", 300, 10, 15}
	classesMap["Archer"] = assetClass{"Archer", 245, 0, 25}
	classesMap["Ninja"] = assetClass{"Ninja", 285, 0, 20}
	classesMap["Cleric"] = assetClass{"Cleric", 210, 20, 5}
	classesMap["Berserk"] = assetClass{"Berserk", 290, 5, 25}
}
func initAttacks() {
	attacksMap["Paladin slash"] = eventAttack{"Paladin slash", "Paladin", 1,4}
	attacksMap["Basic arrow"] = eventAttack{"Basic arrow", "Archer", 1, 4}
	attacksMap["shuriken throw"] = eventAttack{"Shuriken throw", "Ninja", 1, 4}
	attacksMap["Prier"] = eventAttack{"Prier", "Cleric", 0, 4}
	attacksMap["Berserk cut"] = eventAttack{"Berserk cut", "Berserk", 1, 4}
}
func initCombos() {
	combosMap["PaladinCombo"] = eventCombo{"DoubleSlash", "Paladin", attacksMap["Paladin slash"], attacksMap["Paladin slash"], 0, 3}
	combosMap["ArcherCombo"] = eventCombo{"DoubleArrow", "Archer", attacksMap["Basic arrow"], attacksMap["Basic arrow"], 0, 3}
	combosMap["NinjaCombo"] = eventCombo{"DoubleShuriken", "Ninja", attacksMap["Shuriken throw"], attacksMap["Shuriken throw"], 0, 3}
	combosMap["ClericCombo"] = eventCombo{"DoublePrier", "Cleric", attacksMap["Prier"], attacksMap["Prier"], 10, 3}
	combosMap["BerserkCombo"] = eventCombo{"DoubleCut", "Berserk", attacksMap["Berserk cut"], attacksMap["Berserk cut"], 0, 3}
}
func initPlayers() {
}

// Initializing in-game functions
func calculateDamage(attacker assetPlayer, attack eventAttack) uint{
	return (uint)(attack.damage + (uint)(rand.Intn((int)(attacker.class.strength))))
}
func calculateHealing(healer assetPlayer, heal eventAttack) uint{
	return (uint)(heal.damage + (uint)(rand.Intn((int)(healer.class.healing))))
}
func hit(attacker assetPlayer, attack eventAttack, target *assetPlayer) {

	fmt.Println(attacker.name, "uses", attack.name, "on", target.name)
	if attack.effect == 1 {
		doneDamage := calculateDamage(attacker, attack)
		getHit(target, doneDamage)
	} else {
		doneHeal := calculateHealing(attacker, attack)
		getHealed( target, doneHeal)
	}
	fmt.Println(target.name, "has currently", target.class.health, "HP!")
}
func getHit(target *assetPlayer, amount uint){
	target.class.health -= amount
	fmt.Println(target.name, "suffered", amount, "damage !")
}
func getHealed(target *assetPlayer, amount uint){
	target.class.health += amount
	fmt.Println(target.name, "recovered", amount, "HP !")
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