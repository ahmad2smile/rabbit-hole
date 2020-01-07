package validation

import (
	"strings"
)

func IsValidWord(word string, anagram string, anagramCharLen int) bool {
	isValid := true

		// a word can not be bigger than whole anagram
		isValid = len(word) <= anagramCharLen

		if isValid {
			for _, char := range word {
				// every character in word must be in anagram else useless word
				isValid = strings.Contains(anagram, string(char))

				// repetition of a certain character in the word can not be greater than repetition in whole anagram
				isValid = (strings.Count(word, string(char)) <= strings.Count(anagram, string(char)))

				if !isValid {
					break
				}
			}
		}
	
	return isValid
}