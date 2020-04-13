package main

import (
	"flag"
	"os"

	trss "github.com/capric98/t-rss"
)

var (
	config   = flag.String("conf", "config.yml", "config file")
	loglevel = flag.String("log", "info", "log level: warn/info/debug/trace")
	learn    = flag.Bool("learn", false, "learn")

	userConfigDir, _ = os.UserConfigDir()
)

func init() {
	flag.Parse()
	if _, e := os.Stat(*config); os.IsNotExist(e) {
		*config = userConfigDir + "/t-rss/config.yml"
	}
}

func main() {
	trss.WithConfigFile(*config, *loglevel, *learn)
}
