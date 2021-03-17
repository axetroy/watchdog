package main

import (
	"flag"
	"fmt"
	"os"

	"github.com/gookit/color"
)

var (
	version = "dev"
	commit  = "none"
	date    = "unknown"
	// builtBy = "unknown"
)

func printHelp() {
	println(`watchdog - a cli tool for watch service running status
USAGE:
  watchdog [OPTIONS]
OPTIONS:
  --help        Print help information.
  --version     Print version information.
  --config      Specify config file
SOURCE CODE:
  https://github.com/axetroy/watchdog`)
}

func main() {
	var (
		showHelp    bool
		showVersion bool
		configPath  string
		noColor     bool
	)

	flag.StringVar(&configPath, "config", "", "The config file path")
	flag.BoolVar(&showHelp, "help", false, "Print help information")
	flag.BoolVar(&showVersion, "version", false, "Print version information")

	flag.Usage = printHelp

	flag.Parse()

	if showHelp {
		printHelp()
		os.Exit(0)
	}

	if showVersion {
		println(fmt.Sprintf("%s %s %s", version, commit, date))
		os.Exit(0)
	}

	if color.SupportColor() {
		color.Enable = !noColor
	} else {
		color.Enable = false
	}
}
