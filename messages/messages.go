package messages

type TilePosition struct {
	Col int `json:"col"`
	Row int `json:"row"`
}

type Message struct {
	Type    string      `json:"Type"`
	Payload interface{} `json:"Payload"`
}

type BoardUpdate struct {
	Board    [][]string `json:"Board"`
	Score    int        `json:"Score"`
	TimeLeft int        `json:"TimeLeft"`
}

type WordSubmission struct {
	SelectedTiles []TilePosition `json:"SelectedTiles"`
}

type ManualRemoval struct {
	SelectedTiles []TilePosition `json:"SelectedTiles"`
}
