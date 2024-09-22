package game

import (
	"bufio"
	"log"
	"os"
	"strings"
)

func LoadDictionary() map[string]bool {
	dict := make(map[string]bool)
	file, err := os.Open("static/words_alpha.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		word := strings.ToUpper(scanner.Text())
		dict[word] = true
	}

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}
	return dict
}
