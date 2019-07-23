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

// Package cmd represents cobra command
package cmd

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"strings"
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
	Short: "Play with random wikipedia artivle",
	Long: `Pulls a random article from wikipedia
		and plays it immediatly`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("wikiRand called")

		var url = "https://en.wikipedia.org/w/api.php?action=query&prop=info&generator=random&format=json&inprop=url"
		resp, _ := http.Get(url)
		bytes, _ := ioutil.ReadAll(resp.Body)

		fmt.Println("HTML:\n\n", string(bytes))

		resp.Body.Close()

		// We have a random article, now to extract the url
		fullurlStartIndex := strings.Index(string(bytes), "fullurl")
		fullurlEndIndex := strings.Index(string(bytes), "editurl")
		var fullurl = string(bytes)[fullurlStartIndex+10 : fullurlEndIndex-3]
		fmt.Print("\n\n\n\nFULL URL: ")
		fmt.Println(fullurl)

		url = fullurl
		resp, _ = http.Get(url)
		bytes, _ = ioutil.ReadAll(resp.Body)

		// fmt.Println("HTML:\n\n", string(bytes))

		r := strings.NewReader(string(bytes))

		var potentialText []string

		z := html.NewTokenizer(r)
		done := false
		for done != true {
			tt := z.Next()
			switch tt {
			case html.ErrorToken:
				fmt.Println("error")
				fmt.Println(z.Err())
				done = true
			case html.TextToken:
				// emitBytes should copy the []byte it receives,
				// if it doesn't process it immediately.
				var s = string(z.Text())
				if len(s) > 100 {
					fmt.Println("appending")
					potentialText = append(potentialText, s)
				}
				break
			case html.EndTagToken:
				fmt.Println("end")
				break
			}
		}

		var englishSentences []string
		for _, c := range potentialText {
			if utility.IsStringProbablyEnglishSentence(c) {
				englishSentences = append(englishSentences, c)
			}
		}
		fmt.Println(englishSentences)
		fmt.Println(fullurl)
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
