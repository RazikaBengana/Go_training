package main

// Importation du package "helpers" personnalisé pour utiliser la fonction "RandomNumber"
import (
	"github.com/RazikaBengana/12-channels/helpers"
	"log"
)

// Déclare une constante pour la valeur maximale utilisée lors de la génération d'un nombre aléatoire
const numPool = 1000

// Fonction qui écrit une valeur aléatoire dans un channel
func calculateValue(intChan chan int) {

	// Génère un nombre aléatoire en utilisant la fonction du package "helpers"
	randomNumber := helpers.RandomNumber(numPool)
	// Envoie le nombre aléatoire généré dans le channel
	intChan <- randomNumber
}

func main() {
	// Crée un channel pour les entiers
	intChan := make(chan int)
	// "defer" assure la fermeture du channel à la fin de l'exécution de main() --> important !
	defer close(intChan)

	// - Exécute "calculateValue" dans une goroutine, permettant ainsi une exécution concurrente, simultanée;
	// - Pendant que "calculateValue" génère un nombre aléatoire et l'envoie dans le channel "intChan",
	//   le programme principal continue son exécution sans attendre que "calculateValue" termine
	go calculateValue(intChan)

	// Reçoit une valeur du channel et la stocke dans "num"
	num := <-intChan
	log.Println(num)
}

/*

Channel :

- Un channel en Go est un conduit typé utilisé pour communiquer entre goroutines;
- On peut penser à un channel comme à un tube où on peut envoyer et recevoir des valeurs
  entre différentes parties d'un programme exécutant des opérations concurrentes;
- Les channels garantissent la synchronisation et la sécurité des données, évitant ainsi
  les problèmes liés à l'accès concurrent aux ressources partagées


- Par exemple dans ce code :

  - "intChan" est un channel d'entiers;
  - La fonction "calculateValue" génère un nombre aléatoire et l'envoie dans "intChan" avec "intChan <- randomNumber";
  - Dans "main", on reçoit ce nombre aléatoire avec "num := <-intChan" et on l'affiche ensuite

*/

/*

Goroutine :

- Une goroutine en Go est un fil d'exécution très léger géré par le runtime de Go;
- Elle permet d'exécuter des fonctions de manière concurrente, parallèle à d'autres goroutines;
- Pour démarrer une goroutine : il suffit d'utiliser le mot-clé "go" suivi de l'appel de fonction;
- Les goroutines sont au cœur de la programmation concurrente en Go, permettant de réaliser des opérations
  asynchrones sans bloquer l'exécution du programme principal


- Par exemple dans ce code :

  - "go calculateValue(intChan)" lance la fonction "calculateValue" dans une goroutine :
    cela permet à "calculateValue" de s'exécuter de façon concurrente au reste du programme principal;
  - Pendant que "calculateValue" travaille (ici, génère un nombre aléatoire et l'envoie dans le channel "intChan"),
    le programme principal n'est pas interrompu et peut continuer à exécuter d'autres instructions

*/
