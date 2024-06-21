package logic

import (
	"errors"
	"strings"
)

const (
	letterStateRight = iota
	letterStateExists
	letterStateAbsent
)

type WordleGame struct {
	storage  Storage
	word     string
	attempts byte
}

func NewWordle(storage *Storage) *WordleGame {
	wordle := new(WordleGame)
	wordle.storage = *storage
	wordle.Regenerate()
	return wordle
}

func (wordle *WordleGame) Attempt(attemptWord *string) (*[]int, error) {
	if wordle.attempts == 0 {
		return nil, errors.New("попытки закончились")
	}

	if len(*attemptWord) != len(wordle.word) {
		return nil, errors.New("длина слова неправильная")
	}

	if !Contains(*wordle.storage.Words(), *attemptWord) {
		return nil, errors.New("слова нет в словаре")
	}

	result := make([]int, 6)

	for i := 0; i < 6; i++ {
		aw := []rune(*attemptWord)
		ww := []rune(wordle.word)
		if ww[i] == aw[i] {
			result[i] = letterStateRight
		} else if strings.Contains(string(ww), string((aw)[i])) {
			result[i] = letterStateExists
		} else {
			result[i] = letterStateAbsent
		}
	}

	wordle.attempts--
	return &result, nil
}

func (wordle *WordleGame) Regenerate() {
	wordle.word = *GenerateWord(wordle.storage)
	wordle.attempts = 5
}

func Contains(s []string, e string) bool {
	for _, a := range s {
		if a == e {
			return true
		}
	}
	return false
}
