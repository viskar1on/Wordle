package ui

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"main/logic"
	"net/http"
)

type WordleUi struct {
	logic logic.WordleGame
}

func NewWordleUi(logic *logic.WordleGame) *WordleUi {
	ui := new(WordleUi)
	ui.logic = *logic
	return ui
}

func WriteFile(w *http.ResponseWriter, filename string) {
	file := NewFile(filename)
	fmt.Fprint(*w, string(*file.Body))
}

func GameHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/html")
	WriteFile(&w, "resources/index.html")
}

func ScriptHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/javascript")
	WriteFile(&w, "resources/script.js")
}

func (ui *WordleUi) RegenerateHandler(w http.ResponseWriter, r *http.Request) {
	ui.logic.Regenerate()
}

func (ui *WordleUi) TryHandler(w http.ResponseWriter, r *http.Request) {
	buff, _ := io.ReadAll(r.Body)
	attempt := string(buff)
	resultMap, err := ui.logic.Attempt(&attempt)

	if err != nil {
		fmt.Fprintln(w, err)
		return
	}

	jsonString, _ := json.Marshal(resultMap)
	fmt.Fprint(w, string(jsonString))
}

func (ui *WordleUi) Start() {
	http.HandleFunc("/", GameHandler)
	http.HandleFunc("/script.js", ScriptHandler)
	http.HandleFunc("/regenerate", ui.RegenerateHandler)
	http.HandleFunc("/try", ui.TryHandler)
	log.Fatal(http.ListenAndServe(":8080", nil))
}
