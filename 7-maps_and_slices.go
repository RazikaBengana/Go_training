package main

import (
	"log"
	"sort" // Importe le package "sort" pour trier les slices
)

type User2 struct {
	FirstName string
	LastName  string
}

func main() {
	// Initialisation d'une map avec des clés de type "string" et des valeurs de type "User2"
	myMap := make(map[string]User2)

	// Création d'une instance de "User2'
	me := User2{
		FirstName: "Marie",
		LastName:  "JEANNE",
	}

	// Ajoute l'instance "me" dans la map avec la clé "me"
	myMap["me"] = me

	// Affiche le prénom et le nom de l'utilisateur associé à la clé "me"
	log.Println(myMap["me"].FirstName, myMap["me"].LastName)

	// Déclaration d'un slice de "int" vide
	var mySlice []int

	// Ajout d'éléments dans le slice
	mySlice = append(mySlice, 2)
	mySlice = append(mySlice, 1)
	mySlice = append(mySlice, 3)

	sort.Ints(mySlice) // Tri du slice en ordre croissant

	log.Println(mySlice) // Affiche le slice trié

	// Création et initialisation d'un slice de "int"
	numbers := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	log.Println(numbers)

	// Affiche les deux premiers éléments du slice "numbers"
	log.Println(numbers[0:2])

	// Création d'un slice de "strings"
	names := []string{"one", "seven", "fish", "cat"}
	log.Println(names)
}

/* Maps:
   - Structures de données intégrées au langage;
   - Collection de paires clé-valeur où chaque clé est unique;
   - Modifiables : on peut ajouter, modifier et supprimer des paires clé-valeur;
   - Non ordonnées : l'ordre d'insertion ne détermine pas l'ordre de stockage ou d'itération;
   - Usage : idéal pour des accès rapides par clé, comme des lookups ou des dictionnaires
*/

/* Slices:
   - Séquences d'éléments de même type, avec une taille modifiable;
   - Basés sur des tableaux, mais offrent plus de flexibilité car ils peuvent dynamiquement changer de taille;
   - Usage : parfait pour manipuler des collections d'éléments, comme des listes ou des séquences
*/
