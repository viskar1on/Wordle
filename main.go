package main

import (
	"main/logic"
	"main/ui"
)

func main() {
	storage := logic.NewStorage("db.txt")
	logic := logic.NewWordle(storage)
	web := ui.NewWordleUi(logic)
	web.Start()
}
