package main

import "log"

func main() {
	var myString string
	myString = "Green"

	// "log.Println" permet d'écrire un message dans la sortie standard de logs (journal), ce qui est utile surveiller le comportement du programme
	// "log.Println" est une fonction variadique : elle peut prendre zéro, un ou plusieurs paramètres
	log.Println("myString is set to", myString)

	// Appelle la fonction "changeUsingPointer" en passant l'adresse de "myString"
	changeUsingPointer(&myString)

	// Affiche dans le journal que "myString" a été modifié à "Red" après l'appel de fonction
	log.Println("after func call myString is set to", myString)
}

func changeUsingPointer(s *string) {

	// Ajout du déréférencement de "s" pour afficher sa valeur
	log.Println("s is set to", *s)

	// Le signe ":=" est un raccourci qui combine la déclaration et l'initialisation de "newValue" avec "Red"
	newValue := "Red"

	// Modifie la valeur pointée par "s" pour lui affecter "Red"
	*s = newValue
}
