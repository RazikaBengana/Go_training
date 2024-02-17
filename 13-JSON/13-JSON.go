package main

import (
	"encoding/json" // Importation du paquet pour l'encodage et le décodage des données JSON
	"fmt"
	"log"
)

// Définition d'une structure "Person" pour mapper les données JSON
// Les champs sont précédés du tag "json"
type Person struct {
	FirstName string `json:"first_name"`
	LastName  string `json:"last_name"`
	HairColor string `json:"hair_color"`
	HasDog    bool   `json:"has_dog"`
}

func main() {
	// Strings représentant un tableau JSON d'objets "Person"
	myJson := `
[
    {
        "first_name": "Clark",
        "last_name": "Kent",
        "hair_color": "black",
        "has_dog": true
    },
    {
        "first_name": "Bruce",
        "last_name": "Wayne",
        "hair_color": "black",
        "has_dog": false
    }
]`

	// Déclaration d'un slice de "Person" pour stocker les données JSON décodées
	var unmarshalled []Person

	// Décodage du JSON dans le slice de "Person"
	// Gestion d'erreur en cas d'échec
	err := json.Unmarshal([]byte(myJson), &unmarshalled)
	if err != nil {
		log.Println("Error unmarshalling json", err)
	}

	// Affichage des données décodées à l'aide de "unmarshalled"
	log.Printf("unmarshalled: %v", unmarshalled)

	// Préparation de données pour l'encodage en JSON
	// Déclaration d'un slice de "Person" pour encoder en JSON
	var mySlice []Person

	var m1 Person
	m1.FirstName = "Wally"
	m1.LastName = "West"
	m1.HairColor = "red"
	m1.HasDog = false

	mySlice = append(mySlice, m1) // Ajout de la première personne au slice

	var m2 Person
	m2.FirstName = "Diana"
	m2.LastName = "Prince"
	m2.HairColor = "black"
	m2.HasDog = false

	mySlice = append(mySlice, m2) // Ajout de la deuxième personne au slice

	// Encodage du slice en JSON avec indentation pour une meilleure lisibilité
	newJson, err := json.MarshalIndent(mySlice, "", "     ")
	if err != nil {
		log.Println("error marshalling", err)
	}

	// Affichage du JSON encodé
	fmt.Println(string(newJson))
}

/*

JSON (JavaScript Object Notation) :

   - C'est un format de données textuelles léger pour l'échange de données;
   - Il est facilement lisible par les humains et analysable par les machines;
   - JSON est souvent utilisé pour transmettre des données entre un serveur et une application web ou pour stocker
     des configurations et des paramètres;

   - Les données sont organisées en paires clé/valeur et peuvent représenter :
     * des objets
     * des tableaux
     * des nombres
     * des chaînes de caractères
     * des booléens
     * des valeurs nulles

   - En Go, la sérialisation et la désérialisation JSON sont réalisées grâce aux packages standard comme `encoding/json`:

   - Sérialisation en Go (Encodage) :
     * Convertit des structures de données Go (comme des structs, slices, etc...) en chaînes de caractères JSON;
     * Cela est généralement réalisé avec "json.Marshal" ou "json.MarshalIndent" pour un formatage plus lisible;

   - Désérialisation en Go (Décodage) :
     * Convertit des chaînes de caractères JSON en structures de données Go spécifiques;
     * Cela est réalisé avec "json.Unmarshal", permettant d'interpréter les données JSON pour les utiliser dans des applications Go

*/
