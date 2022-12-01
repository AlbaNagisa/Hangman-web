package HangmanModule

type HangManData struct {
	Word     string // Mot affiche
	ToFind   string // Mot complet
	Attempts int    // Essais restants
	Tries    []rune // Lettres testees
	Name     string // Nom de la backup
	Alphabet []string
}

type MenuLoc []struct {
	name string
	x, y int
}

type Settings struct {
	asciiPath string
}
