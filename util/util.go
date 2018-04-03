package util

import (
	"io/ioutil"
	"os"
	"path/filepath"
	"runtime"

	"github.com/inconshreveable/log15"
	"github.com/k0kubun/pp"
)

// GenWorkers generate workders
func GenWorkers(num int) chan<- func() {
	tasks := make(chan func())
	for i := 0; i < num; i++ {
		go func() {
			for f := range tasks {
				f()
			}
		}()
	}
	return tasks
}

// GetDefaultLogDir returns default log directory
func GetDefaultLogDir() string {
	defaultLogDir := "/var/log/vuls"
	if runtime.GOOS == "windows" {
		defaultLogDir = filepath.Join(os.Getenv("APPDATA"), "vuls")
	}
	return defaultLogDir
}

// SetLogger set logger
func SetLogger(logDir string, quiet, debug bool) {
	lvlHundler := log15.LvlFilterHandler(log15.LvlInfo, log15.StderrHandler)
	if debug {
		lvlHundler = log15.LvlFilterHandler(log15.LvlDebug, log15.StdoutHandler)
	}
	if quiet {
		lvlHundler = log15.LvlFilterHandler(log15.LvlDebug, log15.DiscardHandler())
		pp.SetDefaultOutput(ioutil.Discard)
	}

	if _, err := os.Stat(logDir); os.IsNotExist(err) {
		if err := os.Mkdir(logDir, 0700); err != nil {
			log15.Error("Failed to create log directory", "err", err)
		}
	}
	var hundler log15.Handler
	if _, err := os.Stat(logDir); err == nil {
		logPath := filepath.Join(logDir, "goval-dictionary.log")
		hundler = log15.MultiHandler(
			log15.Must.FileHandler(logPath, log15.LogfmtFormat()),
			lvlHundler,
		)
	} else {
		hundler = lvlHundler
	}
	log15.Root().SetHandler(hundler)
}
