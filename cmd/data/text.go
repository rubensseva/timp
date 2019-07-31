package data

import (
	"encoding/json"
	"fmt"
  "log"
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
	return model.TextJSONListToTRegular(textsJSON)
}

// readAllTextsUnsafe returns all texts but does not panic
// if result is null
func readAllTextsUnsafe() []model.Text {
	textfile, fileErr := ioutil.ReadFile("cmd/data/json/texts.json")
  if fileErr != nil {
    panic(fileErr)
  }
	var texts []model.Text
  JSONErr := json.Unmarshal([]byte(textfile), &texts)
  if JSONErr != nil {
    panic(JSONErr)
  }
	return texts
}

// GetAllTexts returns all texts from json file
func GetAllTexts() []model.Text {
	var texts = readAllTexts()
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

	fmt.Println("creating text ", text.GetText())
	var newText = model.NewText(text.GetText(), text.GetAuthor())
	texts = append(texts, newText)
	writefile, JSONErr := json.MarshalIndent(model.TextListToJSON(texts), "", " ")
  if JSONErr != nil {
    log.Fatal(JSONErr)
  }
  fileErr := ioutil.WriteFile("cmd/data/json/texts.json", writefile, 0644)
  if fileErr != nil {
    log.Fatal(fileErr)
  }
	fmt.Println("create text success (hopefully)")
}
