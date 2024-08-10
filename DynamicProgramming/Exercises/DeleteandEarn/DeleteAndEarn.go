package main

import (
	"sort"
)

func maxInt(a, b int) int {
	if a < b {
		return b
	} else {
		return a
	}
}

func deleteAndEarn(nums []int) int {

	var keySlice []int = nil

	sort.SliceStable(nums, func(i, j int) bool {
		return nums[i] < nums[j]
	})

	var maxMemo map[int]int = make(map[int]int)
	for _, val := range nums {
		if _, ok := maxMemo[val]; !ok {
			keySlice = append(keySlice, val)
		}
		maxMemo[val] = maxMemo[val] + val
	}

	var includedSlice []int = make([]int, len(maxMemo))
	var excludedSlice []int = make([]int, len(maxMemo))
	currentVal := 0
	includedSlice[0] = maxMemo[keySlice[0]]
	for i := 1; i < len(keySlice); i++ {

		currentVal = keySlice[i]

		if keySlice[i]-keySlice[i-1] != 1 {
			includedSlice[i] = maxInt(maxMemo[currentVal]+includedSlice[i-1], maxMemo[currentVal]+excludedSlice[i-1])

		} else {
			includedSlice[i] = maxInt(includedSlice[i-1], maxMemo[currentVal]+excludedSlice[i-1])
		}
		excludedSlice[i] = maxInt(excludedSlice[i-1], includedSlice[i-1])
	}

	return includedSlice[len(keySlice)-1]
}

func main() {
	//input := []int{1, 1, 1, 2, 4, 5, 5, 5, 6}
	//input := []int{2, 2, 3, 3, 3, 4}
	//3451
	input := []int{12, 32, 93, 17, 100, 72, 40, 71, 37, 92, 58, 34, 29, 78, 11, 84, 77, 90, 92, 35, 12, 5, 27, 92, 91, 23, 65, 91, 85, 14, 42, 28, 80, 85, 38, 71, 62, 82, 66, 3, 33, 33, 55, 60, 48, 78, 63, 11, 20, 51, 78, 42, 37, 21, 100, 13, 60, 57, 91, 53, 49, 15, 45, 19, 51, 2, 96, 22, 32, 2, 46, 62, 58, 11, 29, 6, 74, 38, 70, 97, 4, 22, 76, 19, 1, 90, 63, 55, 64, 44, 90, 51, 36, 16, 65, 95, 64, 59, 53, 93}

	deleteAndEarn(input)

}
