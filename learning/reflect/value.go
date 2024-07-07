package main

import (
	"fmt"
	"reflect"
)

func main() {
	c := make(chan int, 4)
	v := reflect.ValueOf(c)
	if v.TrySend(reflect.ValueOf(100)) {
		fmt.Println(v.TryRecv())
	}
}
