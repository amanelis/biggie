package main

import (
	"os"
	"path/filepath"

	exchange "github.com/preichenberger/go-coinbase-exchange"
)

// App - hold all app configurations here
type App struct {
	client *exchange.Client
	config ConfigReader
	helper Helper
}

func (a App) createPid() error {
	pid := getPidPaths(a.config)

	err := os.MkdirAll(pid["pidpath"], os.ModePerm)
	if err != nil {
		return err
	}

	f, err := os.Create(pid["pidfull"])
	if err != nil {
		return err
	}
	defer f.Close()

	if _, err := os.Stat(pid["pidfull"]); err == nil {
		prf(" -> [App] File '%s' created\n", pid["pidfull"])
		return nil
	} else {
		prf(" -> [App] File '%s' not created\n", pid["pidfull"])
		return err
	}
}

func (a App) destroyPid() error {
	pid := getPidPaths(a.config)

	prf("\n -> [App] Removing: '%s' file\n", pid["pidfull"])
	err := os.Remove(pid["pidfull"])
	if err != nil {
		return err
	}

	return nil
}

func getPidPaths(config ConfigReader) map[string]string {
	p := config.GetString("pidpath")
	f := config.GetString("pidfile")

	return map[string]string{
		"pidpath": p,
		"pidfile": f,
		"pidfull": filepath.Join(p, f),
	}
}
