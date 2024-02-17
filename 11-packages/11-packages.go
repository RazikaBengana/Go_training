package main

// Import du package "helpers" depuis son emplacement dans le module
// L'import utilise le chemin complet défini dans le fichier go.mod
import (
	"github.com/RazikaBengana/11-packages/helpers"
	"log"
)

func main() {
	// Déclare une variable de type "SomeType", qui est défini dans le package "helpers"
	var myVar helpers.SomeType

	// Attribue une valeur à la propriété "TypeName" de la structure "SomeType"
	myVar.TypeName = "Some name"

	log.Println(myVar.TypeName)
}
