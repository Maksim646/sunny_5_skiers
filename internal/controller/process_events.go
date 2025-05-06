package controller

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"time"

	"github.com/Maksim646/sunny_5_skiers/model"
)

func ProcessEvents(events []model.CompetitorEvent, outputFilePath string, timeFormat string) error {
	outputLogFile, err := os.OpenFile(outputFilePath, os.O_APPEND|os.O_CREATE|os.O_WRONLY|os.O_TRUNC, 0644)
	if err != nil {
		return err
	}
	defer outputLogFile.Close()

	outputLogFileWriter := bufio.NewWriter(outputLogFile)

	for _, event := range events {

		comment := formatEventWithComment(event, model.Comments, timeFormat)

		_, err := outputLogFileWriter.WriteString(comment + "\n")
		if err != nil {
			return fmt.Errorf("could not write to file: %w", err)
		}

	}

	if err := outputLogFileWriter.Flush(); err != nil {
		return fmt.Errorf("could not flush buffer: %w", err)
	}

	return nil
}

func formatEventWithComment(event model.CompetitorEvent, comments map[int]string, timeFormat string) string {
	timeStr := event.Time.Format(timeFormat)
	var msg string

	switch event.ID {
	case model.EventRegistered, model.EventOnTheStartLine, model.EventStart, model.EventLeftFiringRange, model.EventPenaltyLapStart, model.EventPenaltyLapEnd, model.EventLapCompleted:
		msg = fmt.Sprintf("The competitor(%d) %s", event.Competitor, comments[event.ID])
	case model.EventStartTimeSet:
		msg = fmt.Sprintf("The start time for the competitor(%d) was set by a draw to %s", event.Competitor, event.ExtraParams)
	case model.EventOnTheFiringRange:
		msg = fmt.Sprintf("The competitor(%d) %s(%s)", event.Competitor, comments[event.ID], event.ExtraParams)
	case model.EventTargetHit:
		msg = fmt.Sprintf("The target(%s) has been hit by competitor(%d)", event.ExtraParams, event.Competitor)
	case model.EventNotFinished:
		msg = fmt.Sprintf("The competitor(%d) %s: %s", event.Competitor, comments[event.ID], event.ExtraParams)
	default:
		msg = fmt.Sprintf("Unknown event ID (%d) for competitor(%d)", event.ID, event.Competitor)
	}

	return fmt.Sprintf("[%s] %s", timeStr, msg)
}

func formatDuration(d time.Duration, reportTableTimeFormat string) string {
	h := int(d.Hours())
	m := int(d.Minutes()) % 60
	s := int(d.Seconds()) % 60
	ms := int(d.Milliseconds()) % 1000
	return fmt.Sprintf(reportTableTimeFormat, h, m, s, ms)
}

func SortedEvents(events []model.CompetitorEvent) []model.CompetitorEvent {
	sorted := make([]model.CompetitorEvent, len(events))
	copy(sorted, events)

	sort.Slice(sorted, func(i, j int) bool {
		return sorted[i].Time.Before(sorted[j].Time)
	})

	return sorted
}
