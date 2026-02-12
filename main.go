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
	Isbn     string    `json:"isbn"`
	Title    string    `json:"title"`
	Director *Director `json:"director"`
}

type Director struct {
	Firstname string `json:"firstname"`
	Lastname  string `json:"lastname"`
}

var movies []Movie

func getMovies(w http.ResponseWriter, r *http.Request)  {

	w.Header().Set("Content-Type","application/json")
	err := json.NewEncoder(w).Encode(movies)
	if err != nil {
		return
	}
}

func getMovie(w http.ResponseWriter, r *http.Request){

	w.Header().Set("Content-Type","application/json")
	params := mux.Vars(r)

	for _, item := range movies {
		if item.ID == params["id"]{
			err := json.NewEncoder(w).Encode(item)
			if err != nil {
				return
			}
			break
		}
	}
}

func deleteMovie(w http.ResponseWriter, r *http.Request) {

	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r)

	for i, item := range movies {
		if item.ID == params["id"]{
			movies = append(movies [:i], movies[i+1:]...)
			break
		}
	}

	err := json.NewEncoder(w).Encode(movies)
	if err != nil{
		return
	}
}

func creatMovie(w http.ResponseWriter, r *http.Request){

	w.Header().Set("Content-Type","application/json")
	var movie Movie
	_ = json.NewDecoder(r.Body).Decode(&movie)
	movie.ID = strconv.Itoa(rand.Intn(10000000))
	movies = append(movies,movie)
}



func main() {
	r := mux.NewRouter()

	movies = append(movies, Movie{ID: "1", Isbn: "234567", Title: "Jaws", Director: &Director{Firstname: "Steven", Lastname: "Spielberg"}})
	movies = append(movies, Movie{ID: "2", Isbn: "456789", Title: "There will be blood", Director: &Director{Firstname: "Paul-Thomas", Lastname: "Anderson"}})
	movies = append(movies, Movie{ID: "3", Isbn: "678901", Title: "The Godfather", Director: &Director{Firstname: "Francis Ford", Lastname: "Coppola"}})
	movies = append(movies, Movie{ID: "4", Isbn: "890123", Title: "The Dark Knight", Director: &Director{Firstname: "Christopher", Lastname: "Nolan"}})

	r.HandleFunc("/movies", getMovies).Methods("GET")
	r.HandleFunc("/movies:{id}", getMovie).Methods("GET")
	r.HandleFunc("/movies", creatMovie).Methods("POST")
	r.HandleFunc("movies/{id}", updateMovie).Methods("PUT")
	r.HandleFunc("movies/{id}", deleteMovie).Methods("DELETE")

	fmt.Printf("Starting server at port 8080\n")
	log.Fatal(http.ListenAndServe(":8080",r))
}
