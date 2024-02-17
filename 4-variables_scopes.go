package main

import "log"

// Déclaration d'une variable globale "s"
// Cette variable est accessible dans toutes les fonctions du package
var s = "seven"

func main() {

	// Déclaration d'une variable locale "s2" dans le scope de la fonction "main"
	var s2 = "six"

	log.Println("s is", s)   // variable globale
	log.Println("s2 is", s2) // variable locale

	// La fonction "saySomething2" est appelée avec "xxx" comme argument, et elle retourne deux chaînes ("xxx" et "World")
	log.Println(saySomething2("xxxx"))
}

// La fonction "saySomething2" prend une chaîne "s3" en argument et retourne 2 chaînes
func saySomething2(s3 string) (string, string) {

	// Affichage de la variable globale "s" dans le scope de la fonction "saySomething2"
	log.Println("s from the saySomething2 func is", s)

	// Retourne la chaîne passée en argument "s3" et la chaîne "World"
	return s3, "World"
}
