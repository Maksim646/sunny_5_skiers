package _test

import (
	"reflect"
	"testing"
	"time"

	"github.com/Maksim646/sunny_5_skiers/handler"
	"github.com/Maksim646/sunny_5_skiers/model"
	"github.com/stretchr/testify/assert"
)

func TestParseConfig(t *testing.T) {
	config, err := handler.ParseConfig("test_config.json", "15:04:05.000", "15:04:05")
	if err != nil {
		t.Errorf("Error parsing config: %v", err)
		return
	}

	startTime, _ := time.Parse("15:04:05.000", "10:00:00.000")

	testConfig := model.Config{
		Laps:        2,
		LapLen:      3500,
		PenaltyLen:  150,
		FiringLines: 2,
		StartRaw:    "10:00:00.000",
		DeltaRaw:    "00:01:30",
		Start:       startTime,
		StartDelta:  time.Duration(1*time.Minute + 30*time.Second),
	}

	if !reflect.DeepEqual(config, testConfig) {
		t.Errorf("Config does not match test config\nExpected: %+v\nGot: %+v", testConfig, config)
	} else {
		t.Logf("Config matches the expected values")
	}

	t.Logf("Expected config: %+v", testConfig)
	t.Logf("Parsed config: %+v", config)

	assert.NoError(t, err)
}

func TestParseEvents(t *testing.T) {
	events, err := handler.ParseEvents("test_events", "15:04:05.000")
	if err != nil {
		t.Errorf("Error parsing events: %v", err)
		return
	}

	parse := func(ts string) time.Time {
		tm, err := time.Parse("15:04:05.000", ts)
		if err != nil {
			t.Fatalf("Failed to parse time: %v", err)
		}
		return tm
	}

	expected := []model.CompetitorEvent{
		{Time: parse("09:05:59.867"), ID: 1, Competitor: 7},
		{Time: parse("09:15:00.841"), ID: 2, Competitor: 7, ExtraParams: "09:30:00.000"},
		{Time: parse("09:29:45.734"), ID: 3, Competitor: 7},
		{Time: parse("09:30:01.005"), ID: 4, Competitor: 7},
		{Time: parse("09:49:31.659"), ID: 5, Competitor: 7, ExtraParams: "1"},
		{Time: parse("09:49:33.123"), ID: 6, Competitor: 7, ExtraParams: "1"},
		{Time: parse("09:49:34.650"), ID: 6, Competitor: 7, ExtraParams: "2"},
		{Time: parse("09:49:35.937"), ID: 6, Competitor: 7, ExtraParams: "4"},
		{Time: parse("09:49:37.364"), ID: 6, Competitor: 7, ExtraParams: "5"},
		{Time: parse("09:49:38.339"), ID: 7, Competitor: 7},
		{Time: parse("09:49:55.915"), ID: 8, Competitor: 7},
		{Time: parse("09:51:48.391"), ID: 9, Competitor: 7},
		{Time: parse("09:59:03.872"), ID: 10, Competitor: 7},
		{Time: parse("09:59:03.872"), ID: 11, Competitor: 7, ExtraParams: "Lost in the forest"},
	}

	if len(events) != len(expected) {
		t.Fatalf("Expected %d events, got %d", len(expected), len(events))
	}

	for i, ev := range expected {
		assert.Equal(t, ev.ID, events[i].ID, "ID mismatch at index %d", i)
		assert.Equal(t, ev.Competitor, events[i].Competitor, "Competitor mismatch at index %d", i)
		assert.WithinDuration(t, ev.Time, events[i].Time, time.Millisecond, "Time mismatch at index %d", i)
		assert.Equal(t, ev.ExtraParams, events[i].ExtraParams, "ExtraParams mismatch at index %d", i)
	}
}
