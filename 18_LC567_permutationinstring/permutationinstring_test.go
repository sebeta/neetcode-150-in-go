package permutationinstring

import (
	"math/rand"
	"strings"
	"testing"
)

// checkStatement fails the test if the input doesn't follow the constraints
// of the problem (leetcode567.md):
//   - 1 <= s1.length, s2.length <= 10^4
//   - s1 and s2 consist of lowercase English letters
func checkStatement(t *testing.T, s1, s2 string) {
	t.Helper()

	for name, s := range map[string]string{"s1": s1, "s2": s2} {
		if len(s) < 1 || len(s) > 10_000 {
			t.Fatalf("test setup error: %s.length is %d, want between 1 and 10000", name, len(s))
		}
		for i := 0; i < len(s); i++ {
			if s[i] < 'a' || s[i] > 'z' {
				t.Fatalf("test setup error: character %q of %s at position %d is not a lowercase English letter", s[i], name, i)
			}
		}
	}
}

// letterCount returns how many times each letter appears in s.
func letterCount(s string) [26]int {
	var count [26]int
	for i := 0; i < len(s); i++ {
		count[s[i]-'a']++
	}

	return count
}

// bruteForce is an independent way of finding the answer: it counts the
// letters of every substring of s2 that has the length of s1, from scratch,
// and compares them with the letters of s1. Two strings are permutations of
// each other if they have the same letters the same number of times.
func bruteForce(s1, s2 string) bool {
	want := letterCount(s1)
	for start := 0; start+len(s1) <= len(s2); start++ {
		if letterCount(s2[start:start+len(s1)]) == want {
			return true
		}
	}

	return false
}

// shuffle returns a random permutation of s.
func shuffle(r *rand.Rand, s string) string {
	b := []byte(s)
	r.Shuffle(len(b), func(i, j int) { b[i], b[j] = b[j], b[i] })

	return string(b)
}

func randomString(r *rand.Rand, n, letters int) string {
	var sb strings.Builder
	for range n {
		sb.WriteByte(byte('a' + r.Intn(letters)))
	}

	return sb.String()
}

func TestCheckInclusion(t *testing.T) {
	tests := []struct {
		name   string
		s1, s2 string
		want   bool
	}{
		{"leetcode example 1", "ab", "eidbaooo", true},
		{"leetcode example 2", "ab", "eidboaoo", false},

		// One letter.
		{"one letter that is in s2", "a", "a", true},
		{"one letter that is not in s2", "a", "b", false},
		{"one letter at the end of s2", "a", "ba", true},
		{"one letter that never appears", "a", "bbb", false},

		// s1 is longer than s2, or both have the same length.
		{"s1 is longer than s2", "abc", "ab", false},
		{"s1 is longer than s2 by one letter", "ab", "a", false},
		{"s1 and s2 are the same string", "abc", "abc", true},
		{"s2 is the reverse of s1", "abc", "cba", true},
		{"s2 is another order of s1", "abc", "bca", true},
		{"same length, same letters, different counts", "aab", "abb", false},

		// Where the permutation is.
		{"permutation at the start of s2", "ab", "baxyz", true},
		{"permutation at the end of s2", "ab", "xyzba", true},
		{"permutation in the middle of s2", "ab", "xbay", true},
		{"s1 itself is in s2", "abc", "xxabcxx", true},

		// Repeated letters in s1: the counts have to match exactly.
		{"repeated letters, exact permutation", "aab", "baa", true},
		{"repeated letters, found in the middle", "aab", "abab", true},
		{"repeated letters, no window has the same counts", "aab", "abbab", false},
		{"s1 needs a letter twice and s2 has it once", "aa", "a", false},
		{"s1 needs a letter twice but they are not together", "aa", "aba", false},
		{"s1 needs a letter twice and s2 has them together", "aa", "baab", true},
		{"all letters are the same", "aaa", "aaaa", true},

		// The permutation has to be a contiguous substring of s2.
		{"the letters are in s2 but separated", "ab", "axb", false},
		{"the letters are in s2 but separated by several letters", "abc", "axbxc", false},
		{"all the letters appear, but never together", "abc", "aabbcc", false},
		{"the letters are together but one is wrong", "abcd", "abce", false},

		// The whole alphabet.
		{"the whole alphabet and its reverse", "abcdefghijklmnopqrstuvwxyz", "zyxwvutsrqponmlkjihgfedcba", true},
		{"the whole alphabet inside a longer string", "abcdefghijklmnopqrstuvwxyz", "aaaa" + "mnopqrstuvwxyzabcdefghijkl" + "aaaa", true},
		{"the whole alphabet with one letter repeated instead", "abcdefghijklmnopqrstuvwxyz", "abcdefghijklmnopqrstuvwxyy", false},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			checkStatement(t, tt.s1, tt.s2)

			// Check that the expected result of the test case is right.
			if want := bruteForce(tt.s1, tt.s2); want != tt.want {
				t.Fatalf("test setup error: the expected result is %v, but the right answer for (%q, %q) is %v", tt.want, tt.s1, tt.s2, want)
			}

			if got := CheckInclusion(tt.s1, tt.s2); got != tt.want {
				t.Errorf("CheckInclusion(%q, %q) = %v, want %v", tt.s1, tt.s2, got, tt.want)
			}
		})
	}
}

// TestCheckInclusionRandomSmallInputs compares the result with a brute force
// calculation for many small random strings, using a few letters so they
// repeat often. It uses a fixed seed so a failure can be reproduced.
func TestCheckInclusionRandomSmallInputs(t *testing.T) {
	r := rand.New(rand.NewSource(1))

	for range 3000 {
		letters := 1 + r.Intn(4) // 1 to 4 different letters
		s1 := randomString(r, 1+r.Intn(5), letters)
		s2 := randomString(r, 1+r.Intn(12), letters)
		checkStatement(t, s1, s2)

		want := bruteForce(s1, s2)
		if got := CheckInclusion(s1, s2); got != want {
			t.Fatalf("CheckInclusion(%q, %q) = %v, want %v", s1, s2, got, want)
		}
	}
}

// TestCheckInclusionLargeInput uses the maximum allowed length (10^4). A
// solution that sorts or counts the letters of every substring from scratch
// takes much longer here.
func TestCheckInclusionLargeInput(t *testing.T) {
	const n = 10_000
	r := rand.New(rand.NewSource(2))

	s1 := randomString(r, 5_000, 26)

	// s2 with a permutation of s1 planted at a random position.
	plant := func(permutation string) string {
		filler := randomString(r, n-len(permutation), 26)
		at := r.Intn(len(filler) + 1)

		return filler[:at] + permutation + filler[at:]
	}

	// A permutation of s1 where one letter was changed for another one.
	changed := []byte(shuffle(r, s1))
	if changed[0] == 'z' {
		changed[0] = 'a'
	} else {
		changed[0] = 'z'
	}

	almostPlanted := plant(string(changed))

	tests := []struct {
		name   string
		s1, s2 string
		want   bool
	}{
		{"all letters equal, s2 has enough of them", strings.Repeat("a", 5_000), strings.Repeat("a", n), true},
		{"all letters equal, one letter is wrong", strings.Repeat("a", 4_999) + "b", strings.Repeat("a", n), false},
		{"s1 and s2 have the maximum length and are permutations", strings.Repeat("ab", 5_000), shuffle(r, strings.Repeat("ab", 5_000)), true},
		{"s1 and s2 have the maximum length and one letter differs", s1 + s1, string(changed) + shuffle(r, s1), false},
		{"permutation at the end of a long s2", s1, strings.Repeat("z", n-len(s1)) + shuffle(r, s1), true},
		{"permutation at the start of a long s2", s1, shuffle(r, s1) + strings.Repeat("z", n-len(s1)), true},
		{"permutation planted in a random s2", s1, plant(shuffle(r, s1)), true},
		{"almost a permutation planted in a random s2", s1, almostPlanted, false},
		{"short s1 in a random s2", randomString(r, 50, 26), randomString(r, n, 26), false},
		{"short s1 with few letters in a random s2", randomString(r, 20, 3), randomString(r, n, 3), true},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			checkStatement(t, tt.s1, tt.s2)

			if want := bruteForce(tt.s1, tt.s2); want != tt.want {
				t.Fatalf("test setup error: the expected result is %v, but the right answer is %v", tt.want, want)
			}

			if got := CheckInclusion(tt.s1, tt.s2); got != tt.want {
				t.Errorf("CheckInclusion(large input) = %v, want %v", got, tt.want)
			}
		})
	}
}
