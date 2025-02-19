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

type StringScore struct {
	Text              string
	Score             float32
	IsProbablyEnglish bool
}

func IsStringProbablyEnglishSentence(s string) StringScore {
	var stringScore StringScore
	stringScore.Text = s
	stringScore.Score = 0
	stringScore.IsProbablyEnglish = false
	if len(s) < 5 {
		return stringScore
	}
	words := strings.Fields(s)
	dictionary, err := readLines(os.Getenv("HOME") + "/.timp/20k.txt")
	if err != nil {
		log.Fatalf("readlines, %s", err)
	}
	var numOfWords = len(words)
	var numOfWordsMatched = 0
	for _, c := range words {
		if stringInSlice(strings.ToLower(c), dictionary) {
			numOfWordsMatched++
		}
	}
	var score = float32(numOfWordsMatched) / float32(numOfWords)
	stringScore.Score = score
	if score > 0.5 {
		fmt.Println("Potential candidate with score: " + fmt.Sprintf("%.3f", float32(score)))
		fmt.Println(words)
		stringScore.IsProbablyEnglish = true
	}
	return stringScore
}
