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
	fmt.Fprintf(w, "Movies crud api!\n")
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

	// fmt.Fprintf(w, "Deleted the movie by id")
	// if id not found
	http.NotFound(w, r)
}

// create movie
func createMovie(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	// Parse the request body
	var newMovie Movie
	err := json.NewDecoder(r.Body).Decode(&newMovie)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
	}

	// Close the request body
	defer r.Body.Close()

	// Add movie
	movies = append(movies, newMovie)

	// return the new movie
	err = json.NewEncoder(w).Encode(movies)
	if err != nil {
		http.Error(w, err.Error(), 500)
		return
	}

	// fmt.Fprintf(w, "Create a new movie")
}

// Delete movies by id
func deleteMovie(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	params := mux.Vars(r)
	for key, val := range movies {
		if val.ID == params["id"] {
			movies = append(movies[:key], movies[key+1:]...)
			err := json.NewEncoder(w).Encode((movies))
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
	r.HandleFunc("/movies", getMovies).Methods("GET")           // get all movies
	r.HandleFunc("/movies/{id}", getMovie).Methods("GET")       // get movie by id
	r.HandleFunc("/delete/{id}", deleteMovie).Methods("DELETE") // delete movie by id
	r.HandleFunc("/create", createMovie).Methods("PUT")         // create movie

	if err := http.ListenAndServe(":8080", r); err != nil {
		log.Fatal("Srever error", err)
	}
}
