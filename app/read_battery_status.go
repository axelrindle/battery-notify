package app

import (
	"bytes"
	"fmt"
	"os"
	"path"
	"strconv"

	"go.uber.org/zap"
)

func (a *App) readFile(name string) ([]byte, error) {
	file := path.Join(a.Config.DevicePath, name)
	data, err := os.ReadFile(file)
	if err != nil {
		a.Logger.Error("failed reading device data", zap.Error(err), zap.String("file", file))
		return nil, err
	}

	return bytes.Trim(data, "\n"), nil
}

func (a *App) readBatteryStatus() {
	charge_now_s, _ := a.readFile("charge_now")
	charge_full_s, _ := a.readFile("charge_full")

	charge_now, _ := strconv.ParseFloat(string(charge_now_s), 64)
	charge_full, _ := strconv.ParseFloat(string(charge_full_s), 64)

	a.charge = (charge_now / charge_full) * 100.0

	a.Logger.Debug("got battery charge level", zap.String("charge", fmt.Sprintf("%.2f%%", a.charge)))

	a.Logger.Debug("calling notifiers", zap.Int("count", len(a.Config.Notifications)))

	failed := 0
	skipped := 0
	for _, n := range a.Config.Notifications {
		ran, err := a.notifier.Notify(a.charge, n)

		if err != nil {
			a.Logger.Error("notification failed", zap.String("notification", n.ID), zap.Error(err))
			failed++
		}
		if !ran {
			skipped++
		}
	}

	a.Logger.Debug("called notifiers", zap.Int("failed", failed), zap.Int("skipped", skipped))
}
