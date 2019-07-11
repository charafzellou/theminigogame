package main

// Importing necessary dependencies
import (
	"fmt"
	"math/rand"
	"time"
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
func comboHit(attacker assetPlayer, combo eventCombo, target *assetPlayer){
	hit( attacker, combo.attackOne, target)
	fmt.Println("OMG ! it's", combo.name ,"combo attack !")
	hit( attacker, combo.attackTwo, target)
	fmt.Println("Bonus power :")
	if combo.attackOne.effect + combo.attackTwo.effect == 0 {
		getHealed(target, combo.damageBonus)
	} else {
		getHit(target, combo.damageBonus)
	}
}
func createPlayer() assetPlayer{
	name := ""
	fmt.Print("Character Name: ")
	_, _ = fmt.Scanf("%s\n", &name)

	fmt.Println("Select your class :")
	classesList := getClasses(classesMap)
	displayList(classesList)
	class := choiceFromList(classesList)

	fmt.Println("Select your primary attack")
	attacksList := getAttacks(attacksMap, classesList[class])
	displayList(attacksList)
	attack := choiceFromList(attacksList)
	attackChosen := make(map[int]eventAttack)
	attackChosen[1] = attacksMap[attacksList[attack]]

	fmt.Println("Select your special attack")
	combosList := getCombo(combosMap, classesList[class])
	displayList(combosList)
	combo := choiceFromList(combosList)

	return assetPlayer{
		name,
		classesMap[classesList[class]],
		0,
		attackChosen,
		combosMap[combosList[combo]],
		combosMap[combosList[combo]].reloadTime,
		attacksMap[attacksList[attack]],
	}
}
func displayList(list map[int]string){
	for idx, elemName := range list {
		fmt.Println(idx, "", elemName)
	}
}
func getClasses(list map[string]assetClass) map[int]string {
	result := make(map[int]string)
	index := 1
	for className := range list {
		result[index] = className
		index++
	}
	return result
}
func getAttacks(list map[string]eventAttack, classeName string) map[int]string {
	result := make(map[int]string)
	index := 1
	for attackName, attack := range list {
		if attack.class == classeName {
			result[index] = attackName
			index++
		}
	}
	return result
}
func getCombo(list map[string]eventCombo, classeName string) map[int]string {
	result := make(map[int]string)
	index := 1
	for comboName, combo := range list {
		if combo.class == classeName {
			result[index] = comboName
			index++
		}
	}
	return result
}
func choiceFromList(list map[int]string) (choice int){
	for {
		isInt, _ := fmt.Scan(&choice)
		if isInt == 1 {
			if list[choice] == "" {
				fmt.Println("Please enter valid number")
			} else {
				break
			}
		}
	}
	return
}

// Initializing Menu functions
func printMenuUpperPart(){
	fmt.Println("     ___     ___     ___     ___     ___     ___     ___     ___")
	fmt.Println(" ___/   \\___/   \\___/   \\___/   \\___/   \\___/   \\___/   \\___/   \\___")
	fmt.Println("/   \\___/   \\___/   \\___/   \\___/   \\___/   \\___/   \\___/   \\___/   \\")
	fmt.Println("\\___/   \\___/   \\___/   \\___/   \\___/   \\___/   \\___/   \\___/   \\___/")
	fmt.Println("/   \\___/                                                   \\___/   \\")
	fmt.Println("\\___/                                                           \\___/")
}
func printMenuBottomPart(){
	fmt.Println("/   \\___                                                     ___/   \\")
	fmt.Println("\\___/   \\___     ___     ___     ___     ___     ___     ___/   \\___/")
	fmt.Println("/   \\___/   \\___/   \\___/   \\___/   \\___/   \\___/   \\___/   \\___/   \\")
	fmt.Println("\\___/   \\___/   \\___/   \\___/   \\___/   \\___/   \\___/   \\___/   \\___/")
	fmt.Println("    \\___/   \\___/   \\___/   \\___/   \\___/   \\___/   \\___/   \\___/")
}
func mainMenu(){
	for {
		choice := '0'
		exit := 0
		printMenuUpperPart()
		fmt.Println("/   \\                    The Mini Go Game                       /   \\")
		fmt.Println("\\___/                                                           \\___/")
		fmt.Println("/   \\                                                           /   \\")
		fmt.Println("\\___/                 1.    Story Mode                          \\___/")
		fmt.Println("/   \\                                                           /   \\")
		fmt.Println("\\___/                 2.       PvP                              \\___/")
		fmt.Println("/   \\                                                           /   \\")
		fmt.Println("\\___/                 3.     run demo                           \\___/")
		fmt.Println("/   \\                                                           /   \\")
		fmt.Println("\\___/                 q.     Exit Game                          \\___/")
		printMenuBottomPart()
		_, _ = fmt.Scanf("%c\n", &choice)
		switch choice {
		case '1':
			fmt.Println("Work in progress")
			// TODO : insert function that manage story mode
			break
		case '2':
			pvpMenu()
			break
		case '3':
			runDemo()
			break
		case 'q':
			exit = 1
			break
		default:
			fmt.Println("Incorrect input, try again")
		}
		if exit == 1 {
			fmt.Println("Bye Bye")
			time.Sleep(2 * time.Second)
			break
		}
		time.Sleep(3 * time.Second)
	}
}
func pvpMenu(){
	for {
		choice := '0'
		exit := 0
		printMenuUpperPart()
		fmt.Println("/   \\                    The Mini Go Game                       /   \\")
		fmt.Println("\\___/                        PvP mode                           \\___/")
		fmt.Println("/   \\                                                           /   \\")
		fmt.Println("\\___/                 1.      1 vs 1                            \\___/")
		fmt.Println("/   \\                                                           /   \\")
		fmt.Println("\\___/                 2.      2 vs 2                            \\___/")
		fmt.Println("/   \\                                                           /   \\")
		fmt.Println("\\___/                 3.      go back                           \\___/")
		printMenuBottomPart()
		_, _ = fmt.Scanf("%c\n", &choice)
		switch choice {
		case '1':
			fmt.Println("Work in progress")
			// TODO : insert 1 vs 1 mode
			break
		case '2':
			fmt.Println("Work in progress")
			// TODO : inset 2 vs 2 mode
			break
		case '3':
			exit = 1
			break
		default:
			fmt.Println("Incorrect input, try again")
		}
		if exit == 1 {
			time.Sleep(1 * time.Second)
			break
		}
		time.Sleep(3 * time.Second)
	}
}
func runDemo(){
	paladin := assetClass{"Paladin", 300, 20, 10}
	archer := assetClass{"Archer", 245, 5, 25}
	ninja := assetClass{"Ninja", 285, 20, 20}

	attackPaladin := eventAttack{"attackPaladin", "Paladin", 1,40}
	listPaladin := make(map[int]eventAttack)
	listPaladin[0] = attackPaladin
	attackArcher := eventAttack{"attackArcher", "Archer", 1,40}
	listArcher := make(map[int]eventAttack)
	listArcher[0] = attackArcher
	attackNinja := eventAttack{"medicinalHerb", "Ninja", 0,40}
	listNinja := make(map[int]eventAttack)
	listNinja[0] = attackNinja

	combo1 := eventCombo{ "doublePaladin", "Paladin", attackPaladin, attackPaladin, 5, 3}
	combo2 := eventCombo{ "doubleArcher", "Archer", attackArcher, attackArcher, 5, 3}
	combo3 := eventCombo{ "doubleNinja", "Ninja", attackNinja, attackNinja, 5, 3}

	paladinJuan := assetPlayer{"Juanitus", paladin, 0, listPaladin, combo1, 3, attackPaladin}
	archerJuan := assetPlayer{"Juanito", archer, 0, listArcher, combo2, 3, attackArcher}
	ninjaJuan := assetPlayer{"Juan", ninja, 0, listNinja, combo3, 3, attackNinja}

	fmt.Println("starting PvP demo : ")
	time.Sleep(3 * time.Second)
	hit(paladinJuan, listPaladin[0], &archerJuan)
	time.Sleep(3 * time.Second)
	hit(archerJuan, listArcher[0], &paladinJuan)
	time.Sleep(3 * time.Second)
	comboHit( paladinJuan, paladinJuan.specialAttack, &archerJuan)
	time.Sleep(3 * time.Second)
	comboHit( archerJuan, archerJuan.specialAttack, &paladinJuan)
	time.Sleep(3 * time.Second)
	hit(ninjaJuan, listNinja[0], &archerJuan)
	time.Sleep(3 * time.Second)
	hit(ninjaJuan, listNinja[0], &paladinJuan)
	time.Sleep(3 * time.Second)
	comboHit(ninjaJuan, ninjaJuan.specialAttack, &archerJuan)
	time.Sleep(8 * time.Second)

}

// Initializing Main Storyline
func main() {
	mainMenu()
	/*choice := 2
	fmt.Println("How many players ?")
	_, _ = fmt.Scan(&choice)
	for idx := 1; idx <= choice; idx++{
		PvpPlayers[uint(idx)] = createPlayer()
	}
	fmt.Println(PvpPlayers)*/
}