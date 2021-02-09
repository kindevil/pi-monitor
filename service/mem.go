/*
 * @Author: jia
 * @LastEditTime: 2020-12-08 14:12:14
 * @FilePath: /pi-monitor/service/mem.go
 * @Date: 2020-12-08 13:39:02
 * @Software: VS Code
 */
package service

import (
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

func GetMem() *Memory {
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
		UsedPercent: v.UsedPercent,
		SwapTotal:   v.SwapTotal / 1024 / 1024,
		SwapFree:    v.SwapFree / 1024 / 1024,
		SwapCached:  v.SwapCached / 1024 / 1024,
	}

	return mem
}
