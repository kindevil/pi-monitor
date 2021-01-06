package service

import (
	"strconv"

	"github.com/shirou/gopsutil/v3/cpu"
	"github.com/shirou/gopsutil/v3/host"
	"github.com/wonderivan/logger"
)

type CPU struct {
	Cores   int
	Threads int
	Freq    *Freq
	Load    *CPULoad
	Temp    string
}

type Freq struct {
	Maxfreq float64
	Minfreq float64
	Curfreq float64
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

var (
	lastTimesStat cpu.TimesStat
)

//loadTimes 获取运行时间
func loadTimes() cpu.TimesStat {
	t, err := cpu.Times(false)
	if err != nil {
		logger.Error(err)
	}
	return t[0]
}

//initLastTimes 初始化运行时间，以便第一次做比较时使用
func initLastTimes() {
	lastTimesStat = loadTimes()
}

//cpuload 获取占用率
func cpuLoad() *CPULoad {
	t := loadTimes()
	last_total := lastTimesStat.User + lastTimesStat.System + lastTimesStat.Idle + lastTimesStat.Nice + lastTimesStat.Iowait + lastTimesStat.Irq + lastTimesStat.Softirq + lastTimesStat.Steal
	cur_total := t.User + t.System + t.Idle + t.Nice + t.Iowait + t.Irq + t.Softirq + t.Steal
	used_total := cur_total - last_total

	cpuload := &CPULoad{
		Percent: (1 - ((t.Idle - lastTimesStat.Idle) / used_total)) * 100,
		Idle:    (t.Idle - lastTimesStat.Idle) / used_total * 100,
		User:    (t.User - lastTimesStat.User) / used_total * 100,
		Sys:     (t.System - lastTimesStat.System) / used_total * 100,
		Nice:    (t.Nice - lastTimesStat.Nice) / used_total * 100,
		Iowait:  (t.Iowait - lastTimesStat.Iowait) / used_total * 100,
		Irq:     (t.Irq - lastTimesStat.Irq) / used_total * 100,
		Softirq: (t.Softirq - lastTimesStat.Softirq) / used_total * 100,
	}
	lastTimesStat = t
	return cpuload
}

func cpuTemperature() string {
	temperatures, err := host.SensorsTemperatures()
	if err != nil {
		logger.Error(err)
	}
	return floatToString(temperatures[0].Temperature)
}

func cpuCount(logical bool) int {
	c, err := cpu.Counts(logical)
	if err != nil {
		logger.Error(err)
	}
	return c
}

func getFreq(name string) float64 {
	var lines []string
	var err error
	var freq float64
	var value float64

	lines, err = ReadLines("/sys/devices/system/cpu/cpu0/cpufreq/" + name)
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

func freq() *Freq {
	f := &Freq{
		Maxfreq: getFreq("cpuinfo_max_freq"),
		Minfreq: getFreq("cpuinfo_min_freq"),
		Curfreq: getFreq("cpuinfo_cur_freq"),
	}
	return f
}

func GetCPU() *CPU {
	cpu := &CPU{
		Cores:   cpuCount(false),
		Threads: cpuCount(false),
		Freq:    freq(),
		Load:    cpuLoad(),
		Temp:    cpuTemperature(),
	}
	return cpu
}
