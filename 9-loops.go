package main

import (
	"fmt"
	"log"
)

func main() {
	type User struct {
		FirstName string
		LastName  string
		Email     string
		Age       int
	}

	// Déclaration d'un slice de "User" pour stocker les utilisateurs
	var users []User
	users = append(users, User{"John", "Smith", "john@smith.com", 30})
	users = append(users, User{"Mary", "Jones", "mary@jones.com", 20})
	users = append(users, User{"Sally", "Brown", "sally@smith.com", 45})
	users = append(users, User{"Alex", "Anderson", "alex@smith.com", 17})

	// Boucle sur chaque utilisateur du slice "users"
	// "_," --> permet d'ignorer la 1ère valeur
	for _, user := range users {
		log.Println(user.FirstName, user.LastName, user.Email, user.Age)
	}

	fmt.Println() // Ajoute une nouvelle ligne simple

	loop2()

	fmt.Println()

	loop3()

	fmt.Println()

	loop4()

	fmt.Println()

	loop5()
}

func loop2() {
	for i := 0; i <= 10; i++ {
		log.Println(i)
	}
}

func loop3() {
	// Déclaration d'un slice de strings
	animals := []string{"dog", "fish", "horse", "cow", "cat"}

	// Boucle sur chaque élément du slice "animals"
	for i, animal := range animals {
		// Affichage de l'index et de la valeur
		log.Println(i, animal)
	}
}

func loop4() {
	// Création d'une map pour stocker des paires clé-valeur
	animals := make(map[string]string)
	// Ajout d'une entrée dans la map
	animals["dog"] = "Fido"

	// Boucle sur chaque paire clé-valeur de la map
	for animalType, animal := range animals {
		// Affichage de la clé et de la valeur
		log.Println(animalType, animal)
	}
}

func loop5() {
	var firstLine = "Once upon a midnight time"

	// Boucle sur chaque caractère de la chaîne
	// "i"           --> représente l'index de la rune (ou du caractère) dans la chaîne
	// "runeValue"   --> est la valeur de la rune elle-même, c'est-à-dire le point de code Unicode du caractère
	for i, runeValue := range firstLine {
		log.Println(i, ":", runeValue)
	}
}

// En Go, une chaîne est en fait une tranche de runes (bytes), d'où l'affichage de numéros et non de lettres directement;
// rune = point de code Unicode

// --------------------------------------------------------------------------------------------------

// Les chaînes en Go sont immuables, signifiant qu'elles ne peuvent pas être modifiées après leur création

// --------------------------------------------------------------------------------------------------

/* Rappel :
   - "log" est utilisé pour logger des infos, et ajoute automatiquement un timestamp au début de chaque message
   - "fmt" est utilisé pour des affichages simples
*/

// --------------------------------------------------------------------------------------------------

// En Go, la boucle "for" est la seule structure de boucle disponible dans le langage
