// Package validparentheses solves the "Valid parentheses" problem.
package validparentheses

// chPairs maps each closing bracket to the opening bracket it closes.
var chPairs = map[byte]byte{
	')': '(',
	']': '[',
	'}': '{',
}

// IsValid reports whether s, which only has the characters '(', ')', '{',
// '}', '[' and ']', is valid: every open bracket is closed by a bracket of
// the same type, in the correct order, and every close bracket has an open
// bracket to close.
//
// It uses a stack with the open brackets that haven't been closed yet. Every
// open bracket is pushed. A close bracket is only valid if the stack isn't
// empty and the last open bracket (the top of the stack) is the one it closes;
// in that case, the open bracket is popped. At the end, the stack must be
// empty, otherwise some brackets were never closed. Counting how many brackets
// of each type have been opened and closed is not enough, because it doesn't
// keep their order: "([)]" and "([])" would look the same.
//
// For example, with "([)]": '(' and '[' are pushed, but when ')' arrives the
// top of the stack is '[', not '(', so it returns false.
//
// Time complexity: O(n), where n is the length of s, since each character is
// visited once.
// Space complexity: O(n), since the stack can hold all the characters (for
// example, with "((((").
//
// NOTE: This is my solution, it is the most optimal.
func IsValid(s string) bool {
	var stack []byte
	for i := range s {
		switch s[i] {
		case '(', '[', '{':
			stack = append(stack, s[i])
		case ')', ']', '}':
			if len(stack) == 0 || stack[len(stack)-1] != chPairs[s[i]] {
				return false
			}
			stack = stack[:len(stack)-1]
		}
	}

	return len(stack) == 0
}
