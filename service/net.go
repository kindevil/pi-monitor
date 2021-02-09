package service

import (
	"time"

	"github.com/shirou/gopsutil/v3/net"
	log "github.com/sirupsen/logrus"
)

type Net struct {
	Interface []*Inter
}

type Inter struct {
	Name         string
	HardwareAddr string
	Addrs        net.InterfaceAddrList
	BytesSent    uint64
	BytesRecv    uint64
	PacketsSent  uint64
	PacketsRecv  uint64
	Send         string
	Recv         string
}

var (
	netLastStat []net.IOCountersStat
	netLastTime time.Time
)

func loadNetStat() []net.IOCountersStat {
	stat, err := net.IOCounters(true)
	if err != nil {
		log.Error(err)
		return nil
	}
	return stat
}

//initLastTimes 初始化运行时间，以便第一次做比较时使用
// func initNetStat() {
// 	netLastStat = nil
// 	netLastTime = time.Now()
// 	logger.Debug(netLastTime)
// }

func GetNet() []*Inter {

	list, err := net.Interfaces()
	if err != nil {
		log.Error(err)
	}

	stat, err := net.IOCounters(true)
	if err != nil {
		log.Error(err)
	}

	//n := &Net{}
	var Interface []*Inter

	time.Sleep(time.Second)

	timeNow := time.Now()

	for _, val := range list {
		if val.Name == "lo" {
			continue
		}

		s := search(val.Name, stat)
		var recv, send, diff float64 = 0, 0, 0
		if len(netLastStat) != 0 {
			old := search(val.Name, netLastStat)
			diff = float64(timeNow.UnixNano()/1e6-netLastTime.UnixNano()/1e6) / 1000
			recv = float64(s.BytesRecv-old.BytesRecv) / diff
			send = float64(s.BytesSent-old.BytesSent) / diff
		}

		i := &Inter{
			Name:         val.Name,
			HardwareAddr: val.HardwareAddr,
			Addrs:        val.Addrs,
			BytesRecv:    s.BytesRecv,
			BytesSent:    s.BytesSent,
			PacketsRecv:  s.PacketsRecv,
			PacketsSent:  s.PacketsSent,
			Recv:         floatToString(recv),
			Send:         floatToString(send),
		}
		// if key == 1 {
		// 	logger.Info(uint64(timeNow.Unix() - netLastTime.Unix()))
		// 	logger.Info(uint64(timeNow.Unix() - netLastTime.Unix()))
		// 	logger.Info("BytesRecv:", s.BytesRecv)
		// 	logger.Info("BytesRecv:", old.BytesRecv)
		// 	logger.Info("BytesSent:", s.BytesSent)
		// 	logger.Info("BytesSent:", old.BytesSent)
		// }
		Interface = append(Interface, i)
	}

	netLastStat = stat[:]
	netLastTime = timeNow

	return Interface
}

func search(name string, stat []net.IOCountersStat) net.IOCountersStat {
	for _, val := range stat {
		if val.Name == name {
			return val
		}
	}

	return net.IOCountersStat{}
}
