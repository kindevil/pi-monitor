/*
@Time : 2020/12/17 11:11 上午
@Author : jia
@File : mem.go.go
@Software : GoLand
*/

package device

import (
	"pi-monitor/helper"

	"github.com/shirou/gopsutil/v3/mem"
	log "github.com/sirupsen/logrus"
)

type Memory struct {
	Total       uint64
	Used        uint64
	Free        uint64
	Shared      uint64
	Cached      uint64
	Available   uint64
	UsedPercent float64
	SwapTotal   uint64
	SwapFree    uint64
	SwapCached  uint64
}

func GetMemory() *Memory {
	v, err := mem.VirtualMemory()
	if err != nil {
		log.Error(err)
	}

	mem := &Memory{
		Total:       v.Total / 1024 / 1024,
		Used:        v.Used / 1024 / 1024,
		Free:        v.Free / 1024 / 1024,
		Shared:      v.Shared / 1024 / 1024,
		Cached:      v.Cached / 1024 / 1024,
		Available:   v.Available / 1024 / 1024,
		UsedPercent: helper.ToFixed(v.UsedPercent),
		SwapTotal:   v.SwapTotal / 1024 / 1024,
		SwapFree:    v.SwapFree / 1024 / 1024,
		SwapCached:  v.SwapCached / 1024 / 1024,
	}

	return mem
}
