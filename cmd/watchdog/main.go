package main

import (
	"errors"
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
  --config      Specify config file. defaults to 'watchdog.config.json'. allow json/yml file.
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

	defaultPort := os.Getenv("PORT")

	if defaultPort == "" {
		defaultPort = "9999"
	}

	flag.StringVar(&configPath, "config", "watchdog.config.json", "The config file path")
	flag.StringVar(&port, "port", defaultPort, "Specify the port for HTTP listening")
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

	if configPath == "" {
		fmt.Println("--config=<path> is required.")
		os.Exit(1)
	}

	if _, err := os.Stat(configPath); errors.Is(err, os.ErrNotExist) {
		fmt.Printf("config file %s does not exists.\n", configPath)
		os.Exit(1)
	}

	c, err := watchdog.NewConfigFromFile(configPath)

	if err != nil {
		fmt.Printf("%+v\n", err)
		os.Exit(1)
	}

	// 默认 30 s 间隔
	if c.Interval == 0 {
		c.Interval = 30
	}

	// 默认 发送 100 次提醒每天
	if c.MaxNotifyTimesForDay == 0 {
		c.MaxNotifyTimesForDay = 100
	}

	// 默认 发送 100 次提醒每时
	if c.MaxNotifyTimesForHour == 0 {
		c.MaxNotifyTimesForHour = 5
	}

	for _, s := range c.Service {
		interval := getServiceFieldWithDefault(s, func(s watchdog.Service) uint {
			return s.Interval
		}, c.Interval)

		maxNotifyTimesForDay := getServiceFieldWithDefault(s, func(s watchdog.Service) uint {
			return s.MaxNotifyTimesForDay
		}, c.MaxNotifyTimesForDay)

		maxNotifyTimesForHour := getServiceFieldWithDefault(s, func(s watchdog.Service) uint {
			return s.MaxNotifyTimesForHour
		}, c.MaxNotifyTimesForHour)

		scheduler := scheduling.NewScheduling(scheduling.Options{
			Interval:              time.Second * time.Duration(interval),
			Job:                   watchdog.NewRunnerJob(s),
			MaxNotifyTimesForDay:  maxNotifyTimesForDay,
			MaxNotifyTimesForHour: maxNotifyTimesForHour,
		})
		go scheduler.Start()
	}

	watchdog.Serve(port, c)
}

// 获取这个服务器的默认配置字段，如果没有，则使用全局配置
func getServiceFieldWithDefault(s watchdog.Service, fn func(s watchdog.Service) uint, defaultValue uint) uint {
	value := fn(s)

	if value == 0 {
		return defaultValue
	}

	return value
}
