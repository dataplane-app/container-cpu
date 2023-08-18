package main

import (
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
	"time"

	"github.com/shirou/gopsutil/v3/process"
)

func readCgroupFile(path string) (map[string]int64, error) {
	content, err := os.ReadFile(path)
	if err != nil {
		return nil, err
	}

	log.Println(string(content))

	lines := strings.Split(string(content), "\n")
	values := make(map[string]int64)

	for _, line := range lines {
		parts := strings.Fields(line)
		if len(parts) != 2 {
			continue
		}
		key := parts[0]
		value, err := strconv.ParseInt(parts[1], 10, 64)
		if err != nil {
			return nil, err
		}
		values[key] = value
	}

	return values, nil
}

func calculateCPUUsagePercent(stats map[string]int64) float64 {

	log.Println(stats["user_usec"], stats["system_usec"])
	userUsec := stats["user_usec"]
	systemUsec := stats["system_usec"]
	totalUsec := userUsec + systemUsec

	// microseconds to seconds

	// divide by 10 i.e. 1/10 of a second = 100,000 microsoeconds
	return float64(totalUsec) / 10000000000 // Convert to seconds
}

func main() {
	cgroupPath := "/sys/fs/cgroup"
	cpuStatPath := cgroupPath + "/cpu.stat"

	cpuStatValues, err := readCgroupFile(cpuStatPath)
	if err != nil {
		fmt.Println("Error reading cpu.stat:", err)
		return
	}

	cpuPressure := cgroupPath + "/cpu.pressure"
	content, _ := os.ReadFile(cpuPressure)
	log.Println("CPU pressure", content)

	log.Println("PID", os.Getpid())

	p, _ := process.NewProcess(int32(os.Getpid()))

	pCPU, _ := p.CPUPercent()

	log.Println("process cpu", pCPU)

	done := make(chan int)

	for i := 0; i < 1; i++ {
		go func() {
			for {
				select {
				case <-done:
					return
				default:
				}
			}
		}()
	}

	go func() {
		for {
			select {
			case <-done:
				return
			default:
				pCPU, _ := p.CPUPercent()

				log.Println("process cpu sec:", pCPU)
			}
			time.Sleep(1 * time.Second)
		}
	}()

	// log.Println("process cpu", pCPU)

	time.Sleep(10 * time.Second)
	close(done)

	log.Println("process cpu", pCPU)

	cpuUsagePercent := calculateCPUUsagePercent(cpuStatValues)
	fmt.Printf("CPU Usage: %.2f%%\n", cpuUsagePercent)
}
