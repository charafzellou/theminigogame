package main

import (
	"fmt"
	"time"
)

func runDemo() {
	paladin := assetClass{"Paladin", 300, 20, 10}
	archer := assetClass{"Archer", 245, 5, 25}
	ninja := assetClass{"Ninja", 285, 20, 20}

	attackPaladin := eventAttack{"Paladin Slash", "Paladin", 1, 40}
	listPaladin := make(map[int]eventAttack)
	listPaladin[0] = attackPaladin
	attackArcher := eventAttack{"Basic Arrow", "Archer", 1, 40}
	listArcher := make(map[int]eventAttack)
	listArcher[0] = attackArcher
	attackNinja := eventAttack{"Shuriken Throw", "Ninja", 0, 40}
	listNinja := make(map[int]eventAttack)
	listNinja[0] = attackNinja

	combo1 := eventCombo{"LionHeart's Duty", "Paladin", attackPaladin, attackPaladin, 5, 3}
	combo2 := eventCombo{"Hawkeye's Touch", "Archer", attackArcher, attackArcher, 5, 3}
	combo3 := eventCombo{"Shogunat's Honour", "Ninja", attackNinja, attackNinja, 5, 3}

	paladinJuan := assetPlayer{"Juanitus", paladin, 0, listPaladin, combo1, 3, attackPaladin}
	archerJuan := assetPlayer{"Juanito", archer, 0, listArcher, combo2, 3, attackArcher}
	ninjaJuan := assetPlayer{"Juan", ninja, 0, listNinja, combo3, 3, attackNinja}

	fmt.Println("Starting PvP demo : ")
	time.Sleep(3 * time.Second)
	hit(paladinJuan, listPaladin[0], &archerJuan)
	time.Sleep(3 * time.Second)
	hit(archerJuan, listArcher[0], &paladinJuan)
	time.Sleep(3 * time.Second)
	comboHit(paladinJuan, paladinJuan.SpecialAttack, &archerJuan)
	time.Sleep(3 * time.Second)
	comboHit(archerJuan, archerJuan.SpecialAttack, &paladinJuan)
	time.Sleep(3 * time.Second)
	hit(ninjaJuan, listNinja[0], &archerJuan)
	time.Sleep(3 * time.Second)
	hit(ninjaJuan, listNinja[0], &paladinJuan)
	time.Sleep(3 * time.Second)
	comboHit(ninjaJuan, ninjaJuan.SpecialAttack, &archerJuan)
	time.Sleep(8 * time.Second)

}
