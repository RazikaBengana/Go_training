package main

import (
	"log"
	// Importe le paquet "log" pour la journalisation des messages dans la console --> utile pour le débogage ou l'enregistrement d'informations
	"time"
	// Importe le paquet "time" pour manipuler des dates et des heures
)

// "User1" est une structure définissant un utilisateur
// Les champs commencant par une majuscule sont exportés, c'est-à-dire accessibles depuis d'autres paquets
type User1 struct {
	FirstName   string
	LastName    string
	PhoneNumber string
	Age         int       // "Age" non initialisé ici, sa valeur par défaut sera donc 0
	BirthDate   time.Time // "Birthdate" non initialisée ici, sa valeur par défaut sera donc zero value de "time.Time"
}

func main() {

	// Création d'une instance de "User1" avec initialisation de certains champs
	// Les champs non initialisés, comme "Age" et "BirthDate", auront leurs valeurs par défaut
	user := User1{
		FirstName:   "Tom",
		LastName:    "SARVET",
		PhoneNumber: "06 34 45 67 89",
	}

	log.Println(user.FirstName, user.LastName, "Birthdate:", user.BirthDate)
}
