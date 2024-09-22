package game

func GenerateLetterBag() []string {
	frequencyDistribution := map[string]int{
		"A": 9, "B": 2, "C": 2, "D": 4, "E": 12,
		"F": 2, "G": 3, "H": 2, "I": 9, "J": 1,
		"K": 1, "L": 4, "M": 2, "N": 6, "O": 8,
		"P": 2, "QU": 1, "R": 6, "S": 4, "T": 6,
		"U": 4, "V": 2, "W": 2, "X": 1, "Y": 2,
		"Z": 1, "ER": 1, "HE": 1,
	}

	letterBag := []string{}
	for letter, freq := range frequencyDistribution {
		for i := 0; i < freq; i++ {
			letterBag = append(letterBag, letter)
		}
	}
	return letterBag
}
