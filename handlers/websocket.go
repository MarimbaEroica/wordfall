package handlers

import (
	"encoding/json"
	"log"
	"net/http"
	"wordfall/game"
	"wordfall/messages"

	"github.com/gorilla/websocket"
)

var upgrader = websocket.Upgrader{}

func HandleWebSocket(w http.ResponseWriter, r *http.Request) {
	// Upgrade initial GET request to a websocket
	conn, err := upgrader.Upgrade(w, r, nil)
	if err != nil {
		log.Println("Upgrade error:", err)
		return
	}
	defer conn.Close()

	g := game.NewGame()
	g.StartTimer()

	// Send initial board state to client
	sendBoardUpdate(conn, g)

	for {
		_, msgBytes, err := conn.ReadMessage()
		if err != nil {
			log.Println("Read error:", err)
			break
		}

		var msg messages.Message
		err = json.Unmarshal(msgBytes, &msg)
		if err != nil {
			log.Println("Unmarshal error:", err)
			continue
		}

		switch msg.Type {
		case "wordSubmission":
			handleWordSubmission(conn, g, msg.Payload)
		case "manualRemoval":
			handleManualRemoval(conn, g, msg.Payload)
		}
	}
}

func sendBoardUpdate(conn *websocket.Conn, g *game.Game) {
	visibleBoard := g.Board.GetVisibleTiles()
	boardUpdate := messages.BoardUpdate{
		Board:    visibleBoard,
		Score:    g.Score,
		TimeLeft: g.TimeLeft,
	}
	msg := messages.Message{
		Type:    "boardUpdate",
		Payload: boardUpdate,
	}
	msgBytes, _ := json.Marshal(msg)
	conn.WriteMessage(websocket.TextMessage, msgBytes)
}

func handleWordSubmission(conn *websocket.Conn, g *game.Game, payload interface{}) {
	data, _ := json.Marshal(payload)
	var submission messages.WordSubmission
	json.Unmarshal(data, &submission)

	valid, points := g.ValidateWord(submission.SelectedTiles)
	if valid {
		g.Board.RemoveTiles(submission.SelectedTiles)
		g.Score += points
		sendBoardUpdate(conn, g)
	} else {
		// Send invalid word message
		msg := messages.Message{
			Type:    "invalidWord",
			Payload: nil,
		}
		msgBytes, _ := json.Marshal(msg)
		conn.WriteMessage(websocket.TextMessage, msgBytes)
	}
}

func handleManualRemoval(conn *websocket.Conn, g *game.Game, payload interface{}) {
	data, _ := json.Marshal(payload)
	var removal messages.ManualRemoval
	json.Unmarshal(data, &removal)

	g.Board.RemoveTiles(removal.SelectedTiles)
	sendBoardUpdate(conn, g)
}
