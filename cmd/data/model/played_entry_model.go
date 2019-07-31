package model

import "time"

// PlayedEntry represents values for a single play
type PlayedEntry struct {
	id               int
	text             Text
	player           string
	timePlayed       time.Time
	timeSpent        float32
	wpm              float32
	didFinishLegally bool
}

type PlayedEntryJSON struct {
  Id              int
  Text            TextJSON
  Player          string
  TimePlayed      time.Time
  TimeSpent       float32
  Wpm             float32
  DidFinishLegally bool
}

func NewPlayedEntry(id int, text Text, player string, timePlayed time.Time, timeSpent float32, wpm float32, didFinishLegally bool) PlayedEntry {
  return PlayedEntry {
    id,
    text,
    player,
    timePlayed,
    timeSpent,
    wpm,
    didFinishLegally }
}

func NewPlayedEntryCopy(p PlayedEntry) PlayedEntry {
  return PlayedEntry {
    p.id,
    p.text,
    p.player,
    p.timePlayed,
    p.timeSpent,
    p.wpm,
    p.didFinishLegally }
}

func (p PlayedEntry) toJSONobj() PlayedEntryJSON {
  return PlayedEntryJSON{p.id, p.text.ToJSONobj(), p.player, p.timePlayed, p.timeSpent, p.wpm, p.didFinishLegally}
}
func (p PlayedEntryJSON) ToRegularObj() PlayedEntry {
  return PlayedEntry{p.Id, p.Text.ToRegularObj(), p.Player, p.TimePlayed, p.TimeSpent, p.Wpm, p.DidFinishLegally}
}


func (p PlayedEntry) GetID() int {
  return p.id
}

func (p PlayedEntry) GetText() Text {
  return p.text
}

func (p PlayedEntry) GetPlayer() string {
  return p.player
}

func (p PlayedEntry) GetTimePlayed() time.Time {
  return p.timePlayed
}

func (p PlayedEntry) GetTimeSpent() float32 {
  return p.timeSpent
}

func (p PlayedEntry) GetWpm() float32 {
  return p.wpm
}

func (p PlayedEntry) GetDidFinishLegally() bool {
  return p.didFinishLegally
}

func PlayedEntryListToJSON(playedEntries []PlayedEntry) []PlayedEntryJSON {
  var playedEntriesJSON []PlayedEntryJSON
  for _, playedEntry := range playedEntries {
    playedEntriesJSON = append(playedEntriesJSON, playedEntry.toJSONobj())
  }
  return playedEntriesJSON
}

func PlayedEntryJSONListToRegular(playedEntriesJSON []PlayedEntryJSON) []PlayedEntry {
  var playedEntries []PlayedEntry
  for _, playedEntryJSON := range playedEntriesJSON {
    playedEntries = append(playedEntries, playedEntryJSON.ToRegularObj())
  }
  return playedEntries
}
