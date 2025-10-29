package main

import (
	"flag"
	"log"
	"os"
	"os/signal"
	"syscall"

	_ "embed"

	"github.com/axelrindle/battery-notifier/app"
	"github.com/axelrindle/battery-notifier/config"
	"github.com/axelrindle/battery-notifier/version"
	"go.uber.org/zap"
)

//go:embed banner.txt
var banner string

var (
	showVersion bool
	configFile  string
)

func init() {
	println(banner)
}

func main() {
	flag.BoolVar(&showVersion, "version", false, "show program version")
	flag.StringVar(&configFile, "config", "config.yml", "path to the config file")
	flag.Parse()
	if showVersion {
		println(version.BuildVersion())
		return
	}

	config := &config.Config{}
	config.Load(configFile)

	logger, err := makeLogger(config)
	if err != nil {
		log.Fatal(err)
	}

	if err := validateDevice(config.DevicePath); err != nil {
		log.Fatal(err)
	}

	logger.Info("starting program", zap.String("mode", config.Environment))

	app := &app.App{
		Config: config,
		Logger: logger,
	}
	app.Init()

	signalChan := make(chan os.Signal, 1)
	signal.Notify(signalChan, syscall.SIGINT, syscall.SIGTERM)

	<-signalChan

	logger.Info("shutting down")
	app.Shutdown()
}
