package controller

import (
	"bufio"
	"encoding/json"
	"os"
	"strconv"
	"strings"
	"time"

	"github.com/Maksim646/sunny_5_skiers/model"
)

func ParseEvents(eventPath string, eventTimeFormat string) ([]model.CompetitorEvent, error) {
	file, err := os.Open(eventPath)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	var events []model.CompetitorEvent
	scanner := bufio.NewScanner(file)

	for scanner.Scan() {
		line := strings.TrimSpace(scanner.Text())
		if line == "" {
			continue
		}

		parts := strings.Fields(line)
		if len(parts) < 3 {
			return nil, err
		}

		timeStr := strings.Trim(parts[0], "[]")
		eventTime, err := time.Parse(eventTimeFormat, timeStr)
		if err != nil {
			return nil, err
		}

		eventID, err := strconv.Atoi(parts[1])
		if err != nil {
			return nil, err
		}

		competitorID, err := strconv.Atoi(parts[2])
		if err != nil {
			return nil, err
		}

		var extra string
		if len(parts) > 3 {
			extra = strings.Join(parts[3:], " ")
		}

		event := model.CompetitorEvent{
			Time:        eventTime,
			ID:          eventID,
			Competitor:  competitorID,
			ExtraParams: extra,
		}

		events = append(events, event)
	}

	if err := scanner.Err(); err != nil {
		return nil, err
	}

	return events, nil
}

func ParseConfig(path string, timeFormat string, timeDurationFormat string) (model.Config, error) {
	var config model.Config

	file, err := os.Open(path)
	if err != nil {
		return config, err
	}
	defer file.Close()

	decoder := json.NewDecoder(file)
	if err := decoder.Decode(&config); err != nil {
		return config, err
	}

	startTime, err := time.Parse(timeFormat, config.StartRaw)
	if err != nil {
		return config, err
	}
	config.Start = startTime

	deltaParsed, err := time.Parse(timeDurationFormat, config.DeltaRaw)
	if err != nil {
		return config, err
	}

	config.StartDelta = time.Duration(
		deltaParsed.Hour()*int(time.Hour) +
			deltaParsed.Minute()*int(time.Minute) +
			deltaParsed.Second()*int(time.Second),
	)

	return config, nil
}
