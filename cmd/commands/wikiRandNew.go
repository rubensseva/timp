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
	"github.com/rubensseva/timp/cmd/net"
	"github.com/rubensseva/timp/cmd/data/model"
	"github.com/rubensseva/timp/cmd/data"
)

var num_of_threads int

// newTextCmd represents the newText command
var wikiRandNewCmd = &cobra.Command{
	Use:   "new",
	Short: "Add wikipedia texts",
	Long: `Adds a number of wikipedia texts to the
  text database concurrently.`,
	Run: func(cmd *cobra.Command, args []string) {


		fmt.Println("Runnning play function")
    channels := make(chan model.Text)

    if num_of_threads == 0 {
      num_of_threads = 10
    }

    fmt.Printf("Fetching %v texts\n", num_of_threads)
    if num_of_threads > 100 {
      fmt.Println("Warning: fetching a very large amount of texts concurrently. Program might crash.")
    }

    for i := 0; i < num_of_threads; i++ {
      go func() {
        channels <- net.WikiGetRandText()
      }()
    }

    for i := 0; i < num_of_threads; i++ {
      data.AddText(<-channels)
    }

	},
}

func init() {
	wikiRandCmd.AddCommand(wikiRandNewCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// newTextCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// newTextCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
	wikiRandCmd.PersistentFlags().IntVarP(&num_of_threads, "threads", "t", 2, "How many texts to get, and how many threads to run")
}
