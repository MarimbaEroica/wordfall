package messages

type TilePosition struct {
	Col int `json:"col"`
	Row int `json:"row"`
}

type Message struct {
	Type    string      `json:"type"`
	Payload interface{} `json:"payload"`
}

type BoardUpdate struct {
	Board    [][]string `json:"board"`
	Score    int        `json:"score"`
	TimeLeft int        `json:"timeLeft"`
}

type TimeUpdate struct {
	TimeLeft int `json:"timeLeft"`
}

type WordSubmission struct {
	SelectedTiles []TilePosition `json:"selectedTiles"`
}

type ManualRemoval struct {
	SelectedTiles []TilePosition `json:"selectedTiles"`
}
