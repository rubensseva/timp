/*
	Copyright © 2019 Ruben Svanåsbakken Sevaldson <r.sevaldson@gmail.com>

	Permission is hereby granted, free of charge, to any person obtaining a copy
	of this software and associated documentation files (the "Software"), to deal
	in the Software without restriction, including without limitation the rights
	to use, copy, modify, merge, publish, distribute, sublicense, and/or sell
	copies of the Software, and to permit persons to whom the Software is
	furnished to do so, subject to the following conditions:

	The above copyright notice and this permission notice shall be included in
	all copies or substantial portions of the Software.

	THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND, EXPRESS OR
	IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF MERCHANTABILITY,
	FITNESS FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT. IN NO EVENT SHALL THE
	AUTHORS OR COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER
	LIABILITY, WHETHER IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING FROM,
	OUT OF OR IN CONNECTION WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN
	THE SOFTWARE.
*/

// Package commands represents the actual available commands to
// use from command line
package commands

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"regexp"
	"strings"
	"timp/cmd/data/model"
	"timp/cmd/tcell_helpers"
	"timp/cmd/utility"

	"github.com/spf13/cobra"
	"golang.org/x/net/html"
)

type jsonMapper struct {
	fullurl string
}

// newTextCmd represents the newText command
var wikiRandCmd = &cobra.Command{
	Use:   "wikiRand",
	Short: "Play with random wikipedia article",
	Long: `Pulls a random article from wikipedia
		and plays it immediatly`,
	Run: func(cmd *cobra.Command, args []string) {

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

		if len(englishSentences) > 0 {
			fmt.Println("Runnning play function")
			tcell_helpers.Play(text)
		} else {
			fmt.Println("Was not able to find random wiki text, something might be wrong")
		}
	},
}

func init() {
	rootCmd.AddCommand(wikiRandCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// newTextCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// newTextCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
