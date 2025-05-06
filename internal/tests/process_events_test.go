package _test

import (
	"os"
	"testing"
	"time"

	"github.com/Maksim646/sunny_5_skiers/internal/handler"
	"github.com/Maksim646/sunny_5_skiers/model"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestProcessEvents(t *testing.T) {
	baseTime := time.Date(2025, time.May, 6, 10, 0, 0, 0, time.UTC)
	timeFormat := "15:04:05.000"

	t.Run("Valid", func(t *testing.T) {
		events := []model.CompetitorEvent{
			{ID: 1, Competitor: 1, Time: baseTime},
			{ID: 2, Competitor: 1, Time: baseTime.Add(10 * time.Second), ExtraParams: "09:30:00.000"},
			{ID: 3, Competitor: 1, Time: baseTime.Add(20 * time.Second)},
			{ID: 4, Competitor: 1, Time: baseTime.Add(40 * time.Second)},
			{ID: 5, Competitor: 1, Time: baseTime.Add(50 * time.Second), ExtraParams: "1"},
			{ID: 6, Competitor: 1, Time: baseTime.Add(60 * time.Second), ExtraParams: "1"},
			{ID: 6, Competitor: 1, Time: baseTime.Add(70 * time.Second), ExtraParams: "2"},
			{ID: 6, Competitor: 1, Time: baseTime.Add(80 * time.Second), ExtraParams: "4"},
			{ID: 6, Competitor: 1, Time: baseTime.Add(90 * time.Second), ExtraParams: "5"},
			{ID: 7, Competitor: 1, Time: baseTime.Add(100 * time.Second)},
			{ID: 8, Competitor: 1, Time: baseTime.Add(110 * time.Second)},
			{ID: 9, Competitor: 1, Time: baseTime.Add(120 * time.Second)},
			{ID: 10, Competitor: 1, Time: baseTime.Add(130 * time.Second)},
			{ID: 5, Competitor: 1, Time: baseTime.Add(140 * time.Second), ExtraParams: "1"},
			{ID: 6, Competitor: 1, Time: baseTime.Add(150 * time.Second), ExtraParams: "1"},
			{ID: 6, Competitor: 1, Time: baseTime.Add(160 * time.Second), ExtraParams: "2"},
			{ID: 6, Competitor: 1, Time: baseTime.Add(170 * time.Second), ExtraParams: "4"},
			{ID: 6, Competitor: 1, Time: baseTime.Add(180 * time.Second), ExtraParams: "5"},
			{ID: 7, Competitor: 1, Time: baseTime.Add(190 * time.Second)},
			{ID: 8, Competitor: 1, Time: baseTime.Add(200 * time.Second)},
			{ID: 9, Competitor: 1, Time: baseTime.Add(210 * time.Second)},
			{ID: 10, Competitor: 1, Time: baseTime.Add(220 * time.Second)},
		}

		actualPath := "test_process_events/test_valid_log_actual.txt"
		expectedPath := "test_process_events/test_valid_log_expected.txt"

		defer os.Remove(actualPath)

		err := handler.ProcessEvents(events, actualPath, timeFormat)
		require.NoError(t, err)

		actualContent, err := os.ReadFile(actualPath)
		require.NoError(t, err, "Cannot read actual log file")

		expectedContent, err := os.ReadFile(expectedPath)
		require.NoError(t, err, "Cannot read expected log file")

		assert.Equal(t, string(expectedContent), string(actualContent), "Log output does not match expected result")
	})

	t.Run("EveryLog", func(t *testing.T) {
		actualPath := "test_process_events/test_every_log.txt"
		expectedPath := "test_process_events/test_every_log_expected.txt"

		defer os.Remove(actualPath)

		events := []model.CompetitorEvent{
			{ID: 1, Competitor: 1, Time: baseTime},
			{ID: 2, Competitor: 1, Time: baseTime.Add(10 * time.Second), ExtraParams: "09:30:00.000"},
			{ID: 3, Competitor: 1, Time: baseTime.Add(20 * time.Second)},
			{ID: 4, Competitor: 1, Time: baseTime.Add(30 * time.Second)},
			{ID: 5, Competitor: 1, Time: baseTime.Add(40 * time.Second), ExtraParams: "1"},
			{ID: 6, Competitor: 1, Time: baseTime.Add(50 * time.Second), ExtraParams: "1"},
			{ID: 7, Competitor: 1, Time: baseTime.Add(60 * time.Second)},
			{ID: 8, Competitor: 1, Time: baseTime.Add(70 * time.Second)},
			{ID: 9, Competitor: 1, Time: baseTime.Add(80 * time.Second)},
			{ID: 10, Competitor: 1, Time: baseTime.Add(90 * time.Second)},
			{ID: 11, Competitor: 1, Time: baseTime.Add(100 * time.Second), ExtraParams: "Lost in the forest"},
		}

		err := handler.ProcessEvents(events, actualPath, timeFormat)
		require.NoError(t, err)

		actualContent, err := os.ReadFile(actualPath)
		require.NoError(t, err)

		expectedContent, err := os.ReadFile(expectedPath)
		require.NoError(t, err, "Cannot read expected log file")

		assert.Equal(t, string(expectedContent), string(actualContent), "Log output does not match expected result")
	})

}

func TestGenerateResultingTable(t *testing.T) {
	baseTime := time.Date(2025, time.May, 6, 10, 0, 0, 0, time.UTC)

	t.Run("Valid", func(t *testing.T) {
		events := []model.CompetitorEvent{
			{ID: 1, Competitor: 1, Time: baseTime},
			{ID: 2, Competitor: 1, Time: baseTime.Add(10 * time.Second), ExtraParams: "09:30:00.000"},
			{ID: 3, Competitor: 1, Time: baseTime.Add(20 * time.Second), ExtraParams: "09:30:00.000"},
			{ID: 4, Competitor: 1, Time: baseTime.Add(30 * time.Second)},
			{ID: 5, Competitor: 1, Time: baseTime.Add(40 * time.Second), ExtraParams: "1"},
			{ID: 6, Competitor: 1, Time: baseTime.Add(50 * time.Second), ExtraParams: "1"},
			{ID: 6, Competitor: 1, Time: baseTime.Add(60 * time.Second), ExtraParams: "2"},
			{ID: 6, Competitor: 1, Time: baseTime.Add(70 * time.Second), ExtraParams: "4"},
			{ID: 6, Competitor: 1, Time: baseTime.Add(80 * time.Second), ExtraParams: "5"},
			{ID: 7, Competitor: 1, Time: baseTime.Add(90 * time.Second)},
			{ID: 8, Competitor: 1, Time: baseTime.Add(100 * time.Second)},
			{ID: 9, Competitor: 1, Time: baseTime.Add(110 * time.Second)},
			{ID: 10, Competitor: 1, Time: baseTime.Add(120 * time.Second)},
			{ID: 5, Competitor: 1, Time: baseTime.Add(130 * time.Second), ExtraParams: "1"},
			{ID: 6, Competitor: 1, Time: baseTime.Add(140 * time.Second), ExtraParams: "1"},
			{ID: 6, Competitor: 1, Time: baseTime.Add(150 * time.Second), ExtraParams: "2"},
			{ID: 6, Competitor: 1, Time: baseTime.Add(160 * time.Second), ExtraParams: "4"},
			{ID: 6, Competitor: 1, Time: baseTime.Add(170 * time.Second), ExtraParams: "5"},
			{ID: 7, Competitor: 1, Time: baseTime.Add(180 * time.Second)},
			{ID: 8, Competitor: 1, Time: baseTime.Add(190 * time.Second)},
			{ID: 9, Competitor: 1, Time: baseTime.Add(200 * time.Second)},
			{ID: 10, Competitor: 1, Time: baseTime.Add(210 * time.Second)},
		}

		config := model.Config{
			Laps:        2,
			LapLen:      3500,
			PenaltyLen:  150,
			FiringLines: 2,
			Start:       baseTime,
			StartDelta:  1 * time.Minute,
		}

		actualPath := "test_process_events/test_result_table_actual.txt"
		expectedPath := "test_process_events/test_result_table_expected.txt"

		defer os.Remove(actualPath)

		err := handler.GenerateResultingTable(events, actualPath, "%02d:%02d:%02d.%03d", config, 5)
		require.NoError(t, err, "GenerateResultingTable returned error")

		actualContent, err := os.ReadFile(actualPath)
		require.NoError(t, err, "Cannot read actual result file")

		expectedContent, err := os.ReadFile(expectedPath)
		require.NoError(t, err, "Cannot read expected result file")

		assert.Equal(t, string(expectedContent), string(actualContent), "Generated result table does not match expected output")
	})

	t.Run("NotStarted", func(t *testing.T) {
		events := []model.CompetitorEvent{
			{ID: 1, Competitor: 1, Time: baseTime},
			{ID: 2, Competitor: 1, Time: baseTime.Add(10 * time.Second), ExtraParams: "09:30:00.000"},
			{ID: 3, Competitor: 1, Time: baseTime.Add(20 * time.Second), ExtraParams: "09:30:00.000"},
		}

		config := model.Config{
			Laps:        2,
			LapLen:      3500,
			PenaltyLen:  150,
			FiringLines: 2,
			Start:       baseTime,
			StartDelta:  1 * time.Minute,
		}

		actualPath := "test_process_events/test_result_table_not_started_actual.txt"
		expectedPath := "test_process_events/test_result_table_not_started_expected.txt"

		defer os.Remove(actualPath)

		err := handler.GenerateResultingTable(events, actualPath, "%02d:%02d:%02d.%03d", config, 5)
		require.NoError(t, err, "GenerateResultingTable returned error")

		actualContent, err := os.ReadFile(actualPath)
		require.NoError(t, err, "Cannot read actual result file")

		expectedContent, err := os.ReadFile(expectedPath)
		require.NoError(t, err, "Cannot read expected result file")

		assert.Equal(t, string(expectedContent), string(actualContent), "Generated result table does not match expected output")

	})

	t.Run("NotFinished", func(t *testing.T) {
		events := []model.CompetitorEvent{
			{ID: 1, Competitor: 1, Time: baseTime},
			{ID: 2, Competitor: 1, Time: baseTime.Add(10 * time.Second), ExtraParams: "09:30:00.000"},
			{ID: 3, Competitor: 1, Time: baseTime.Add(20 * time.Second), ExtraParams: "09:30:00.000"},
			{ID: 4, Competitor: 1, Time: baseTime.Add(30 * time.Second)},
			{ID: 11, Competitor: 1, Time: baseTime.Add(40 * time.Second), ExtraParams: "He was too tired and went home"},
		}

		config := model.Config{
			Laps:        2,
			LapLen:      3500,
			PenaltyLen:  150,
			FiringLines: 2,
			Start:       baseTime,
			StartDelta:  1 * time.Minute,
		}

		actualPath := "test_process_events/test_result_table_not_finished_actual.txt"
		expectedPath := "test_process_events/test_result_table_not_finished_expected.txt"

		defer os.Remove(actualPath)

		err := handler.GenerateResultingTable(events, actualPath, "%02d:%02d:%02d.%03d", config, 5)
		require.NoError(t, err, "GenerateResultingTable returned error")

		actualContent, err := os.ReadFile(actualPath)
		require.NoError(t, err, "Cannot read actual result file")

		expectedContent, err := os.ReadFile(expectedPath)
		require.NoError(t, err, "Cannot read expected result file")

		assert.Equal(t, string(expectedContent), string(actualContent), "Generated result table does not match expected output")

	})

}
