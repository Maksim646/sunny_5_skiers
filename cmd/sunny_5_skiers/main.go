package main

import (
	"github.com/Maksim646/sunny_5_skiers/config"
	"github.com/Maksim646/sunny_5_skiers/handler"
	logger "github.com/Maksim646/sunny_5_skiers/pkg"
	"github.com/kelseyhightower/envconfig"
	"go.uber.org/zap"
)

var cfg config.Config

func main() {
	envconfig.MustProcess("", &cfg)

	logger.InitLogger()
	defer zap.L().Sync()
	zap.L().Info("Logger started")

	parsedConfig, err := handler.ParseConfig(cfg.ConfigPath, cfg.TimeFormat, cfg.TimeDurationFormat)
	if err != nil {
		zap.L().Error("error load config", zap.Error(err))
	}

	events, err := handler.ParseEvents(cfg.EventsPath, cfg.TimeFormat)
	if err != nil {
		zap.L().Error("error parse events", zap.Error(err))
	}

	err = handler.ProcessEvents(events, cfg.OutputFilePath, cfg.TimeFormat)
	if err != nil {
		zap.L().Error("error process events", zap.Error(err))
	}

	err = handler.GenerateResultingTable(events, cfg.ResultTablePath, cfg.ReportTableTimeFormat, parsedConfig, cfg.TargetsInFireLine)
	if err != nil {
		zap.L().Error("error process events", zap.Error(err))
	}
}
