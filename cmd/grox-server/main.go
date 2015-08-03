package main

import (
	"flag"

	"github.com/golang-vietnam/grox/server"
	"github.com/golang-vietnam/grox/utils/load-config"
	"github.com/golang-vietnam/grox/utils/logs"
)

var (
	flConfigFile = flag.String("config-file", "config-default.json", "Load config from file")

	l = logs.New("grox-server")
)

func main() {
	flag.Parse()

	var cfg server.Config
	err := loadConfig.FromFileAndEnv(&cfg, *flConfigFile)
	if err != nil {
		l.Fatalln("Error loading config:", err)
	}

	server.Start(cfg)
}
