package main

import (
	"fmt"
	"html/template"
	"log"
	"net/http"
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

	http.HandleFunc("/", h1)

	log.Fatal(http.ListenAndServe(":8080", nil))
}
