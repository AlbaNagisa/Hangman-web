package HangmanModule

import (
	"log"
	"math/rand"
	"strings"
	"time"
)

/*
	Random est une fonction recursive qui prend un tableau de int pointer et le renvoie avec les nombres généré

les parametres :

	max est le nombre max qui peut etre générer
	n est le nombre de chiffre a générer avec le meme max
	r est un tableau de int pointé
	lorsque unique est sur false la fonction peut alors généré plusieurs fois le meme nombre
*/
func Random(max, n int, r *[]int, unique bool) {
	if unique && (n > max-1) {
		log.Fatalf("Error: I can't generate more than %d different numbers", max-1)
	}
	//génére une seed par rapport au temps car c'est une valeur qui change en permanence
	seed := rand.NewSource(time.Now().UnixNano())
	random := rand.New(seed)
	//génére un nombre entier entre 0 est le max
	nbGen := random.Intn(max)
	//si on ne veut pas deux fois le meme nombre fait de recurcivité jusqu'à trouver un nombre différent
	if unique {
		for _, v := range *r {
			if nbGen == v {
				Random(max, n, r, unique)
			}
		}
	}
	if len(*r) != n {
		*r = append(*r, nbGen)
		Random(max, n, r, unique)
	}
}

/*
Cette fonction choisi un mot random dans le fichier mis en argument
*/
func RandomWord(difficulties string) string {
	var words []string
	var randomNumber []int

	FileReader(&words, "HangmanModule/Dicos/"+difficulties+".txt") // Mise du fichier en argument de la commande dans words
	Random(len(words), 1, &randomNumber, true)                     // Choix d'un chiffre entre 0 et len(word) et mise dans randomNumber
	return strings.ToLower(words[randomNumber[0]])                 // Renvoi du mot selectionne par Random()
}
