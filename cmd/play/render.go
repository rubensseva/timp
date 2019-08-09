package play

import (
	"github.com/gdamore/tcell"

	"strings"
)

/**
 * -----> LEGACY <------
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
	var lineList []string

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
			lineList = append(lineList, tmpStr)
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
			lineList = append(lineList, tmpStr)
		}

		// Adding 1 here because we need to count whitespace
		widthCount += len(word) + 1
		i++
	}
	return lineList
}

/*
* ------> Legacy <------
* PutText Draws text to screen
* Params: tcell screen, the actual text and progress on text,
* position and width of textbox.
*
* Handles what part of text is already typed and
* what part of text is not complete
 */
func putText(s tcell.Screen, text string, progressIndex int, rowStart int, colStart int, textBoxWidth int) {
	var row = rowStart
	var style = tcell.StyleDefault
	var greenStyle = tcell.StyleDefault.Foreground(tcell.NewRGBColor(50, 250, 50))
	var formattedText = textBoxFormatString(text, textBoxWidth)
	var currentLength = 0
	var isAfter = false
	var boxPadding = 10

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
	for _, stringLine := range formattedText {
		puts(s, style, colStart, row, "|")
		var colStartInc = colStart + boxPadding
		var stringLineLength = len([]rune(stringLine))

		// Draw a line of text on three conditions, either before the line that is
		// being type, on the line that is being typed, or after the line
		// that is being typed respectively.
		if currentLength+stringLineLength < progressIndex && !isAfter {
			puts(s, greenStyle, colStartInc, row, stringLine)
		} else if !isAfter {
			puts(s, greenStyle, colStartInc, row, string([]rune(stringLine)[:progressIndex-currentLength]))
			puts(s, style, colStartInc+(progressIndex-currentLength), row, string([]rune(stringLine)[progressIndex-currentLength:]))
			puts(s, style, colStart+textBoxWidth+boxPadding, row, "|")
			isAfter = true
			row++
			continue
		} else if isAfter {
			puts(s, style, colStartInc, row, stringLine)
		}

		currentLength = currentLength + stringLineLength
		puts(s, style, colStart+textBoxWidth+boxPadding, row, "|")
		row++
	}

	// Draw last part of the box part of the textbox
	puts(s, style, colStart, row, "|")
	puts(s, style, colStart+textBoxWidth+boxPadding, row, "|")
	row++
	puts(s, style, colStart, row, "|")
	puts(s, style, colStart+textBoxWidth+boxPadding, row, "|")
	row++
	puts(s, style, colStart, row, "|")
	puts(s, style, colStart+textBoxWidth+boxPadding, row, "|")
	row++
	for i := 0; i <= textBoxWidth+boxPadding; i++ {
		puts(s, style, colStart+i, row, "-")
	}
}
