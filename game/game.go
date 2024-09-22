package game

import (
	"time"
	"wordfall/messages"
)

type Game struct {
	Board      *Board
	Score      int
	TimeLeft   int
	Dictionary map[string]bool
}

func NewGame() *Game {
	g := &Game{
		Board:      NewBoard(),
		Score:      0,
		TimeLeft:   180,
		Dictionary: LoadDictionary(),
	}
	return g
}

func (g *Game) StartTimer() {
	ticker := time.NewTicker(1 * time.Second)
	go func() {
		for range ticker.C {
			g.TimeLeft--
			if g.TimeLeft <= 0 {
				ticker.Stop()
				// Handle game over if necessary
			}
		}
	}()
}

func (g *Game) ValidateWord(selectedTiles []messages.TilePosition) (bool, int) {
	word := ""
	for _, tilePos := range selectedTiles {
		tile := g.Board.GetTile(tilePos.Col, tilePos.Row)
		word += tile
	}

	// Check if the word exists
	if !g.Dictionary[word] {
		return false, 0
	}

	// Check if tiles are adjacent
	if !g.Board.ValidatePath(selectedTiles) {
		return false, 0
	}

	// Calculate score
	points := CalculateScore(len(word))
	return true, points
}

func CalculateScore(length int) int {
	switch length {
	case 4:
		return 1
	case 5:
		return 2
	case 6:
		return 4
	case 7:
		return 7
	default:
		if length >= 8 {
			return 5*(length-8) + 12
		}
	}
	return 0
}
