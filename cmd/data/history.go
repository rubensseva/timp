package data

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"time"
	"timp/cmd/model"
	"timp/cmd/utility"
)

func max(i int, j int) int {
	if i > j {
		return i
	}
	return j
}

func readAllHistoryEntries() []model.PlayedEntry {
	historyFile, _ := ioutil.ReadFile("cmd/resources/history.json")
	var historyEntries []model.PlayedEntry
	_ = json.Unmarshal([]byte(historyFile), &historyEntries)
	return historyEntries
}

// GetAllHistoryEntries returns all history entries from json file
func GetAllHistoryEntries() []model.PlayedEntry {
	return readAllHistoryEntries()
}

// AppendToHistory appends one history entry to history json file.
// Auto increments the id
func AppendToHistory(text model.Text, player string, timeSpent time.Duration, didFinishLegally bool) {
	fmt.Println("In apend to history")
	var playedEntries = readAllHistoryEntries()
	var maxID = 0
	for _, entry := range playedEntries {
		maxID = max(maxID, entry.Id)
	}

	// wpm is float32
	var wpm = utility.CalcWPM(text, timeSpent)

	var newHistory = model.PlayedEntry{Id: maxID + 1, Text: text, Player: player, TimePlayed: time.Now(), TimeSpent: float32(timeSpent.Seconds()), Wpm: wpm, DidFinishLegally: didFinishLegally}
	playedEntries = append(playedEntries, newHistory)
	writefile, _ := json.MarshalIndent(playedEntries, "", " ")
	_ = ioutil.WriteFile("cmd/data/json/history.json", writefile, 0644)
	fmt.Println("create text success (hopefully)")
}
