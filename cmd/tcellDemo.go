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
	"os"

	"github.com/gdamore/tcell"
	"github.com/gdamore/tcell/encoding"
	"github.com/mattn/go-runewidth"
	"github.com/spf13/cobra"
)

var rowDemo = 0
var styleDemo = tcell.StyleDefault

func putlnDemo(s tcell.Screen, str string) {

	putsDemo(s, style, 1, rowDemo, str)
	rowDemo++
}

func putsDemo(s tcell.Screen, style tcell.Style, x, y int, str string) {
	i := 0
	var deferred []rune
	dwidth := 0
	zwj := false
	for _, r := range str {
		if r == '\u200d' {
			if len(deferred) == 0 {
				deferred = append(deferred, ' ')
				dwidth = 1
			}
			deferred = append(deferred, r)
			zwj = true
			continue
		}
		if zwj {
			deferred = append(deferred, r)
			zwj = false
			continue
		}
		switch runewidth.RuneWidth(r) {
		case 0:
			if len(deferred) == 0 {
				deferred = append(deferred, ' ')
				dwidth = 1
			}
		case 1:
			if len(deferred) != 0 {
				s.SetContent(x+i, y, deferred[0], deferred[1:], style)
				i += dwidth
			}
			deferred = nil
			dwidth = 1
		case 2:
			if len(deferred) != 0 {
				s.SetContent(x+i, y, deferred[0], deferred[1:], style)
				i += dwidth
			}
			deferred = nil
			dwidth = 2
		}
		deferred = append(deferred, r)
	}
	if len(deferred) != 0 {
		s.SetContent(x+i, y, deferred[0], deferred[1:], style)
		i += dwidth
	}
}

// playCmd represents the play command
var tcellDemo = &cobra.Command{
	Use:   "tcellDemo",
	Short: "Demo for tcell.",
	Long: `Just a demo for unicode in tcell.
Safe to remove`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("tcell demo called")

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
			Foreground(tcell.ColorBlack).
			Background(tcell.ColorWhite))
		s.Clear()

		quit := make(chan struct{})

		style = bold
		putlnDemo(s, "Press ESC to Exit")
		putlnDemo(s, "Character set: "+s.CharacterSet())
		style = plain

		putlnDemo(s, "English:   October")
		putlnDemo(s, "Icelandic: október")
		putlnDemo(s, "Arabic:    أكتوبر")
		putlnDemo(s, "Russian:   октября")
		putlnDemo(s, "Greek:     Οκτωβρίου")
		putlnDemo(s, "Chinese:   十月 (note, two double wide characters)")
		putlnDemo(s, "Combining: A\u030a (should look like Angstrom)")
		putlnDemo(s, "Emoticon:  \U0001f618 (blowing a kiss)")
		putlnDemo(s, "Airplane:  \u2708 (fly away)")
		putlnDemo(s, "Command:   \u2318 (mac clover key)")
		putlnDemo(s, "Enclose:   !\u20e3 (should be enclosed exclamation)")
		putlnDemo(s, "ZWJ:       \U0001f9db\u200d\u2640 (female vampire)")
		putlnDemo(s, "ZWJ:       \U0001f9db\u200d\u2642 (male vampire)")
		putlnDemo(s, "Family:    \U0001f469\u200d\U0001f467\u200d\U0001f467 (woman girl girl)\n")
		putlnDemo(s, "Region:    \U0001f1fa\U0001f1f8 (USA! USA!)\n")
		putlnDemo(s, "")
		putlnDemo(s, "Box:")
		putlnDemo(s, string([]rune{
			tcell.RuneULCorner,
			tcell.RuneHLine,
			tcell.RuneTTee,
			tcell.RuneHLine,
			tcell.RuneURCorner,
		}))
		putlnDemo(s, string([]rune{
			tcell.RuneVLine,
			tcell.RuneBullet,
			tcell.RuneVLine,
			tcell.RuneLantern,
			tcell.RuneVLine,
		})+"  (bullet, lantern/section)")
		putlnDemo(s, string([]rune{
			tcell.RuneLTee,
			tcell.RuneHLine,
			tcell.RunePlus,
			tcell.RuneHLine,
			tcell.RuneRTee,
		}))
		putlnDemo(s, string([]rune{
			tcell.RuneVLine,
			tcell.RuneDiamond,
			tcell.RuneVLine,
			tcell.RuneUArrow,
			tcell.RuneVLine,
		})+"  (diamond, up arrow)")
		putlnDemo(s, string([]rune{
			tcell.RuneLLCorner,
			tcell.RuneHLine,
			tcell.RuneBTee,
			tcell.RuneHLine,
			tcell.RuneLRCorner,
		}))

		s.Show()
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
	rootCmd.AddCommand(tcellDemo)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// playCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// playCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
