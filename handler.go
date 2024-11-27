package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"sync"
)

type KeyValueStore struct {
	Key map[string]string `json:"key"`
	mu  sync.Mutex
}

var requestCount int

const (
	limit = 10
)

func NewKeyValueStore() *KeyValueStore {
	return &KeyValueStore{
		Key: make(map[string]string),
		mu:  sync.Mutex{},
	}
}

func GetValueByKey(w http.ResponseWriter, r *http.Request) {
	key := r.URL.Query().Get("key")

	val, err := repo.GetValueByKey(key)
	if err != nil {
		http.Error(w, "That key does not exist.", http.StatusNotFound)
		return
	}

	w.WriteHeader(http.StatusOK)
	fmt.Fprintln(w, val)
}

func CreateKey(w http.ResponseWriter, r *http.Request) {
	var requestBody KeyValueStore

	err := json.NewDecoder(r.Body).Decode(&requestBody)
	if err != nil {
		http.Error(w, "Failed to decode request body", http.StatusBadRequest)
		return
	}

	data, err := repo.CreateKey(requestBody.Key)
	if err != nil {
		http.Error(w, "Failed to add key.", http.StatusInternalServerError)
		return
	}

	requestCount++

	// Only persist data when the request count reaches the limit
	if requestCount >= limit {
		err = PersistData(data)
		if err != nil {
			http.Error(w, "Data save in db failed.", http.StatusInternalServerError)
			return
		}

		requestCount = 0
	}

	fmt.Fprintln(w, "File updated successfully.")
}
