package main

import (
	"fmt"
)

func main() {

	haystack := "mississippi"
	needle := "issipi"

	fmt.Println(strStr(haystack, needle))
}

func strStr(haystack string, needle string) int {

	if needle == "" {
		return 0
	}

	if len(needle) > len(haystack) {
		return -1
	}

	needles := []rune(needle)
	hays := []rune(haystack)

	for i, c := range hays {

		if c == needles[0] {

			possibleFound := i
			matchCount := len(needles)
			z := i
			for x := 0; x < len(needles); x++ {

				if z <= len(hays)-1 {

					if needles[x] == hays[z] {
						matchCount--
					} else {
						break
					}

					z++
				}

			}
			if matchCount == 0 {
				return possibleFound
			}

		}
	}

	return -1

}
