package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

// request parameter
type Movie struct {
	ID       string    `json:id`
	Title    string    `json:title`
	Director *Director `json:director`
}

type Director struct {
	Firstname string `json:firstname`
	Lastname  string `json:lastname`
}

var movies []Movie

func HomeHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hello!\n")
}

// Display all movies
func getMovies(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	err := json.NewEncoder(w).Encode(movies)
	if err != nil {
		http.Error(w, err.Error(), 500)
		return
	}
}

// Display movie matching with id
func getMovie(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	params := mux.Vars(r)
	for _, val := range movies {
		if val.ID == params["id"] {
			res := []Movie{val}
			err := json.NewEncoder(w).Encode(res)
			if err != nil {
				http.Error(w, err.Error(), 500)
				return
			}
			return
		}
	}
	// if id not found
	http.NotFound(w, r)
}

// Delete movies by id
func deleteMovie(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	params := mux.Vars(r)
	for key, val := range movies {
		if val.ID == params["id"] {
			res := []Movie{val}
			res = append(res[:key], res[key+1:]...)
			err := json.NewEncoder(w).Encode((res))
			if err != nil {
				http.Error(w, err.Error(), 500)
				return
			}
			return
		}
	}
	http.NotFound(w, r)
}

func main() {
	fmt.Println("Movies crud api")

	r := mux.NewRouter()

	movies = append(movies, Movie{ID: "1", Title: "Yeh Jawaani Hai Deewani", Director: &Director{Firstname: "Ayan", Lastname: "Mukerji"}})
	movies = append(movies, Movie{ID: "2", Title: "Sanju", Director: &Director{Firstname: "Rajkumar", Lastname: "Hirani"}})
	movies = append(movies, Movie{ID: "3", Title: " Barfi!", Director: &Director{Firstname: "Anurag", Lastname: "Basu"}})

	r.HandleFunc("/", HomeHandler)
	r.HandleFunc("/movies", getMovies).Methods("GET")
	r.HandleFunc("/movies/{id}", getMovie).Methods("GET")
	r.HandleFunc("/delete/{id}", deleteMovie).Methods("DELETE")

	if err := http.ListenAndServe(":8080", r); err != nil {
		log.Fatal("Srever error", err)
	}
}
