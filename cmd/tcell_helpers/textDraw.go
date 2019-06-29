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
	i := 0;
	var deferred []rune;
	dwidth := 0;
	zwj := false;
	for _, r := range str {
		if r == '\u200d' {
			if len(deferred) == 0 {
				deferred = append(deferred, ' ');
				dwidth = 1;
			}
			deferred = append(deferred, r);
			zwj = true;
			continue;
		}
		if zwj {
			deferred = append(deferred, r);
			zwj = false;
			continue;
		}
		switch runewidth.RuneWidth(r) {
		case 0:
			if len(deferred) == 0 {
				deferred = append(deferred, ' ');
				dwidth = 1;
			}
		case 1:
			if len(deferred) != 0 {
				s.SetContent(x+i, y, deferred[0], deferred[1:], style);
				i += dwidth;
			}
			deferred = nil;
			dwidth = 1;
		case 2:
			if len(deferred) != 0 {
				s.SetContent(x+i, y, deferred[0], deferred[1:], style);
				i += dwidth;
			}
			deferred = nil;
			dwidth = 2;
		}
		deferred = append(deferred, r);
	}
	if len(deferred) != 0 {
		s.SetContent(x+i, y, deferred[0], deferred[1:], style);
		i += dwidth;
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
func textBoxFormatString(text string, textBoxWidth int) ([]string) {
  words := strings.Fields(text);
  var wordsList []string;

  var widthCount = 0;
  var lineCount = 0;
  var lastWord = 0;

  /**
    * Iterate over words, and split into correct list of strings
    * If adding the next word exeeds text box width we enter the if condition, and
    * make the words so far to a single string and append.
  */
  var i = 0
  for _, word := range words {
    if (widthCount + len(word) + 1 >= textBoxWidth) {
      var tmpStr = "";
      for j := lastWord; j < i; j++ {
        tmpStr = tmpStr + words[j];
        if (j != i) {
          tmpStr = tmpStr + " ";
        }
      }
      wordsList = append(wordsList, tmpStr);
      widthCount = 0;
      lineCount++;
      lastWord = i;
    }

    // Add the last word(s)
    if (i == len(words) - 1) {
      var tmpStr = "";
      for j := lastWord; j < len(words); j++ {
        tmpStr = tmpStr + words[j];
        if (j != i) {
          tmpStr = tmpStr + " ";
        }
      }
      wordsList = append(wordsList, tmpStr);
    }

    // Adding 1 here because we need to count whitespace
    widthCount += len(word) + 1;
    i++;
  }

  return wordsList;
}




func PutText(s tcell.Screen, text string, progressIndex int, rowStart int, colStart int, textBoxWidth int) {
  var row = rowStart;
  var style = tcell.StyleDefault;
  var greenStyle = tcell.StyleDefault.Foreground(tcell.NewRGBColor(50, 250, 50));

  /**
    * First, we find the substring that is 
    * already typed
  */
  subStr := text[:progressIndex];
	var finStr = strings.Replace(text, subStr, "", 1);
  var stringList []string;
  var currentString string;
  var widthCount = 0;
  var currentLineNum = 0;
  for _, r := range subStr {
    if widthCount <= textBoxWidth {
      currentString = currentString + string(r);
      widthCount++;
    } else {
      currentString = currentString + string(r);
      stringList = append(stringList, currentString);
      currentString = "";
      widthCount = 0;
      currentLineNum++;
    }
  }

  /**
    * Draw all lines that are aldready typed
  */
  for _, r := range stringList {
    puts(s, greenStyle, colStart, row, r);
    row++;
  }

  /**
    * Handle the singular line that is partially typed
  */
  var subLength = 0;
  for _, r := range stringList {
    subLength += len(r)
  }
  var cutSubStr = subStr[subLength:];
  puts(s, greenStyle, colStart, row, cutSubStr);
  cutStrFromFin := "";
  if (len(finStr) >= (textBoxWidth - len(cutSubStr))) {
    cutStrFromFin = finStr[:(textBoxWidth) - len(cutSubStr)];
  }
  puts(s, style, colStart + len(cutSubStr) + 1, row, cutStrFromFin);
	row++;

  /**
    * Draw the lines that are not started on
  */
  trimmedFinStr := finStr[len(cutStrFromFin):];
  var stringListFin []string;
  var currentStringFin string;
  var widthCountFin = 0;
  var currentLineNumFin = 0;
  for _, r := range trimmedFinStr {
    if widthCountFin <= textBoxWidth {
      currentStringFin = currentStringFin + string(r);
      widthCountFin++;
    } else {
      currentStringFin = currentStringFin + string(r);
      stringListFin = append(stringListFin, currentStringFin);
      currentStringFin = "";
      widthCountFin = 0;
      currentLineNumFin++;
    }
  }
  for _, r := range stringListFin {
    puts(s, style, colStart, row, r);
    row++;
  }

  /**
    * Draw the last line
  */
  var finLength = 0;
  for _, r := range stringListFin {
    finLength += len(r);
  }
  lastOfTrimmedFin := trimmedFinStr[finLength:];
  puts(s, style, colStart, row, lastOfTrimmedFin);
}
