package main

/*
表达式计算。
如一下输入:
4
   xx = 5
yy=4
zz   = 3
   d = xx +yy+ zz

输出d的值为12
*/

import (
	"fmt"
	"strconv"
	"strings"
)

var m = make(map[string]int64, 0)

func readLine(line string) string {
	var v []string
	var s []int64
	isVar := true
	line = strings.Trim(line, " ")
	//fmt.Print64ln(line)
	tempV := ""
	for i := 0; i < len(line); i++ {
		if line[i] == ' ' || line[i] == '+' || line[i] == '=' {
			if tempV != "" {
				if !isVar {
					val, _ := strconv.Atoi(tempV)
					s = append(s, int64(val))
					isVar = true
				} else {
					v = append(v, tempV)
				}
				tempV = ""
			}
		} else if line[i] >= '0' && line[i] <= '9' {
			if tempV == "" {
				isVar = false
			}
			tempV += string(line[i])
			if i == len(line)-1 {
				if tempV != "" {
					if !isVar {
						val, _ := strconv.Atoi(tempV)
						s = append(s, int64(val))
					} else {
						v = append(v, tempV)
					}
					tempV = ""
				}
			}
		} else {
			tempV += string(line[i])
			if i == len(line)-1 {
				if tempV != "" {
					if !isVar {
						val, _ := strconv.Atoi(tempV)
						s = append(s, int64(val))
					} else {
						v = append(v, tempV)
					}
					tempV = ""
				}
			}
		}
	}
	if len(v) == 1 {
		sum := int64(0)
		for i := 0; i < len(s); i++ {
			sum += s[i]
		}
		m[v[0]] = sum
	} else {
		sum := int64(0)
		for i := 0; i < len(s); i++ {
			sum += s[i]
		}
		for i := 1; i < len(v); i++ {
			sum += m[v[i]]
		}
		m[v[0]] = sum
	}
	//fmt.Printf("%s = %d\n", v[0], m[v[0]])
	return v[0]
}

func scanLine() string {
	var c byte
	var b []byte
	for {
		fmt.Scanf("%c", &c)
		if c != '\n' {
			b = append(b, c)
		} else {
			break
		}
	}
	return string(b)
}

func main() {

	var n int64
	fmt.Scanln(&n)
	var line string
	for i := 0; i < int(n); i++ {
		//fmt.Scanln(&line)
		line = scanLine()
		vari := readLine(line)
		if i == int(n-1) {
			fmt.Println(m[vari])
		}
	}
}
