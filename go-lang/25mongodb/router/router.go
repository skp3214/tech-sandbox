package router

import (
	"mongoskp3214/controller"

	"github.com/gorilla/mux"
)

func Router() *mux.Router{
	router:=mux.NewRouter()

    router.HandleFunc("/api/movies",controller.GetMyAllMovies).Methods("GET")
	router.HandleFunc("/api/movie",controller.CreateMovie).Methods("POST")
	router.HandleFunc("/api/movie/{id}",controller.MarksAsWatched).Methods("PUT")
	router.HandleFunc("/api/movie/{id}",controller.DeleteOneMovie).Methods("DELETE")
	router.HandleFunc("/api/deleteall",controller.DeleteAllMovie).Methods("DELETE")

	return router
}