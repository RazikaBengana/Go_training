package main

import "log"

func main() {
	// utilisation de "if"

	// Déclaration d'une variable "isTrue" de type "booléen"
	var isTrue bool

	// Initialisation de "isTrue"
	// Cette variable est utilisée comme un drapeau (flag) pour contrôler l'exécution conditionnelle des blocs de code suivants
	isTrue = false

	if isTrue == true {
		log.Println("isTrue is", isTrue)
	} else {
		log.Println("isTrue is", isTrue)
	}

	// Déclaration et initialisation d'une variable string "cat"
	cat := "cat2"

	if cat == "cat" {
		log.Println("Cat is cat")

	} else {
		log.Println("Cat is not cat")
	}

	// Déclaration et initialisation d'une variable int "myNum"
	myNum := 100

	if myNum > 99 && !isTrue {
		// Cette seule ligne sera validée car "myNum" est > 99 et "isTrue" est false
		// Puisque "!isTrue" évalue à true, cela rend cette condition vraie si "myNum" est également > 99
		log.Println("myNum is greater than 99 and isTrue is false")

	} else if myNum < 100 && isTrue {
		log.Println("1")

	} else if myNum == 101 || isTrue {
		log.Println("2")

	} else if myNum > 1000 && isTrue {
		log.Println("3")
	}

	myVar := "cat"

	// Utilisation de "switch"
	switch myVar {
	case "cat":
		log.Println("myVar is set to cat")
	case "dog":
		log.Println("myVar is set to dog")
	case "horse":
		log.Println("myVar is set to horse")
	case "fish":
		log.Println("myVar is set to fish")
	default:
		log.Println("myVar is something else")
	}
}

/* Dans d'autres langages de programmation utilisant l'instruction "switch", l'exécution continue
   pour tester les autres cas même après une correspondance;
   Si trois cas correspondent, les trois actions correspondantes seraient exécutées;

   Go se comporte différemment : dès qu'un cas correspond, Go sort du "switch";
   Ce comportement évite l'utilisation de l'instruction "break" après chaque cas, comme c'est souvent nécessaire dans d'autres langages
*/
