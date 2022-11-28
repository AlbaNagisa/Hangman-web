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
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		var templateshtml = template.Must(template.ParseGlob("./web/static/*.html"))

		switch r.URL.Path {
		case "/":
			templateshtml.ExecuteTemplate(w, "index.html", d)
		case "/setup":
			d = HangmanModule.SetHangman()
			http.Redirect(w, r, "/jeu", http.StatusFound)
		default:
			dir, _ := os.Getwd()
			files, _ := os.ReadDir(dir + "/web/static")
			exist := false
			for _, file := range files {
				if strings.Contains(r.URL.Path, ".html") {
					if strings.Contains(r.URL.Path, file.Name()) {
						exist = true
						break
					}
				} else {
					if strings.Contains(r.URL.Path+".html", file.Name()) {
						exist = true
						break
					}
				}
			}
			if !exist {
				templateshtml.ExecuteTemplate(w, "404.html", "")
			} else {
				if strings.Contains(r.URL.Path, ".html") {
					templateshtml.ExecuteTemplate(w, r.URL.Path[1:], d)
				} else {
					templateshtml.ExecuteTemplate(w, r.URL.Path[1:]+".html", d)
				}
			}

		}
	})

	log.Println("Starting server at port 5050")

	if err := http.ListenAndServe(":5050", nil); err != nil {
		log.Fatal(err)
	}
}
