package tcell_helpers

import (
	"testing"
)

func TestTextBoxFormatString(t *testing.T) {
	var result = textBoxFormatString("hello, this is a text to format", 10)

	for i, thing := range result {
		switch i {
		case 0:
			if thing != "hello, " {
				t.Error("Failed on 0, was: " + thing)
			}
			break
		case 1:
			if thing != "this is " {
				t.Error("Failed on 1, was: " + thing)
			}
			break
		case 2:
			if thing != "a text " {
				t.Error("Failed 2, was: " + thing)
			}
			break
		case 3:
			if thing != "to " {
				t.Error("Failed on 3, was: " + thing)
			}
			break
		case 4:
			if thing != "format" {
				t.Error("Failed on 4, was: " + thing)
			}
			break
		default:
			t.Fail()
			break
		}
	}
}

func TestTextBoxFormatString_StringWidthSmallerThanTextBox(t *testing.T) {
	var result = textBoxFormatString("hello, my friend", 20)

	for i, thing := range result {
		switch i {
		case 0:
			if thing != "hello, my friend" {
				t.Error("Failed on 0, was: " + thing)
			}
			break
		default:
			t.Fail()
			break
		}
	}
}
