package game

import (
	"encoding/json"
	"fmt"
	"net/http"
	"sync"
	"time"

	"github.com/gorilla/websocket"
)

var upgrader = websocket.Upgrader{
	CheckOrigin: func(r *http.Request) bool {
		return true
	},
}

type Game struct {
	Board *Board
	Score int
	Timer *time.Timer
	Mutex sync.Mutex
}

func HandleGame(w http.ResponseWriter, r *http.Request) {
	conn, err := upgrader.Upgrade(w, r, nil)
	if err != nil {
		http.Error(w, "Could not upgrade to WebSocket", http.StatusInternalServerError)
		return
	}
	defer conn.Close()

	game := &Game{
		Board: InitializeBoard(),
		Score: 0,
		Timer: time.NewTimer(180 * time.Second),
	}

	go func() {
		<-game.Timer.C
		conn.WriteMessage(websocket.TextMessage, []byte("Game over! Your final score is: "+fmt.Sprint(game.Score)))
		conn.Close()
	}()

	for {
		_, message, err := conn.ReadMessage()
		if err != nil {
			return
		}

		var command map[string]interface{}
		err = json.Unmarshal(message, &command)
		if err != nil {
			continue
		}

		game.Mutex.Lock()

		switch command["action"] {
		case "submit_word":
			path := parsePath(command["path"].([]interface{}))
			if ValidateInput(path, game.Board) {
				word := formWord(path, game.Board)
				game.Score += CalculateScore(word)
				UpdateBoard(game.Board, path)
				conn.WriteMessage(websocket.TextMessage, []byte("Score: "+fmt.Sprint(game.Score)))
			} else {
				conn.WriteMessage(websocket.TextMessage, []byte("Invalid word"))
			}
		case "manual_remove":
			path := parsePath(command["path"].([]interface{}))
			UpdateBoard(game.Board, path)
			conn.WriteMessage(websocket.TextMessage, []byte("Tiles removed"))
		case "get_board":
			data := PrepareBoardData(game.Board)
			conn.WriteJSON(data)
		}

		game.Mutex.Unlock()
	}
}

// Helper function to parse path from the command
func parsePath(rawPath []interface{}) [][2]int {
	path := make([][2]int, len(rawPath))
	for i, pos := range rawPath {
		position := pos.([]interface{})
		path[i][0] = int(position[0].(float64))
		path[i][1] = int(position[1].(float64))
	}
	return path
}
