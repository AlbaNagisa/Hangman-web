package Functions

import (
	"Hangman-web/HangmanModule"
	"strconv"
)

func Scoreboard() [][]string {
	players := CsvReader()
	Sort(&players)
	return players
}

func Podium(data HangmanModule.Session) {
	var podiumPlayers [][]string
	players := Scoreboard()
	podiumPlayers = players[:3]
	alreadyAdded := false
	for _, i := range podiumPlayers {
		AddToScorboard(i, data)
		if data.Email == i[0] {
			alreadyAdded = true
		}
	}
	for x, i := range players {
		if data.Email == i[0] {
			if !alreadyAdded {
				AddToScorboard(i, data)
				if x != 0 {
					AddToScorboard(players[x-1], data)
				}
			}

		}
	}
}

func Sort(a *[][]string) {
	for i := 0; i < len((*a)); i++ {
		for j := 0; j < len((*a)); j++ {
			b, _ := strconv.Atoi((*a)[i][7])
			c, _ := strconv.Atoi((*a)[j][7])
			if c < b {
				(*a)[i], (*a)[j] = (*a)[j], (*a)[i]
			}
		}
	}
}

/*
	 func SortStruct(a *[]HangmanModule.Player) {
		for i := 0; i < len((*a)); i++ {
			for j := 0; j < len((*a)); j++ {
				b := (*a)[i]
				c := (*a)[j]
				if c.Points < b.Points {
					(*a)[i], (*a)[j] = (*a)[j], (*a)[i]
				}
			}
		}
	}
*/
func AddToScorboard(i []string, data HangmanModule.Session) {
	var p HangmanModule.Player
	pts, _ := strconv.Atoi(i[7])
	p = HangmanModule.Player{
		Pseudo: i[2],
		Points: pts,
	}
	data.Scoreboard = append(data.Scoreboard, p)
}
