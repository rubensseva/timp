package utility

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strings"
)

func readLines(path string) ([]string, error) {
	file, err := os.Open(path)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	var lines []string
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}
	return lines, scanner.Err()
}

func stringInSlice(a string, list []string) bool {
	for _, b := range list {
		if b == a {
			return true
		}
	}
	return false
}

func IsStringProbablyEnglishSentence(s string) bool {
	if len(s) < 5 {
		return false
	}
	words := strings.Fields(s)
	dictionary, err := readLines("cmd/resources/words.txt")
	if err != nil {
		log.Fatalf("readlines, %s", err)
	}
	var numOfWords = len(words)
	var numOfWordsMatched = 0
	for _, c := range words {
		if stringInSlice(c, dictionary) {
			numOfWordsMatched++
		}
	}
	var score = float32(numOfWordsMatched) / float32(numOfWords)
	fmt.Println("\n\nCalculated score for: ")
	fmt.Println(s)
	fmt.Println(words)
	fmt.Println(len(words))
	fmt.Println(score)
	if score > 0.5 {
		return true
	}
	return false
}
