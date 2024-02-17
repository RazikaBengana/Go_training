// Déclaration du package "helpers"
// Tous les fichiers dans le même répertoire appartiennent au même package --> ici, helpers.go appartient donc au package helpers
package helpers

// Définit une structure publique "SomeType", accessible depuis d'autres packages
// Les structures et leurs champs commençant par une majuscule sont exportés et donc accessibles en dehors du package

type SomeType struct {
	TypeName   string // Champ exporté, accessible en dehors du package
	TypeNumber int    // Autre champ exporté
}
