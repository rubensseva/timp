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

	"github.com/spf13/cobra"

	"encoding/json"
	"io/ioutil"

	"timp/cmd/model"
)

// newTextCmd represents the newText command
var textNewCmd = &cobra.Command{
	Use:   "new",
	Short: "Adds new text",
	Long: `Adds a new text that may be run.
example: timp newText /path/to/file`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("newText called")
		if len(args) != 1 {
			fmt.Println("args not equal to one.. Did you remember to specify text to add?")
			return
		}

		textfile, _ := ioutil.ReadFile("cmd/resources/texts.json")
		var texts []model.Text
		_ = json.Unmarshal([]byte(textfile), &texts)

		var isAText = false
		for _, text := range texts {
			if text.Text == args[0] {
				isAText = true
			}
		}

		if isAText {
			fmt.Println("specified text " + args[0] + " is already a text")
			return
		}

		fmt.Println("creating text ", args[0])
		var newText = model.Text{Text: args[0], Author: ""}
		texts = append(texts, newText)
		writefile, _ := json.MarshalIndent(texts, "", " ")
		_ = ioutil.WriteFile("cmd/resources/texts.json", writefile, 0644)
		fmt.Println("create text success (hopefully)")
	},
}

func init() {
	textCmd.AddCommand(textNewCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// newTextCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// newTextCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
