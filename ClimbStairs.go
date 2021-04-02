//You are climbing a staircase. It takes n steps to reach the top.
//
//Each time you can either climb 1 or 2 steps. In how many distinct ways can you climb to the top?

package main

import "fmt"

func main() {

	stairs := 3

	fmt.Println(climbStairs(stairs))
}

func climbStairs(n int) int {

	return climbStairsRec(0, n)
}

func climbStairsRec(i int, n int) int {

	if i > n {
		return 0
	}

	if i == n {
		return 1
	}

	return climbStairsRec(i+1, n) + climbStairsRec(i+2, n)
}
