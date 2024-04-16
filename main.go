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

	for k := range wordsListNormalized {
		if len(k) != wordNeededLen {
			continue
		}
		for i := range k {
			firstWordPart := k[:i]
			secondWordPart := k[i:]

			if _, ok := wordsListNormalized[firstWordPart]; ok {
				if _, ok := wordsListNormalized[secondWordPart]; ok {

					fmt.Println(k)

				}
			}
		}
	}

	if err := scanner.Err(); err != nil {
		fmt.Println(err.Error())
	}

	os.Exit(0)
}