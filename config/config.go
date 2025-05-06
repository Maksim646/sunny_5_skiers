package config

type Config struct {
	ConfigPath            string `envconfig:"CONFIG_PATH" default:"../config.json"`
	EventsPath            string `envconfig:"EVENTS_PATH" default:"../events"`
	OutputFilePath        string `envconfig:"OUTPUT_FILE_PATH" default:"../output_events_log.txt"`
	ResultTablePath       string `envconfig:"RESULT_TABLE_PATH" default:"../result_table.txt"`
	TimeFormat            string `envconfig:"TIME_FORMAT" default:"15:04:05.000"`
	TimeDurationFormat    string `envconfig:"TIME_DURATION_FORMAT" default:"15:04:05"`
	ReportTableTimeFormat string `envconfig:"REPORT_TABLE_TIME_FORMAT" default:"%02d:%02d:%02d.%03d"`
	TargetsInFireLine     int    `envconfig:"TARGETS_IN_FIRE_LINE" default:"5"`
}
