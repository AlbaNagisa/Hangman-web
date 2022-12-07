package HangmanModule

/*
Cree le jeu
*/
func SetHangman(difficulties string) HangManData {
	var d HangManData
	if difficulties == "" {
		var indexDif []int
		Random(2, 1, &indexDif, true)
		difficulties = []string{"easy", "medium", "hard"}[indexDif[0]]
	}
	d = HangManData{
		ToFind:     RandomWord(difficulties),
		Attempts:   10,
		Word:       CreateWordWith_(d.ToFind),
		Alphabet:   []string{"A", "B", "C", "D", "E", "F", "G", "H", "I", "J", "K", "L", "M", "N", "O", "P", "Q", "R", "S", "T", "U", "V", "W", "X", "Y", "Z"},
		Win:        false,
		Difficulty: difficulties,
		TriesContain: func(letter string) bool {
			for _, elt := range d.Tries {
				if elt == letter {
					return true
				}
			}
			return false
		},
	}

	return d
}

/*
Transforme le mot en mot avec lettre avec des tirets
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
