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

	d.Alphabet = []string{"A", "B", "C", "D", "E", "F", "G", "H", "I", "J", "K", "L", "M", "N", "O", "P", "Q", "R", "S", "T", "U", "V", "W", "X", "Y", "Z"}
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		var templateshtml = template.Must(template.ParseGlob("./web/static/*.html"))

		switch r.URL.Path {
		case "/":
			templateshtml.ExecuteTemplate(w, "index.html", d)
		case "/setup":
			d = HangmanModule.SetHangman()
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
