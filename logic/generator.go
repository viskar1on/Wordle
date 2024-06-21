package logic

import (
	"math/rand"
)

func GenerateWord(storage Storage) *string {
	words := *storage.Words()
	return &words[rand.Intn(len(words))]
}
