package main

import (
	"fmt"
	"log"
	"os"
	"time"

	"github.com/shirou/gopsutil/v3/cpu"
	"github.com/shirou/gopsutil/v3/mem"
	"github.com/shirou/gopsutil/v3/process"
	// "github.com/shirou/gopsutil/mem"  // to use v2
)

func main() {

	var (
		// cpuPercentUsage    atomic.Value //  avg cpu percent in multi cpu core 100% is the max percent
		// memoryPercentUsage atomic.Value // 100% is the max percent
		// memoryUsagebytes   atomic.Value

		// memoryStatCollectorOnce sync.Once
		// cpuStatCollectorOnce    sync.Once

		CurrentPID = os.Getpid()
		// currentProcess     atomic.Value
	)
	v, _ := mem.VirtualMemory()

	// almost every return value is a struct
	fmt.Printf("Total: %v, Free:%v, UsedPercent:%f%%\n", v.Total, v.Free, v.UsedPercent)

	// convert to JSON. String() is also implemented
	fmt.Println(v)

	p, _ := process.NewProcess(int32(241))

	pCPU, _ := p.CPUPercent()

	log.Println(CurrentPID, p, pCPU)

	c, _ := cpu.Percent(time.Second, true)

	log.Println(c)
}
