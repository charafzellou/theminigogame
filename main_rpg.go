package main

import "fmt"

func main() {
	paladin := fighterStruct{"juan", "paladin", 150, 10}
	ninja := fighterStruct{"juan", "ninja", 100, 35}

	fmt.Println(paladin)
	fmt.Println(ninja)

	hit( &ninja, paladin.strength)

	fmt.Println(paladin)
	fmt.Println(ninja)
}

func hit(target *fighterStruct, strength int){
	target.health -= strength
}

//func getHit()

type fighterStruct struct {
	name string
	class string
	health int
	strength int
}

