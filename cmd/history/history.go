package history

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

func AppendToHistory(text model.Text, player string, timeSpent time.Duration, didFinishLegally bool) {
	fmt.Println("In apend to history")

	textfile, _ := ioutil.ReadFile("cmd/resources/history.json")
	var playedEntries []model.PlayedEntry
	_ = json.Unmarshal([]byte(textfile), &playedEntries)
	var maxId int = 0
	for _, entry := range playedEntries {
		maxId = max(maxId, entry.Id)
	}

	var wpm float32 = utility.CalcWPM(text, timeSpent)
	var newHistory = model.PlayedEntry{Id: maxId + 1, Text: text, Player: player, TimePlayed: time.Now(), TimeSpent: float32(timeSpent.Seconds()), Wpm: wpm, DidFinishLegally: didFinishLegally}

	playedEntries = append(playedEntries, newHistory)
	writefile, _ := json.MarshalIndent(playedEntries, "", " ")
	_ = ioutil.WriteFile("cmd/resources/history.json", writefile, 0644)
	fmt.Println("create text success (hopefully)")
}
