package app

import "go.uber.org/zap"

func (a *App) Shutdown() {
	err := a.scheduler.Shutdown()
	if err != nil {
		a.Logger.Fatal("failed to shutdown the scheduler", zap.Error(err))
	}
}
