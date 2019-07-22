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
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
	"strings"
	"time"

	"github.com/gdamore/tcell"
	"github.com/gdamore/tcell/encoding"
	"github.com/spf13/cobra"

	"timp/cmd/model"
	"timp/cmd/tcell_helpers"
	"timp/cmd/utility"
)

const textBoxWidth = 50

const playTextBoxPos_x = 30
const playTextBoxPos_y = 20

var style = tcell.StyleDefault
var greenStyle = tcell.StyleDefault.Foreground(tcell.NewRGBColor(50, 250, 50))

// playCmd represents the play command
var playCmd = &cobra.Command{
	Use:   "play",
	Short: "Play timp with given input text",
	Long: `This command takes control of the terminal and starts the 
  main feature of timp. Tcell is used to start a session 
  where you may input the given text on screen, and progress 
  is shown. NOT IMPLEMENTED YET: Stats will be shown after 
  completion.`,

	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("play called")

		s, e := tcell.NewScreen()
		if e != nil {
			fmt.Fprintf(os.Stderr, "%v\n", e)
			os.Exit(1)
		}

		encoding.Register()

		if e = s.Init(); e != nil {
			fmt.Fprintf(os.Stderr, "%v\n", e)
			os.Exit(1)
		}

		plain := tcell.StyleDefault
		bold := style.Bold(true)

		s.SetStyle(tcell.StyleDefault.
			Foreground(tcell.ColorWhite).
			Background(tcell.ColorBlack))
		s.Clear()

		quit := make(chan struct{})

		style = bold
		tcell_helpers.PutText(s, "Press ESC to Exit", 0, 0, 0, 25)
		tcell_helpers.PutText(s, "Character set: "+s.CharacterSet(), 0, 2, 0, 25)
		style = plain

		textfile, _ := ioutil.ReadFile("cmd/resources/texts.json")
		var texts []model.Text
		_ = json.Unmarshal([]byte(textfile), &texts)

		var randIndex = utility.RandomGen(0, len(texts))

		textToRun := texts[randIndex]
		totNumOfWords := len(strings.Fields(textToRun.Text))
		start := time.Now()

		tcell_helpers.PutText(s, textToRun.Text, 0, playTextBoxPos_y, playTextBoxPos_x, 40)
		var stringTyped = ""
		var numCorrect = 0

		go func() {
			for {
				ev := s.PollEvent()
				switch ev := ev.(type) {
				case *tcell.EventKey:
					switch ev.Key() {
					case tcell.KeyEscape, tcell.KeyEnter:
						close(quit)
						return
					case tcell.KeyCtrlL:
						s.Sync()
					case tcell.KeyBackspace2:
						if numCorrect > 0 {
							numCorrect--
							tcell_helpers.PutText(s, textToRun.Text, numCorrect, playTextBoxPos_y, playTextBoxPos_x, 40)
							s.Show()
						}
					case tcell.KeyRune:
						stringTyped = stringTyped + string(ev.Rune())
						if []rune(textToRun.Text)[numCorrect] == ev.Rune() {
							numCorrect++
						}
						tcell_helpers.PutText(s, textToRun.Text, numCorrect, playTextBoxPos_y, playTextBoxPos_x, 40)
						s.Show()
					}
				case *tcell.EventResize:
					s.Sync()
				}
				if numCorrect >= len(textToRun.Text) {
					close(quit)
					return
				}
			}
		}()
		<-quit
		s.Fini()
		t := time.Now()
		elapsed := t.Sub(start)
		println("\n\nGame complete!, with text: \n")
		println(textToRun.Text)
		println("\nSTATS")
		print("Elapsed time: ")
		println(fmt.Sprintf("%.2fs", float32((elapsed.Seconds()))))
		print("Words completed: ")
		println(totNumOfWords)
		print("Words per minute: ")
		println(fmt.Sprintf("%.3f", float32(totNumOfWords)/float32(float32(elapsed.Seconds())/60.0)))
		println()
		println()
	},
}

func init() {
	rootCmd.AddCommand(playCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// playCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// playCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
