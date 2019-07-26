package main

import (
	"fmt"
	"time"
)

// Initializing Menu functions
func mainMenu() {
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
		fmt.Println("\\___/                 q.      log out                           \\___/")
		printMenuBottomPart()
		_, _ = fmt.Scanf("%c\n", &choice)
		switch choice {
		case '1':
			storyMode()
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
			loggedAccount = -1
			fmt.Println("logging out, bye bye !")
			time.Sleep(2 * time.Second)
			break
		}
		time.Sleep(3 * time.Second)
	}
}
func pvpMenu() {
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
			start1vs1()
			break
		case '2':
			start2vs2()
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

func displayWinnerScreen(winner int) {
	printMenuUpperPart()
	fmt.Println("/   \\                                                           /   \\")
	fmt.Println("\\___/                                                           \\___/")
	fmt.Println("/   \\                                                           /   \\")
	if winner == 0 {
		fmt.Println("\\___/                     RED TEAM WINS                         \\___/")
	} else {
		fmt.Println("\\___/                    BLUE TEAM WINS                         \\___/")
	}
	fmt.Println("/   \\                                                           /   \\")
	fmt.Println("\\___/                                                           \\___/")
	fmt.Println("/   \\                                                           /   \\")
	fmt.Println("\\___/                                                           \\___/")
	printMenuBottomPart()
}
func printMenuUpperPart() {
	fmt.Println("     ___     ___     ___     ___     ___     ___     ___     ___")
	fmt.Println(" ___/   \\___/   \\___/   \\___/   \\___/   \\___/   \\___/   \\___/   \\___")
	fmt.Println("/   \\___/   \\___/   \\___/   \\___/   \\___/   \\___/   \\___/   \\___/   \\")
	fmt.Println("\\___/   \\___/   \\___/   \\___/   \\___/   \\___/   \\___/   \\___/   \\___/")
	fmt.Println("/   \\___/                                                   \\___/   \\")
	fmt.Println("\\___/                                                           \\___/")
}
func printMenuBottomPart() {
	fmt.Println("/   \\___                                                     ___/   \\")
	fmt.Println("\\___/   \\___     ___     ___     ___     ___     ___     ___/   \\___/")
	fmt.Println("/   \\___/   \\___/   \\___/   \\___/   \\___/   \\___/   \\___/   \\___/   \\")
	fmt.Println("\\___/   \\___/   \\___/   \\___/   \\___/   \\___/   \\___/   \\___/   \\___/")
	fmt.Println("    \\___/   \\___/   \\___/   \\___/   \\___/   \\___/   \\___/   \\___/")
}
func showText(text string, dur int) {
	fmt.Print(text)
	time.Sleep(time.Duration(dur) * (time.Second))
}
func showTextLn(text string, dur int) {
	fmt.Println(text)
	time.Sleep(time.Duration(dur) * (time.Second))
}
func showSeparator() {
	fmt.Println("_______________________________________________________________________________")
}
