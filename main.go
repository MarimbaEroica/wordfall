package main

import (
	"log"
	"net/http"

	"github.com/MarimbaEroica/wordfall/handlers"
)

func main() {
	http.HandleFunc("/ws", handlers.HandleWebSocket)
	fs := http.FileServer(http.Dir("./static"))
	http.Handle("/", fs)

	log.Println("Server started on :8080")
	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		log.Fatal("ListenAndServe: ", err)
	}
}
