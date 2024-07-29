package main

import (
	"fmt"
)

func lengthOfLIS(nums []int) int {

	sub = append(sub, nums[0])
	var currentIndex int = 0
	var minIndex int = 0
	for i := 1; i < len(nums); i++ {
		fmt.Println(sub)
		if nums[i] > sub[currentIndex] {
			currentIndex++
			sub = append(sub, nums[i])

		} else {
			minIndex = binarySearch(sub, nums[i])
			sub[minIndex] = nums[i]
		}
	}

	return len(sub)
}

func binarySearch(subArr []int, val int) int {
	left := 0
	right := len(subArr) - 1
	var middle int = (left + right) / 2
	for left < right {
		if val == subArr[middle] {
			return middle
		} else if val > subArr[middle] {
			left = middle + 1
		} else {
			right = middle
		}
		middle = (left + right) / 2
	}

	return left
}

var inarr []int = []int{7, 5, 6, 8, 2, 1, 2, 3, 4, 5, 6, 7, 8}
var sub []int = make([]int, 0, cap(inarr))

func main() {

	var res int = lengthOfLIS(inarr)
	fmt.Printf("****Longest increasing Subsequence %v **** \n ", res)

}
