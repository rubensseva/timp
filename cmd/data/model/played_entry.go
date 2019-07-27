package model

import "time"

// PlayedEntry represents values for a single play
type PlayedEntry struct {
	ID               int
	Text             Text
	Player           string
	TimePlayed       time.Time
	TimeSpent        float32
	Wpm              float32
	DidFinishLegally bool
}
