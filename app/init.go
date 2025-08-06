package app

import (
	"time"

	"github.com/axelrindle/battery-notifier/notify"
	"github.com/go-co-op/gocron/v2"
	"go.uber.org/zap"
)

func (a *App) Init() {
	a.initScheduler()
	a.initNotifier()
}

func (a *App) initScheduler() {
	scheduler, err := gocron.NewScheduler()
	if err != nil {
		a.Logger.Fatal("failed to create a scheduler", zap.Error(err))
	}
	a.scheduler = scheduler

	a.scheduler.NewJob(
		gocron.OneTimeJob(gocron.OneTimeJobStartImmediately()),
		gocron.NewTask(a.readBatteryStatus),
	)
	a.scheduler.NewJob(
		gocron.DurationJob(time.Second*time.Duration(a.Config.Interval)),
		gocron.NewTask(a.readBatteryStatus),
	)
	a.scheduler.Start()
	a.Logger.Info("scheduler started", zap.Int64("interval", a.Config.Interval))
}

func (a *App) initNotifier() {
	a.notifier = notify.Notifier{
		Config: a.Config,
		Logger: a.Logger,
	}
}
