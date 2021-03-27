package main

import (
	"flag"
	"fmt"
	"os"
	"time"

	"github.com/axetroy/watchdog"
	"github.com/axetroy/watchdog/internal/scheduling"
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
  --config      Specify config file. defaults to 'watchdog.config.json'.
  --port        Specify the port for HTTP listening. defaults to '9999'.
SOURCE CODE:
  https://github.com/axetroy/watchdog`)
}

func main() {
	var (
		showHelp    bool
		showVersion bool
		configPath  string
		port        string
		noColor     bool
	)

	flag.StringVar(&configPath, "config", "watchdog.config.json", "The config file path")
	flag.StringVar(&port, "port", "9999", "Specify the port for HTTP listening")
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

	c, err := watchdog.NewConfigFromFile(configPath)

	if err != nil {
		fmt.Printf("%+v\n", err)
		os.Exit(1)
	}

	for _, s := range c.Service {
		interval := s.Interval

		// 如果服务没有单独设置间隔时间，则获取全局配置
		if interval == 0 {
			interval = c.Interval
		}

		// 默认 30 s 间隔
		if interval == 0 {
			interval = 5
		}

		scheduler := scheduling.NewScheduling(time.Second*time.Duration(interval), watchdog.NewRunnerJob(s))
		go scheduler.Start()
	}

	watchdog.Serve(port, c)
}
