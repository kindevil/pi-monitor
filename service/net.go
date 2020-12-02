package service

import (
	"github.com/shirou/gopsutil/v3/net"
	"github.com/wonderivan/logger"
)

func GetNet() {

	list, _ := net.Interfaces()
	logger.Info(list)

	stat, _ := net.IOCounters(true)
	logger.Info(stat)
}
