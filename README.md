# TheMiniGoGame

TheMiniGoGame is a... well, everything is in the title 🙄

It is a school project to learn basic *`Go`* / *`GoLang`* skills based on The Witcher lore.

The minigame is a turn-based RPG allowing you to control a character and **fight your friends** or play a **multiple-choice story** 📊💬.

Choose one of the **five available classes** ♞-🏹-🐱‍👤-🙏-💪, explore the **world** 🗺️, make the right choices *(or not?)*, brew **potions** ⚗️, learn **combo attacks** 🛡️, and defeat the **dangerous boss** 👹 waiting at the end of this beautiful pile of code 🎉!

## The game

```go
// work in progress
```

## Project Roadmap
- [x] Create two functions: "getHit" and "hit" which respectively remove life points from a user and remove them on a monster. Parameters are to be defined freely.
- [x] Create a function which, depending on the fate suffered, allows you to restore life points.
- [X] Create the corresponding Structs (at least two Structs: Player and Attack), which would allow two players to compete in PVP.
- [X] The Struct must allow players to have predefined attacks as well as basic life points.
- [X] Create a Combo function (corresponding to the Player Struct) allowing you to chain several attacks together at the same time, and therefore to subject an opponent to several attacks.
- [x] Implement a way to allow groups of players to compete against each other (hint: use Collections)
- [x] Each player must be able to register when accessing the program using a pseudonym and a password. Set up a “register” function.
- [x] Once this is done, ask for the player's username and password and offer them players with +- 20% of their level to be able to play with them and/or face them.
- [X] Create a "Monster" Struct which will allow players to do PVM (Player Vs Machine).
- [ ] Once connected to the program, a monster can attack a player if the player remains inactive for more than 20 seconds.
- [ ] BONUS: Develop a mechanic allowing several players to face several monsters at the same time.

## Contributing
Pull requests are always welcome 🤓, though we have no desire to turn this school project into a full game.

For bug fixes or improvements, please open an issue and we will get back to you as soon as we can! 🔜

## License
[GNU Affero General Public License v3.0](https://choosealicense.com/licenses/agpl-3.0/) 🥐