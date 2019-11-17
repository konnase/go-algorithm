package main

import (
	"fmt"
	"strconv"
	"strings"
)

func main()  {
	var str string
	fmt.Scanln(&str)
	var n int
	fmt.Scanf("%d", &n)
	str = strings.TrimLeft(str, "{")
	str = strings.TrimRight(str, "}")
	numStr := strings.Split(str, ",")
	var count int
	for _, s := range numStr {

		num, _ := strconv.Atoi(s)
		for i:=0;i<num/2;i++ {
			if
		}
	}
}
