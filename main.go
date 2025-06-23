package main

import(
	"fmt"
	"net/http"
	"log"
	"encoding/json"
	"math/rand"
	"github.com/gorilla/mux"
	"strconv"
)

type movie struct{
	ID	 string  `json:"id"`		// ID of the movie
	Isbn string  `json:"isbn"`	// ISBN of the movie
	Title string  `json:"title"`	// Title of the movie
	Director *Director `json:"director"` // Director of the movie
}

type Director struct{
	Firstname string `json:"firstname"` // First name of the director
	Lastname  string `json:"lastname"`  // Last name of the director
}

var movies []movie

func main(){
	r:= mux.NewRouter()
	// Initializing some movies
	movies = append(movies, movie{ID: "1", Isbn: "438743", Title: "Movie One", Director: &Director{Firstname: "John", Lastname: "Doe"}})
	movies = append(movies, movie{ID: "2", Isbn: "438744", Title: "Movie Two", Director: &Director{Firstname: "Jane", Lastname: "Doe"}})

	r.HandleFunc("/movies", getMovies).Methods("GET")
	r.HandleFunc("/movies/{id}", getMovie).Methods("GET")
	r.HandleFunc("/movies", createMovie).Methods("POST")
	r.HandleFunc("/movies/{id}", updateMovie).Methods("PUT")
	r.HandleFunc("/movies/{id}", deleteMovie).Methods("DELETE")

	fmt.Println("Starting server on port 8000")
	log.Fatal(http.ListenAndServe(":8000", r))

}

func getMovies(w http.ResponseWriter, r *http.Request){
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(movies)
}

func getMovie(w http.ResponseWriter, r *http.Request){
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r)
	for _, item := range movies {
		if item.ID == params["id"] {
			json.NewEncoder(w).Encode(item)
			return
		}
	}
	http.Error(w, "Movie not found", http.StatusNotFound)
}

func createMovie(w http.ResponseWriter, r *http.Request){
	w.Header().Set("Content-Type", "application/json")
	var movie movie
	_ = json.NewDecoder(r.Body).Decode(&movie)
	movie.ID = strconv.Itoa(rand.Intn(1000000)) // Mock ID - not safe for production
	movies = append(movies, movie)
	json.NewEncoder(w).Encode(movie)
}

func updateMovie(w http.ResponseWriter, r *http.Request){
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r)
	for index, item := range movies {
		if item.ID == params["id"] {
			movies = append(movies[:index], movies[index+1:]...) // Remove the movie
			var movie movie
			_ = json.NewDecoder(r.Body).Decode(&movie)
			movie.ID = item.ID // Keep the same ID
			movies = append(movies, movie) // Add the updated movie
			json.NewEncoder(w).Encode(movie)
			return
		}
	}
	http.Error(w, "Movie not found", http.StatusNotFound)
}

func deleteMovie(w http.ResponseWriter, r *http.Request){
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r)
	for index, item := range movies {
		if item.ID == params["id"] {
			movies = append(movies[:index], movies[index+1:]...) // Remove the movie
			break
		}
	}
	json.NewEncoder(w).Encode(movies)
}