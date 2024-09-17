package main

import "fmt"

//https://leetcode.com/problems/uncommon-words-from-two-sentences/description/?envType=daily-question&envId=2024-09-17

func uncommonFromSentences(s1 string, s2 string) []string {

	var wordCount map[string]int = make(map[string]int)
	var currentWord string = ""
	for i := 0; i < len(s1); i++ {

		if string(s1[i]) == " " {
			wordCount[currentWord]++
			currentWord = ""
			continue
		}
		currentWord = currentWord + string(s1[i])
	}
	if currentWord != "" {
		wordCount[currentWord]++
		currentWord = ""
	}
	for i := 0; i < len(s2); i++ {

		if string(s2[i]) == " " {
			wordCount[currentWord]++
			currentWord = ""
			continue
		}
		currentWord = currentWord + string(s2[i])
	}
	if currentWord != "" {
		wordCount[currentWord]++
		currentWord = ""
	}

	uniqSlice := make([]string, 0)

	for i, v := range wordCount {
		if v == 1 {
			uniqSlice = append(uniqSlice, i)
		}
	}

	return uniqSlice
}

func main() {
	s1 := "apple apple"
	s2 := "banana"
	res := uncommonFromSentences(s1, s2)

	fmt.Println(res)
}
