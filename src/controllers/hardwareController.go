package controller

import (
	"github.com/gin-gonic/gin"
	"github.com/shirou/gopsutil/cpu"
	"github.com/shirou/gopsutil/disk"
	"github.com/shirou/gopsutil/host"
	"github.com/shirou/gopsutil/mem"
	"github.com/shirou/gopsutil/net"
	"net/http"
	"strconv"
)

func MemoryInfo() gin.HandlerFunc {
	return func(context *gin.Context) {
		vmStat, err := mem.VirtualMemory()
		if err != nil {
			return
		}
		context.JSON(http.StatusOK, gin.H{
			"total_memory": strconv.FormatUint(vmStat.Total, 10),
			"free_memory":  strconv.FormatUint(vmStat.Free, 10),
			"percentage":   strconv.FormatFloat(vmStat.UsedPercent, 'f', 2, 64),
			"raw_value":    vmStat,
		})
	}
}
func CPUInfo() gin.HandlerFunc {
	return func(context *gin.Context) {
		cpuStat, _ := cpu.Info()
		percentage, _ := cpu.Percent(0, true)
		context.JSON(http.StatusOK, gin.H{
			"percentage": percentage,
			"info":       cpuStat,
		})
	}
}

func HostInfo() gin.HandlerFunc {
	return func(context *gin.Context) {
		hostStat, _ := host.Info()
		context.JSON(http.StatusOK, gin.H{
			"raw_value": hostStat,
		})
	}
}
func InterfaceInfo() gin.HandlerFunc {
	return func(context *gin.Context) {
		interfStat, _ := net.Interfaces()
		context.JSON(http.StatusOK, gin.H{
			"raw_value": interfStat,
		})
	}
}
func DiskInfo() gin.HandlerFunc {
	return func(context *gin.Context) {
		diskStat, _ := disk.Usage("/")
		context.JSON(http.StatusOK, gin.H{
			"raw_value":        diskStat,
			"total_disk_space": strconv.FormatUint(diskStat.Total, 10),
			"usage_percent":    strconv.FormatFloat(diskStat.UsedPercent, 'f', 2, 64),
		})
	}
}
