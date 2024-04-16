package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

const (
	wordNeededLen = 6
	fileName      = "wordlist.txt"
)

func main() {
	file, err := os.Open(fileName)
	if err != nil {
		fmt.Println(err.Error())
		os.Exit(1)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)

	wordsListNormalized := make(map[string]bool)
	for scanner.Scan() {
		word := scanner.Text()

		wordToLower := strings.ToLower(word)
		wordsListNormalized[wordToLower] = true
	}

	for word := range wordsListNormalized {
		if !isValidWord(word) {
			continue
		}

		for i := range word {
			if i == 0 {
				continue
			}
			firstWordPart := word[:i]
			secondWordPart := word[i:]

			if isSubStringInWordList(firstWordPart, wordsListNormalized) {
				if isSubStringInWordList(secondWordPart, wordsListNormalized) {
					fmt.Printf("%s + %s => %s\n", firstWordPart, secondWordPart, word)
				}
			}
		}
	}

	if err := scanner.Err(); err != nil {
		fmt.Println(err.Error())
	}

	os.Exit(0)
}

func isValidWord(word string) bool {
	return len(word) == wordNeededLen
}

func isSubStringInWordList(subString string, wordList map[string]bool) bool {
	if _, ok := wordList[subString]; ok {
		return true
	}

	return false
}
