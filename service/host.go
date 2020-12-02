package service

import (
	"bufio"
	"os"
	"os/exec"
	"strings"
	"time"

	"github.com/shirou/gopsutil/v3/host"
	"github.com/wonderivan/logger"
)

type Host struct {
	Hostname string
	Uptime   *UpTime
	BootTime string
	OS       string
	Platform string
	Kernal   string
	Version  string
	Hardware string
	Serial   string
	Model    string
}

type UpTime struct {
	Minute uint64
	Hour   uint64
	Day    uint64
	Year   uint64
}

func readKernal() string {
	cmd := exec.Command("uname", "-a")
	stdout, err := cmd.Output()
	if err != nil {
		logger.Info(err)
		return ""
	}
	return strings.Trim(string(stdout), "\n")
}

func readInfo() map[string]string {
	var info = make(map[string]string)
	info["hardware"] = ""
	info["serial"] = ""
	info["model"] = ""

	f, err := os.Open("/proc/cpuinfo")
	if err != nil {
		logger.Info(err)
	}
	defer f.Close()

	scanner := bufio.NewScanner(f)
	for scanner.Scan() {
		line := scanner.Text()
		if strings.Contains(line, "Hardware") {
			info["hardware"] = strings.Trim(strings.Split(line, ":")[1], " ")
		}

		if strings.Contains(line, "Serial") {
			info["serial"] = strings.Trim(strings.Split(line, ":")[1], " ")
		}

		if strings.Contains(line, "Model") {
			info["model"] = strings.Trim(strings.Split(line, ":")[1], " ")
		}
	}

	if scanner.Err() != nil {
		logger.Error(scanner.Err())
	}

	return info
}

func runningTime(t uint64) *UpTime {
	var mins = (t - (t % 60)) / 60
	var min = mins % 60
	var hours = (mins - min) / 60
	var hour = hours % 24
	var day = hours / 24
	var year = hours / 24 / 365

	upTime := &UpTime{
		Minute: min,
		Hour:   hour,
		Day:    day,
		Year:   year,
	}

	return upTime
}

func GetHost() *Host {
	info, err := host.Info()
	if err != nil {
		logger.Error(err)
	}

	boardInfo := readInfo()

	host := &Host{
		Hostname: info.Hostname,
		OS:       info.OS,
		Platform: info.Platform,
		Hardware: boardInfo["hardware"],
		Serial:   boardInfo["serial"],
		Model:    boardInfo["model"],
		Uptime:   runningTime(info.Uptime),
		BootTime: time.Unix(int64(info.BootTime), 0).Format("2006-01-02 15:04:05"),
		Kernal:   readKernal(),
	}
	return host
}
