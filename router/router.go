package router

import (
	"go-fiber/controller"

	"github.com/gorilla/mux"
)

func Router() *mux.Router {
	router := mux.NewRouter()
	router.HandleFunc("/movies", controller.GetAllMovies).Methods("GET")
	router.HandleFunc("/movie", controller.CreateMovies).Methods("POST")
	router.HandleFunc("/movie/{movieId}", controller.MarkedAsWatched).Methods("PUT")
	router.HandleFunc("/movie/{movieId}", controller.DeleteOne).Methods("DELETE")
	router.HandleFunc("/movie", controller.DeleteAll).Methods("DELETE")
	return router
}
