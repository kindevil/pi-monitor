package service

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"os/exec"
	"strings"
	"time"

	"github.com/shirou/gopsutil/v3/host"
	"github.com/wonderivan/logger"
)

type Host struct {
	Hostname     string    // 主机名
	OS           string    // 系统版本
	Vendor       string    // 厂家
	Model        string    // 硬件版本
	Serial       string    // 序列号
	BootTime     string    // 启动时间
	Kernal       string    // 内核信息
	InterfaceNum int       // 网卡数
	Disks        []*device // 存储设备
}

type UpTime struct {
	Minute uint64
	Hour   uint64
	Day    uint64
	Year   uint64
}

type netInter struct {
	Count int
	Names []string
}

var hostInfo *Host

//hostname 获取主机名
func hostname() {
	var name string
	name, err := os.Hostname()
	if err != nil {
		logger.Error(err)
	}
	hostInfo.Hostname = name
}

//osVersion 获取操作系统版本
func osVersion() {
	hostInfo.OS = readOSRelease("PRETTY_NAME")
}

//getHostInfo 获取主机版本与厂家信息
func hostVendor() {
	model := readLine("/proc/device-tree/model")
	switch {
	case strings.Contains(model, "Raspberry"):
		hostInfo.Model = model
		hostInfo.Vendor = "Raspberry Pi"
	case strings.Contains(model, "Radxa"):
		hostInfo.Model = model
		hostInfo.Vendor = "Rock Pi"
	case strings.Contains(model, "FriendlyARM"):
		hostInfo.Model = model
		hostInfo.Vendor = "Nano Pi"
	default:
		hostInfo.Model = "unknown"
		hostInfo.Vendor = "unknown"
	}
}

//getSerial 获取序列号
func serial() {
	if PathExists("/proc/device-tree/serial-number") {
		hostInfo.Serial = readLine("/proc/device-tree/serial-number")
	} else if PathExists("/proc/cpuinfo") {
		hostInfo.Serial = scanCpuInfo("Serial")
	} else {
		hostInfo.Serial = "unknown"
	}
}

//bootTime 启动时间
func bootTime(t uint64) {
	hostInfo.BootTime = time.Unix(int64(t), 0).Format("2006-01-02 15:04:05")
}

//kernel 读取内核信息
func kernel() {
	cmd := exec.Command("uname", "-a")
	stdout, err := cmd.Output()
	if err != nil {
		logger.Info(err)
	}
	hostInfo.Kernal = strings.Trim(string(stdout), "\n")
}

func netInterface() {
	n := GetNet().Interface
	nt := &netInter{}
	nt.Count = len(n)
	for _, val := range n {
		nt.Names = append(nt.Names, val.Name)
	}
}

func diskDevice() {
	hostInfo.Disks = GetDisk()
}

//readLine 读取文件第一行
func readLine(path string) string {
	file, err := os.Open(path)
	if err != nil {
		fmt.Println(err)
		return ""
	}
	defer file.Close()

	var lineText string
	scanner := bufio.NewScanner(file)
	scanner.Scan()
	lineText = scanner.Text()
	return lineText[:len(lineText)-1]
}

//readOSRelease 读取/etc/os-release
func readOSRelease(keyward string) string {
	file, err := os.Open("/etc/os-release")
	if err != nil {
		fmt.Println(err)
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
	file, err := os.Open("/proc/cpuinfo")
	if err != nil {
		log.Fatal(err)
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

func getInfo(path string) string {

	file, err := os.Open(path)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()
	var lineText string
	scanner := bufio.NewScanner(file)
	scanner.Scan()
	lineText = scanner.Text()

	return lineText[:len(lineText)-1]
}

func readInfo() map[string]string {
	var info = make(map[string]string)
	info["hardware"] = ""
	info["serial"] = getInfo("/proc/device-tree/serial-number")
	info["model"] = getInfo("/proc/device-tree/model")

	// f, err := os.Open("/proc/cpuinfo")
	// if err != nil {
	// 	logger.Info(err)
	// }
	// defer f.Close()

	// scanner := bufio.NewScanner(f)
	// for scanner.Scan() {
	// 	line := scanner.Text()
	// 	if strings.Contains(line, "Hardware") {
	// 		info["hardware"] = strings.Trim(strings.Split(line, ":")[1], " ")
	// 	}

	// 	if strings.Contains(line, "Serial") {
	// 		info["serial"] = strings.Trim(strings.Split(line, ":")[1], " ")
	// 	}

	// 	if strings.Contains(line, "Model") {
	// 		info["model"] = strings.Trim(strings.Split(line, ":")[1], " ")
	// 	}
	// }

	// if scanner.Err() != nil {
	// 	logger.Error(scanner.Err())
	// }

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

func getHost() {
	info, err := host.Info()
	if err != nil {
		logger.Error(err)
	}

	hostInfo = &Host{}

	osVersion()
	hostname()
	hostVendor()
	serial()
	bootTime(info.BootTime)
	kernel()
	hostInfo.InterfaceNum = len(GetNet().Interface)
	diskDevice()
}

func GetHost() *Host {
	return hostInfo
}
