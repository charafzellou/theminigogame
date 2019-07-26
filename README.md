# TheMiniGoGame

TheMiniGoGame is a.. well, everything is in the title 🙄

It is a school project to learn basic *`Go`* / *`GoLang`* skills.

The minigame is a turn-based RPG allowing you to control a character and **fight your friends** or play a **multiple-choice story** 📊💬.

Choose one of the **five available classes** ♞-🏹-🐱‍👤-🙏-💪, explore the **world** 🗺️, make the right choices *(or not?)*, brew **potions** ⚗️, learn **combo attacks** 🛡️, and defeat the **dangerous boss** 👹 waiting at the end of this beautiful pile of code 🎉!

## The game

```go
// work in progress
```

## Contributing
Pull requests are always welcome 🤓, though we have no desire to turn this school project into a full game.

For bug fixes or improvements, please open an issue and we will get back to you as soon as we can! 🔜

## Project BIG STEPS (in french sorry)
- [x] Créer deux fonctions : "getHit" et "hit" qui respectivement permettent de retirer des points de vie et d'en faire subir. Libre à vous de déterminer les paramètres.
- [x] Créer une fonction qui, en fonction du sort subi, permettent de redonner des points de vie (pas de joueurs pour le moment)
- [X] Créer les Structs correspondants (au minimum deux Structs : Joueur et Attaque), qui permettraient à deux joueurs de s'affronter en PVP.
- [X] La Struct devra permettre aux joueurs d'avoir des attaques prédéfinies ainsi que des points de vie de base.
- [X] Créer une fonction Combo (correspondant à la Struct Joueur) permettant d'enchainer plusieurs attaques
en même temps, et donc de faire subir plusieurs attaques à un adversaire.
- [x] Implémenter une manière permettant à des groupes de joueurs de s'affronter entre-eux (hint: utilisez les collections)
- [x] Chaque joueur doit pouvoir s'enregister lorsqu'il accède au programme grâce à un pseudonyme et un mot de passe. Mettre en place une fonction "register".
- [x] Une fois cela fait, demander le nom d'utilisateur et le mot de passe du joueur et lui proposer des joueurs ayant +- 20% de son niveau pour pouvoir jouer avec eux et / ou les affronter
- [ ] Créer une Struct "Monstre" qui permettra aux joueurs de faire du PVM (Player Vs Machine).
- [ ] Une fois dans le programme, un monstre peut attaquer un joueur si celui-ci, une fois connecté au jeu, reste inactif pendant plus de 20 secondes.
- [ ] BONUS: Développer une mécanique permettant à plusieurs joueurs d'affronter plusieurs monstres en même temps.

## License
[GNU General Public License v3.0](https://choosealicense.com/licenses/gpl-3.0/) 🥐