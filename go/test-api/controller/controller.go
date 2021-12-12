package controller

import (
	"encoding/json"
	"log"
	"net/http"
	"test-api/helper"
	"test-api/model"

	"github.com/gorilla/mux"
)

func FakeApi(w http.ResponseWriter,r *http.Request){
	w.Header().Set("Content-Type","application/json")
	json.NewEncoder(w).Encode("Hello World")
}
func GetAllMovies(w http.ResponseWriter,r *http.Request) {
	w.Header().Set("Content-Type","application/json")
	allMovies := helper.GetAllMovies()
	json.NewEncoder(w).Encode(allMovies)
}

func GetOneMovie(w http.ResponseWriter,r *http.Request) {
	w.Header().Set("Content-Type","application/json")
	params := mux.Vars(r)
	movie := helper.GetOneMovie(params["id"])
	json.NewEncoder(w).Encode(movie)

}

func InsertOneMovie(w http.ResponseWriter,r *http.Request) {
	w.Header().Set("Content-Type","application/json")
	var movie model.Netflix
	err := json.NewDecoder(r.Body).Decode(&movie)
	if err!=nil{
		log.Fatal(err)
	}
	data := helper.InsertOneMovie(movie)
	json.NewEncoder(w).Encode(data)
}

func UpdateOneMovie(w http.ResponseWriter,r *http.Request) {
	w.Header().Set("Content-Type","application/json")
	params := mux.Vars(r)
	count := helper.UpdateOneMovie(params["id"])
	json.NewEncoder(w).Encode(count)
}

func DeleteOneMovie(w http.ResponseWriter,r *http.Request) {
	w.Header().Set("Content-Type","application/json")
	params := mux.Vars(r)
	count := helper.DeleteOneMovie(params["id"])
	json.NewEncoder(w).Encode(count)
}

func DeleteAllMovies(w http.ResponseWriter,r *http.Request) {
	w.Header().Set("Content-Type","application/json")
	count := helper.DeleteAllMovies()
	json.NewEncoder(w).Encode(count)
}
