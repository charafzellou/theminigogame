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
	Attack eventAttack
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
	attacksMap["Paladin slash"] = eventAttack{"Paladin slash", "Paladin", 1, 4}
	attacksMap["Basic arrow"] = eventAttack{"Basic arrow", "Archer", 1, 4}
	attacksMap["Shuriken throw"] = eventAttack{"Shuriken throw", "Ninja", 1, 4}
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
	getAccounts()
}
