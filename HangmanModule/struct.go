package HangmanModule

type HangManData struct {
	Word     string   // Mot affiche
	ToFind   string   // Mot complet
	Attempts int      // Essais restants
	Tries    []string // Lettres testees
	Name     string   // Nom de la backup
	Alphabet []string //un alphabet wesh
}

type MenuLoc []struct {
	name string
	x, y int
}

type Settings struct {
	asciiPath string
}
