package main

import (
	"fmt"
	"sync"
	"time"
)

func multiProcess(num []int, n int) int {
	batchSize := 10000
	iters := make(chan int, n/batchSize)

	wg := &sync.WaitGroup{}
	for i := 0; i < n/batchSize; i++ {
		wg.Add(1)
		go func(num []int) {
			sum := 0
			for _, val := range num {
				sum += val
			}
			//fmt.Printf("goroutine: %d\n", sum)
			iters <- sum
			defer wg.Done()
		}(num[i*batchSize : (i+1)*batchSize])
	}
	wg.Wait()
	sum := 0
	close(iters)
	for v := range iters {
		sum += v
	}
	return sum
}

func loop(num []int, n int) int {
	sum := 0
	for _, val := range num {
		sum += val
	}
	return sum
}

func main() {
	n := 500000000
	var num []int
	for i := 0; i < n; i++ {
		num = append(num, 6)
	}
	startTime := time.Now()
	sum := multiProcess(num, n)
	endTime := time.Now()
	fmt.Printf("multiProcess: time: %d, sum: %d\n", (endTime.UnixNano()-startTime.UnixNano())/1000000, sum)

	startTime = time.Now()
	sum = loop(num, n)
	endTime = time.Now()
	fmt.Printf("loop: time: %d, sum: %d\n", (endTime.UnixNano()-startTime.UnixNano())/1000000, sum)
}
