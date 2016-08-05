package main

import (
	"ginserver/env"
	"ginserver/routers"

	"github.com/cihub/seelog"
)

func main() {
	defer seelog.Flush()
	logConfig := env.Get("LOG_CONFIG")
	logger, err := seelog.LoggerFromConfigAsFile(logConfig)
	if err != nil {
		panic(err)
	}
	seelog.ReplaceLogger(logger)

	host := env.Get("APP_HOST")
	if host == "" {
		panic("APP_HOST is empty")
	}

	routers.Handle(host)
}
