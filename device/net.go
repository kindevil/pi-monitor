/*
@Time : 2020/12/17 9:41 上午
@Author : jia
@File : net.go.go
@Software : GoLand
*/

package device

import (
	"pi-monitor/helper"
	"time"

	"github.com/shirou/gopsutil/v3/net"
	log "github.com/sirupsen/logrus"
)

type Net struct {
	Interface []*Interface
	Names     []string
	Count     int
}

type Interface struct {
	Name         string
	HardwareAddr string
	Addrs        net.InterfaceAddrList
	BytesSent    uint64
	BytesRecv    uint64
	PacketsSent  uint64
	PacketsRecv  uint64
	Send         float64
	Recv         float64
}

var (
	lastInterface map[string]*Interface
	lastTime      time.Time
)

func netInit() {
	lastInterface = interfaceStat()
	lastTime = time.Now()
}

func GetInterfaceStat() map[string]*Interface {
	currentInterface := interfaceStat()
	timeNow := time.Now()
	for _, interfaceStat := range currentInterface {
		last := lastInterface[interfaceStat.Name]
		var recv, send, diff float64 = 0, 0, 0

		diff = float64(timeNow.UnixNano()/1e6-lastTime.UnixNano()/1e6) / 1000
		recv = float64(interfaceStat.BytesRecv-last.BytesRecv) / diff / 1024
		send = float64(interfaceStat.BytesSent-last.BytesSent) / diff / 1024
		currentInterface[interfaceStat.Name].Recv = helper.ToFixed(recv)
		currentInterface[interfaceStat.Name].Send = helper.ToFixed(send)
	}

	lastInterface = currentInterface
	lastTime = timeNow

	return currentInterface
}

func GetNetCount() int {
	return len(lastInterface)
}

func GetNetNames() []string {
	var names []string
	for _, v := range lastInterface {
		names = append(names, v.Name)
	}
	return names
}

func interfaceStat() map[string]*Interface {
	var interfaceStatList = make(map[string]*Interface)
	InterfaceStatList := loadInterfaceStatList()
	ioCounters := loadCounters()
	for _, interfaceStat := range InterfaceStatList {
		if interfaceStat.Name == "lo" {
			continue
		}

		ioCounter := getCounter(interfaceStat.Name, ioCounters)

		i := &Interface{
			Name:         interfaceStat.Name,
			HardwareAddr: interfaceStat.HardwareAddr,
			Addrs:        interfaceStat.Addrs,
			BytesRecv:    ioCounter.BytesRecv,
			BytesSent:    ioCounter.BytesSent,
			PacketsRecv:  ioCounter.PacketsRecv,
			PacketsSent:  ioCounter.PacketsSent,
			Recv:         0,
			Send:         0,
		}

		interfaceStatList[interfaceStat.Name] = i
	}
	return interfaceStatList
}

func loadInterfaceStatList() net.InterfaceStatList {
	interfaceList, err := net.Interfaces()
	if err != nil {
		log.Error(err)
	}
	return interfaceList
}

func loadCounters() []net.IOCountersStat {
	ioCountersStat, err := net.IOCounters(false)
	if err != nil {
		log.Error(err)
	}
	return ioCountersStat
}

func getCounter(name string, ioCounters []net.IOCountersStat) net.IOCountersStat {
	for _, value := range ioCounters {
		if value.Name == name {
			return value
		}
	}

	return net.IOCountersStat{}
}
