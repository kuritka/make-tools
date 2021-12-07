package main

import (
	"github.com/kuritka/make-tools/internal/tools"
	"github.com/kuritka/make-tools/internal/utils/zerolog"
	z "github.com/rs/zerolog"
	"os"
	"path/filepath"

	"gopkg.in/alecthomas/kingpin.v2"
)

var (
	app            = kingpin.New(filepath.Base(os.Args[0]), "MakeFile Tools support").DefaultEnvars()
	debug          = app.Flag("debug", "Run with debug logging.").Short('d').Bool()
	message		   = app.Flag("message", "Prints message.").Short('m').String()
	envExists = app.Flag("env-exists",
		"Checks if the environment variable exists and if not terminates the makefile. " +
		"If the 'message' argument is filled in, prints message.").Short('e').String()
	save = app.Flag("save","").Short('s').String()
	load = app.Flag("load","").Short('l').String()
	initialize = app.Flag("init","").Short('i').Bool()
	helmver = app.Flag("get-helm-version","Prints helm version from Chart.yaml").String()
	appver = app.Flag("get-helm-appversion","Prints helm app version from Chart.yaml").String()
	failif = app.Flag("fail-if","Fails if expression is true").Short('f').String()
)

var log = zerolog.Logger()
func main(){
	kingpin.MustParse(app.Parse(os.Args[1:]))
	if *debug {
		z.SetGlobalLevel(z.DebugLevel)
		log.Debug().Msg("Debug mode enabled")
	}
	if tools.EnvVarExists(*envExists, *message) {
		return
	}
	if tools.Save(*save, *message) {
		return
	}
	if tools.Load(*load, *message) {
		return
	}
	if tools.Initialize(*initialize, *message) {
		return
	}
	if tools.HelmVer(*helmver) {
		return
	}
	if tools.HelmAppVer(*appver) {
		return
	}
	if tools.FailIf(*failif, *message) {
		os.Exit(1)
	}
}
