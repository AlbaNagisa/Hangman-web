package HangmanModule

/*
Cree le jeu
*/
func SetHangman() HangManData {
	var d HangManData
	d.ToFind = RandomWord()
	d.Attempts = 10
	d.Word = CreateWordWith_(d.ToFind)
	d.Alphabet = []string{"A", "B", "C", "D", "E", "F", "G", "H", "I", "J", "K", "L", "M", "N", "O", "P", "Q", "R", "S", "T", "U", "V", "W", "X", "Y", "Z"}
	return d
}

/*
Transforme le mot en mot a deux lettre avec des tirets
*/
func CreateWordWith_(w string) string {
	var ts []rune
	var RandLetters []int

	for i := 0; i < len([]rune(w)); i++ {
		ts = append(ts, '_')
	}

	Random(len(ts), len(ts)/2-1, &RandLetters, true)
	for i := 0; i < len(ts); i++ {
		for _, v := range RandLetters {
			if i == v {
				ts[i] = []rune(w)[i]
			}
		}
	}

	return string(ts)
}
