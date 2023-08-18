// test.go
package main

import (
	"fmt"
	"runtime"
	"time"

	_ "go.uber.org/automaxprocs"
)

func main() {
	fmt.Println(runtime.NumCPU())
	fmt.Println(runtime.GOMAXPROCS(0))
	time.Sleep(time.Second)

	fmt.Println("test")
}
