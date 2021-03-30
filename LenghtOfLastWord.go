//Given a string s consists of some words separated by spaces,
//return the length of the last word in the string.
//If the last word does not exist, return 0.

package main

import (
	"fmt"
	"strings"
)

func main() {

	words := "aaasdf as    "

	fmt.Println(lengthOfLastWord(words))

}

func lengthOfLastWord(s string) int {

	s = strings.TrimRight(s, " ")

	words := strings.Split(s, " ")

	if len(words) == 0 {
		return 0
	}

	lastWord := words[len(words)-1]

	return len(lastWord)
}
