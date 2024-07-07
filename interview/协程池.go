// 构造一个协程池，能消费并处理一个数组中的数据，当所有数据消费完成后，关闭协程池。数组中的每个数据，只能被一个协程处理。

package main

import (
	"fmt"
	"sync"
	"time"
)

func main() {
	ch := make(chan int, 5)
	a := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	routineSize := 3

	wg := sync.WaitGroup{}
	for i := 0; i < routineSize; i++ {
		wg.Add(1)
		go consumer(i, ch, &wg)
	}

	for _, v := range a {
		ch <- v
	}

	close(ch)
	wg.Wait()
}

func consumer(id int, ch chan int, wg *sync.WaitGroup) {
	defer wg.Done()
	for i := range ch {
		fmt.Printf("Consumer %d: %d\n", id, i)
		time.Sleep(10 * time.Millisecond)
	}
}
