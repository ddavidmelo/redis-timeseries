package system

import (
	"time"

	log "github.com/sirupsen/logrus"

	"github.com/shirou/gopsutil/v3/cpu"
	"github.com/shirou/gopsutil/v3/disk"
	"github.com/shirou/gopsutil/v3/mem"
	"github.com/shirou/gopsutil/v3/net"
)

// SysInfo saves the basic system information
type SysInfo struct {
	CPUPercentage    []uint8
	RAMPercentage    uint8
	DiskPercentage   uint8
	NetworkReceiving uint64
	NetworkSending   uint64
}

func GetSystemStatus() SysInfo {
	var info SysInfo

	vmStat, _ := mem.VirtualMemory()
	diskStat, _ := disk.Usage("/") // If you're in Unix change this "\\" for "/"

	cpu.Info()
	cpuPercentage, _ := cpu.Percent(1*time.Second, true)
	for _, c := range cpuPercentage {
		info.CPUPercentage = append(info.CPUPercentage, uint8(c))
	}

	network, _ := net.IOCounters(true)
	time.Sleep(time.Second * 1)
	networkNew, _ := net.IOCounters(true)

	for i, v := range networkNew {
		if v.Name == "wlp3s0" {
			info.NetworkReceiving = v.BytesRecv - network[i].BytesRecv
			info.NetworkSending = v.BytesSent - network[i].BytesSent
		}
	}

	info.RAMPercentage = uint8(vmStat.UsedPercent)
	info.DiskPercentage = uint8(diskStat.UsedPercent)

	log.Debugf("metric: %+v", info)
	return info
}
