package main

import (
	"fmt"
)

//Write a function to find the longest common prefix string amongst an array of strings.
//
//If there is no common prefix, return an empty string "".



func main() {


	input := []string{"flower","flow","flight"}

	longestCommonPrefix := longestCommonPrefix(input)

	fmt.Print(longestCommonPrefix)
}


func longestCommonPrefix(strs []string) string {

	if len(strs) == 0 {
		return ""
	}

	if len(strs) == 1  {
		return strs[0]
	}

	longestCommon := ""
	prefixes := strs[0]
	for x:=1; x< len(strs); x++ {

			prefixes := longestPrefixesOfTwoWords(prefixes, strs[x])

			if prefixes != "" {
					longestCommon = prefixes

			} else {
				break
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