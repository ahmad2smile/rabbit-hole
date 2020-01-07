package main

import (
	"bufio"
	"os"
	"fmt"
	"strings"
	"rabbit-hole/errorutils"
	"rabbit-hole/dedupeList"
	"rabbit-hole/validation"
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
		
		if validation.IsValidWord(word, anagram, anagramCharLen) {
			list = append(list, word)
		}
	}

	uniqueList := dedupeList.RemoveDupe(list)

	fmt.Println("Final Words:", len(uniqueList))

	for _, word := range uniqueList {
		optimizedFile.WriteString(word + "\n")
	}
	
}
