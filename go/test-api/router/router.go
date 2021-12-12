package router

import (
	"test-api/controller"

	"github.com/gorilla/mux"
)


func Router() *mux.Router{
	router := mux.NewRouter()
	router.HandleFunc("/",controller.FakeApi).Methods("GET")
	router.HandleFunc("/get/{id}",controller.GetOneMovie).Methods("GET")
	router.HandleFunc("/getAll",controller.GetAllMovies).Methods("GET")
	router.HandleFunc("/insert",controller.InsertOneMovie).Methods("POST")
	router.HandleFunc("/update/{id}",controller.UpdateOneMovie).Methods("PUT")
	router.HandleFunc("/deleteOne/{id}",controller.DeleteOneMovie).Methods("DELETE")
	router.HandleFunc("/deleteAll",controller.DeleteAllMovies).Methods("DELETE")
	return router
}