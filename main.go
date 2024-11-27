package main

import (
	"fmt"
	"net/http"

	"github.com/go-chi/chi/v5"
)

var repo *Repo

func main() {
	kvStore := NewKeyValueStore()
	repo = NewRepository("db.json", kvStore)

	r := chi.NewRouter()
	RegisterRoutes(r)

	apiServer := &http.Server{
		Handler: r,
		Addr:    ":8080",
	}

	fmt.Println("Starting server on :8080")
	err := apiServer.ListenAndServe()
	if err != nil {
		fmt.Printf("Error starting server: %v\n", err)
	}

}
