package controller

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strings"
	"time"

	"github.com/Maksim646/sunny_5_skiers/model"
	"go.uber.org/zap"
)

func GenerateResultingTable(events []model.CompetitorEvent, resultTablePath string, timeFormat string, config model.Config, targetsInFireLine int) error {
	resultTableFile, err := os.OpenFile(resultTablePath, os.O_APPEND|os.O_CREATE|os.O_WRONLY|os.O_TRUNC, 0644)
	if err != nil {
		return err
	}
	defer resultTableFile.Close()

	resultTableFileWriter := bufio.NewWriter(resultTableFile)

	competitorEvents := make(map[int][]model.CompetitorEvent)
	for _, event := range events {
		if event.Competitor == 0 {
			zap.L().Info(fmt.Sprintf("warning: event without competitor ID: %+v", event))
		}
		competitorEvents[event.Competitor] = append(competitorEvents[event.Competitor], event)
	}

	var sortedReports []model.CompetitorReport
	for competitorID, events := range competitorEvents {
		sort.Slice(events, func(i, j int) bool { return events[i].Time.Before(events[j].Time) })
		report := processCompetitorEvents(competitorID, events, config, timeFormat, targetsInFireLine)
		sortedReports = append(sortedReports, report)
	}

	sort.Slice(sortedReports, func(i, j int) bool {
		a, b := sortedReports[i], sortedReports[j]

		if (a.Status == model.CompetitorNotFinished || a.Status == model.CompetitorNotStarted) &&
			(b.Status != model.CompetitorNotFinished && b.Status != model.CompetitorNotStarted) {
			return false
		}
		if (b.Status == model.CompetitorNotFinished || b.Status == model.CompetitorNotStarted) &&
			(a.Status != model.CompetitorNotFinished && a.Status != model.CompetitorNotStarted) {
			return true
		}

		return a.TotalTime < b.TotalTime
	})

	for _, report := range sortedReports {
		reportLine := formatCompetitorReport(report, timeFormat, config)
		_, err := resultTableFileWriter.WriteString(reportLine + "\n")
		if err != nil {
			return fmt.Errorf("could not write report to file: %w", err)
		}
	}

	if err := resultTableFileWriter.Flush(); err != nil {
		return fmt.Errorf("could not flush report buffer: %w", err)
	}

	return nil
}

func processCompetitorEvents(competitorID int, events []model.CompetitorEvent, config model.Config, timeFormat string, targetsInFireLine int) model.CompetitorReport {

	var (
		lapsInfo            []model.LapInfo
		penaltyLapsInfo     []model.LapInfo
		startTime           time.Time
		finishTime          time.Time
		lapStartTime        time.Time
		penaltyLapStartTime time.Time
		laps                int
		hits                int
		speed               float64
	)
	status := model.CompetitorNotStarted

	for _, e := range events {
		switch e.ID {
		case model.EventStart:
			startTime = e.Time
			status = model.CompetitorStarted
			lapStartTime = e.Time
		case model.EventTargetHit:
			hits += 1
		case model.EventPenaltyLapStart:
			penaltyLapStartTime = e.Time
		case model.EventPenaltyLapEnd:
			currentPenaltyLapEnd := e.Time
			penaltyLapDuration := currentPenaltyLapEnd.Sub(penaltyLapStartTime)

			penaltyLapInfo := model.LapInfo{
				Time:  penaltyLapDuration,
				Speed: float64(config.PenaltyLen) / penaltyLapDuration.Seconds(),
			}

			penaltyLapsInfo = append(penaltyLapsInfo, penaltyLapInfo)
			penaltyLapStartTime = currentPenaltyLapEnd
		case model.EventLapCompleted:
			laps += 1
			if laps == config.Laps {
				finishTime = e.Time
			}

			currentLapEnd := e.Time
			lapDuration := currentLapEnd.Sub(lapStartTime)

			seconds := lapDuration.Seconds()
			if seconds > 0 {
				speed = float64(config.LapLen) / seconds
			}

			lapInfo := model.LapInfo{
				Time:  lapDuration,
				Speed: speed,
			}
			lapsInfo = append(lapsInfo, lapInfo)
			lapStartTime = currentLapEnd

		case model.EventNotFinished:
			status = model.CompetitorNotFinished
		}
	}

	if status == model.CompetitorStarted && finishTime.IsZero() {
		status = model.CompetitorNotFinished
	}

	totalTime := time.Duration(0)
	if !startTime.IsZero() && !finishTime.IsZero() {
		totalTime = finishTime.Sub(startTime)
	}

	return model.CompetitorReport{
		CompetitorID: competitorID,
		TotalTime:    totalTime,
		Status:       status,
		Laps:         lapsInfo,
		PenaltyLaps:  penaltyLapsInfo,
		Hits:         hits,
		Shots:        targetsInFireLine * config.FiringLines,
	}
}

func formatCompetitorReport(report model.CompetitorReport, reportTableTimeFormat string, config model.Config) string {
	var sb strings.Builder

	if report.Status == "NotStarted" || report.Status == "NotFinished" {
		sb.WriteString(fmt.Sprintf("[%s] %d ", report.Status, report.CompetitorID))
	} else {
		sb.WriteString(fmt.Sprintf("[%s] %d ", formatDuration(report.TotalTime, reportTableTimeFormat), report.CompetitorID))
	}

	sb.WriteString(formatLapList(report.Laps, config.Laps, reportTableTimeFormat))
	sb.WriteString(" ")
	sb.WriteString(formatLapList(report.PenaltyLaps, config.FiringLines, reportTableTimeFormat))
	sb.WriteString(" ")

	sb.WriteString(fmt.Sprintf("%d/%d", report.Hits, report.Shots))

	return sb.String()
}

func formatLapList(laps []model.LapInfo, expectedCount int, timeFmt string) string {
	var sb strings.Builder
	sb.WriteString("[")

	for i := 0; i < expectedCount; i++ {
		if i > 0 {
			sb.WriteString(", ")
		}
		if i < len(laps) {
			sb.WriteString(fmt.Sprintf("{%s, %.3f}", formatDuration(laps[i].Time, timeFmt), laps[i].Speed))
		} else {
			sb.WriteString("{,}")
		}
	}

	sb.WriteString("]")
	return sb.String()
}
