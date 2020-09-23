package main

import (
	"fmt"
	"sort"
	"strings"
)
////:## Are words anagram
////:Create a function that returns true or false if two words are `anagrams`
////:An anagram is a type of word play, the result of rearranging the letters of a word or phrase to produce a new word or phrase, using all the original letters exactly once; for example, the word anagram can be rearranged into nag-a-ram.
//
//
///**
// Are word anagrams
//
// - parameter firstWord:  first work to check
// - parameter secondWord: second word to check
//
// - returns: if anagram or not
// */
func main() {

fmt.Print(areWordsAnagram("test", "tset"))

}

func areWordsAnagram(firstWord string, secondWord string) bool {

	returnValue := false

	if len(firstWord) == len(secondWord) {
		firstWordCharacters :=  strings.Split(firstWord, "")
		secondWordCharacters := strings.Split(secondWord, "")
		sort.Strings(firstWordCharacters)
		sort.Strings(secondWordCharacters)

		returnValue = strings.Join(firstWordCharacters,"") == strings.Join(secondWordCharacters,"")
	}

	return returnValue
}