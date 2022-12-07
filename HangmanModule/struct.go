package HangmanModule

type HangManData struct {
	Word     string   // Mot affiche
	ToFind   string   // Mot complet
	Attempts int      // Essais restants
	Tries    []string // Lettres testees
	Name     string   // Nom de la backup
	Alphabet []string //un alphabet wesh
	Win      bool
}

type Session struct {
	Logged bool
	Game   HangManData
	Email  string
	Pseudo string
	Mdp    string
	NWin   int
	NLoose int
	Ratio  int
}

type MenuLoc []struct {
	name string
	x, y int
}

type Settings struct {
	asciiPath string
}
