package main

import "fmt"

func main() {

	input := 2012
	fmt.Println(intToRoman(input))

}

func intToRoman(num int) string {

	m := make(map[string]int)

	m["M"] = 1000
	m["CM"] = 900
	m["D"] = 500
	m["CD"] = 400
	m["C"] = 100
	m["XC"] = 90
	m["L"] = 50
	m["XL"] = 40
	m["X"] = 10
	m["IX"] = 9
	m["V"] = 5
	m["IV"] = 4
	m["I"] = 1

	notFinish := true

	computedNum := num
	roman := ""
	for ok := true; ok; ok = notFinish {

		//for k,v := range m {
		//
		//	if v == computedNum {
		//		roman = roman + k
		//		computedNum = computedNum - v
		//	} else if computedNum > v
		//}

		if computedNum == 0 {
			notFinish = false
		}
	}

	return roman

}
