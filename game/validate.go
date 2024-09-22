package game

import (
	"sort"
	"strings"
)

// Assume that `dictionary` is a sorted slice of strings
var dictionary = []string{"apple", "banana", "orange", "quilt", "hero", "quick"}

func ValidateInput(path [][2]int, board *Board) bool {
	if !isValidPath(path) {
		return false
	}

	word := formWord(path, board)

	word = strings.ToLower(word)
	index := sort.SearchStrings(dictionary, word)

	return index < len(dictionary) && dictionary[index] == word
}

// Helper function to check adjacency and ensure each tile is used only once
func isValidPath(path [][2]int) bool {
	visited := make(map[[2]int]bool)

	for i, pos := range path {
		if visited[pos] {
			return false
		}
		visited[pos] = true

		if i == 0 {
			continue
		}

		dx := abs(path[i][0] - path[i-1][0])
		dy := abs(path[i][1] - path[i-1][1])
		if dx > 1 || dy > 1 || (dx == 0 && dy == 0) {
			return false
		}
	}

	return true
}

// Helper function to form a word from the path
func formWord(path [][2]int, board *Board) string {
	var word strings.Builder
	for _, pos := range path {
		word.WriteString(board.Columns[pos[0]].Tiles[pos[1]].Letter)
	}
	return word.String()
}

func abs(a int) int {
	if a < 0 {
		return -a
	}
	return a
}

// Scoring function based on word length
func CalculateScore(word string) int {
	length := len(word)
	switch {
	case length == 4:
		return 1
	case length == 5:
		return 2
	case length == 6:
		return 4
	case length == 7:
		return 7
	case length >= 8:
		return 5*(length-8) + 12
	default:
		return 0
	}
}
