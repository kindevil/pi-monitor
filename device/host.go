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
	"github.com/shirou/gopsutil/v3/host"
	log "github.com/sirupsen/logrus"
	"github.com/wonderivan/logger"
	"os"
	"os/exec"
	"pi-monitor/helper"
	"strings"
	"time"
)

type Host struct {
	Hostname       string   // 主机名
	OS             string   // 系统版本
	Vendor         string   // 厂家
	Model          string   // 硬件版本
	Serial         string   // 序列号
	BootTime       string   // 启动时间
	Kernal         string   // 内核信息
	InterfaceNum   int      // 网卡数
	InterfaceNames []string // 网卡名称
	infostat       *host.InfoStat
}

var hostinfo *Host

func init() {
	hostinfo = new(Host)
	hostinfo.getInfostat()
	hostinfo.hostname()
	hostinfo.osVersion()
	hostinfo.hostVendor()
	hostinfo.serial()
	hostinfo.bootTime()
	hostinfo.kernel()
	hostinfo.InterfaceNum = GetNetCount()
	hostinfo.InterfaceNames = GetNetNames()
}

func GetHost() *Host {
	return hostinfo
}

func (this *Host) getInfostat() {
	stat, err := host.Info()
	if err != nil {
		log.Error(err)
	}
	this.infostat = stat
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
		log.Fatal(err)
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
		log.Fatal(err)
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
