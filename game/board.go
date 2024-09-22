package game

type Tile struct {
	Letter string `json:"letter"`
}

type Column struct {
	Tiles []Tile `json:"tiles"`
}

type Board struct {
	Columns [5]Column `json:"columns"`
}

// Initialize the board with tiles
func InitializeBoard() *Board {
	var board Board
	for i := 0; i < 5; i++ {
		board.Columns[i] = Column{Tiles: drawTiles(1000)}
	}
	return &board
}

func drawTiles(n int) []Tile {
	// Define the tile distribution here
	// This is a simple example; you should modify it to match Boggle's distribution
	possibleTiles := []string{"A", "E", "I", "O", "U", "QU", "HE", "R", "S", "T", "N"}
	tiles := make([]Tile, n)
	for i := 0; i < n; i++ {
		tiles[i] = Tile{Letter: possibleTiles[i%len(possibleTiles)]}
	}
	return tiles
}

// Helper function to prepare board data for the frontend
func PrepareBoardData(board *Board) [][]Tile {
	data := make([][]Tile, 5)

	for i := 0; i < 5; i++ {
		column := board.Columns[i]
		data[i] = column.Tiles[:7] // Send the top 7 tiles, 5 visible and 2 in reserve
	}

	return data
}

// Updates board by removing the tiles in path and making the tiles from the columns fall into place
func UpdateBoard(board *Board, path [][2]int) {
	tilesToRemove := make(map[[2]int]bool)
	for _, pos := range path {
		tilesToRemove[pos] = true
	}

	for _, pos := range path {
		col := pos[0]
		row := pos[1]
		board.Columns[col].Tiles = removeTile(board.Columns[col].Tiles, row)
	}
}

// Helper function to remove a tile from a column and let the above tiles fall
func removeTile(tiles []Tile, row int) []Tile {
	return append(tiles[:row], tiles[row+1:]...)
}
