package main

import (
	"log"
	"net/http"

	"github.com/auenc/simple-rest/endpoints"
	"github.com/auenc/simple-rest/repositories"
	"github.com/auenc/simple-rest/services"
	"github.com/gorilla/mux"
)

func main() {
	inMemoryRepo := repositories.NewInMemoryJobRepository()
	jobService := &services.JobService{JobRepo: inMemoryRepo}
	jobAPI := endpoints.NewJobREST(jobService)

	router := mux.NewRouter()

	router.HandleFunc("/", jobAPI.Create).Methods("POST")
	router.HandleFunc("/", jobAPI.GetAll).Methods("GET")
	router.HandleFunc("/{id}", jobAPI.Get).Methods("GET")
	router.HandleFunc("/{id}", jobAPI.Update).Methods("PUT")
	router.HandleFunc("/{id}", jobAPI.Delete).Methods("DELETE")

	log.Fatalf("%s\n", http.ListenAndServe(":8080", router))
}
