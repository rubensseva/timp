package net

import (
	"fmt"
	"golang.org/x/net/html"
	"io/ioutil"
	"log"
	"net/http"
	"regexp"
	"strings"
	"github.com/rubensseva/timp/cmd/data/model"
	"github.com/rubensseva/timp/cmd/utility"
)

func WikiGetRandText() model.Text {
	var englishSentences []utility.StringScore
	for i := 0; i < 10; i++ {
		fmt.Println("Fetching random text from wikipedia...")
		var url = "https://en.wikipedia.org/wiki/Special:Random"
		var resp, _ = http.Get(url)
		var bytes, _ = ioutil.ReadAll(resp.Body)

		r := strings.NewReader(string(bytes))

		var potentialText []string

		z := html.NewTokenizer(r)
		done := false
		for done != true {
			tt := z.Next()
			switch tt {
			case html.ErrorToken:
				done = true
			case html.TextToken:
				var s = string(z.Text())
				if len(s) > 100 {
					potentialText = append(potentialText, s)
					// potentialText = append(potentialText, strings.Trim(s, " ,.)(][}{"))
				}
				break
			case html.EndTagToken:
				break
			}
		}

		for _, c := range potentialText {
			var stringScore = utility.IsStringProbablyEnglishSentence(c)
			if stringScore.IsProbablyEnglish {
				englishSentences = append(englishSentences, stringScore)
			}
		}
		if len(englishSentences) > 0 {
			break
		}
	}

	fmt.Println("Found " + string(len(englishSentences)) + " sentences. Finding most suitable sentence.")

	var maxScore float32 = 0.0
	var bestString string
	for _, c := range englishSentences {
		if c.Score > maxScore && len(c.Text) < 500 {
			maxScore = c.Score
			bestString = c.Text
		}
	}

	reg, err := regexp.Compile("[^a-zA-Z0-9,._ -%!?/]+")
	if err != nil {
		log.Fatal(err)
	}

	var processedText string = reg.ReplaceAllString(bestString, "")
	processedText = strings.Trim(processedText, " ,-/")

	reg, err = regexp.Compile(`[\s\p{Zs}]{2,}`)
	if err != nil {
		log.Fatal(err)
	}
	processedText = reg.ReplaceAllString(processedText, " ")

	text := model.NewText(processedText, "Wikipedia")
	return text
}
