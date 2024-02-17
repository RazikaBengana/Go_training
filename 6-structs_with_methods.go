package main

import "log"

// Définition d'une structure de type "myStruct" avec un champ "FirstName" de type string
type myStruct struct {
	FirstName string
}

// Déclaration d'une méthode "printFirstName" pour la structure "myStruct"
// "m *myStruct" est appelé le récepteur de la méthode, permettant d'accéder aux champs de "myStruct"
func (m *myStruct) printFirstName() string {

	return m.FirstName // Retourne la valeur du champ "FirstName" de l'instance de "myStruct"
}

func main() {
	// Déclaration d'une variable "myVar" de type "myStruct" avec allocation de mémoire
	var myVar myStruct

	myVar.FirstName = "John"

	myVar2 := myStruct{
		FirstName: "Mary",
	}

	// Utilisation de la méthode "printFirstName"
	log.Println("myVar is set to", myVar.printFirstName())
	log.Println("myVar2 is set to", myVar2.printFirstName())
}

// Une méthode est une fonction spéciale attachée à une structure (ou un type)
// Elle peut accéder aux données de la structure sur laquelle elle est appelée et les manipuler, à l'aide d'un récepteur

/* Ici, "m *myStruct" est le récepteur de la méthode, indiquant que cette méthode est liée à la structure "myStruct";
   Lorsque "printFirstName()" est appelée sur une instance de "myStruct", elle utilise le récepteur "m" pour accéder au champ
   "FirstName" de cette instance et retourner sa valeur
*/
