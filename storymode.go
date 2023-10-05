package main

func storyMode() {
	player := createPlayer()

	showTextLn("-- THIS IS STORY MODE --", 2)
	showTextLn("We will accompany you in a short but very challenging story!", 2)
	showTextLn("We hope you are ready for the numerous battles awaiting you", 2)
	showText("Good luck...", 2)
	showTextLn("You will need it :)", 2)
	showSeparator()
	storyActOne(player)
	showSeparator()
	storyActTwo(player)
	showSeparator()
	storyActThree(player)
	showSeparator()
	showTextLn("-- YOU REACHED THE END OF STORY MODE --", 3)
	showTextLn("Bye, and thank you for playing !", 5)
}

func storyActOne(player assetPlayer) {
	golem := assetMonster{"Golem", 80, 5}
	showTextLn("-- ACT ONE --", 2)
	showTextLn(("You are travelling through a forest when you meet a " + golem.Name), 2)
	getHit(&player, golem.Damage+random.Intn(5))
	showTextLn("You manage to fight back, and run away", 2)
	showTextLn("You are bleeding, but you manage to hide under some bushes", 2)
}
func storyActTwo(player assetPlayer) {
	golem := assetMonster{"Golem", 80, 5}
	wraith := assetMonster{"Wraith", 60, 20}
	showTextLn("-- ACT TWO --", 2)
	showTextLn("The bleeding did not stop.. You are feeling worse!", 2)
	getHit(&player, golem.Damage+random.Intn(5))
	showTextLn(("It would have been easier... If not for the " + wraith.Name), 2)
	getHit(&player, wraith.Damage+random.Intn(5))
	showTextLn("Today is not your day buddy..", 2)
}
func storyActThree(player assetPlayer) {
	witcher := assetMonster{"Witcher", 120, 35}
	showTextLn("-- ACT THREE --", 2)
	showTextLn(("Just outside of the damned forest, you come across a village. You meet a " + witcher.Name), 2)
	showTextLn("He does not seem happy to see you...", 2)
	getHit(&player, witcher.Damage)
	getHit(&player, witcher.Damage+random.Intn(5))
	getHit(&player, witcher.Damage-random.Intn(5))
	showTextLn("Wow, a triple attack?", 2)
	showTextLn("You suck at this game..", 2)
	showTextLn("But it is not your fault... We had less time for the Story Mode :)", 2)
}
