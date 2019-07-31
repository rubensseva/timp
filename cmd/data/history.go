package data

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"time"
	"timp/cmd/data/model"
	"timp/cmd/utility"
)

func max(i int, j int) int {
	if i > j {
		return i
	}
	return j
}

func readAllHistoryEntries() []model.PlayedEntry {
	historyFile, fileErr := ioutil.ReadFile("cmd/data/json/history.json")
	if fileErr != nil {
		panic(fileErr)
	}
	var historyEntries []model.PlayedEntry
	JSONErr := json.Unmarshal([]byte(historyFile), &historyEntries)
	if JSONErr != nil {
		panic(JSONErr)
	}
	if len(historyEntries) == 0 {
		panic("Trying to get history, but no history exists. Generate some history first.")
	}
	return historyEntries
}

func readAllHistoryEntriesUnsafe() []model.PlayedEntry {
	historyFile, fileErr := ioutil.ReadFile("cmd/data/json/history.json")
	if fileErr != nil {
		fmt.Println(fileErr)
	}
	var historyEntriesJSON []model.PlayedEntryJSON
	JSONErr := json.Unmarshal([]byte(historyFile), &historyEntriesJSON)
	if JSONErr != nil {
		fmt.Println(JSONErr)
	}
	if len(historyEntriesJSON) == 0 {
		fmt.Println("Trying to get history, but no history exists. Generate some history first.")
	}
	return model.PlayedEntryJSONListToRegular(historyEntriesJSON)
}

// GetAllHistoryEntries returns all history entries from json file
func GetAllHistoryEntries() []model.PlayedEntry {
	return readAllHistoryEntriesUnsafe()
}

// AppendToHistory appends one history entry to history json file.
// Auto increments the id
func AppendToHistory(text model.Text, player string, timeSpent time.Duration, didFinishLegally bool) {
	fmt.Println("In apend to history")
	var playedEntries = readAllHistoryEntriesUnsafe()
	var maxID = 0
	for _, entry := range playedEntries {
		maxID = max(maxID, entry.GetID())
	}

	// wpm is float32
	var wpm = utility.CalcWPM(text, timeSpent)

	var newHistory = model.NewPlayedEntry(maxID+1, text, player, time.Now(), float32(timeSpent.Seconds()), wpm, didFinishLegally)
	playedEntries = append(playedEntries, newHistory)
	writefile, JSONErr := json.MarshalIndent(model.PlayedEntryListToJSON(playedEntries), "", " ")
	if JSONErr != nil {
		panic(JSONErr)
	}
	fileErr := ioutil.WriteFile("cmd/data/json/history.json", writefile, 0644)
	if fileErr != nil {
		panic(fileErr)
	}
	fmt.Println("create text success (hopefully)")
}
