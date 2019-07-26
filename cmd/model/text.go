package model

import (
	"time"
)

// Represented a text
type Text struct {
	Text   string
	Author string
}

// Represents values for a single play
type PlayedEntry struct {
	Id         int
	Text       Text
	Player     string
	TimePlayed time.Time
	TimeSpent  time.Duration
	Wpm        float32
}
