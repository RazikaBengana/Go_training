package main

import (
	"errors" // Pour la gestion des erreurs
	"log"
)

func main() {
	// Tentative de division de 100 par 0
	result, err := divide(100, 0)
	// Si une erreur est retournée, elle est affichée dans les logs et la fonction main est terminée
	if err != nil {
		log.Println(err)
		return
	}

	// Si aucune erreur, affichage du résultat de la division
	log.Println("result of division is", result)
}

// Fonction pour diviser deux nombres flottants
func divide(x, y float32) (float32, error) {
	var result float32
	// Vérification si le diviseur est égal à 0
	if y == 0 {
		// Retourne une erreur si tentative de division par 0
		return result, errors.New("cannot divide by 0")
	}
	result = x / y     // Effectue la division si y != 0
	return result, nil // Retourne le résultat de la division et aucune erreur
}
