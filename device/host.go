/*
 * @Author: jia
 * @LastEditTime: 2020-12-17 09:21:59
 * @FilePath: /pi-monitor/device/host.go
 * @Date: 2020-12-17 09:17:06
 * @Software: VS Code
 */
package device

import (
	"bufio"
	"os"
	"os/exec"
	"pi-monitor/helper"
	"strings"
	"time"

	"github.com/shirou/gopsutil/v3/cpu"
	"github.com/shirou/gopsutil/v3/host"
	"github.com/shirou/gopsutil/v3/load"
	log "github.com/sirupsen/logrus"
	"github.com/wonderivan/logger"
)

type Host struct {
	Hostname string // 主机名
	OS       string // 系统版本
	Vendor   string // 厂家
	Model    string // 硬件版本
	Serial   string // 序列号
	BootTime string // 启动时间
	Kernal   string // 内核信息
	CPU      struct {
		InfoStat    cpu.InfoStat
		Core        int     // CPU核心数
		Freq        *Freq   // CPU频率
		Temperature float64 // CPU温度
	}
	Mem       *Memory
	Disks     []*Disk
	Interface struct {
		Interfaces map[string]*Interface // 网卡名称
		Count      int                   // 网卡数
	}
	Load     *load.AvgStat
	infostat *host.InfoStat
}

type UPtime struct {
	Minute uint64
	Hour   uint64
	Day    uint64
	Year   uint64
}

var (
	hostinfo *Host
)

func init() {
	netInit()

	hostinfo = new(Host)
	hostinfo.infostat = getInfostat()
	hostinfo.hostname()
	hostinfo.osVersion()
	hostinfo.hostVendor()
	hostinfo.serial()
	hostinfo.bootTime()
	hostinfo.kernel()
	hostinfo.CPU.InfoStat = GetInfoStat()
	hostinfo.CPU.Core = GetCounts(false)
	hostinfo.CPU.Freq = GetFreqs()
	hostinfo.CPU.Temperature = GetTemperature()
	hostinfo.Mem = GetMemory()
	hostinfo.Disks = GetDisk()
	hostinfo.Interface.Count = GetNetCount()
	hostinfo.Interface.Interfaces = GetInterfaceStat()
	hostinfo.Load = GetLoad()
}

func GetHost() *Host {
	return hostinfo
}

func GetUptime() *UPtime {
	t := getInfostat().Uptime

	var mins = (t - (t % 60)) / 60
	var min = mins % 60
	var hours = (mins - min) / 60
	var hour = hours % 24
	var day = hours / 24
	var year = hours / 24 / 365

	upTime := &UPtime{
		Minute: min,
		Hour:   hour,
		Day:    day,
		Year:   year,
	}

	return upTime
}

func getInfostat() *host.InfoStat {
	stat, err := host.Info()
	if err != nil {
		log.Error(err)
	}
	return stat
}

func (this *Host) hostname() {
	this.Hostname = this.infostat.Hostname
}

func (this *Host) osVersion() {
	this.OS = this.readOSRelease("PRETTY_NAME")
	if this.OS == "" {
		this.OS = this.infostat.OS
	}
}

//getHostInfo 获取主机版本与厂家信息
func (this *Host) hostVendor() {
	model := helper.ReadLine("/proc/device-tree/model")
	switch {
	case strings.Contains(model, "Raspberry"):
		this.Model = model
		this.Vendor = "Raspberry Pi"
	case strings.Contains(model, "Radxa"):
		this.Model = model
		this.Vendor = "Rock Pi"
	case strings.Contains(model, "FriendlyARM"):
		this.Model = model
		this.Vendor = "Nano Pi"
	default:
		this.Model = "unknown"
		this.Vendor = "unknown"
	}
}

//getSerial 获取序列号
func (this *Host) serial() {
	if helper.PathExists("/proc/device-tree/serial-number") {
		this.Serial = helper.ReadLine("/proc/device-tree/serial-number")
	} else if helper.PathExists("/proc/cpuinfo") {
		this.Serial = scanCpuInfo("Serial")
	} else {
		this.Serial = "unknown"
	}
}

//bootTime 启动时间
func (this *Host) bootTime() {
	this.BootTime = time.Unix(int64(this.infostat.BootTime), 0).Format("2006-01-02 15:04:05")
}

//kernel 读取内核信息
func (this *Host) kernel() {
	cmd := exec.Command("uname", "-a")
	stdout, err := cmd.Output()
	if err != nil {
		logger.Info(err)
	}
	this.Kernal = strings.Trim(string(stdout), "\n")
}

//readOSRelease 读取/etc/os-release
func (this *Host) readOSRelease(keyward string) string {
	if !helper.PathExists("/etc/os-release") {
		return ""
	}

	file, err := os.Open("/etc/os-release")
	if err != nil {
		log.Error(err)
		return ""
	}
	defer file.Close()

	var lineText string
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		lineText = scanner.Text()
		if strings.Contains(lineText, keyward) {
			lineText = strings.Trim(strings.Split(lineText, "=")[1], "\"")
			break
		}
	}

	return lineText
}

func scanCpuInfo(keyward string) string {
	if !helper.PathExists("/proc/cpuinfo") {
		return ""
	}

	file, err := os.Open("/proc/cpuinfo")
	if err != nil {
		log.Error(err)
		return ""
	}
	defer file.Close()

	var lineText string
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		lineText = scanner.Text()
		if strings.Contains(lineText, keyward) {
			lineText = strings.Trim(strings.Split(lineText, ":")[1], " ")
			break
		}
	}

	return lineText
}
