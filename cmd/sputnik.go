package main

import (
	"flag"
	log "github.com/Sirupsen/logrus"
	"github.com/fsnotify/fsnotify"
	"github.com/spf13/viper"
	"os"
	"os/signal"
	"path/filepath"
	"sputnik/core"
	"sputnik/handle"
)

var (
	Version string
	Ver     bool
	Build   string
	LogFile string
	err     error
	logFile *os.File
)

func init() {
	flag.BoolVar(&Ver, "version", false, "print version")
}

func main() {
	// Parse CLI flags
	flag.Parse()

	logger := log.StandardLogger()

	if Ver {
		logger.Infof("version: %s %s\n", Version, Build)
		return
	}

	// Create channel to receive message when main process should be stopped
	StopChan := make(chan bool)

	configInit()

	c := make(chan os.Signal, 1)
	signal.Notify(c, os.Interrupt)
	go func() {
		for sig := range c {
			logger.Infof("Signal %#v", sig)
			break
		}
		StopChan <- true
	}()

	coreInit()

	// Wait to stop
	<-StopChan

	logger.Infoln("Sputnik stopped normaly")

	return
}

func configInit() {
	viper.SetConfigName(filepath.Base(os.Args[0]))

	// get the path of this file location (it's not the same as 'current path')
	arg0path, err := filepath.Abs(filepath.Dir(os.Args[0]))
	if err != nil {
		// trying to looking for at the 'current path'
		viper.AddConfigPath(".")
	} else {
		log.Printf("%#v", arg0path)
		viper.AddConfigPath(arg0path)
	}

	// extend the looking area
	viper.AddConfigPath("/etc/sputnik/")
	viper.AddConfigPath("$HOME/.sputnik")

	configReload()

	// hot reload on any modification
	viper.WatchConfig()
	viper.OnConfigChange(func(e fsnotify.Event) {
		log.Println("Config file changed:", e.Name)
		configReload()
	})
}

func configReload() {
	err := viper.ReadInConfig()
	if err != nil {
		return
	}

	logfile := viper.Get("log")
	if logFile != nil {
		logFile.Close()
		log.SetOutput(os.Stderr)
	}
	if logfile != nil && logfile.(string) != "" {
		if logFile, err = os.OpenFile(logfile.(string), os.O_RDWR|os.O_CREATE, 0644); err != nil {
			log.Errorln("%s", err)
		}
		log.SetOutput(logFile)
	}
}

func coreInit() {
	cache := core.Cache{}
	cache.Init("cache.db")

	if handle.ScenarioSupported() {
		log.Info("Scenario support is enabled")
		scenario := handle.Scenario{}
		scenario.Init()
	} else {
		log.Info("Scenario support is disabled")
	}

	if handle.ControlEnabled() {
		log.Info("Control interface is enabled")
		control := handle.Control{}
		control.Init()
	} else {
		log.Info("Control interface is disabled")
	}

	core.Bootstrap()
	return
}
