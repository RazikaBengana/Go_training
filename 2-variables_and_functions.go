package main

import "fmt"

func main() {
	fmt.Println("Hello, world.")

	// Déclaration de la variable "whatToSay" de type string
	var whatToSay string

	// Déclaration de la variable "i" de type int

	/* En Go, la taille d'une variable de type "int" dépend de l'architecture de l'ordinateur sur lequel le programme est compilé et exécuté;
	   Par exemple, sur un ordinateur 64 bits, int est de 64 bits par défaut;
	   Sur un ordinateur 32 bits, int est de 32 bits par défaut;
	   etc...
	*/
	var i int

	whatToSay = "Goodbye, cruel world"

	fmt.Println(whatToSay)

	i = 7

	fmt.Println("i is set to", i)

	// Appel de la fonction "saySomething"
	// Assignation de ses valeurs de retour à "whatWasSaid" et "theOtherThingThatWasSaid"
	whatWasSaid, theOtherThingThatWasSaid := saySomething()

	// Affichage des valeurs retournées par la fonction "saySomething"
	fmt.Println("The function returned", whatWasSaid, theOtherThingThatWasSaid)
}

// Définition de la fonction "saySomething"
// Cette fonction retourne 2 chaînes de caractères
func saySomething() (string, string) {
	return "something", "else"
}
