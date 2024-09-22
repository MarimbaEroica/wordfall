package main

import (
	"fmt"
	"net/http"

	"github.com/MarimbaEroica/wordfall/game"
)

func main() {
	http.HandleFunc("/game", game.HandleGame)
	fmt.Println("Starting server on :8080")
	http.ListenAndServe(":8080", nil)
}
