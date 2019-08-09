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
	"timp/cmd/net"
	"timp/cmd/data/model"
	"timp/cmd/data"
)


// newTextCmd represents the newText command
var wikiRandNewCmd = &cobra.Command{
	Use:   "new",
	Short: "Add 10 wikipedia texts to text",
	Long: `Adds 10 wikipedia texts to text`,
	Run: func(cmd *cobra.Command, args []string) {


		fmt.Println("Runnning play function")
    channel_1 := make(chan struct{})
    channel_2 := make(chan struct{})
    channel_3 := make(chan struct{})
    channel_4 := make(chan struct{})
    channel_5 := make(chan struct{})
    channel_6 := make(chan struct{})
    channel_7 := make(chan struct{})
    var text_1 model.Text
    var text_2 model.Text
    var text_3 model.Text
    var text_4 model.Text
    var text_5 model.Text
    var text_6 model.Text
    var text_7 model.Text

    go func() {
      text_1 = net.WikiGetRandText()
      channel_1 <- struct{}{}
    }()
    go func() {
      text_2 = net.WikiGetRandText()
      channel_2 <- struct{}{}
    }()
    go func() {
      text_3 = net.WikiGetRandText()
      channel_3 <- struct{}{}
    }()
    go func() {
      text_4 = net.WikiGetRandText()
      channel_4 <- struct{}{}
    }()
    go func() {
      text_5 = net.WikiGetRandText()
      channel_5 <- struct{}{}
    }()
    go func() {
      text_6 = net.WikiGetRandText()
      channel_6 <- struct{}{}
    }()
    go func() {
      text_7 = net.WikiGetRandText()
      channel_7 <- struct{}{}
    }()




    <-channel_1
    <-channel_2
    <-channel_3
    <-channel_4
    <-channel_5
    <-channel_6
    <-channel_7
		data.AddText(text_1)
		data.AddText(text_2)
		data.AddText(text_3)
		data.AddText(text_4)
		data.AddText(text_5)
		data.AddText(text_6)
		data.AddText(text_7)
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
}
