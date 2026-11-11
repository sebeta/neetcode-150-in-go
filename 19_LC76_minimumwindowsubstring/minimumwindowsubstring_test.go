package minimumwindowsubstring

import (
	"math/rand"
	"strings"
	"testing"
)

// checkStatement fails the test if the input doesn't follow the constraints
// of the problem (leetcode76.md):
//   - 1 <= s.length, t.length <= 10^5
//   - s and t consist of uppercase and lowercase English letters
//   - the answer is unique (there is no other window of the same length with
//     different letters)
func checkStatement(t *testing.T, s, tt string) {
	t.Helper()

	for name, str := range map[string]string{"s": s, "t": tt} {
		if len(str) < 1 || len(str) > 100_000 {
			t.Fatalf("test setup error: %s.length is %d, want between 1 and 100000", name, len(str))
		}
		for i := 0; i < len(str); i++ {
			c := str[i]
			if !(c >= 'a' && c <= 'z' || c >= 'A' && c <= 'Z') {
				t.Fatalf("test setup error: character %q of %s at position %d is not an English letter", c, name, i)
			}
		}
	}

	if _, distinct := oracle(s, tt); distinct > 1 {
		t.Fatalf("test setup error: the answer for (%q, %q) is not unique: there are %d different minimum windows", s, tt, distinct)
	}
}

// startsOfValidWindows returns the starting positions of the substrings of s
// with length size that include every character of tt (duplicates included).
func startsOfValidWindows(s, tt string, size int) []int {
	var need, have [128]int
	for i := 0; i < len(tt); i++ {
		need[tt[i]]++
	}

	missing := len(tt) // characters of tt that the window still needs
	var starts []int
	for end := 0; end < len(s); end++ {
		if have[s[end]] < need[s[end]] {
			missing--
		}
		have[s[end]]++

		if end >= size {
			leaving := s[end-size]
			have[leaving]--
			if have[leaving] < need[leaving] {
				missing++
			}
		}

		if end >= size-1 && missing == 0 {
			starts = append(starts, end-size+1)
		}
	}

	return starts
}

// oracle is an independent way of finding the answer: a window that includes
// all the characters of tt can always be made longer, so the minimum length is
// found with a binary search, checking windows of a fixed size. It returns the
// minimum window (or "" if there is none) and how many different minimum
// windows there are.
func oracle(s, tt string) (best string, distinct int) {
	if len(tt) > len(s) || len(startsOfValidWindows(s, tt, len(s))) == 0 {
		return "", 0
	}

	low, high := len(tt), len(s)
	for low < high {
		mid := (low + high) / 2
		if len(startsOfValidWindows(s, tt, mid)) > 0 {
			high = mid
		} else {
			low = mid + 1
		}
	}

	windows := make(map[string]struct{})
	for _, start := range startsOfValidWindows(s, tt, low) {
		windows[s[start:start+low]] = struct{}{}
	}
	for w := range windows {
		best = w
	}

	return best, len(windows)
}

// bruteForce is a second independent calculation: it tries every substring
// of s. It is too slow for big inputs, so it is only used for small ones.
func bruteForce(s, tt string) string {
	var need [128]int
	for i := 0; i < len(tt); i++ {
		need[tt[i]]++
	}

	best := ""
	for i := 0; i < len(s); i++ {
		var have [128]int
		for j := i; j < len(s); j++ {
			have[s[j]]++

			covered := true
			for c := range need {
				if have[c] < need[c] {
					covered = false
					break
				}
			}
			if covered {
				if best == "" || j-i+1 < len(best) {
					best = s[i : j+1]
				}
				break
			}
		}
	}

	return best
}

func TestMinWindow(t *testing.T) {
	tests := []struct {
		name string
		s, t string
		want string
	}{
		{"leetcode example 1", "ADOBECODEBANC", "ABC", "BANC"},
		{"leetcode example 2", "a", "a", "a"},
		{"leetcode example 3", "a", "aa", ""},

		// There is no window.
		{"t is longer than s", "ab", "abc", ""},
		{"t has a letter that s doesn't have", "abc", "d", ""},
		{"s has the letters of t, but not enough times", "abc", "aab", ""},
		{"the letters are only in the other case", "a", "A", ""},
		{"all the letters but one are in s", "abcdef", "abcdeg", ""},

		// The window is the whole string, or part of it.
		{"s and t are the same string", "abc", "abc", "abc"},
		{"s is a permutation of t", "cba", "abc", "cba"},
		{"t is one letter", "abc", "b", "b"},
		{"t is one letter and s has it several times", "xbxbx", "b", "b"},

		// Position of the window.
		{"window at the start", "abcxxx", "abc", "abc"},
		{"window at the end", "xxxabc", "abc", "abc"},
		{"window in the middle", "xxabcxx", "abc", "abc"},
		{"window with other letters inside", "aXbXc", "abc", "aXbXc"},

		// Duplicates in t: the counts have to be included.
		{"t repeats a letter", "aaab", "aab", "aab"},
		{"t repeats a letter and s has it at the start", "bbaa", "aba", "baa"},
		{"t is made of one repeated letter", "aXaaa", "aaa", "aaa"},
		{"t repeats several letters", "xaabbcxabbc", "aabbc", "aabbc"},

		// The window has to shrink from the left.
		{"a longer window first and a shorter one later", "aaabbbcccxabc", "abc", "abc"},
		{"a long window that contains a shorter one", "abxxxxxbac", "bac", "bac"},
		{"the first valid window is not the shortest", "bbaacxbac", "abc", "bac"},

		// Uppercase and lowercase are different letters.
		{"uppercase and lowercase letters", "aAbB", "AB", "AbB"},
		{"t has uppercase and lowercase letters", "xaAx", "Aa", "aA"},
		{"only the uppercase one is needed", "Aa", "A", "A"},

		// The whole alphabet.
		{
			"the whole alphabet and its reverse",
			"abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ",
			"ZYXWVUTSRQPONMLKJIHGFEDCBAzyxwvutsrqponmlkjihgfedcba",
			"abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ",
		},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			checkStatement(t, tc.s, tc.t)

			// Check that the expected result of the test case is right.
			if want, _ := oracle(tc.s, tc.t); want != tc.want {
				t.Fatalf("test setup error: the expected result is %q, but the right answer for (%q, %q) is %q", tc.want, tc.s, tc.t, want)
			}
			if want := bruteForce(tc.s, tc.t); want != tc.want {
				t.Fatalf("test setup error: the expected result is %q, but the brute force for (%q, %q) is %q", tc.want, tc.s, tc.t, want)
			}

			if got := MinWindow(tc.s, tc.t); got != tc.want {
				t.Errorf("MinWindow(%q, %q) = %q, want %q", tc.s, tc.t, got, tc.want)
			}
		})
	}
}

// TestMinWindowRandomSmallInputs compares the result with the brute force for
// many small random strings, using a few letters so they repeat often. The
// statement says the answer is unique, so the inputs where several different
// windows have the minimum length are skipped. It uses a fixed seed so a
// failure can be reproduced.
func TestMinWindowRandomSmallInputs(t *testing.T) {
	r := rand.New(rand.NewSource(1))
	random := func(n, letters int) string {
		var sb strings.Builder
		for range n {
			c := byte('a' + r.Intn(letters))
			if r.Intn(2) == 0 {
				c = c - 'a' + 'A' // some uppercase letters too
			}
			sb.WriteByte(c)
		}

		return sb.String()
	}

	tested := 0
	for range 20_000 {
		letters := 1 + r.Intn(3)
		s := random(1+r.Intn(14), letters)
		tt := random(1+r.Intn(4), letters)

		if _, distinct := oracle(s, tt); distinct > 1 {
			continue
		}
		tested++

		want := bruteForce(s, tt)
		if got := MinWindow(s, tt); got != want {
			t.Fatalf("MinWindow(%q, %q) = %q, want %q", s, tt, got, want)
		}
	}

	if tested < 1_000 {
		t.Fatalf("test setup error: only %d random inputs had a unique answer", tested)
	}
}

// TestMinWindowLargeInput uses the maximum allowed length (10^5). A solution
// that checks every window from scratch takes far too long here.
func TestMinWindowLargeInput(t *testing.T) {
	const n = 100_000
	r := rand.New(rand.NewSource(2))

	// randomT is a random string of the letters a to e.
	randomT := func(size int) string {
		var sb strings.Builder
		for range size {
			sb.WriteByte(byte('a' + r.Intn(5)))
		}

		return sb.String()
	}

	tBlock := randomT(5_000)
	partialBlock := tBlock[:len(tBlock)-1] // one letter is missing

	randomOther := func(size int) string {
		var sb strings.Builder
		for range size {
			sb.WriteByte(byte('f' + r.Intn(20))) // letters that are not in t
		}

		return sb.String()
	}

	tests := []struct {
		name string
		s, t string
		want string
	}{
		{
			"t is the only block of letters of t inside s",
			strings.Repeat("z", 50_000) + tBlock + strings.Repeat("z", n-50_000-len(tBlock)),
			tBlock,
			tBlock,
		},
		{
			"a block that is not enough comes before the right one",
			strings.Repeat("z", 30_000) + partialBlock + strings.Repeat("z", 30_000) + tBlock + strings.Repeat("z", 20_000),
			tBlock,
			tBlock,
		},
		{
			"random letters that are not in t around the block",
			randomOther(40_000) + tBlock + randomOther(n-40_000-len(tBlock)),
			tBlock,
			tBlock,
		},
		{"s and t are made of the same letter", strings.Repeat("a", n), strings.Repeat("a", 50_000), strings.Repeat("a", 50_000)},
		{"s and t are the same string of the maximum length", strings.Repeat("a", n), strings.Repeat("a", n), strings.Repeat("a", n)},
		{"s doesn't have the letter of t", strings.Repeat("a", n), "b", ""},
		{"t needs one more letter than s has", strings.Repeat("a", n-1) + "b", strings.Repeat("a", n), ""},
		{"the window is where two blocks meet", strings.Repeat("a", 50_000) + strings.Repeat("b", 50_000), "ab", "ab"},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			checkStatement(t, tc.s, tc.t)

			if want, _ := oracle(tc.s, tc.t); want != tc.want {
				t.Fatalf("test setup error: the expected result has %d letters, but the right answer has %d", len(tc.want), len(want))
			}

			if got := MinWindow(tc.s, tc.t); got != tc.want {
				t.Errorf("MinWindow(large input) has %d letters, want %d", len(got), len(tc.want))
			}
		})
	}
}
