package main

import (
	"fmt"
	"html/template"
	"log"
	"net/http"
	"time"
)

type Film struct {
	Title    string
	Director string
}

func main() {
	fmt.Println("Hello, World!")

	h1 := func(w http.ResponseWriter, r *http.Request) {

		temp := template.Must(template.ParseFiles("index.html"))

		films := map[string][]Film{
			"Films": {
				{Title: "The Shawshank Redemption", Director: "Frank Darabont"},
				{Title: "The Godfather", Director: "Francis Ford Coppola"},
				{Title: "The Dark Knight", Director: "Christopher Nolan"},
				{Title: "The amazing spider man ", Director: "Christopher Nolan"},
			},
		}

		temp.Execute(w, films)
	}
	h2 := func(w http.ResponseWriter, r *http.Request) {
		time.Sleep(1 * time.Second)
		title := r.PostFormValue("title")
		director := r.PostFormValue("director")
		htmlStr := fmt.Sprintf("<li class='list-group-item bg-primary text-white'> %s - %s </li>", title, director)
		tmpl, _ := template.New("t").Parse(htmlStr)
		// if r.Method == "POST" {
		// 	r.ParseForm()
		// 	title := r.FormValue("title")
		// 	director := r.FormValue("director")
		// 	log.Println("Title:", title)
		// 	log.Println("Director:", director)
		// }
		tmpl.Execute(w, nil)
	}

	http.HandleFunc("/", h1)
	http.HandleFunc("/add-film/", h2)

	log.Fatal(http.ListenAndServe(":8080", nil))
}
