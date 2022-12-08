package Functions

import (
	"Hangman-web/HangmanModule"
	"log"
	"strconv"
)

func Scoreboard() [][]string {
	players := CsvReader()
	Sort(&players)
	return players
}

func Podium(data HangmanModule.Session) []HangmanModule.Player {
	var podiumPlayers [][]string
	players := Scoreboard()
	data.Scoreboard = []HangmanModule.Player{}
	podiumPlayers = players[:3]
	indexPodiumPlayers := []int{1, 2, 3}

	for x, i := range players {
		if x <= 4 {
			if data.Email == i[0] {
				podiumPlayers = players[:5]
				indexPodiumPlayers = []int{1, 2, 3, 4, 5}
				break
			}
		} else {
			if data.Email == i[0] {
				podiumPlayers = append(podiumPlayers, players[x-1])
				indexPodiumPlayers = append(indexPodiumPlayers, x-1)
				podiumPlayers = append(podiumPlayers, i)
				indexPodiumPlayers = append(indexPodiumPlayers, x)
				break
			}
		}
	}
	if !data.Logged {
		podiumPlayers = players[:5]
		indexPodiumPlayers = []int{1, 2, 3, 4, 5}
	}
	for x, i := range podiumPlayers {
		data.Scoreboard = AddToScorboard(i, data, indexPodiumPlayers[x])
	}

	SortStruct(&data.Scoreboard)
	log.Println(data.Scoreboard)
	return data.Scoreboard
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

func SortStruct(a *[]HangmanModule.Player) {
	for i := 0; i < len((*a)); i++ {
		for j := 0; j < len((*a)); j++ {
			b := (*a)[i]
			c := (*a)[j]

			if c.Points == b.Points {
				continue
			}
			if c.Points < b.Points {
				(*a)[i], (*a)[j] = (*a)[j], (*a)[i]
			}
		}
	}
}

func AddToScorboard(i []string, data HangmanModule.Session, pos int) []HangmanModule.Player {
	var p HangmanModule.Player
	pts, _ := strconv.Atoi(i[7])
	p = HangmanModule.Player{
		Pseudo:   i[2],
		Points:   pts,
		Position: pos,
	}
	data.Scoreboard = append(data.Scoreboard, p)
	return data.Scoreboard
}
