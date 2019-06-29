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
package cmd

import (
	"fmt"
	"os"
	"strings"

	"github.com/gdamore/tcell"
	"github.com/gdamore/tcell/encoding"
	"github.com/spf13/cobra"

  "timp/cmd/tcell_helpers"
)

var TEXT_BOX_WIDTH = 50

var row = 0
var style = tcell.StyleDefault
var greenStyle = tcell.StyleDefault.Foreground(tcell.NewRGBColor(50, 250, 50))

func putln(s tcell.Screen, str string) {
	tcell_helpers.Puts(s, style, 1, row, str)
	row++
}

func putlnGreen(s tcell.Screen, str string) {
	tcell_helpers.Puts(s, greenStyle, 1, row, str)
	row++
}

func putText(s tcell.Screen, text string, progressIndex int) {
  subStr := text[:progressIndex]

	var finStr = strings.Replace(text, subStr, "", 1)
  var stringList []string
  var currentString string
  var widthCount = 0
  var currentLineNum = 0

  for _, r := range subStr {
    if widthCount <= TEXT_BOX_WIDTH {
      currentString = currentString + string(r)
      widthCount++
    } else {
      currentString = currentString + string(r)
      stringList = append(stringList, currentString)
      currentString = ""
      widthCount = 0
      currentLineNum++
    }

  }

  for _, r := range stringList {
    tcell_helpers.Puts(s, greenStyle, 1, row, r)
    row++
  }

  var subLength = 0
  for _, r := range stringList {
    subLength += len(r)
  }

  var cutSubStr = subStr[subLength:]

  tcell_helpers.Puts(s, greenStyle, 1, row, cutSubStr)

  cutStrFromFin := ""

  if (len(finStr) >= (TEXT_BOX_WIDTH - len(cutSubStr))) {
    cutStrFromFin = finStr[:(TEXT_BOX_WIDTH) - len(cutSubStr)]
  }

  tcell_helpers.Puts(s, style, len(cutSubStr) + 1, row, cutStrFromFin)
	row++

  trimmedFinStr := finStr[len(cutStrFromFin):]

  var stringListFin []string
  var currentStringFin string
  var widthCountFin = 0
  var currentLineNumFin = 0
  for _, r := range trimmedFinStr {
    if widthCountFin <= TEXT_BOX_WIDTH {
      currentStringFin = currentStringFin + string(r)
      widthCountFin++
    } else {
      currentStringFin = currentStringFin + string(r)
      stringListFin = append(stringListFin, currentStringFin)
      currentStringFin = ""
      widthCountFin = 0
      currentLineNumFin++
    }

  }

  for _, r := range stringListFin {
    tcell_helpers.Puts(s, style, 1, row, r)
    row++
  }

  var finLength = 0
  for _, r := range stringListFin {
    finLength += len(r)
  }

  lastOfTrimmedFin := trimmedFinStr[finLength:]


  tcell_helpers.Puts(s, style, 1, row, lastOfTrimmedFin)
}

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
		putln(s, "Press ESC to Exit")
		putln(s, "Character set: "+s.CharacterSet())
		style = plain

    finString := "Hello, this is a text I just wrote. I like this text and I like to program. Do you like to program? I would like to know very much. bye bye. In addition I would like to say that the world is nice and that I like ice cream! Ice cream is very nice and so are apples."
    putText(s,finString,25)

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
					}
				case *tcell.EventResize:
					s.Sync()
				}
			}
		}()

		<-quit

		s.Fini()

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
