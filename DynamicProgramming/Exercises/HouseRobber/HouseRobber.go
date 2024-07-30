package main

import (
	"fmt"
)

var input []int = []int{2, 7, 9, 3, 1}
var houseValue []int
var maxProfit map[int]int = make(map[int]int)

func maxInt(a, b int) int {
	if a < b {
		return b
	} else {
		return a
	}
}

func calculatePotentialValue(i int) int {

	if i == 0 {
		fmt.Println("Index is 0")

		return maxProfit[0]
	} else if i == 1 {
		fmt.Println("Index is 1")
		return maxProfit[1]
	} else if _, ok := maxProfit[i]; !ok {
		fmt.Printf("Index is  %v \n", i)
		maxProfit[i] = maxInt(calculatePotentialValue(i-2)+houseValue[i], calculatePotentialValue(i-1))
	}

	return maxProfit[i]
}

func rob(nums []int) int {

	if len(nums) >= 2 {
		maxProfit[0] = nums[0]
		maxProfit[1] = maxInt(nums[0], nums[1])
		houseValue = make([]int, len(nums))
		_ = copy(houseValue, nums)
		calculatePotentialValue(len(nums) - 1)
		fmt.Println(maxProfit)

		return maxInt(maxProfit[len(nums)-2], maxProfit[len(nums)-1])
	}
	return nums[0]

}
func main() {

	fmt.Println(rob(input))

}
