package main

import (
	"fmt"
)

//Write a function to find the longest common prefix string amongst an array of strings.
//
//If there is no common prefix, return an empty string "".



func main() {


	input := []string{"dog","racecar","car"}

	longestCommonPrefix := longestCommonPrefix(input)

	fmt.Print(longestCommonPrefix)
}


func longestCommonPrefix(strs []string) string {

	longestCommonMap := make(map[string] int)

	if len(strs) == 1  {
		return strs[0]
	}

	for x:=0; x< len(strs); x++ {

		for y:=x+1; y<len(strs); y++ {

			prefixes := longestPrefixesOfTwoWords(strs[x], strs[y])
			if longestCommonMap[prefixes] == 0 {
				longestCommonMap[prefixes] = 1
			} else {
				longestCommonMap[prefixes] = longestCommonMap[prefixes] + 1
			}
		}

	}

	value := 0
	longestCommon := ""
	for key, element := range longestCommonMap {
		if element > value {
			longestCommon = key
			value = element
		}
	}

	return longestCommon
}

func longestPrefixesOfTwoWords(firstWord string, secondWord string) string {

	splitedFirstWord := []rune(firstWord)
	splitedSecondWord := []rune(secondWord)

	longest := ""

	for x := 0; x < len(splitedFirstWord); x++ {

		if x <= len(splitedSecondWord) - 1 {
			if splitedFirstWord[x] == splitedSecondWord[x] {
				longest += string(splitedFirstWord[x])
			 } else {
			 	break
			}
		}
	}
	return longest
}