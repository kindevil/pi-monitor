/*
 * @Author: jia
 * @LastEditTime: 2020-12-08 13:55:14
 * @FilePath: /pi-monitor/service/mem.go
 * @Date: 2020-12-08 13:39:02
 * @Software: VS Code
 */
package service

import (
	"github.com/shirou/gopsutil/v3/mem"
	"github.com/wonderivan/logger"
)

type Memory struct {
	Total       uint64
	Used        uint64
	Free        uint64
	Shared      uint64
	Buffer      uint64
	Available   uint64
	UsedPercent float64
	SwapTotal   uint64
	SwapFree    uint64
	SwapCached  uint64
}

func GetMem() *Memory {
	v, err := mem.VirtualMemory()
	if err != nil {
		logger.Error(err)
	}

	mem := &Memory{
		Total:       v.Total,
		Used:        v.Used,
		Free:        v.Free,
		Shared:      v.Shared,
		Buffer:      v.Buffers,
		Available:   v.Available,
		UsedPercent: v.UsedPercent,
		SwapTotal:   v.SwapTotal,
		SwapFree:    v.SwapFree,
		SwapCached:  v.SwapCached,
	}

	return mem
}
