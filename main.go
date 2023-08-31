package main

import (
	"fmt"
	"time"
)

func main() {

	benchmarkTime := time.Now()

	// Run code for execution time measurement
	time.Sleep(1 * time.Second)

	timeElapsed := time.Since(benchmarkTime)
	fmt.Println("Time elapsed: ", timeElapsed)

}
