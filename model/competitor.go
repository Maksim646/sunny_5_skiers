package model

import "time"

var (
	CompetitorStarted     = "started"
	CompetitorNotStarted  = "NotStarted"
	CompetitorNotFinished = "NotFinished"
)

const (
	EventRegistered       = 1
	EventStartTimeSet     = 2
	EventOnTheStartLine   = 3
	EventStart            = 4
	EventOnTheFiringRange = 5
	EventTargetHit        = 6
	EventLeftFiringRange  = 7
	EventPenaltyLapStart  = 8
	EventPenaltyLapEnd    = 9
	EventLapCompleted     = 10
	EventNotFinished      = 11
)

var (
	Comments = map[int]string{
		1:  "registered",
		2:  "The start time was set by a draw",
		3:  "is on the start line",
		4:  "has started",
		5:  "is on the firing range",
		6:  "The target has been hit",
		7:  "left the firing range",
		8:  "entered the penalty laps",
		9:  "left the penalty laps",
		10: "ended the main lap",
		11: "can`t continue",
	}
)

type CompetitorEvent struct {
	Time        time.Time
	ID          int
	Competitor  int
	ExtraParams string
}

type LapInfo struct {
	Time  time.Duration
	Speed float64
}

type CompetitorReport struct {
	CompetitorID int
	Status       string // "OK", "NotStarted", "NotFinished"
	TotalTime    time.Duration
	Laps         []LapInfo
	PenaltyLaps  []LapInfo
	Hits         int
	Shots        int
}
