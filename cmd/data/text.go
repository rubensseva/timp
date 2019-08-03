package data

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"timp/cmd/data/model"
	"timp/cmd/utility"
)

func readAllTexts() []model.Text {
	textfile, fileErr := ioutil.ReadFile("cmd/data/json/texts.json")
	if fileErr != nil {
		panic(fileErr)
	}
	var textsJSON []model.TextJSON
	JSONErr := json.Unmarshal([]byte(textfile), &textsJSON)
	if JSONErr != nil {
		panic(JSONErr)
	}
	if len(textsJSON) == 0 {
		panic("Trying to get texts, but no text is created. Create a text first.")
	}
	return model.TextJSONListToRegular(textsJSON)
}

// readAllTextsUnsafe returns all texts but does not panic
// if result is null
func readAllTextsUnsafe() []model.Text {
	textfile, fileErr := ioutil.ReadFile("cmd/data/json/texts.json")
	if fileErr != nil {
		fmt.Println(fileErr)
	}
	var textsJSON []model.TextJSON
	JSONErr := json.Unmarshal([]byte(textfile), &textsJSON)
	if JSONErr != nil {
		fmt.Println(JSONErr)
	}
	if len(textsJSON) == 0 {
		fmt.Println("Trying to get texts, but no text is created. Create a text first.")
	}
	return model.TextJSONListToRegular(textsJSON)
}

// GetAllTexts returns all texts from json file
func GetAllTexts() []model.Text {
	var texts = readAllTextsUnsafe()
	return texts
}

// GetRandomText gets all texts from json file, and returns a
// random one of these
func GetRandomText() model.Text {
	var texts = readAllTexts()
	var randIndex = utility.RandomGen(0, len(texts))
	return texts[randIndex]
}

// AddText adds a single text to text json file
func AddText(text model.Text) {
	var texts = readAllTextsUnsafe()

	var isAText = false
	for _, textEntry := range texts {
		if textEntry.GetText() == text.GetText() {
			isAText = true
		}
	}

	if isAText {
		fmt.Println("specified text " + text.GetText() + " is already a text")
		return
	}

  var author = "unknown"
  if text.GetAuthor() != "" {
    author = text.GetAuthor()
  }

	fmt.Println("creating text ", text.GetText(), " with author: ", author)
	var newText = model.NewText(text.GetText(), author)
	texts = append(texts, newText)
	writefile, JSONErr := json.MarshalIndent(model.TextListToJSON(texts), "", " ")
	if JSONErr != nil {
		fmt.Println(JSONErr)
	}
	fileErr := ioutil.WriteFile("cmd/data/json/texts.json", writefile, 0644)
	if fileErr != nil {
		fmt.Println(fileErr)
	}
	fmt.Println("create text success (hopefully)")
}
