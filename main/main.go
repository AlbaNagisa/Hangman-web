package main

import (
	"Hangman-web/HangmanModule"
	"fmt"
	"log"
	"net/http"
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
			if strings.Contains(r.URL.Path, ".html") {
				templateshtml.ExecuteTemplate(w, r.URL.Path[1:], d)
			} else {
				templateshtml.ExecuteTemplate(w, r.URL.Path[1:]+".html", d)

			}

		}
	})

	fmt.Printf("Starting server at port 8080\n")
	if err := http.ListenAndServe(":8080", nil); err != nil {
		log.Fatal(err)
	}
}
