package main

import "fmt"

func calcTribonacci(memo map[int]int, index int) int {

	if _, ok := memo[index]; !ok {
		memo[index] = calcTribonacci(memo, index-1) + calcTribonacci(memo, index-2) + calcTribonacci(memo, index-3)
		return memo[index]
	} else {
		return memo[index]
	}

}

func tribonacci(n int) int {
	var memoization map[int]int = make(map[int]int)
	memoization[0] = 0
	memoization[1] = 1
	memoization[2] = 1
	if n == 0 {
		return 0
	} else if n <= 2 {
		return 1
	} else {
		return calcTribonacci(memoization, n)
	}
}

func main() {
	fmt.Println(tribonacci(25))
}
