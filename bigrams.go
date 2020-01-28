package nlputils

// ToBigrams will convert a string into a series of bigram values
func ToBigrams(str string) (bigrams []string) {
	words := ToWords(str)
	end := len(words) - 1
	for i, word := range words {
		var bigram string
		switch i {
		case 0:
			bigram = "<START>-=-" + word
		case end:
			bigram = word + "-=-<END>"
		default:
			nextWord := words[i+1]
			bigram = word + "-=-" + nextWord
		}

		bigrams = append(bigrams, bigram)
	}

	return
}
