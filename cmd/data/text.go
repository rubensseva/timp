package data

import (
	"encoding/json"
	"io/ioutil"
	"timp/cmd/model"
	"timp/cmd/utility"
)

func readAllTexts() []model.Text {
	textfile, _ := ioutil.ReadFile("cmd/data/json/texts.json")
	var texts []model.Text
	_ = json.Unmarshal([]byte(textfile), &texts)
	if len(texts) == 0 {
		panic("Trying to get texts, but no text is created. Create a text first.")
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
