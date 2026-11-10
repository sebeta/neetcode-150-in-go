package longestsubstring

import (
	"math/rand"
	"strings"
	"testing"
)

// printable holds every character the input can have: English letters,
// digits, symbols and spaces (the printable ASCII characters, 32 to 126).
const printable = " !\"#$%&'()*+,-./0123456789:;<=>?@ABCDEFGHIJKLMNOPQRSTUVWXYZ[\\]^_`abcdefghijklmnopqrstuvwxyz{|}~"

// checkStatement fails the test if the input doesn't follow the constraints
// of the problem (leetcode.md):
//   - 0 <= s.length <= 10^5
//   - s consists of English letters, digits, symbols and spaces
func checkStatement(t *testing.T, s string) {
	t.Helper()

	if len(s) > 100_000 {
		t.Fatalf("test setup error: s.length is %d, want at most 100000", len(s))
	}
	for i := 0; i < len(s); i++ {
		if s[i] < 32 || s[i] > 126 {
			t.Fatalf("test setup error: character %q at position %d is not a printable ASCII character", s[i], i)
		}
	}
}

// bruteForce is an independent way of finding the answer: from each starting
// position it extends the substring until a character repeats. A substring
// can't have more than 95 characters (the number of different ones), so it
// is fast enough even for big inputs.
func bruteForce(s string) int {
	best := 0
	for start := 0; start < len(s); start++ {
		var seen [128]bool
		length := 0
		for end := start; end < len(s) && !seen[s[end]]; end++ {
			seen[s[end]] = true
			length++
		}
		best = max(best, length)
	}

	return best
}

func TestLengthOfLongestSubstring(t *testing.T) {
	tests := []struct {
		name string
		s    string
		want int
	}{
		{"leetcode example 1", "abcabcbb", 3},
		{"leetcode example 2", "bbbbb", 1},
		{"leetcode example 3", "pwwkew", 3},

		// Very short strings.
		{"empty string", "", 0},
		{"single character", "a", 1},
		{"single space", " ", 1},
		{"two different characters", "ab", 2},
		{"two equal characters", "aa", 1},
		{"two spaces", "  ", 1},

		// All characters are the same or all are different.
		{"all characters are equal", "aaaa", 1},
		{"all characters are different", "abcdef", 6},
		{"the whole alphabet", "abcdefghijklmnopqrstuvwxyz", 26},

		// Where the repeated character is.
		{"the first character repeats at the end", "abca", 3},
		{"the last character repeats an earlier one", "abcb", 3},
		{"the repeated character is next to the first one", "aab", 2},
		{"the repeated characters are at the end", "abcdee", 5},

		// The start of the substring must never move backwards.
		{"a repeated character before the start of the window", "abba", 2},
		{"a repeated character that was already left behind", "tmmzuxt", 5},
		{"the best substring starts after a repeat", "dvdf", 3},
		{"a longer string where the window moves several times", "anviaj", 5},

		// Substring, not subsequence.
		{"a subsequence would be longer", "pwwkew", 3},
		{"the best substring is in the middle", "aabcdea", 5},

		// Case, digits, symbols and spaces.
		{"uppercase and lowercase are different characters", "aA", 2},
		{"mixed case", "abcABC", 6},
		{"digits", "12321", 3},
		{"symbols", "!@#!@#", 3},
		{"spaces are characters too", "a b a", 3},
		{"every printable character once", printable, 95},
		{"every printable character twice", printable + printable, 95},
		{"a long substring and then a repeat", "abcdefghijklmnopqrstuvwxyza", 26},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			checkStatement(t, tt.s)

			// Check that the expected result of the test case is right.
			if want := bruteForce(tt.s); want != tt.want {
				t.Fatalf("test setup error: the expected result is %d, but the right answer for %q is %d", tt.want, tt.s, want)
			}

			if got := LengthOfLongestSubstring(tt.s); got != tt.want {
				t.Errorf("LengthOfLongestSubstring(%q) = %d, want %d", tt.s, got, tt.want)
			}
		})
	}
}

// TestLengthOfLongestSubstringRandomSmallInputs compares the result with a
// brute force calculation for many small random strings, using a small set
// of characters so they repeat often. It uses a fixed seed so a failure can
// be reproduced.
func TestLengthOfLongestSubstringRandomSmallInputs(t *testing.T) {
	r := rand.New(rand.NewSource(1))

	for range 1000 {
		alphabet := printable[:2+r.Intn(8)] // 2 to 9 different characters
		n := r.Intn(21)                     // 0 to 20 characters
		var sb strings.Builder
		for range n {
			sb.WriteByte(alphabet[r.Intn(len(alphabet))])
		}
		s := sb.String()
		checkStatement(t, s)

		want := bruteForce(s)
		if got := LengthOfLongestSubstring(s); got != want {
			t.Fatalf("LengthOfLongestSubstring(%q) = %d, want %d", s, got, want)
		}
	}
}

// TestLengthOfLongestSubstringLargeInput uses the maximum allowed length
// (10^5). A solution that looks at every possible substring (O(n^3)) takes
// far too long here.
func TestLengthOfLongestSubstringLargeInput(t *testing.T) {
	const n = 100_000
	r := rand.New(rand.NewSource(2))

	random := func(alphabet string) string {
		var sb strings.Builder
		for range n {
			sb.WriteByte(alphabet[r.Intn(len(alphabet))])
		}

		return sb.String()
	}

	tests := []struct {
		name string
		s    string
		want int // -1 means: use the brute force calculation
	}{
		{"all characters are equal", strings.Repeat("a", n), 1},
		{"a short pattern repeated", strings.Repeat("abc", n/3), 3},
		{"every printable character repeated", strings.Repeat(printable, n/len(printable)+1)[:n], 95},
		{"random characters from a small set", random("abcde"), -1},
		{"random lowercase letters", random("abcdefghijklmnopqrstuvwxyz"), -1},
		{"random printable characters", random(printable), -1},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			checkStatement(t, tt.s)

			want := bruteForce(tt.s)
			if tt.want != -1 && tt.want != want {
				t.Fatalf("test setup error: the expected result is %d, but the right answer is %d", tt.want, want)
			}

			if got := LengthOfLongestSubstring(tt.s); got != want {
				t.Errorf("LengthOfLongestSubstring(large input) = %d, want %d", got, want)
			}
		})
	}
}
