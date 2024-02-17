package main

import "fmt"

// "Animal" définit l'interface pour le type "Animal"
type Animal interface {
	Says() string      // Retourne le son que fait l'animal
	NumberOfLegs() int // Retourne le nombre de jambes de l'animal
}

type Dog struct {
	Name  string
	Breed string
}

type Gorilla struct {
	Name          string
	Color         string
	NumberOfTeeth int
}

func main() {
	dog := Dog{
		Name:  "Samson",
		Breed: "German Shepherd",
	}

	// Nous pouvons passer "dog" à "PrintInfo()", car le type "Dog" implémente l'interface "Animal" en ayant toutes les
	// fonctions nécessaires;
	// Le paramètre est passé comme un pointeur puisque les fonctions pour le type ont des récepteurs pointeurs --> ce qui est la norme
	// Voir https://tour.golang.org/methods/4 et https://tour.golang.org/methods/8
	PrintInfo(&dog)

	gorilla := Gorilla{
		Name:          "Jock",
		Color:         "grey",
		NumberOfTeeth: 38,
	}

	// Même chose pour "gorilla" à "PrintInfo()"
	PrintInfo(&gorilla)
}

func PrintInfo(a Animal) {
	fmt.Println("Cet animal dit", a.Says(), "et a", a.NumberOfLegs(), "jambes")
}

// La plupart des récepteurs en Go sont de type pointeur
// "Says" a un récepteur de type "*Dog", donc il satisfait une partie des exigences de l'interface "Animal" pour le type "Dog"
func (d *Dog) Says() string {
	return "Woof"
}

// "NumberOfLegs" satisfait le reste des exigences de l'interface "Animal" pour le type "Dog"
func (d *Dog) NumberOfLegs() int {
	return 4
}

// "Says" a un récepteur de type "*Gorilla", donc il satisfait une partie des exigences de l'interface "Animal" pour le type "Gorilla"
func (d *Gorilla) Says() string {
	return "Ugh"
}

// "NumberOfLegs" satisfait le reste des exigences de l'interface "Animal" pour le type "Gorilla"
func (d *Gorilla) NumberOfLegs() int {
	return 2
}

/*

Interface :

- En Go, c'est une façon de décrire un ensemble de méthodes;
- Quand un type fournit des implémentations pour toutes les méthodes décrites dans une interface, on dit qu'il "implémente" cette interface;
- Cela permet d'utiliser des types variés de manière interchangeable, tant qu'ils partagent le même ensemble de méthodes;

- L'interface est comme un contrat : si notre type "signe" ce contrat en implémentant les méthodes requises,
  alors il peut être utilisé partout où cette interface est demandée;

- En pratique, cela signifie qu'on peut écrire des fonctions qui acceptent des interfaces comme types de paramètres;
- Ces fonctions peuvent ensuite être appelées avec n'importe quel type qui implémente l'interface,
  rendant le code plus flexible et réutilisable;

- Par exemple, si on a une interface "Animal" qui exige une méthode "Speak()", n'importe quel type qui implémente "Speak()"
  peut être passé à une fonction qui attend un "Animal", que ce soit un "Dog", un "Cat", ou tout autre type qui sait "Speak()";


- Cela démontre le polymorphisme en Go, où une interface permet de traiter différents types (comme "Dog" et "Gorilla")
  de manière uniforme à travers les méthodes (comportements) qu'ils partagent via l'interface

*/
