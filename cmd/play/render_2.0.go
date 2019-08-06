package play

import (
	"fmt"
	"github.com/gdamore/tcell"
	"github.com/mattn/go-runewidth"
)

// Returns length to next whitespace
// If whitespace was found, boolean is true, else false
func getLengthToWhitespace(str string, index int) (int, bool) {
	for i, c := range str[index:] {
		if c == []rune(" ")[0] {
			return i, true
		}
	}
	return 0, false
}

/**
 * This function was found in the tcell repo
 * zwj stands for zero-width-joiner
 */
func puts(s tcell.Screen, style tcell.Style, x, y int, str string) {
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

func putText2(s tcell.Screen, text string, subText string, rowStart int, colStart int, textBoxWidth int) {

	var row = rowStart
	var style = tcell.StyleDefault
	var greenStyle = tcell.StyleDefault.Foreground(tcell.NewRGBColor(50, 250, 50))
	var redStyle = tcell.StyleDefault.Foreground(tcell.NewRGBColor(250, 50, 50))
	// var formattedText = textBoxFormatString(text, textBoxWidth)
	var boxPadding = 10
	var currentLength = colStart + boxPadding
	// var isAfter = false

	// Draw first part of the box part of the textbox
	for i := 0; i <= textBoxWidth+boxPadding; i++ {
		puts(s, style, colStart+i, row, "-")
	}
	row++
	puts(s, style, colStart, row, "|")
	puts(s, style, colStart+textBoxWidth+boxPadding, row, "|")
	row++
	puts(s, style, colStart, row, "|")
	puts(s, style, colStart+textBoxWidth+boxPadding, row, "|")
	row++
	puts(s, style, colStart, row, "|")
	puts(s, style, colStart+textBoxWidth+boxPadding, row, "|")

	// Main logic, handle and draw text input, also draw box in between.
	// The isAfter variable indicates if we hace exceede the lines
	// that are typed or are being typed.
	for i, letterCharacter := range text {
		puts(s, style, colStart, row, "|")
		if len(subText) > i {
			// TODO: this can be made simpler by just assigning style var
			if rune(subText[i]) == letterCharacter {
				puts(s, greenStyle, currentLength, row, string(letterCharacter))
			} else {
				puts(s, redStyle, currentLength, row, string(letterCharacter))
			}
		} else {
			puts(s, style, currentLength, row, string(letterCharacter))
		}

		lengthToSpace, isInString := getLengthToWhitespace(text, i+1)

		puts(s, redStyle, 1, 20, "lts: "+fmt.Sprintf("%d", lengthToSpace))
		puts(s, redStyle, 1, 21, "ilw: "+fmt.Sprintf("%t", isInString))
		puts(s, redStyle, 1, 22, "i: "+fmt.Sprintf("%d", i))
		puts(s, redStyle, 1, 23, "textboxwidth: "+fmt.Sprintf("%d", textBoxWidth))
		puts(s, redStyle, 1, 24, "i: "+fmt.Sprintf("%d", currentLength))

		puts(s, redStyle, 1, 50, "typing: "+subText)

		if isInString {
			if lengthToSpace+currentLength-colStart > textBoxWidth {
				row++
				currentLength = colStart + boxPadding
				continue
			}
		}
		currentLength++
		puts(s, style, colStart+textBoxWidth+boxPadding, row, "|")
	}

	// Draw last part of textbox
	row++
	puts(s, style, colStart, row, "|")
	puts(s, style, colStart+textBoxWidth+boxPadding, row, "|")
	row++
	puts(s, style, colStart, row, "|")
	puts(s, style, colStart+textBoxWidth+boxPadding, row, "|")
	row++
	puts(s, style, colStart, row, "|")
	puts(s, style, colStart+textBoxWidth+boxPadding, row, "|")
	for i := 0; i <= textBoxWidth+boxPadding; i++ {
		puts(s, style, colStart+i, row, "-")
	}
}
