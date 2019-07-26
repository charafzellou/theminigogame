package main

import "fmt"

// Initializing in-game functions

// Calculations functions
func calculateDamage(attacker assetPlayer, attack eventAttack) int {
	return int(attack.Damage) + (random.Intn((int)(attacker.Class.Strength)))
}
func calculateHealing(healer assetPlayer, heal eventAttack) int {
	return int(heal.Damage) + (random.Intn((int)(healer.Class.Healing)))
}

// Events functions
func hit(attacker assetPlayer, attack eventAttack, target *assetPlayer) {

	fmt.Println(attacker.Name, "uses", attack.Name, "on", target.Name)
	if attack.Effect == 1 {
		doneDamage := calculateDamage(attacker, attack)
		getHit(target, doneDamage)
	} else {
		doneHeal := calculateHealing(attacker, attack)
		getHealed(target, doneHeal)
	}
	fmt.Println(target.Name, "has currently", target.Class.Health, "HP!")
}
func getHit(target *assetPlayer, amount int) {
	target.Class.Health -= amount
	fmt.Println(target.Name, "suffered", amount, "damage !")
}
func getHealed(target *assetPlayer, amount int) {
	target.Class.Health += amount
	fmt.Println(target.Name, "recovered", amount, "HP !")
}
func comboHit(attacker assetPlayer, combo eventCombo, target *assetPlayer) {
	hit(attacker, combo.AttackOne, target)
	fmt.Println("OMG !", attacker.Name, "called in", combo.Name, "!")
	hit(attacker, combo.AttackTwo, target)
	fmt.Println("Bonus power :")
	if combo.AttackOne.Effect+combo.AttackTwo.Effect == 0 {
		getHealed(target, combo.DamageBonus)
	} else {
		getHit(target, combo.DamageBonus)
	}
}

// Asset functions
func createPlayer() assetPlayer {
	name := ""
	fmt.Print("Character Name: ")
	_, _ = fmt.Scanf("%s\n", &name)

	fmt.Println("Select your Class :")
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
		combosMap[combosList[combo]].ReloadTime,
		attacksMap[attacksList[attack]],
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
		if attack.Class == classeName {
			result[index] = attackName
			index++
		}
	}
	return result
}
func getPlayerAttack(list map[int]eventAttack) map[int]string {
	result := make(map[int]string)
	for idx, attack := range list {
		result[idx] = attack.Name
	}
	return result
}
func getCombo(list map[string]eventCombo, classeName string) map[int]string {
	result := make(map[int]string)
	index := 1
	for comboName, combo := range list {
		if combo.Class == classeName {
			result[index] = comboName
			index++
		}
	}
	return result
}
func getTeamsPlayers(list map[uint]*assetPlayer) map[int]string {
	result := make(map[int]string)
	for idx, player := range list {
		result[int(idx)] = player.Name + " (" + player.Class.Name + ")"
	}
	return result
}
func choiceFromList(list map[int]string) (choice int) {
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

func displayList(list map[int]string) {
	for idx, elemName := range list {
		fmt.Println(idx, "", elemName)
	}
}
func setPvpParams(redTeamSize int, blueTeamSize int) {
	fmt.Println("RED TEAM :")
	createTeam(redTeamSize, pvpRedTeam)
	fmt.Println("BLUE TEAM :")
	createTeam(blueTeamSize, pvpBlueTeam)
}
func createTeam(teamSize int, teamMap map[uint]*assetPlayer) {
	for idx := 1; idx <= teamSize; idx++ {
		fmt.Println("Player", idx, "selection")
		player := createPlayer()
		teamMap[uint(idx)] = &player
	}
}
func startPvpFight(redTeamSize int, blueTeamSize int) {
	coinToss := random.Intn(2)
	if coinToss == 0 {
		fmt.Println("RedTeam Starts !")
	} else {
		fmt.Println("BlueTeam Starts !")
	}
	winner := -1
	for winner == -1 {
		switch coinToss {
		case 0:
			if teamTurn(redTeamSize, pvpRedTeam, pvpBlueTeam) == false {
				winner = 1
			}
			coinToss = 1
			break
		case 1:
			if teamTurn(blueTeamSize, pvpBlueTeam, pvpRedTeam) == false {
				winner = 0
			}
			coinToss = 0
			break
		default:
			panic("Error with random function")
		}
	}
	displayWinnerScreen(winner)
}
func teamTurn(teamSize int, teamPlaying map[uint]*assetPlayer, enemyTeam map[uint]*assetPlayer) bool {
	deadCounter := 0
	for playerTurn := 1; playerTurn <= teamSize; playerTurn++ {
		if _, ok := teamPlaying[uint(playerTurn)]; ok {
			action, attack := playerSelectAction(teamPlaying[uint(playerTurn)])
			if action == 1 {
				if teamPlaying[uint(playerTurn)].AttackList[attack].Effect == 0 {
					target := playerSelectTarget(teamPlaying)
					hit(*teamPlaying[uint(playerTurn)], teamPlaying[uint(playerTurn)].AttackList[attack], teamPlaying[uint(target)])
				} else {
					target := playerSelectTarget(enemyTeam)
					hit(*teamPlaying[uint(playerTurn)], teamPlaying[uint(playerTurn)].AttackList[attack], enemyTeam[uint(target)])
					if enemyTeam[uint(target)].Class.Health <= 0 {
						setPlayerAsDead(enemyTeam, uint(target))
					}
				}
			} else {
				if teamPlaying[uint(playerTurn)].SpecialAttack.AttackOne.Effect == 0 {
					target := playerSelectTarget(teamPlaying)
					comboHit(*teamPlaying[uint(playerTurn)], teamPlaying[uint(playerTurn)].SpecialAttack, teamPlaying[uint(target)])
				} else {
					target := playerSelectTarget(enemyTeam)
					comboHit(*teamPlaying[uint(playerTurn)], teamPlaying[uint(playerTurn)].SpecialAttack, enemyTeam[uint(target)])
					if enemyTeam[uint(target)].Class.Health <= 0 {
						setPlayerAsDead(enemyTeam, uint(target))
					}
				}
			}
			if teamPlaying[uint(playerTurn)].ReloadTime > 0 {
				teamPlaying[uint(playerTurn)].ReloadTime--
			} else if action == 2 {
				teamPlaying[uint(playerTurn)].ReloadTime = teamPlaying[uint(playerTurn)].SpecialAttack.ReloadTime
			}
		} else {
			deadCounter++
		}
	}
	if deadCounter == teamSize {
		return false
	}
	return true
}
func playerSelectAction(player *assetPlayer) (action int, attack int) {
	fmt.Println("it's", player.Name+"'s", "turn.")
	fmt.Println("Actions : ")
	actionChoice := make(map[int]string)
	actionChoice[1] = "Attack"
	actionChoice[2] = "SpecialAttack"
	exit := 0
	for exit != 1 {
		displayList(actionChoice)
		action = choiceFromList(actionChoice)
		switch action {
		case 1:
			playerAttackList := getPlayerAttack(player.AttackList)
			displayList(playerAttackList)
			attack = choiceFromList(playerAttackList)
			exit = 1
			break
		case 2:
			if player.ReloadTime != 0 {
				fmt.Println(player.ReloadTime, "turn(s) until usable.")
				fmt.Println("Please select an other action.")
			} else {
				attack = -1
				exit = 1
			}
			break
		default:
			panic("Unknown error !")
		}
	}
	return
}
func playerSelectTarget(team map[uint]*assetPlayer) int {
	fmt.Println("Use on :")
	playersList := getTeamsPlayers(team)
	displayList(playersList)
	target := choiceFromList(playersList)
	return target
}
func setPlayerAsDead(playerTeam map[uint]*assetPlayer, playerID uint) {
	fmt.Println("oh...", playerTeam[playerID].Name, "is dead... †")
	delete(playerTeam, playerID)
}
