package model

import "time"

type Config struct {
	Laps        int `json:"laps"`
	LapLen      int `json:"lapLen"`
	PenaltyLen  int `json:"penaltyLen"`
	FiringLines int `json:"firingLines"`

	StartRaw string `json:"start"`
	DeltaRaw string `json:"startDelta"`

	Start      time.Time     `json:"-"`
	StartDelta time.Duration `json:"-"`
}
