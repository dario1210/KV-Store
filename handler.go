package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"os"
)

type KeyValueStore struct {
	Key map[string]string `json:"key"`
}

func GetValueByKey(w http.ResponseWriter, r *http.Request) {
	key := r.URL.Query().Get("key")

	val, err := repo.GetValueByKey(key)
	if err != nil {
		log.Println("Key does not exist.")
		return
	}

	w.WriteHeader(http.StatusOK)
	fmt.Fprintln(w, val)
}

func CreateKey(w http.ResponseWriter, r *http.Request) {
	var KeyValueStore KeyValueStore
	err := json.NewDecoder(r.Body).Decode(&KeyValueStore)
	if err != nil {
		http.Error(w, "Failed to decode request body", http.StatusBadRequest)
		return
	}

	data, err := repo.CreateKey(KeyValueStore.Key)
	if err != nil {
		log.Println("Cannot add key sorry best.")
	}

	updatedJSON, err := json.MarshalIndent(data, "", "  ")
	if err != nil {
		http.Error(w, "Failed to serialize updated data", http.StatusInternalServerError)
		return
	}

	err = os.WriteFile("db.json", updatedJSON, 0644)
	if err != nil {
		http.Error(w, "Failed to write updated data to the database file", http.StatusInternalServerError)
		return
	}

	fmt.Fprintln(w, "File Updated succesfully.")
}
