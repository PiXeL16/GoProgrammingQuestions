package main

import "fmt"

//Given an integer numRows, return the first numRows of Pascal's triangle.
//
//In Pascal's triangle, each number is the sum of the two numbers directly above it as shown:

//Example 1:
//
//Input: numRows = 5
//Output: [[1],[1,1],[1,2,1],[1,3,3,1],[1,4,6,4,1]]
//Example 2:
//
//Input: numRows = 1
//Output: [[1]]

func main() {

	numRows := 4
	fmt.Println(getRow(numRows))

}

func generate(numsRows int) [][]int {

	triangle := make([][]int, numsRows)

	if numsRows == 0 {
		return triangle
	}

	for x := 0; x < numsRows; x++ {
		var row []int
		if x == 0 {
			row = buildRow(x, nil)
		} else {
			row = buildRow(x, triangle[x-1])
		}

		triangle[x] = row
	}

	return triangle

}

func buildRow(rowIndex int, pastRow []int) []int {
	row := make([]int, rowIndex+1)
	if rowIndex == 0 {
		row = []int{1}
		return row
	}

	//Add 1 on Left Edge
	row[0] = 1
	newRowIndex := 1
	//Iterate over array and check pairs
	for x := 0; x < len(pastRow); x++ {
		if x+1 < len(pastRow) {
			rowValue := pastRow[x] + pastRow[x+1]
			row[newRowIndex] = rowValue
			newRowIndex++
		}
	}
	//Add 1 on the Right Edge
	row[newRowIndex] = 1

	return row
}

func getRow(rowIndex int) []int {

	rowIndex = rowIndex + 1

	triangle := make([][]int, rowIndex)

	if rowIndex == 0 {
		return triangle[0]
	}

	for x := 0; x < rowIndex; x++ {
		var row []int
		if x == 0 {
			row = buildRow(x, nil)
		} else {
			row = buildRow(x, triangle[x-1])
		}

		triangle[x] = row
	}

	return triangle[rowIndex-1]

}
