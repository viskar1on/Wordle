package logic

import (
	"os"
	"strings"
)

type Storage struct {
	words *[]string
}

func NewStorage(filePath string) *Storage {
	storage := new(Storage)
	storage.words = LoadWordsFromFile(filePath)
	return storage
}

func (storage Storage) Words() *[]string {
	return storage.words
}

func LoadWordsFromFile(filePath string) *[]string {
	var words []string
	file, err := os.ReadFile(filePath)

	if err != nil {
		return &words
	}

	data := string(file)
	words = append(words, strings.Split(data, "\n")...)

	for i := range words {
		words[i] = strings.TrimRight(words[i], string(rune(13)))
	}

	return &words
}
