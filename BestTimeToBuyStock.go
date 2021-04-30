//You are given an array prices where prices[i] is the price of a given stock on the ith day.
//
//You want to maximize your profit by choosing a single day to buy one stock and choosing a different day in the future to sell that stock.
//
//Return the maximum profit you can achieve from this transaction. If you cannot achieve any profit, return 0.
//
//
//
//Example 1:
//
//Input: prices = [7,1,5,3,6,4]
//Output: 5
//Explanation: Buy on day 2 (price = 1) and sell on day 5 (price = 6), profit = 6-1 = 5.
//Note that buying on day 2 and selling on day 1 is not allowed because you must buy before you sell.
//Example 2:
//
//Input: prices = [7,6,4,3,1]
//Output: 0
//Explanation: In this case, no transactions are done and the max profit = 0.

package main

import "fmt"

func main() {

	prices := []int{7, 6, 4, 3, 1}

	fmt.Println(maxProfit(prices))
}

func maxProfit(prices []int) int {

	//Find best buy day
	//Find the lowest

	lowestPrice := 0
	lowestPriceIndex := 0
	for x := 0; x < len(prices); x++ {
		if x == 0 {
			lowestPrice = prices[x]
			lowestPriceIndex = 0
		} else {
			if prices[x] < lowestPrice {
				lowestPrice = prices[x]
				lowestPriceIndex = x
			}
		}
	}

	highestPrice := 0
	highestPriceIndex := lowestPriceIndex
	for x := lowestPriceIndex; x < len(prices); x++ {
		if x == 0 {
			highestPrice = prices[x]
			highestPriceIndex = 0
		} else {
			if prices[x] > highestPrice && x > lowestPriceIndex {
				highestPrice = prices[x]
				highestPriceIndex = x
			}
		}
	}

	return prices[highestPriceIndex] - prices[lowestPriceIndex]
}
