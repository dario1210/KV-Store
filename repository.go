package main

import (
	"encoding/json"
	"errors"
	"log"
	"os"
)

type Repository interface {
	GetValueByKey(key string) (string, error)
	CreateKey(data map[string]string) error
}

type Repo struct {
	filePath string
	kvStore  *KeyValueStore
}

func NewRepository(filePath string, kvStore *KeyValueStore) *Repo {
	return &Repo{filePath: filePath, kvStore: kvStore}
}

func (repo *Repo) GetValueByKey(key string) (string, error) {
	repo.kvStore.mu.Lock()
	defer repo.kvStore.mu.Unlock()

	val, ok := repo.kvStore.Key[key]
	if ok {
		return val, nil
	}

	log.Println("Key does not exist on the map, checking file.")

	// If the key is not found in the in-memory map, check the file
	data, err := readAndUnmarshalFile()
	if err != nil {
		return "", err
	}

	val, ok = data[key]
	if !ok {
		return "", errors.New("key does not exist")
	}

	return val, nil
}

func (repo *Repo) CreateKey(requestBody map[string]string) (map[string]string, error) {
	repo.kvStore.mu.Lock()
	defer repo.kvStore.mu.Unlock()

	for k, v := range requestBody {
		repo.kvStore.Key[k] = v
	}

	return repo.kvStore.Key, nil
}

func readAndUnmarshalFile() (map[string]string, error) {
	jsonData, err := os.ReadFile("db.json")
	if err != nil {
		log.Fatal("Cannot load db.json")
	}

	var data map[string]string

	err = json.Unmarshal(jsonData, &data)
	if err != nil {
		log.Println("Json Unmarshal failed.")
		return nil, errors.New("json Unmarshal Failed")
	}

	return data, nil
}

func PersistData(data map[string]string) error {

	existingData, err := readAndUnmarshalFile()
	if err != nil {
		log.Println("Failed to read file.")
		return err
	}

	for k, v := range data {
		existingData[k] = v
	}

	updatedJSON, err := json.MarshalIndent(existingData, "", "  ")
	if err != nil {
		log.Println("Failed to serialize updated data")
		return err
	}

	err = os.WriteFile(repo.filePath, updatedJSON, 0644)
	if err != nil {
		log.Println("Data did not get saved.")
		return err
	}

	return nil
}
