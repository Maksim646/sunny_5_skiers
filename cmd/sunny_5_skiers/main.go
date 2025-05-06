package main

import (
	"github.com/Maksim646/sunny_5_skiers/config"
	"github.com/Maksim646/sunny_5_skiers/internal/controller"
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

	parsedConfig, err := controller.ParseConfig(cfg.ConfigPath, cfg.TimeFormat, cfg.TimeDurationFormat)
	if err != nil {
		zap.L().Error("error load config", zap.Error(err))
	}

	events, err := controller.ParseEvents(cfg.EventsPath, cfg.TimeFormat)
	if err != nil {
		zap.L().Error("error parse events", zap.Error(err))
	}

	err = controller.ProcessEvents(events, cfg.OutputFilePath, cfg.TimeFormat)
	if err != nil {
		zap.L().Error("error process events", zap.Error(err))
	}

	err = controller.GenerateResultingTable(events, cfg.ResultTablePath, cfg.ReportTableTimeFormat, parsedConfig, cfg.TargetsInFireLine)
	if err != nil {
		zap.L().Error("error process events", zap.Error(err))
	}
}
