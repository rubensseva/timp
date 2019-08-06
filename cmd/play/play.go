package play

import (
	"fmt"
	"os"
	"strings"
	"time"
	"timp/cmd/data"
	"timp/cmd/data/model"

	"github.com/gdamore/tcell"
	"github.com/gdamore/tcell/encoding"
)

const maxStringLength = 500

const textBoxWidth = 50

const playTextBoxPos_x = 30
const playTextBoxPos_y = 20

var style = tcell.StyleDefault
var greenStyle = tcell.StyleDefault.Foreground(tcell.NewRGBColor(50, 250, 50))

// Play renders text and starts game
// Also prints stats and calls history append function
func Play(text model.Text) {

	if len(text.GetText()) > maxStringLength {
		fmt.Println("Text to play is to long!")
		fmt.Println("Max accepted length is: " + string(maxStringLength) + " but was: " + string(len(text.GetText())))
		return
	}

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
	putText(s, "Press ESC to Exit", 0, 0, 0, 25)
	putText(s, "Character set: "+s.CharacterSet(), 0, 2, 0, 25)
	style = plain

	textToRun := text.GetText()
	totNumOfWords := len(strings.Fields(textToRun))
	start := time.Now()

	putText(s, textToRun, 0, playTextBoxPos_y, playTextBoxPos_x, 40)
	var stringTyped = ""
	var numCorrect = 0

	var didFinishLegally = false
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
						putText(s, textToRun, numCorrect, playTextBoxPos_y, playTextBoxPos_x, 40)
						s.Show()
					}
				case tcell.KeyRune:
					stringTyped = stringTyped + string(ev.Rune())
					if []rune(textToRun)[numCorrect] == ev.Rune() {
						numCorrect++
					}
					putText(s, textToRun, numCorrect, playTextBoxPos_y, playTextBoxPos_x, 40)
					s.Show()
				}
			case *tcell.EventResize:
				s.Sync()
			}
			if numCorrect >= len(textToRun) {
				didFinishLegally = true
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
	println(textToRun)
	println("\nSTATS")
	print("Elapsed time: ")
	println(fmt.Sprintf("%.2fs", float32((elapsed.Seconds()))))
	print("Words completed: ")
	println(totNumOfWords)
	print("Words per minute: ")
	println(fmt.Sprintf("%.3f", float32(totNumOfWords)/float32(float32(elapsed.Seconds())/60.0)))
	println()
	println()

	var currentUser = data.GetLoggedInUser()

	var playername = "not logged in"
	if currentUser.GetUser().GetUsername() != "" {
		playername = currentUser.GetUser().GetUsername()
	}
	data.AppendToHistory(text, playername, elapsed, didFinishLegally)
}
