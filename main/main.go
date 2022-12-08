package main

import (
	"Hangman-web/Functions"
	"Hangman-web/HangmanModule"
	"encoding/base64"
	"log"
	"net/http"
	"os"
	"strconv"
	"strings"
	"text/template"
)

func main() {

	fileServer := http.FileServer(http.Dir("./web/assets"))
	http.Handle("/assets/", http.StripPrefix("/assets/", fileServer))

	var d HangmanModule.Session

	dir, _ := os.Getwd()
	files, _ := os.ReadDir(dir + "/web/static")
	var fichier []string
	for _, file := range files {
		fichier = append(fichier, file.Name()[:len(file.Name())-5])
	}

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		var templateshtml = template.Must(template.ParseGlob("./web/static/*.html"))
		if (r.Method == "POST") && (r.URL.Path == "/") {
			d.Logged = true
			d.Email = r.FormValue("email")
			d.Mdp = base64.StdEncoding.EncodeToString([]byte(r.FormValue("password")))
			d.Pseudo = r.FormValue("pseudo")
			Functions.CsvWritter(d)
		}
		if (r.Method == "POST") && (r.URL.Path == "/login") {
			d.Mdp = base64.StdEncoding.EncodeToString([]byte(r.FormValue("password")))
			d.Email = r.FormValue("email")
			csvFile := Functions.CsvReader()
			for _, line := range csvFile {
				if strings.EqualFold(strings.ToLower(line[0]), strings.ToLower(d.Email)) {
					if d.Mdp == line[1] {
						d.Logged = true
						d.Email = line[0]
						d.Mdp = line[1]
						d.NLoose, _ = strconv.Atoi(line[4])
						d.Pseudo = line[2]
						d.NWin, _ = strconv.Atoi(line[3])
						d.Ratio, _ = strconv.Atoi(line[5])
						d.Points, _ = strconv.Atoi(line[7])
						http.Redirect(w, r, "/", http.StatusFound)
					}
				}
			}
		}
		switch r.URL.Path {
		case "/":
			templateshtml.ExecuteTemplate(w, "index.html", d)
		case "/setup":
			d.Scoreboard = Functions.Podium(d)
			d.Game = HangmanModule.SetHangman(r.URL.Query().Get("level"))
			http.Redirect(w, r, "/jeu", http.StatusFound)
		case "/hangman":
			if r.Method == "POST" {
				word := r.FormValue("word")
				if word != "" {
					check(d.Game.ToFind, word, &d.Game)
				}
			} else {
				letter := r.URL.Query().Get("letter")
				if letter != "" {
					check(d.Game.ToFind, letter, &d.Game)
				}
			}
			if d.Game.Attempts <= 0 {
				d.Game.Loose = true
				if d.Logged {
					d.NLoose += 1
					if d.NLoose != 0 {
						d.Ratio = d.NWin / d.NLoose
					}
					Functions.CsvEditor(d)
				}
			}
			if d.Game.ToFind == d.Game.Word {
				d.Game.Win = true
				if d.Logged {
					if d.NLoose != 0 {
						d.Ratio = d.NWin / d.NLoose
					}
					switch d.Game.Difficulty {
					case "easy":
						d.Points += 2
					case "medium":
						d.Points += 5
					case "hard":
						d.Points += 10
					}
					d.NWin += 1
					Functions.CsvEditor(d)
				}
			}
			d.Scoreboard = Functions.Podium(d)
			http.Redirect(w, r, "/jeu", http.StatusFound)
		default:
			exist := false
			for i := 0; i < len(fichier); i++ {
				if strings.Contains(r.URL.Path, fichier[i]) {
					exist = true
					break
				}
			}
			if !exist {
				http.Redirect(w, r, "/404", http.StatusFound)
			}

			if len(d.Game.Alphabet) == 0 && (r.URL.Path == "/jeu" || r.URL.Path == "/fin") {
				http.Redirect(w, r, "/setup", http.StatusFound)
			}

			if strings.Contains(r.URL.Path, ".html") {
				templateshtml.ExecuteTemplate(w, r.URL.Path[1:], d)
			} else {
				templateshtml.ExecuteTemplate(w, r.URL.Path[1:]+".html", d)
			}

		}
	})

	log.Println("Starting server at port 8080")

	if err := http.ListenAndServe(":8080", nil); err != nil {
		log.Fatal(err)
	}
}

func check(word, input string, d *HangmanModule.HangManData) {
	ts := []rune(d.Word)
	word = strings.ToLower(word)
	found := false
	foundUsed := false
	input = strings.ToLower(input)
	if len(input) == 1 {
		for i, x := range word {
			for _, u := range d.Tries {
				if u == input {
					foundUsed = true
				}
			}

			if x == []rune(input)[0] {
				ts[i] = []rune(input)[0]
				found = true
			} else {
				if i == len(word)-1 && !found && !foundUsed {
					d.Attempts -= 1
					d.Tries = append(d.Tries, input)
				}
			}
			for i := 0; i < len(d.Alphabet); i++ {
				if input == strings.ToLower(d.Alphabet[i].Letter) {
					d.Alphabet[i].Used = true
				}
			}
		}

	} else {
		if len(input) > 1 {
			if word == input {
				ts = []rune(d.ToFind)

			} else {
				d.Attempts -= 2

			}
		}
	}
	d.Word = string(ts)
}
