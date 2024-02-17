package main

// Importation du paquet "testing" qui fournit le support nécessaire pour écrire des tests unitaires
import "testing"

// Tests manuels

// "TestDivide" vérifie que la fonction "divide" retourne correctement sans erreur pour une division valide
func TestDivide(t *testing.T) {
	_, err := divide(10.0, 1.0)
	// "nil" = absence de valeur
	if err != nil {
		t.Error("Got an error when should not have")
	}
}

// "TestBadDivide" vérifie que la fonction "divide" retourne une erreur pour une division par zéro
func TestBadDivide(t *testing.T) {
	_, err := divide(10.0, 0)
	// Ici, on signale une erreur si aucune erreur n'est retournée alors qu'elle est attendue
	if err == nil {
		t.Error("Did not get an error when should have")
	}
}

// Tests en tableaux

// Définit un tableau de structures pour tester plusieurs cas
var tests = []struct {
	name     string
	dividend float32
	divisor  float32
	expected float32
	isErr    bool
}{
	{"valid-data", 100.0, 10.0, 10.0, false}, // "false" signifie qu'on n'attend pas d'erreur
	{"invalid-data", 100.0, 0.0, 0.0, true},  // "true" signifie qu'on attend une erreur
	{"expect-5", 50.0, 10.0, 5.0, false},
	{"expect-fraction", -1.0, -777.0, 0.0012870013, false},
}

// "TestDivision" teste la fonction "divide" avec plusieurs cas de figure définis dans le tableau "tests"
func TestDivision(t *testing.T) {
	for _, tt := range tests {
		got, err := divide(tt.dividend, tt.divisor) // Teste la division avec les valeurs du tableau
		if tt.isErr {
			if err == nil {
				t.Error("Expected an error but did not get one") // Signale une erreur si une erreur était attendue mais non retournée
			}
		} else {
			if err != nil {
				t.Error("Did not expect an error, but got one", err.Error()) // Signale une erreur si une erreur est retournée alors qu'elle n'était pas attendue
			}
		}

		if got != tt.expected {
			t.Errorf("Expected %f but got %f", tt.expected, got) // Compare le résultat obtenu avec le résultat attendu
		}
	}
}

/*

Commandes utiles pour les tests :

	go test = Exécute les tests sans afficher de détails

	go test -v = Exécute les tests et affiche les détails des tests exécutés

	go test -cover = Affiche le pourcentage de code couvert par les tests

	go test -coverprofile=coverage.out && go tool cover -html=coverage.out = Génère et ouvre un rapport de couverture dans le navigateur

*/
