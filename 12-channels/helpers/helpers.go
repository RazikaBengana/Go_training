// Package helpers déclaration
package helpers

import (
	"math/rand" // Import du package "rand" pour la génération de nombres aléatoires
	"time"      // Import du package "time" pour accéder à la fonction "UnixNano"
)

// RandomNumber génère un nombre aléatoire jusqu'à "n"
func RandomNumber(n int) int {
	rand.Seed(time.Now().UnixNano()) // Initialise le générateur de nombres aléatoires
	value := rand.Intn(n)            // Génère un nombre aléatoire entre 0 et n-1
	return value                     // Retourne le nombre aléatoire généré
}
