package nlputils

import (
	"strings"

	"golang.org/x/text/unicode/norm"
)

// ToWords will convert a string into a series of word values
func ToWords(str string) (words []string) {
	var wordBuf []rune
	// Normalize string before working with it
	str = norm.NFC.String(str)
	// Lowercase string before working with it
	str = strings.ToLower(str)

	for _, char := range str {
		switch char {
		case ' ', '?', '.', '!', ',':
			if len(wordBuf) == 0 {
				break
			}

			word := string(wordBuf)
			words = append(words, word)
			wordBuf = wordBuf[:0]

		default:
			wordBuf = append(wordBuf, char)
		}
	}

	if len(wordBuf) > 0 {
		word := string(wordBuf)
		words = append(words, word)
	}

	return
}
