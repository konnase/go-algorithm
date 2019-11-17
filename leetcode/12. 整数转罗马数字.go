package main

import (
	"bytes"
	"fmt"
)

func intToRoman(num int) string {
	nums := []int{1000,900,500,400,100,90,50,40,10,9,5,4,1}
	romans := []string{"M", "CM", "D", "CD", "C", "XC", "L", "XL", "X", "IX", "V", "IV", "I"}
	result := bytes.Buffer{}
	i := 0
	for ;num > 0;{
		if num >= nums[i] {
			result.WriteString(romans[i])
			num -= nums[i]
		} else {
			i++
		}
	}
	return result.String()
}

func main()  {
	a := 1994
	fmt.Println(intToRoman(a))
}