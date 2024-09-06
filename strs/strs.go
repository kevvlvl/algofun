package strs

import (
	"regexp"
	"strings"
)

var (
	openParenthesis  = "("[0]
	closeParenthesis = ")"[0]
	openSquare       = "["[0]
	closeSquare      = "]"[0]
	openBrace        = "{"[0]
	closeBrace       = "}"[0]
)

// IsValid returns true if the string contains valid parenthesis openings and closings.
// The only valid characters for 's' are: (){}[]
func IsValid(s string) bool {

	var d []byte

	for i := 0; i < len(s); i++ {

		cb := s[i]

		if len(d) > 0 && matchingBracket(cb, d[len(d)-1]) {
			//fmt.Printf("Bracket is matching %v Pop from stack. %v\n", string(cb), string(d[len(d)-1]))
			d = d[:len(d)-1]
		} else {
			//fmt.Printf("Add char to stack %v\n", string(cb))
			d = append(d, cb)
		}
	}

	return len(d) == 0
}

func matchingBracket(c byte, p byte) bool {
	return (c == closeParenthesis && p == openParenthesis) || (c == closeSquare && p == openSquare) || (c == closeBrace && p == openBrace)
}

// LongestCommonPrefix returns the longest matching prefix for the list of words. empty string if none found
func LongestCommonPrefix(strs []string) string {

	m := make(map[string]int)
	next := 1
	prefix := ""
	wordCount := len(strs)

	for {

		letterProcessed := false

		for i := 0; i < wordCount; i++ {
			currentStr := strs[i]

			if len(currentStr) >= next {
				letterProcessed = true
				p := currentStr[0:next]
				m[p]++

				if m[p] == wordCount && len(p) > len(prefix) {
					prefix = p
				}
			}
		}

		if letterProcessed {
			next++
		} else {
			break
		}
	}

	return prefix
}

// GetIndexForStr is an implementation of "strstr"
func GetIndexForStr(haystack string, needle string) int {

	if len(haystack) > 0 {
		return strings.Index(haystack, needle)
	}

	return -1
}

func LengthOfLastWord(s string) int {

	if strings.IndexAny(s, " ") < 0 {
		return len(s)
	}

	s = strings.Trim(s, " ")

	rx := regexp.MustCompile("\\s+")
	tokens := rx.Split(s, -1)

	return len(tokens[len(tokens)-1])
}
