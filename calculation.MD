

### Links

https://docs.kernel.org/admin-guide/cgroup-v2.html#:~:text=A%20read%2Dwrite%20two%20value,default%20is%20%22max%20100000%22.

https://github.com/uber-go/automaxprocs/blob/master/internal/cgroups/cgroups2.go

https://octopus.com/blog/kubernetes-pod-cpu-memory

https://stackoverflow.com/questions/45196440/cpupercent-metric-from-docker-stats-vs-cgroups

In cgroup v2, the `cpu.max` parameter sets a limit on the percentage of CPU time a cgroup can use over a period of time. It's a way to control and limit CPU usage for a cgroup. The value you set for `cpu.max` specifies how much CPU time the cgroup can utilize relative to the total available CPU time.

If you want to compare `cpu.max` to CPU clock time, you need to understand how `cpu.max` and the actual CPU usage work in cgroup v2:

1. **`cpu.max` Value:**
   `cpu.max` is a value in the range of `[0-100000]` where `100000` represents 100% of a single CPU core's time. For example, setting `cpu.max` to `50000` would mean that the cgroup is allowed to use up to 50% of the available CPU time on a single core over the defined period.

2. **CPU Clock Time:**
   CPU time is a measurement of the time that a CPU has spent processing tasks. It's usually measured in "clock ticks" or "jiffies." In Linux, a jiffy is defined as the time between two timer interrupts. The `clock()` function in the C standard library can be used to get the time spent by a process in CPU time.

    https://medium.com/@betz.mark/understanding-resource-limits-in-kubernetes-cpu-time-9eff74d3161b
   The bandwidth control system defines a period, which is usually 1/10 of a second, or 100000 microseconds,

To compare `cpu.max` with CPU clock time, you'd need to do some calculations. Keep in mind that this is a simplification and that there might be some variability due to various factors, including the underlying OS scheduling and CPU architecture.

Here's a simplified approach in Go:

```go
package main

import (
	"fmt"
	"time"
	"syscall"
)

func main() {
	// Assume cpu.max is set to 50000 (50%)
	cpuMaxPercent := 50.0
	cpuMaxJiffies := int64(cpuMaxPercent * 1000) // Convert percentage to jiffies

	startTime := syscall.Times(nil).Utime

	// Simulate CPU-bound work
	for i := 0; i < 100000000; i++ {
	}

	endTime := syscall.Times(nil).Utime

	cpuTime := endTime - startTime

	fmt.Printf("CPU Time: %d jiffies\n", cpuTime)
	fmt.Printf("CPU Max (jiffies): %d jiffies\n", cpuMaxJiffies)

	if cpuTime > cpuMaxJiffies {
		fmt.Println("Exceeded CPU limit.")
	} else {
		fmt.Println("Within CPU limit.")
	}
}
```

This example uses `syscall.Times` to measure the CPU time spent by the process and compares it to the `cpu.max` value. Please note that this is a simplified illustration and might not accurately reflect the actual behavior due to various factors affecting system performance and scheduling.

For accurate and production-level monitoring, you might need to use more advanced tools and libraries that offer detailed resource usage statistics.