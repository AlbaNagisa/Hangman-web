package HangmanModule

type HangManData struct {
	Word       string     // Mot affiche
	ToFind     string     // Mot complet
	Attempts   int        // Essais restants
	Tries      []string   // Lettres testees
	Name       string     // Nom de la backup
	Alphabet   []Alphabet //un alphabet wesh
	Win        bool
	Loose      bool
	Difficulty string
}

type Session struct {
	Logged           bool
	Game             HangManData
	Email            string
	Pseudo           string
	Mdp              string
	NWin             int
	NLoose           int
	Ratio            int
	Points           int
	Scoreboard       []Player
	EasterEgg        int
	FunnyModeEnabled bool
	HomeEasterEgg    bool
}

type Player struct {
	Pseudo   string
	Points   int
	Position int
}

type Alphabet struct {
	Letter string
	Used   bool
}
