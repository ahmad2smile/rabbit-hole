package main

import (
	"bufio"
	"os"
	"strings"
	"rabbit-hole/errorutils"
	"rabbit-hole/dedupeList"
)

const anagram string = "poultry outwits ants"

var anagramCharLen int = len(strings.ReplaceAll(anagram, " ", ""))

const hashToFind string = "e4820b45d2277f3844eac66c903e84be"

func main() {
	file, err := os.Open("./wordlist")
	optimizedFile, createErr := os.OpenFile("./optimized-wordlist", os.O_WRONLY|os.O_CREATE, 0666)

	errorutils.ThrowError(err)
	errorutils.ThrowError(createErr)

	defer file.Close()
	defer optimizedFile.Close()

	scanner := bufio.NewScanner(file)

	var list []string

	for scanner.Scan() {
		word := scanner.Text()
		isValidLine := true

		isValidLine = len(word) <= anagramCharLen

		if isValidLine {
			for _, char := range word {
				isValidLine = strings.Contains(anagram, string(char))
				isValidLine = (strings.Count(word, string(char)) <= strings.Count(anagram, string(char)))

				if !isValidLine {
					break
				}
			}
		}

		if isValidLine {
			list = append(list, word)
		}
	}

	uniqueList := dedupeList.RemoveDupe(list)

	for _, word := range uniqueList {
		optimizedFile.WriteString(word + "\n")
	}
	
}
