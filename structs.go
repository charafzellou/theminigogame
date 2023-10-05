package main

// Initializing assets and events Structures
type assetPlayer struct {
	Name          string
	Class         assetClass
	Inventory     uint
	AttackList    map[int]eventAttack
	SpecialAttack eventCombo
	ReloadTime    uint
	LastAttack    eventAttack
}
type assetClass struct {
	Name     string
	Health   int
	Healing  uint
	Strength uint
}
type assetMonster struct {
	Name   string
	Health int
	Damage int
}

/*
 * Name   : attack Name
 * Class  : Class that can use the attack
 * Effect : 0 => Healing 1 => Damage
 * Damage : amount of Damage without Class' Strength
 */
type eventAttack struct {
	Name   string
	Class  string
	Effect uint
	Damage uint
}
type eventCombo struct {
	Name        string
	Class       string
	AttackOne   eventAttack
	AttackTwo   eventAttack
	DamageBonus int
	ReloadTime  uint
}
type objectConsumable struct {
	name        string
	amountBonus uint
}

func initClasses() {
	classesMap["Paladin"] = assetClass{"Paladin", 300, 10, 15}
	classesMap["Archer"] = assetClass{"Archer", 245, 0, 25}
	classesMap["Ninja"] = assetClass{"Ninja", 285, 0, 20}
	classesMap["Cleric"] = assetClass{"Cleric", 210, 20, 5}
	classesMap["Berserk"] = assetClass{"Berserk", 290, 5, 25}
}
func initAttacks() {
	// Paladin Class
	attacksMap["Paladin Slash"] = eventAttack{"Paladin Slash", "Paladin", 1, 4}
	attacksMap["Long Sword Swing"] = eventAttack{"Long Sword Swing", "Paladin", 1, 3}
	attacksMap["Shield Block"] = eventAttack{"Shield Block", "Paladin", 1, 2}
	// Archer Class
	attacksMap["Basic Arrow"] = eventAttack{"Basic Arrow", "Archer", 1, 4}
	attacksMap["Bow Charge"] = eventAttack{"Bow Charge", "Archer", 1, 3}
	attacksMap["Quick Punch"] = eventAttack{"Quick Punch", "Archer", 1, 2}
	// Ninja Class
	attacksMap["Shuriken Throw"] = eventAttack{"Shuriken Throw", "Ninja", 1, 4}
	attacksMap["Hidden Dagger"] = eventAttack{"Hidden Dagger", "Ninja", 1, 4}
	attacksMap["Lightning Dash"] = eventAttack{"Lightning Dash", "Ninja", 1, 2}
	// Cleric Class
	attacksMap["Prayer to Gods"] = eventAttack{"Prayer to Gods", "Cleric", 0, 4}
	attacksMap["Sunday Preach"] = eventAttack{"Sunday Preach", "Cleric", 0, 3}
	attacksMap["Random Gibber"] = eventAttack{"Random Gibber", "Cleric", 1, 1}
	// Berserk Class
	attacksMap["Berserk Cut"] = eventAttack{"Berserk Cut", "Berserk", 1, 4}
	attacksMap["Call to Arms"] = eventAttack{"Call to Arms", "Berserk", 0, 3}
	attacksMap["Spartian Rage"] = eventAttack{"Spartian Rage", "Berserk", 1, 24}
}
func initCombos() {
	combosMap["Lionheart's Duty"] = eventCombo{"Lionheart's Duty", "Paladin", attacksMap["Shield Block"], attacksMap["Paladin Slash"], 5, 4}
	combosMap["Hawkeye's Touch"] = eventCombo{"Hawkeye's Touch", "Archer", attacksMap["Basic Arrow"], attacksMap["Basic Arrow"], 5, 4}
	combosMap["Shogunat's Honor"] = eventCombo{"Shogunat's Honor", "Ninja", attacksMap["Lightning Dash"], attacksMap["Hidden Dagger"], 5, 4}
	combosMap["Athena's Blessing"] = eventCombo{"Athena's Blessing", "Cleric", attacksMap["Prayer to Gods"], attacksMap["Random Gibber"], 15, 3}
	combosMap["Thor's Wrath"] = eventCombo{"Thor's Wrath", "Berserk", attacksMap["Berserk cut"], attacksMap["Call to Arms"], 0, 5}
}
func initPlayers() {
	getAccounts()
}
