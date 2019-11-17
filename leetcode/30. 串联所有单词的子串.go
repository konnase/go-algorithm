package main

import "fmt"

func findSubstring(s string, words []string) []int {
	if len(s) == 0 || len(words) == 0 {
		return []int{}
	}
	wordsCount := make(map[string]int)
	for _, word := range words {
		count, _ := wordsCount[word]
		wordsCount[word] = count + 1
	}
	//fmt.Println(wordsCount)
	lenW := len(words)
	lenW0 := len(words[0])
	var result []int
	for start := 0; start < lenW0; start++ {
		i := start
		count := 0
		tempWordsCount := make(map[string]int)
		for _, word := range words {
			tempWordsCount[word] = 0
		}
		for ; ; {
			count += 1
			if i+lenW0 > len(s) {
				break
			}
			subStr := s[i:i+lenW0]
			tempWordCount, ok := tempWordsCount[subStr]
			if ok {
				if count == lenW && tempWordCount+1 == wordsCount[subStr] {
					tempWordsCount[subStr] += 1
					item := i-lenW0*(lenW-1)
					result = append(result, item)
					count -= 1
					i += lenW0
					dropStr := s[item:item+lenW0]
					tempWordsCount[dropStr] -= 1
					continue
				}
				if tempWordCount+1 > wordsCount[subStr] {
					j := i - (count-1) * lenW0
					dropStr := s[j:j+lenW0]
					tempWordsCount[dropStr] -= 1
					count -= 2
					continue
				}
				tempWordsCount[subStr] += 1
				i += lenW0
			} else {
				for _, word := range words {
					tempWordsCount[word] = 0
				}
				i += lenW0
				count = 0
			}
		}
	}
	return result
}

func main()  {
	//s := "barfoothefoobarman"
	//words := []string{"foo","bar"}
	s := "wordgoodgoodgoodbestword"
	//words := []string{"word","good", "best", "word"}
	words := []string{"word","good", "best", "good"}
	result := findSubstring(s, words)
	fmt.Println(result)
}