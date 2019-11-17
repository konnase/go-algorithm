package main

import (
	"fmt"
	"math"
)

func divide(dividend int, divisor int) int {
	dvd := int32(dividend)
	dvs := int32(divisor)
	if dvd == dvs {
		return 1
	}
	if dvs == math.MinInt32 {
		return 0
	}
	// sign == true: same sign
	sign := (dvd >= 0) == (dvs >= 0)
	if dvd > 0 {
		dvd = -dvd
	}
	if dvs > 0 {
		dvs = -dvs
	}
	result := int32(0)
	for ; dvd <= dvs;  {
		tempDvs := dvs
		tempResult := int32(-1)
		for ; dvd <= (tempDvs << 1); {
			if tempDvs <= math.MinInt32 >> 1 {
				break
			}
			tempResult = tempResult << 1
			tempDvs = tempDvs << 1
		}
		result += tempResult
		dvd -= tempDvs
	}
	if sign {
		if result == math.MinInt32 {
			return math.MaxInt32
		}
		result = -result
	}
	return int(result)
}

func main() {
	//dividend := 2147483647
	//divisor := 3
	dividend := -2147483648
	divisor := 1
	fmt.Println(divide(dividend, divisor))
}