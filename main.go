package main

import (
	"pi-monitor/service"

	"github.com/wonderivan/logger"
)

func main() {
	logger.Info(service.GetHost())
	//route.Run()
}
