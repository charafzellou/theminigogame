package main

import (
	"fmt"
)

type assetFighter struct {
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

func hit(target *assetFighter, strength uint) {
	target.class.health -= strength
}

func getHit() {

}

func main() {
	paladin := assetClass{"Paladin", 100, 20}
	ninja := assetClass{"Ninja", 80, 35}

	fighterJuan := assetFighter{"Juan", paladin}

	fmt.Println(paladin)
	fmt.Println(ninja)
	fmt.Println(fighterJuan)

	hit(&fighterJuan, fighterJuan.class.strength)

	fmt.Println(fighterJuan)
}
