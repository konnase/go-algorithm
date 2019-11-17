package main

/*
catsanddog
cat,dog,and,cats,sand

pineapplepenapple
pine,applepen,apple,pen,pineapple
*/

import (
	"fmt"
	"strings"
)

var count = 0

func combination(str string, prev string, phraseList []string) {
	if len(str) == 0 {
		return
	}
	newStr := prev + string(str[0])
	fmt.Println(newStr)
	newLen := len(newStr)
	//equalCount := 0
	for i := 0; i < len(phraseList); i++ {
		if newLen > len(phraseList[i]) {
			continue
		}
		if newStr == phraseList[i][:newLen] {
			if len(str) == 1 {
				count++
				return
			}
			if newLen == len(phraseList[i]) {
				combination(str[1:], "", phraseList)
			}
			combination(str[1:], newStr, phraseList)
		}
	}
	//todo: 如果没有一个phrase与newStr相等
}

func main() {
	var str, phrase string
	fmt.Scanln(&str)
	fmt.Scanln(&phrase)
	phraseList := strings.Split(phrase, ",")
	fmt.Println(phraseList)
	combination(str, "", phraseList)
	fmt.Println(count)
}
