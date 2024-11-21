package main

import "github.com/go-chi/chi/v5"

func RegisterRoutes(r chi.Router) {
	r.Get("/get", GetValueByKey)
	r.Put("/create", CreateKey)
}
