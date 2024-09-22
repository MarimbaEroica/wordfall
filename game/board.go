package game

import (
	"fmt"
	"math"
	"math/rand"
	"sort"

	"github.com/MarimbaEroica/wordfall/messages"
)

type Board struct {
	Columns [][]string
}

const (
	NumColumns   = 5
	ColumnHeight = 7 // 5 visible + 2 hidden
)

func NewBoard() *Board {
	b := &Board{
		Columns: make([][]string, NumColumns),
	}
	for i := 0; i < NumColumns; i++ {
		b.Columns[i] = GenerateColumn()
	}
	return b
}

func GenerateColumn() []string {
	column := make([]string, 1000)
	letterBag := GenerateLetterBag()
	for i := 0; i < 1000; i++ {
		idx := rand.Intn(len(letterBag))
		column[i] = letterBag[idx]
	}
	return column
}

func (b *Board) GetVisibleTiles() [][]string {
	visibleTiles := make([][]string, NumColumns)
	for i := 0; i < NumColumns; i++ {
		columnLen := len(b.Columns[i])
		startIdx := columnLen - ColumnHeight
		if startIdx < 0 {
			startIdx = 0
		}
		visibleTiles[i] = b.Columns[i][startIdx:columnLen]
	}
	return visibleTiles
}

func (b *Board) GetTile(col, row int) string {
	columnLen := len(b.Columns[col])
	tileIdx := columnLen - row - 1
	if tileIdx >= 0 && tileIdx < len(b.Columns[col]) {
		return b.Columns[col][tileIdx]
	}
	return ""
}

func (b *Board) RemoveTiles(selectedTiles []messages.TilePosition) {
	// Group tiles by column
	tilesByColumn := make(map[int][]int)
	for _, tilePos := range selectedTiles {
		tilesByColumn[tilePos.Col] = append(tilesByColumn[tilePos.Col], tilePos.Row)
	}

	for col, rows := range tilesByColumn {
		// Sort rows in ascending order
		sort.Ints(rows)
		for i, row := range rows {
			columnLen := len(b.Columns[col])
			tileIdx := columnLen - row - 1 - i
			if tileIdx >= 0 && tileIdx < len(b.Columns[col]) {
				b.Columns[col] = append(b.Columns[col][:tileIdx], b.Columns[col][tileIdx+1:]...)
			}
		}
	}
}

func (b *Board) ValidatePath(selectedTiles []messages.TilePosition) bool {
	if len(selectedTiles) == 0 {
		return false
	}

	// Create a map to keep track of visited positions to prevent reusing the same tile
	visited := make(map[string]bool)
	for _, tile := range selectedTiles {
		key := positionKey(tile.Col, tile.Row)
		if visited[key] {
			// Cannot reuse the same tile
			return false
		}
		visited[key] = true
	}

	// Start from the first tile and check adjacency
	for i := 1; i < len(selectedTiles); i++ {
		prev := selectedTiles[i-1]
		curr := selectedTiles[i]
		if !areAdjacent(prev, curr) {
			return false
		}
	}

	return true
}

func areAdjacent(a, b messages.TilePosition) bool {
	colDiff := math.Abs(float64(a.Col - b.Col))
	rowDiff := math.Abs(float64(a.Row - b.Row))
	return colDiff <= 1 && rowDiff <= 1 && !(colDiff == 0 && rowDiff == 0)
}

func positionKey(col, row int) string {
	return fmt.Sprintf("%d,%d", col, row)
}
