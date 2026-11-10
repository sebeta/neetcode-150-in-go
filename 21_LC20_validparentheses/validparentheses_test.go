package validparentheses

import (
	"math/rand"
	"strings"
	"testing"
)

// checkStatement fails the test if the input doesn't follow the constraints
// of the problem (leetcode20.md):
//   - 1 <= s.length <= 10^4
//   - s consists of parentheses only '()[]{}'
func checkStatement(t *testing.T, s string) {
	t.Helper()

	if len(s) < 1 || len(s) > 10_000 {
		t.Fatalf("test setup error: s.length is %d, want between 1 and 10000", len(s))
	}
	for i := 0; i < len(s); i++ {
		if !strings.ContainsRune("()[]{}", rune(s[i])) {
			t.Fatalf("test setup error: character %q at position %d is not one of ()[]{}", s[i], i)
		}
	}
}

// byReduction is an independent way of finding the answer: it keeps removing
// the pairs "()", "[]" and "{}" that are next to each other. The string is
// valid if nothing is left at the end. It is slower than a stack, but it is
// fast enough for the inputs of the tests.
func byReduction(s string) bool {
	for {
		reduced := strings.NewReplacer("()", "", "[]", "", "{}", "").Replace(s)
		if reduced == s {
			return s == ""
		}
		s = reduced
	}
}

// randomValid returns a random valid string of exactly n characters (n has
// to be even), built by choosing at each step to open or to close a bracket.
func randomValid(r *rand.Rand, n int) string {
	const open, closing = "([{", ")]}"

	var stack []int
	var sb strings.Builder
	for i := range n {
		remaining := n - i
		switch {
		case len(stack) == remaining: // the rest has to be closing brackets
			kind := stack[len(stack)-1]
			stack = stack[:len(stack)-1]
			sb.WriteByte(closing[kind])
		case len(stack) == 0 || r.Intn(2) == 0:
			kind := r.Intn(3)
			stack = append(stack, kind)
			sb.WriteByte(open[kind])
		default:
			kind := stack[len(stack)-1]
			stack = stack[:len(stack)-1]
			sb.WriteByte(closing[kind])
		}
	}

	return sb.String()
}

func TestIsValid(t *testing.T) {
	tests := []struct {
		name string
		s    string
		want bool
	}{
		{"leetcode example 1", "()", true},
		{"leetcode example 2", "()[]{}", true},
		{"leetcode example 3", "(]", false},
		{"leetcode example 4", "([])", true},
		{"leetcode example 5", "([)]", false},

		// A single bracket is never valid.
		{"single opening parenthesis", "(", false},
		{"single opening square bracket", "[", false},
		{"single opening curly bracket", "{", false},
		{"single closing parenthesis", ")", false},
		{"single closing square bracket", "]", false},
		{"single closing curly bracket", "}", false},

		// Each kind of bracket by itself.
		{"parentheses", "()", true},
		{"square brackets", "[]", true},
		{"curly brackets", "{}", true},
		{"nested parentheses", "((()))", true},
		{"consecutive parentheses", "()()", true},

		// Brackets in the wrong order.
		{"closing before opening", ")(", false},
		{"closing before opening, square brackets", "][", false},
		{"closing before opening, curly brackets", "}{", false},
		{"closing bracket first and then a valid pair", ")()", false},

		// Brackets that are not closed or not opened.
		{"missing closing bracket", "(()", false},
		{"missing opening bracket", "())", false},
		{"missing closing bracket at the end", "()(", false},
		{"missing closing bracket of a different kind", "([]", false},
		{"only opening brackets", "({[", false},
		{"only closing brackets", ")}]", false},
		{"only opening brackets of one kind", "((((", false},
		{"only closing brackets of one kind", "))))", false},

		// The kinds of brackets don't match.
		{"closing with the wrong kind", "(}", false},
		{"opening one kind and closing another", "{)", false},
		{"wrong kinds inside valid ones", "{[}]", false},
		{"correct kinds in the wrong order", "(}{)", false},
		{"interleaved brackets", "[(])", false},
		{"interleaved brackets with curly ones", "{[(])}", false},

		// Valid combinations.
		{"the three kinds side by side", "(){}[]", true},
		{"the three kinds nested", "{[()]}", true},
		{"the three kinds nested in another order", "([{}])", true},
		{"nested and consecutive", "{[]}()[{}]", true},
		{"deep nesting of different kinds", "[({})]", true},
		{"repeated sequence", "()[]{}()[]{}", true},
		{"nested missing the last closing bracket", "{[()]", false},

		// Longer strings.
		{"nested then consecutive", "((()))()[]{{}}", true},
		{"one extra closing bracket at the end", "()[]{}]", false},
		{"one extra opening bracket at the start", "[()[]{}", false},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			checkStatement(t, tt.s)

			// Check that the expected result of the test case is right.
			if want := byReduction(tt.s); want != tt.want {
				t.Fatalf("test setup error: the expected result is %v, but the right answer for %q is %v", tt.want, tt.s, want)
			}

			if got := IsValid(tt.s); got != tt.want {
				t.Errorf("IsValid(%q) = %v, want %v", tt.s, got, tt.want)
			}
		})
	}
}

// TestIsValidRandomSmallInputs compares the result with an independent
// calculation for many small random strings: some are completely random and
// some are valid strings with one bracket changed. It uses a fixed seed so a
// failure can be reproduced.
func TestIsValidRandomSmallInputs(t *testing.T) {
	r := rand.New(rand.NewSource(1))
	const brackets = "()[]{}"

	for range 2000 {
		var s string
		switch r.Intn(3) {
		case 0: // any string
			n := 1 + r.Intn(12)
			b := make([]byte, n)
			for i := range b {
				b[i] = brackets[r.Intn(len(brackets))]
			}
			s = string(b)
		case 1: // a valid string
			s = randomValid(r, 2*(1+r.Intn(6)))
		default: // a valid string with one bracket changed
			b := []byte(randomValid(r, 2*(1+r.Intn(6))))
			b[r.Intn(len(b))] = brackets[r.Intn(len(brackets))]
			s = string(b)
		}
		checkStatement(t, s)

		want := byReduction(s)
		if got := IsValid(s); got != want {
			t.Fatalf("IsValid(%q) = %v, want %v", s, got, want)
		}
	}
}

// TestIsValidLargeInput uses the maximum allowed length (10^4), with deep
// nesting and long sequences.
func TestIsValidLargeInput(t *testing.T) {
	const n = 10_000
	r := rand.New(rand.NewSource(2))

	// A valid string with the type of one bracket changed can't be valid.
	mismatched := []byte(randomValid(r, n))
	for i, c := range mismatched {
		if c == '(' {
			mismatched[i] = '['
			break
		}
	}

	random := make([]byte, n)
	for i := range random {
		random[i] = "()[]{}"[r.Intn(6)]
	}

	tests := []struct {
		name string
		s    string
		want bool
	}{
		{"deeply nested parentheses", strings.Repeat("(", n/2) + strings.Repeat(")", n/2), true},
		{"deeply nested with one closing bracket of another kind", strings.Repeat("(", n/2) + strings.Repeat(")", n/2-1) + "]", false},
		{"deeply nested with one closing bracket missing", strings.Repeat("(", n/2) + strings.Repeat(")", n/2-1), false},
		{"deeply nested with the kinds in the wrong order", strings.Repeat("([", n/4) + strings.Repeat(")]", n/4), false},
		{"deeply nested with the kinds in the right order", strings.Repeat("([", n/4) + strings.Repeat("])", n/4), true},
		{"long sequence of consecutive pairs", strings.Repeat("()", n/2), true},
		{"long sequence of the three kinds", strings.Repeat("(){}[]", n/6) + "()()", true},
		{"only opening brackets", strings.Repeat("(", n), false},
		{"only closing brackets", strings.Repeat("}", n), false},
		{"odd length", strings.Repeat("()", n/2-1) + "(", false},
		{"random valid string", randomValid(r, n), true},
		{"valid string with one bracket of another kind", string(mismatched), false},
		{"random string", string(random), byReduction(string(random))},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			checkStatement(t, tt.s)

			if got := IsValid(tt.s); got != tt.want {
				t.Errorf("IsValid(large input) = %v, want %v", got, tt.want)
			}
		})
	}
}
