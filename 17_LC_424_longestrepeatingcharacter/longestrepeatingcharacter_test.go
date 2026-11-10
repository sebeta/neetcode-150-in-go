package longestrepeatingcharacter

import (
	"math/rand"
	"strings"
	"testing"
)

// implementations are all the solutions of the problem. Every test runs each
// one of them, and the benchmark compares them.
var implementations = []struct {
	name string
	fn   func(s string, k int) int
}{
	{"CharacterReplacement", CharacterReplacement},
	{"CharacterReplacement2", CharacterReplacement2},
}

// checkStatement fails the test if the input doesn't follow the constraints
// of the problem (leetcode424.md):
//   - 1 <= s.length <= 10^5
//   - s consists of only uppercase English letters
//   - 0 <= k <= s.length
func checkStatement(t *testing.T, s string, k int) {
	t.Helper()

	if len(s) < 1 || len(s) > 100_000 {
		t.Fatalf("test setup error: s.length is %d, want between 1 and 100000", len(s))
	}
	for i := 0; i < len(s); i++ {
		if s[i] < 'A' || s[i] > 'Z' {
			t.Fatalf("test setup error: character %q at position %d is not an uppercase English letter", s[i], i)
		}
	}
	if k < 0 || k > len(s) {
		t.Fatalf("test setup error: k is %d, want between 0 and %d", k, len(s))
	}
}

// bruteForce is an independent way of finding the answer: it checks every
// substring. A substring can become a run of one letter if the number of
// characters that are not its most frequent letter is at most k. It is too
// slow for big inputs, so it is only used to check the test cases and small
// random inputs.
func bruteForce(s string, k int) int {
	best := 0
	for i := 0; i < len(s); i++ {
		var count [26]int
		mostFrequent := 0
		for j := i; j < len(s); j++ {
			count[s[j]-'A']++
			mostFrequent = max(mostFrequent, count[s[j]-'A'])
			if (j-i+1)-mostFrequent <= k {
				best = max(best, j-i+1)
			}
		}
	}

	return best
}

// perLetter calculates the same answer in O(26 * n): for each letter, it
// finds the longest window with at most k characters that are not that
// letter, using two pointers. It is used to get the expected result of big
// inputs.
func perLetter(s string, k int) int {
	best := 0
	for c := byte('A'); c <= 'Z'; c++ {
		left, others := 0, 0
		for right := 0; right < len(s); right++ {
			if s[right] != c {
				others++
			}
			for others > k {
				if s[left] != c {
					others--
				}
				left++
			}
			best = max(best, right-left+1)
		}
	}

	return best
}

func TestCharacterReplacement(t *testing.T) {
	tests := []struct {
		name string
		s    string
		k    int
		want int
	}{
		{"leetcode example 1", "ABAB", 2, 4},
		{"leetcode example 2", "AABABBA", 1, 4},

		// A single character.
		{"single character, k is 0", "A", 0, 1},
		{"single character, k is 1", "A", 1, 1},

		// All the characters are the same.
		{"all characters are equal, k is 0", "AAAA", 0, 4},
		{"all characters are equal, k is not needed", "AAAA", 2, 4},

		// No replacements allowed.
		{"k is 0 and two different characters", "AB", 0, 1},
		{"k is 0, the longest run is in the middle", "ABBBBC", 0, 4},
		{"k is 0, several runs", "AABBBCC", 0, 3},
		{"k is 0, a repeated letter that is not together", "ABAA", 0, 2},
		{"k is 0 and all characters are different", "ABCD", 0, 1},

		// Few replacements.
		{"k is 1 and two different characters", "AB", 1, 2},
		{"one replacement joins two runs", "ABAA", 1, 4},
		{"one replacement at the end of a run", "AAAB", 1, 4},
		{"one replacement at the start of a run", "ABBB", 1, 4},
		{"one replacement is not enough to join two runs", "AAAB", 0, 3},
		{"k is 1 and all characters are different", "ABCD", 1, 2},
		{"k is 2 and all characters are different", "ABCD", 2, 3},

		// k is as big as the string.
		{"k is the length of the string", "ABCDE", 5, 5},
		{"k is the length minus one", "ABCD", 3, 4},
		{"k is the length of a string with two letters", "AB", 2, 2},

		// The letter that wins is not the first one.
		{"the letter to keep is in the middle", "BAAAB", 2, 5},
		{"two runs of different letters", "AAAAABBBBB", 3, 8},
		{"replacing on both sides of a run", "AABBBBAA", 2, 6},
		{"the best window changes letter", "AABABBA", 0, 2},

		// Repeated patterns.
		{"three letters repeated", "ABCABCABC", 2, 4},
		{"alternating letters", "ABABABAB", 1, 3},
		{"alternating letters with more replacements", "ABABABAB", 3, 7},

		// The whole alphabet.
		{"the whole alphabet, k is 0", "ABCDEFGHIJKLMNOPQRSTUVWXYZ", 0, 1},
		{"the whole alphabet, k is 25", "ABCDEFGHIJKLMNOPQRSTUVWXYZ", 25, 26},
		{"the whole alphabet, k is 5", "ABCDEFGHIJKLMNOPQRSTUVWXYZ", 5, 6},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			checkStatement(t, tt.s, tt.k)

			// Check that the expected result of the test case is right.
			if want := bruteForce(tt.s, tt.k); want != tt.want {
				t.Fatalf("test setup error: the expected result is %d, but the right answer for (%q, %d) is %d", tt.want, tt.s, tt.k, want)
			}

			for _, impl := range implementations {
				if got := impl.fn(tt.s, tt.k); got != tt.want {
					t.Errorf("%s(%q, %d) = %d, want %d", impl.name, tt.s, tt.k, got, tt.want)
				}
			}
		})
	}
}

// TestCharacterReplacementRandomSmallInputs compares the result with a brute
// force calculation for many small random strings, using a few letters so
// they repeat often. It uses a fixed seed so a failure can be reproduced.
func TestCharacterReplacementRandomSmallInputs(t *testing.T) {
	r := rand.New(rand.NewSource(1))

	for range 2000 {
		letters := 1 + r.Intn(4) // 1 to 4 different letters
		n := 1 + r.Intn(20)      // 1 to 20 characters
		var sb strings.Builder
		for range n {
			sb.WriteByte(byte('A' + r.Intn(letters)))
		}
		s := sb.String()
		k := r.Intn(n + 1) // 0 to n
		checkStatement(t, s, k)

		want := bruteForce(s, k)
		for _, impl := range implementations {
			if got := impl.fn(s, k); got != want {
				t.Fatalf("%s(%q, %d) = %d, want %d", impl.name, s, k, got, want)
			}
		}
	}
}

// TestCharacterReplacementLargeInput uses the maximum allowed length (10^5).
// A solution that checks every substring (O(n^2) or worse) takes far too
// long here.
func TestCharacterReplacementLargeInput(t *testing.T) {
	const n = 100_000
	r := rand.New(rand.NewSource(2))

	random := func(letters int) string {
		var sb strings.Builder
		for range n {
			sb.WriteByte(byte('A' + r.Intn(letters)))
		}

		return sb.String()
	}

	allEqual := strings.Repeat("A", n)
	alternating := strings.Repeat("AB", n/2)

	tests := []struct {
		name string
		s    string
		k    int
		want int // -1 means: use the O(26 * n) calculation
	}{
		{"all characters are equal", allEqual, 0, n},
		{"all characters are equal, k is the length", allEqual, n, n},
		{"alternating letters, k is 0", alternating, 0, 1},
		{"alternating letters, half of them can be replaced", alternating, n / 2, n},
		{"alternating letters, one less than needed", alternating, n/2 - 1, n - 1},
		{"random with 2 letters, k is 10", random(2), 10, -1},
		{"random with 2 letters, k is 1000", random(2), 1000, -1},
		{"random with 5 letters, k is 0", random(5), 0, -1},
		{"random with 5 letters, k is 100", random(5), 100, -1},
		{"random with 26 letters, k is 10", random(26), 10, -1},
		{"random with 26 letters, k is 5000", random(26), 5000, -1},
		{"random with 26 letters, k is the length", random(26), n, n},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			checkStatement(t, tt.s, tt.k)

			want := perLetter(tt.s, tt.k)
			if tt.want != -1 && tt.want != want {
				t.Fatalf("test setup error: the expected result is %d, but the right answer is %d", tt.want, want)
			}

			for _, impl := range implementations {
				if got := impl.fn(tt.s, tt.k); got != want {
					t.Errorf("%s(large input, %d) = %d, want %d", impl.name, tt.k, got, want)
				}
			}
		})
	}
}

// BenchmarkCharacterReplacement compares the implementations with inputs of
// the maximum allowed length. Run it with:
//
//	go test -bench . -benchmem ./17_LC_424_longestrepeatingcharacter/
func BenchmarkCharacterReplacement(b *testing.B) {
	const n = 100_000
	r := rand.New(rand.NewSource(3))

	random := func(letters int) string {
		var sb strings.Builder
		for range n {
			sb.WriteByte(byte('A' + r.Intn(letters)))
		}

		return sb.String()
	}

	inputs := []struct {
		name string
		s    string
		k    int
	}{
		{"all equal, k=0", strings.Repeat("A", n), 0},
		{"random 2 letters, k=1000", random(2), 1000},
		{"random 26 letters, k=10", random(26), 10},
		{"random 26 letters, k=5000", random(26), 5000},
	}

	for _, in := range inputs {
		for _, impl := range implementations {
			b.Run(impl.name+"/"+in.name, func(b *testing.B) {
				b.ReportAllocs()
				for range b.N {
					impl.fn(in.s, in.k)
				}
			})
		}
	}
}
