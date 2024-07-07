package main

import "fmt"

func main() {
	cat, dog, fish, stopCh := make(chan struct{}), make(chan struct{}), make(chan struct{}), make(chan struct{})
	iter := 3
	go printStr("cat", false, iter, cat, dog, stopCh)
	go printStr("dog", false, iter, dog, fish, stopCh)
	go printStr("fish", true, iter, fish, cat, stopCh)
	cat <- struct{}{}

	<-stopCh
}

func printStr(str string, last bool, iter int, in, out chan struct{}, stopCh chan struct{}) {
	count := 0
	for {
		select {
		case <-in:
			fmt.Println(str)
			count++
			if count == iter && last {
				stopCh <- struct{}{}
				return // 注意要 return
			}
			out <- struct{}{}
		case <-stopCh:
			return
		}
	}
}
