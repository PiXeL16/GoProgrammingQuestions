//Given a string s, find the length of the longest substring without repeating characters.
//
//Example 1:
//
//Input: s = "abcabcbb"
//Output: 3
//Explanation: The answer is "abc", with the length of 3.
//Example 2:
//
//Input: s = "bbbbb"
//Output: 1
//Explanation: The answer is "b", with the length of 1.
//Example 3:
//
//Input: s = "pwwkew"
//Output: 3
//Explanation: The answer is "wke", with the length of 3.
//Notice that the answer must be a substring, "pwke" is a subsequence and not a substring.
//Example 4:
//
//Input: s = ""
//Output: 0

package main

import "fmt"

func main() {

	input := "au"
	fmt.Println(lengthOfLongestSubstring(input))

}

func lengthOfLongestSubstring(s string) int {

	if s == "" {
		return 0
	}

	var characters []rune
	maxCount := 0
	splitedCharacters := []rune(s)
	for x := 0; x < len(splitedCharacters); x++ {

		if len(characters) > maxCount {
			maxCount = len(characters)
		}

		characters = nil

		for y := x; y < len(splitedCharacters); y++ {

			c := splitedCharacters[y]

			if containsCharacter(characters, c) {
				if len(characters) > maxCount {
					maxCount = len(characters)
				}
				characters = nil
				characters = append(characters, c)
			} else {
				characters = append(characters, c)
			}

		}
	}

	if len(characters) > maxCount {
		maxCount = len(characters)
	}

	return maxCount
}

func containsCharacter(characters []rune, character rune) bool {

	for _, c := range characters {
		if c == character {
			return true
		}

	}
	return false
}
