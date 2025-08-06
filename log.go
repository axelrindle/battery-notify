package main

import (
	"github.com/axelrindle/battery-notifier/config"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

func makeLogger(c *config.Config) (*zap.Logger, error) {
	conf := zap.NewDevelopmentConfig()

	if c.Environment == "production" {
		conf.Level = zap.NewAtomicLevelAt(zap.InfoLevel)
	}

	conf.EncoderConfig.EncodeLevel = zapcore.CapitalColorLevelEncoder

	return conf.Build()
}
