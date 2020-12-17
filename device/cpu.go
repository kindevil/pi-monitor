/*
@Time : 2020/12/17 10:39 上午
@Author : jia
@File : cpu.go.go
@Software : GoLand
*/

package device

import (
	"pi-monitor/helper"
	"strconv"

	"github.com/shirou/gopsutil/v3/cpu"
	"github.com/shirou/gopsutil/v3/host"
	log "github.com/sirupsen/logrus"
	"github.com/wonderivan/logger"
)

type CPU struct {
	InfoStat cpu.InfoStat
	Cores    int
	Threads  int
	Freq     *Freq
	Load     *CPULoad
	Temp     float64
}

type CPULoad struct {
	Percent float64
	Idle    float64
	User    float64
	Sys     float64
	Nice    float64
	Iowait  float64
	Irq     float64
	Softirq float64
}

type Freq struct {
	Maxfreq float64
	Minfreq float64
	Curfreq float64
}

var lastTimesStat cpu.TimesStat

func init() {
	lastTimesStat = getTimeStat()
}

func GetCpu() *CPU {
	cpu := &CPU{
		InfoStat: GetInfoStat(),
		Cores:    GetCounts(false),
		Threads:  GetCounts(false),
		Freq:     GetFreqs(),
		Load:     cpuLoad(),
		Temp:     GetTemperature(),
	}
	return cpu
}

func GetInfoStat() cpu.InfoStat {
	infoStat, err := cpu.Info()
	if err != nil {
		log.Error(err)
	}
	return infoStat[0]
}

func getTimeStat() cpu.TimesStat {
	t, err := cpu.Times(false)
	if err != nil {
		logger.Error(err)
	}
	return t[0]
}

func cpuLoad() *CPULoad {
	currentTimeStat := getTimeStat()
	last_total := lastTimesStat.User + lastTimesStat.System + lastTimesStat.Idle + lastTimesStat.Nice + lastTimesStat.Iowait + lastTimesStat.Irq + lastTimesStat.Softirq + lastTimesStat.Steal
	cur_total := currentTimeStat.User + currentTimeStat.System + currentTimeStat.Idle + currentTimeStat.Nice + currentTimeStat.Iowait + currentTimeStat.Irq + currentTimeStat.Softirq + currentTimeStat.Steal
	used_total := cur_total - last_total

	cpuload := &CPULoad{
		Percent: (1 - ((currentTimeStat.Idle - lastTimesStat.Idle) / used_total)) * 100,
		Idle:    (currentTimeStat.Idle - lastTimesStat.Idle) / used_total * 100,
		User:    (currentTimeStat.User - lastTimesStat.User) / used_total * 100,
		Sys:     (currentTimeStat.System - lastTimesStat.System) / used_total * 100,
		Nice:    (currentTimeStat.Nice - lastTimesStat.Nice) / used_total * 100,
		Iowait:  (currentTimeStat.Iowait - lastTimesStat.Iowait) / used_total * 100,
		Irq:     (currentTimeStat.Irq - lastTimesStat.Irq) / used_total * 100,
		Softirq: (currentTimeStat.Softirq - lastTimesStat.Softirq) / used_total * 100,
	}
	lastTimesStat = currentTimeStat

	return cpuload
}

func GetTemperature() float64 {
	temperatures, err := host.SensorsTemperatures()
	if err != nil {
		logger.Error(err)
	}
	return helper.ToFixed(temperatures[0].Temperature)
}

func GetCounts(logical bool) int {
	c, err := cpu.Counts(logical)
	if err != nil {
		logger.Error(err)
	}
	return c
}

func GetFreqs() *Freq {
	f := &Freq{
		Maxfreq: getFreq("cpuinfo_max_freq"),
		Minfreq: getFreq("cpuinfo_min_freq"),
		Curfreq: getFreq("cpuinfo_cur_freq"),
	}
	return f
}

func getFreq(name string) float64 {
	var lines []string
	var err error
	var freq float64
	var value float64

	lines, err = helper.ReadLines("/sys/devices/system/cpu/cpu0/cpufreq/" + name)
	if err != nil {
		logger.Error(err)
	}
	value, err = strconv.ParseFloat(lines[0], 64)
	if err != nil {
		logger.Error(err)
	}
	freq = value / 1000.0
	if freq > 9999 {
		freq = freq / 1000.0 // value in Hz
	}
	return freq
}
