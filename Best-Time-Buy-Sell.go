package main

import "fmt"

func maxProfit(prices []int) int {
	minTillDate, maxProfit := prices[0], 0
	var profit int
	for _, v := range prices {
		profit = v - minTillDate
		if profit > maxProfit {
			maxProfit = profit
		}
		if v < minTillDate {
			minTillDate = v
		}
	}
	return maxProfit
}

func main() {
	prices := []int{7, 6, 4, 3, 1}
	fmt.Println(maxProfit(prices))
}
