package device

import (
	"github.com/shirou/gopsutil/v3/disk"
	"github.com/wonderivan/logger"
)

type Disk struct {
	Name        string
	Mountpoint  string
	Fstype      string
	Total       uint64
	Free        uint64
	Used        uint64
	UsedPercent float64
}

var Disks []*Disk

func GetDisk() []*Disk {
	partitions, err := disk.Partitions(false)
	if err != nil {
		logger.Error(err)
	}

	for _, dev := range partitions {

		d := &Disk{}
		d.Name = dev.Device
		d.Mountpoint = dev.Mountpoint
		d.Fstype = dev.Fstype

		usage, err := disk.Usage(dev.Mountpoint)
		if err != nil {
			logger.Error(err)
			continue
		}

		d.Total = usage.Total / 1024 / 1024
		d.Free = usage.Free / 1024 / 1024
		d.Used = usage.Used / 1024 / 1024
		d.UsedPercent = usage.UsedPercent

		Disks = append(Disks, d)
	}

	return Disks
}
