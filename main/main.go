package main

import (
	"Hangman-web/HangmanModule"
	"log"
	"net/http"
	"os"
	"strings"
	"text/template"
)

func main() {

	fileServer := http.FileServer(http.Dir("./web/assets"))
	http.Handle("/assets/", http.StripPrefix("/assets/", fileServer))

	var d HangmanModule.HangManData

	dir, _ := os.Getwd()
	files, _ := os.ReadDir(dir + "/web/static")
	var fichier []string
	for _, file := range files {
		fichier = append(fichier, file.Name()[:len(file.Name())-5])
	}

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		var templateshtml = template.Must(template.ParseGlob("./web/static/*.html"))
		switch r.URL.Path {
		case "/":
			templateshtml.ExecuteTemplate(w, "index.html", d)
		case "/setup":
			d = HangmanModule.SetHangman()
			http.Redirect(w, r, "/jeu", http.StatusFound)
		case "/hangman":
			log.Println(r.Method)
			if r.Method == "POST" {
				word := r.FormValue("word")
				if word != "" {
					log.Println(word)
				}
			} else {
				letter := r.URL.Query().Get("letter")

				if letter != "" {
					check(d.ToFind, letter, &d)
				}
			}
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

			if len(d.Alphabet) == 0 && r.URL.Path == "/jeu" {
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
	found := false
	foundUsed := false
	input = strings.ToLower(input)
	log.Println(d)
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
		}
	}
	d.Word = string(ts)
}
