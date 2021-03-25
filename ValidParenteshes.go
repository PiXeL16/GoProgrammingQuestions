package main

import (
	"fmt"
	"strings"
)

//Given a string s containing just the characters '(', ')', '{', '}', '[' and ']', determine if the input string is valid.
//
//An input string is valid if:
//
//Open brackets must be closed by the same type of brackets.
//Open brackets must be closed in the correct order.
func main() {

	s := "{[]}"

	fmt.Print(isValid(s))

}

func isValid(s string) bool {

	stack := make([]rune, 0)
	m := map[rune]rune{
		')': '(',
		']': '[',
		'}': '{',
	}

	for _, c := range s {
		switch c {
		case '(', '{', '[':
			stack = append(stack, c)
		case ')', '}', ']':
			if len(stack) == 0 || stack[len(stack)-1] != m[c] {
				return false
			}
			stack = stack[:len(stack)-1]

		}
	}

	return len(stack) == 0
}

func isValidParentheses(s string) bool {

	squareBrackets := []string{"[", "]"}
	curlyBrackets := []string{"{", "}"}
	circularBrackets := []string{"(", ")"}

	if len(s) == 0 || len(s) == 1 {
		return false
	}

	var stack []string

	parentheses := strings.Split(s, "")

	for x := 0; x < len(parentheses); x++ {
		c := parentheses[x]
		if c == squareBrackets[0] {
			stack = append(stack, c)
		}
		if c == squareBrackets[1] {
			stackTop := stack[len(stack)-1]
			if stackTop == squareBrackets[0] {

				stack = stack[:len(stack)-1]
			}
		}
		if c == curlyBrackets[0] {
			if !checkOpositeBrackedValid(curlyBrackets[1], x, parentheses) {
				return false
			}
		}
		if c == circularBrackets[0] {
			if !checkOpositeBrackedValid(circularBrackets[1], x, parentheses) {
				return false
			}
		}
	}

	return true
}

func checkOpositeBrackedValid(opositeBracket string, currentPosition int, entireCharacters []string) bool {

	if currentPosition < len(entireCharacters)-1 {
		//Check next character
		if entireCharacters[currentPosition+1] == opositeBracket {
			return true
			// Special case for 0
		} else if currentPosition == 0 {
			if entireCharacters[len(entireCharacters)-1] == opositeBracket {
				return true
			}
			// Check opposite character
		} else if entireCharacters[len(entireCharacters)-(currentPosition)] == opositeBracket {
			return true
		} else {
			return false
		}
	}

	return false
}
