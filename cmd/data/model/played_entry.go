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
