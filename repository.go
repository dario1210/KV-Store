package main

import (
	"encoding/json"
	"errors"
	"log"
	"os"
)

type Repository interface {
	GetValueByKey(key string) (interface{}, error)
	CreateKey(data map[string]string) error
}

type FileRepo struct {
	filePath string
}

func NewFileRepository(filePath string) *FileRepo {
	return &FileRepo{filePath: filePath}
}

func (repo *FileRepo) GetValueByKey(key string) (interface{}, error) {
	data, err := readAndUnmarshalFile()
	if err != nil {
		return nil, err
	}

	val, ok := data[key]
	if !ok {
		log.Println("The key does not exist.")
		return val, err
	}

	return val, nil
}

func (repo *FileRepo) CreateKey(requestBody map[string]string) (map[string]interface{}, error) {
	existingData, err := readAndUnmarshalFile()
	if err != nil {
		return nil, err
	}

	for k, v := range requestBody {
		existingData[k] = v
	}

	return existingData, nil
}

func readAndUnmarshalFile() (map[string]interface{}, error) {
	jsonData, _ := os.ReadFile("db.json")

	var data map[string]interface{}

	err := json.Unmarshal(jsonData, &data)
	if err != nil {
		log.Println("Json Unmarshal failed.")
		return nil, errors.New("json Unmarshal Failed")
	}

	return data, nil
}
