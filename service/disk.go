package service

import (
	"github.com/shirou/gopsutil/v3/disk"
	log "github.com/sirupsen/logrus"
)

type Device struct {
	Name        string
	Mountpoint  string
	Fstype      string
	Total       uint64
	Free        uint64
	Used        uint64
	UsedPercent float64
}

//var Disks []*Device

func GetDisk() []*Device {
	var Disks []*Device

	partitions, err := disk.Partitions(false)
	if err != nil {
		log.Error(err)
	}

	for _, dev := range partitions {

		d := &Device{}
		d.Name = dev.Device
		d.Mountpoint = dev.Mountpoint
		d.Fstype = dev.Fstype

		usage, err := disk.Usage(dev.Mountpoint)
		if err != nil {
			log.Error(err)
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
