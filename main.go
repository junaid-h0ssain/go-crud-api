package main

import (
	"encoding/json"
	"fmt"
	"github.com/gorilla/mux"
	"log"
	"math/rand"
	"net/http"
	"strconv"
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


// changes to check git remote

func main() {
	r := mux.NewRouter()

	r.HandleFunc("/movies", getMovies).Methods("GET")
	r.HandleFunc("/movies:{id}", getMovie).Methods("GET")
	r.HandleFunc("/movies", creatMovie).Methods("POST")
	r.HandleFunc("movies/{id}", updateMovie).Methods("PUT")
	r.HandleFunc("movies/{id}", deleteMovie).Methods("DELETE")

	fmt.Printf("Starting server at port 8080\n")
	log.Fatal(http.ListenAndServe(":8080",r))
}
