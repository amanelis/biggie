package main

import (
	"os"
	"path/filepath"

	"github.com/Sirupsen/logrus"
	exchange "github.com/preichenberger/go-coinbase-exchange"
)

type App struct {
	client *exchange.Client
	config ConfigReader
	helper Helper
	logger *logrus.Logger
}

func LoadLogger(config ConfigReader) *logrus.Logger {
	log := logrus.New()
	env := config.GetString(Environment)

	if env == Production {
		log.Formatter = &logrus.JSONFormatter{}
	} else {
		log.Formatter = &logrus.TextFormatter{}
	}

	log.Out = os.Stdout

	log.SetLevel(logrus.InfoLevel)
	log.WithField(Env, env)

	return log
}

func (a App) createPid() error {
	pid := getPidPaths(a.config)

	err := os.MkdirAll(pid[PidPath], os.ModePerm)
	if err != nil {
		return err
	}

	f, err := os.Create(pid[PidFull])
	if err != nil {
		return err
	}
	defer f.Close()

	if _, err := os.Stat(pid[PidFull]); err == nil {
		prf(" -> [App] File '%s' created\n", pid[PidFull])
		return nil
	} else {
		prf(" -> [App] File '%s' not created\n", pid[PidFull])
		return err
	}
}

func (a App) destroyPid() error {
	pid := getPidPaths(a.config)

	err := os.Remove(pid[PidFull])
	if err != nil {
		return err
	}

	prf("\n -> [App] Removing: '%s' file\n", pid[PidFull])

	return nil
}

func getPidPaths(config ConfigReader) map[string]string {
	p := config.GetString(PidPath)
	f := config.GetString(PidFile)

	return map[string]string{
		PidPath: p,
		PidFile: f,
		PidFull: filepath.Join(p, f),
	}
}
