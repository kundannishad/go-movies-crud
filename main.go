package main

import (
	"encoding/json"
	"fmt"
	"log"
	"math/rand"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

type Movie struct {
	ID       string    `json:"id"`
	ISBN     string    `json:"isbn"`
	Title    string    `json:"title"`
	Director *Director `json:"director"`
}

type Director struct {
	ID        int32  `json:"id"`
	FirstName string `json:"firstname"`
	LastName  string `json:"lastname"`
}

func getMovies(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(movies)
}

func deleteMovies(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r)

	for index, item := range movies {

		if item.ID == params["id"] {
			movies = append(movies[:index], movies[index+1:]...)

			break
		}

	}

	json.NewEncoder(w).Encode(movies)
}

func getMovie(w http.ResponseWriter, r *http.Request) {

	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r)

	for _, item := range movies {

		if item.ID == params["id"] {
			json.NewEncoder(w).Encode(item)

			return
		}
	}
}

func createMovie(w http.ResponseWriter, r *http.Request) {

	w.Header().Set("Content-Type", "application/json")

	var movie Movie
	_ = json.NewDecoder(r.Body).Decode(&movie)

	movie.ID = strconv.Itoa(rand.Intn(123445))

	movies = append(movies, movie)

	json.NewEncoder(w).Encode(movie)

}

func updateMovies(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-type", "application/json")
	params := mux.Vars(r)

	for index, item := range movies {

		if item.ID == params["id"] {
			movies = append(movies[:index], movies[index+1:]...)

			var movie Movie

			_ = json.NewDecoder(r.Body).Decode(&movie)
			movie.ID = params["id"]
			movies = append(movies, movie)

			json.NewEncoder(w).Encode(movie)
		}
	}
}

var movies []Movie

func main() {
	routes := mux.NewRouter()

	movies = append(movies, Movie{
		ID:       "1",
		ISBN:     "345345",
		Title:    "3 Idiots",
		Director: &Director{ID: 1, FirstName: "Amir", LastName: "Khan"}})

	movies = append(movies, Movie{
		ID:       "2",
		ISBN:     "113",
		Title:    "Sandeep Singh",
		Director: &Director{FirstName: "Jhone", LastName: "Doe"}})

	routes.HandleFunc("/movies", getMovies).Methods("GET")
	routes.HandleFunc("/movies/{id}", getMovie).Methods("GET")
	routes.HandleFunc("/movies", createMovie).Methods("POST")
	routes.HandleFunc("/movies/{id}", updateMovies).Methods("POST")
	routes.HandleFunc("/movies/{id}", deleteMovies).Methods("DELETE")

	fmt.Println("starting server at port 8000")

	log.Fatal(http.ListenAndServe(":8000", routes))

}
