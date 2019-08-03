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
	"github.com/spf13/cobra"
	"timp/cmd/data"
	"timp/cmd/data/model"
)

// textCmd represents the text command
var historyCmd = &cobra.Command{
	Use:   "history",
	Short: "Lists history",
	Long: `Lists entire history
		example: timp histoyr`,

	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("history called")

		var historyEntries []model.PlayedEntry = data.GetAllHistoryEntries()

		for _, historyEntry := range historyEntries {
			fmt.Println()
			fmt.Println("ID: " + fmt.Sprintf("%d", historyEntry.GetID()))
			fmt.Println("Player: " + historyEntry.GetPlayer())
			fmt.Println("Text: " + historyEntry.GetText().GetText())
			fmt.Println("Author: " + historyEntry.GetText().GetAuthor())
			fmt.Println("Date of play: " + historyEntry.GetTimePlayed().String())
			fmt.Println("Time spent: " + fmt.Sprintf("%f", historyEntry.GetTimeSpent()))
			fmt.Println("WPM: " + fmt.Sprintf("%f", historyEntry.GetWpm()))
			fmt.Println("Valid play: " + fmt.Sprintf("%t", historyEntry.GetDidFinishLegally()))
		}
	},
}

func init() {
	rootCmd.AddCommand(historyCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// textCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// textCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
