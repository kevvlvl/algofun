package strs

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

	for {

		letterProcessed := false

		for i := 0; i < len(strs); i++ {
			currentStr := strs[i]

			if len(currentStr) >= next {
				letterProcessed = true
				m[currentStr[0:next]]++
			}
		}

		if letterProcessed {
			next++
		} else {
			break
		}
	}

	prefix := ""
	wordCount := len(strs)
	for key, val := range m {

		if val == wordCount && len(key) > len(prefix) {
			prefix = key
		}
	}

	return prefix
}
