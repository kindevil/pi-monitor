/*
@Time : 2020/12/17 5:47 下午
@Author : jia
@File : load.go
@Software : GoLand
*/

package device

import (
	"pi-monitor/helper"

	"github.com/shirou/gopsutil/v3/load"
	log "github.com/sirupsen/logrus"
)

func GetLoad() *load.AvgStat {
	avgStat, err := load.Avg()
	if err != nil {
		log.Error(err)
	}

	avgStat.Load1 = helper.ToFixed(avgStat.Load1)
	avgStat.Load5 = helper.ToFixed(avgStat.Load5)
	avgStat.Load15 = helper.ToFixed(avgStat.Load15)

	return avgStat
}
