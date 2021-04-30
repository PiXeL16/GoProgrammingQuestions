package main

import "fmt"

func main() {

	countTill1000()
}

func countTill1000() {

	countRecursive(0, 1000)
}

func countRecursive(num int, limit int) {

	if num <= limit {
		fmt.Println(num)
		num++
		countRecursive(num, limit)
	}
}
