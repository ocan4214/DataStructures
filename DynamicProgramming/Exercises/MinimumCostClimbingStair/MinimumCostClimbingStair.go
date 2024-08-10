package main

import "fmt"

func minInt(a, b int) int {
	if a < b {
		return a
	} else {
		return b
	}

}

func Climb(i int, cost []int, minCostPerStair []int) int {

	// i. adım ı geçmen lazım veya i. adıma gelip tırmanman
	minVal := 0
	if len(cost) == 3 {
		return minInt(cost[0]+cost[2], cost[1])
	}

	if i+1 > len(cost)-1 {
		minVal = minInt(minCostPerStair[i-1], cost[i]+minCostPerStair[i-2])
		return minVal
	}

	if minCostPerStair[i-1] < minCostPerStair[i-2] {
		minVal = cost[i] + minCostPerStair[i-1]

	} else {
		minVal = cost[i] + minCostPerStair[i-2]
	}

	return minVal

}
func minCostClimbingStairs(cost []int) int {

	if len(cost) == 2 {
		return minInt(cost[0], cost[1])
	} else if len(cost) > 2 {
		minCostPerStair := make([]int, len(cost))
		minCostPerStair[0] = cost[0]
		minCostPerStair[1] = cost[1]

		for i := 2; i < len(cost); i++ {
			minCostPerStair[i] = Climb(i, cost, minCostPerStair)
		}

		fmt.Println(minCostPerStair)
		return minCostPerStair[len(cost)-1]
	}

	return -1
}

func main() {
	input := []int{1, 100, 1, 5, 6, 100, 99, 1, 100, 1}
	//input := []int{10, 15, 25}
	//input := []int{0, 0, 1, 1}
	//input := []int{0, 1, 1, 0}
	//input := []int{0, 1, 1, 1}
	minCostClimbingStairs(input)
}
