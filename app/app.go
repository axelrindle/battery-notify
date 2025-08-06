package app

import (
	"github.com/axelrindle/battery-notifier/config"
	"github.com/axelrindle/battery-notifier/notify"
	"github.com/go-co-op/gocron/v2"
	"go.uber.org/zap"
)

type App struct {
	Config *config.Config
	Logger *zap.Logger

	scheduler gocron.Scheduler
	notifier  notify.Notifier

	charge float64
}
