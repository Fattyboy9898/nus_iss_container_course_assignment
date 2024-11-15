package main

import (
	"html/template"
	"log"
	"math/rand"
	"net/http"
)

// Define a struct to hold the dynamic data
type PageData struct {
	Message string
}

func get_random_quote() (string, error) {
	quotes := []string{
		"Logic will get you from A to B. Imagination will take you everywhere ",
		"There are 10 kinds of people. Those who know binary and those who don't.",
		"There are two ways of constructing a software design. One way is to make it so simple that there are obviously no deficiencies and  the other is to make it so complicated that there are no obvious defiencies",
		"It's not that I'm so smart, it's just that I stay with problems longer.",
		"It is pitch dark. You are likely to be eaten by a grue",
	}
	quote_idx := rand.Int() % len(quotes)
	return quotes[quote_idx], nil
}

func main() {

	// Parse the template file
	tmpl, err := template.ParseFiles("index.html")
	if err != nil {
		log.Fatalf("Error parsing template: %v", err)
	}
	http.Handle("/static/", http.StripPrefix("/static/", http.FileServer(http.Dir("./static"))))
	// Set up the handler
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		quote, err := get_random_quote()
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
		}
		log.Println("Server started at http://localhost:8080")
		data := PageData{
			Message: quote,
		}
		err = tmpl.Execute(w, data)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
		}
	})

	log.Println("Server started at http://localhost:8080")
	err = http.ListenAndServe(":8080", nil)
	if err != nil {
		log.Fatal(err)
	}
}
