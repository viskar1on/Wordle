package ui

import "os"

type File struct {
	Body *[]byte
}

func NewFile(filename string) *File {
	page := new(File)
	page.Body, _ = LoadFile(&filename)
	return page
}

func LoadFile(filename *string) (*[]byte, error) {
	body, err := os.ReadFile(*filename)

	if err != nil {
		return nil, err
	}

	return &body, nil
}
