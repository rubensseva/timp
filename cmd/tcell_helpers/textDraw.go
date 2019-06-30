package tcell_helpers

import (
	"github.com/gdamore/tcell"
	"github.com/mattn/go-runewidth"

	"strings"
)

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
 * Private helper function
 * Takes in a string to format and a width
 * Returns a list of strings
 *
 * Converts string to one large list of words based on whitespace, then
 * splits this list into several lists of strings that are now representing
 * correct formatting for textbox. This is done to prevent the final textbox
 * from doing a new line inside a word.
 */
func textBoxFormatString(text string, textBoxWidth int) []string {
	words := strings.Fields(text)
	var wordsList []string

	var widthCount = 0
	var lineCount = 0
	var lastWord = 0

	/**
	 * Iterate over words, and split into correct list of strings
	 * If adding the next word exeeds text box width we enter the if condition, and
	 * make the words so far to a single string and append.
	 */
	var i = 0
	for _, word := range words {
		if widthCount+len(word)+1 >= textBoxWidth {
			var tmpStr = ""
			for j := lastWord; j < i; j++ {
				tmpStr = tmpStr + words[j]
				if j != i {
					tmpStr = tmpStr + " "
				}
			}
			wordsList = append(wordsList, tmpStr)
			widthCount = 0
			lineCount++
			lastWord = i
		}

		// Add the last word(s)
		if i == len(words)-1 {
			var tmpStr = ""
			for j := lastWord; j < len(words); j++ {
				tmpStr = tmpStr + words[j]
				if j != i {
					tmpStr = tmpStr + " "
				}
			}
			wordsList = append(wordsList, tmpStr)
		}

		// Adding 1 here because we need to count whitespace
		widthCount += len(word) + 1
		i++
	}

	return wordsList
}

/*
* PutText Draws text to screen
* Params: tcell screen, the actual text and progress on text,
* position and width of textbox.
*
* Handles what part of text is already typed and
* what part of text is not complete
 */
func PutText(s tcell.Screen, text string, progressIndex int, rowStart int, colStart int, textBoxWidth int) {
	var row = rowStart
	var style = tcell.StyleDefault
	var greenStyle = tcell.StyleDefault.Foreground(tcell.NewRGBColor(50, 250, 50))

	var formattedText = textBoxFormatString(text, textBoxWidth)

	var currentLength = 0
	var isAfter = false
	for _, stringLine := range formattedText {
		if currentLength+len(stringLine) < progressIndex && !isAfter {
			puts(s, greenStyle, colStart, row, stringLine)
		} else if !isAfter {
			puts(s, greenStyle, colStart, row, stringLine[:progressIndex-currentLength])
			puts(s, style, colStart+(progressIndex-currentLength), row, stringLine[progressIndex-currentLength:])
			isAfter = true
			row++
			continue
		}
		if isAfter {
			puts(s, style, colStart, row, stringLine)
		}
		row++
		currentLength = currentLength + len(stringLine)
	}
}
