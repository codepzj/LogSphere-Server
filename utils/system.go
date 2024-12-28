package utils

import (
	"fmt"
	"github.com/shirou/gopsutil/cpu"
	"github.com/shirou/gopsutil/disk"
	"github.com/shirou/gopsutil/mem"
	"strconv"
	"time"
)

func GetSystemUsage() map[string]float64 {
	usageMap := map[string]float64{
		"cpu":    0.00,
		"memory": 0.00,
		"disk":   0.00,
	}

	// 获取cpu使用情况
	if CpuPercent, err := cpu.Percent(1*time.Second, false); err == nil {
		usageMap["cpu"], _ = strconv.ParseFloat(fmt.Sprintf("%.2f", CpuPercent[0]), 64)
	}
	// 获取内存使用情况
	if vmStat, err := mem.VirtualMemory(); err == nil {
		usageMap["memory"], _ = strconv.ParseFloat(fmt.Sprintf("%.2f", vmStat.UsedPercent), 64)
	}
	// 获取硬盘使用情况
	if diskStat, err := disk.Usage("/"); err == nil {
		usageMap["memory"], _ = strconv.ParseFloat(fmt.Sprintf("%.2f", diskStat.UsedPercent), 64)
	}

	return usageMap
}
