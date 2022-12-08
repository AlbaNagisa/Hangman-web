package HangmanModule

/*
Cree le jeu
*/
func SetHangman(difficulties string, session *Session) HangManData {
	var d HangManData
	if difficulties == "easteregg" && session.EasterEgg < 10 {
		session.EasterEgg++
		return d
	}

	d.ToFind = RandomWord(difficulties)
	d.Attempts = 10
	d.Word = CreateWordWith_(d.ToFind)
	for i := 'A'; i <= 'Z'; i++ {
		var a Alphabet = Alphabet{
			Letter: string(i),
			Used:   false,
		}
		d.Alphabet = append(d.Alphabet, a)
	}
	d.Win = false
	d.Difficulty = difficulties

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
