package play

import (
	"fmt"
	"github.com/gdamore/tcell"
	"github.com/mattn/go-runewidth"
)

/**
 * Returns length to next whitespace
 * If whitespace was found, boolean is true, else false
*/
func getLengthToWhitespace(str []rune, index int) (int, bool) {
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


/**
 * Draws text on terminal screen, given text and subtext
*/
func putText2(s tcell.Screen, text []rune, subText []rune, rowStart int, colStart int, textBoxWidth int) {

  // Variables for main text box
	var row = rowStart
	var defaultStyle = tcell.StyleDefault
	var greenStyle = tcell.StyleDefault.Foreground(tcell.NewRGBColor(50, 250, 50))
	var redStyle = tcell.StyleDefault.Foreground(tcell.NewRGBColor(200, 70, 70))
  var currentLetterStyle = tcell.StyleDefault.Foreground(tcell.NewRGBColor(50, 250, 50)).Background(tcell.NewRGBColor(50, 150, 50))
	var boxPadding = 10
	var currentLength = colStart + boxPadding

  // Variables for alternative typing box
  var continous_box_offset = 10
  var continous_box_length = 50
  var continous_box_row_start = 50
  var contionous_box_col_start = 1
  var contionous_box_col = contionous_box_col_start

	// Draw first part of the box part of the main textbox
	for i := 0; i <= textBoxWidth+boxPadding; i++ {
		puts(s, style, colStart+i, row, "-")
	}
	row++
	puts(s, defaultStyle, colStart, row, "|")
	puts(s, defaultStyle, colStart+textBoxWidth+boxPadding, row, "|")
	row++
	puts(s, defaultStyle, colStart, row, "|")
	puts(s, defaultStyle, colStart+textBoxWidth+boxPadding, row, "|")
	row++
	puts(s, defaultStyle, colStart, row, "|")
	puts(s, defaultStyle, colStart+textBoxWidth+boxPadding, row, "|")

	// Main logic, loop through all characters in text and 
  // draw the text boxes
	for i, letterCharacter := range text {
		puts(s, defaultStyle, colStart, row, "|")


    // Main text box
    var style tcell.Style
		if len(subText) > i {
			if rune(subText[i]) == letterCharacter {
        style = greenStyle
			} else {
        style = redStyle
			}
		} else {
      style = tcell.StyleDefault
		}
		puts(s, style, currentLength, row, string(letterCharacter))
		lengthToSpace, isInString := getLengthToWhitespace(text, i+1)
		if isInString {
			if lengthToSpace+currentLength-colStart > textBoxWidth {
				row++
				currentLength = colStart + boxPadding
				continue
			}
		}
		currentLength++
    puts(s, style, colStart+textBoxWidth+boxPadding, row, "|")

    // Alternative scrolling textbox
    if i < len(subText) - continous_box_offset || i > len(subText) + continous_box_length {
      // Pass
    } else {
      var style tcell.Style
      if len(subText) == i + 1 {
        style = currentLetterStyle
      } else if len(subText) > i {
        if rune(subText[i]) == letterCharacter {
          style = greenStyle
        } else {
          style = redStyle
        }
      } else {
        style = redStyle
      }
			puts(s, style, contionous_box_col, continous_box_row_start, string(letterCharacter))
      contionous_box_col++
    }

	}

	// Draw last part of textbox
	row++
	puts(s, defaultStyle, colStart, row, "|")
	puts(s, defaultStyle, colStart+textBoxWidth+boxPadding, row, "|")
	row++
	puts(s, defaultStyle, colStart, row, "|")
	puts(s, defaultStyle, colStart+textBoxWidth+boxPadding, row, "|")
	row++
	puts(s, defaultStyle, colStart, row, "|")
	puts(s, defaultStyle, colStart+textBoxWidth+boxPadding, row, "|")
	for i := 0; i <= textBoxWidth+boxPadding; i++ {
		puts(s, defaultStyle, colStart+i, row, "-")
	}

  // General info
  puts(s, redStyle, 1, 23, "textboxwidth: "+fmt.Sprintf("%d", textBoxWidth))
  puts(s, redStyle, 1, 24, "i: "+fmt.Sprintf("%d", currentLength))
}
