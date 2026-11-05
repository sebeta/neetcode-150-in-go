package encodeanddecodestrings

import (
	"slices"
	"strings"
	"testing"
)

// allByteValues returns a string with every one of the 256 possible byte
// values, which covers "any possible characters out of 256 valid ASCII
// characters" from the constraints.
func allByteValues() string {
	b := make([]byte, 256)
	for i := range b {
		b[i] = byte(i)
	}

	return string(b)
}

// manyStrings returns n strings of the given length, each one different.
func manyStrings(n, length int) []string {
	strs := make([]string, n)
	for i := range strs {
		strs[i] = strings.Repeat(string(rune('a'+i%26)), length)
	}

	return strs
}

func TestEncodeAndDecode(t *testing.T) {
	tests := []struct {
		name string
		strs []string
	}{
		{"leetcode example 1", []string{"Hello", "World"}},
		{"leetcode example 2: single empty string", []string{""}},
		{"empty list", []string{}},
		{"nil list", nil},
		{"single string", []string{"a"}},
		{"single character strings", []string{"a", "b", "c"}},
		{"several empty strings", []string{"", "", ""}},
		{"empty strings mixed with others", []string{"", "a", "", "bc", ""}},
		{"empty string at the start", []string{"", "abc"}},
		{"empty string at the end", []string{"abc", ""}},
		{"duplicate strings", []string{"abc", "abc", "abc"}},
		{"order is preserved", []string{"z", "a", "m", "b"}},

		// The encoded format must not get confused by the characters it uses.
		{"string with the separator", []string{"#"}},
		{"strings with several separators", []string{"#", "##", "###"}},
		{"string that looks like an encoded string", []string{"3#abc"}},
		{"several strings that look like encoded strings", []string{"1#a", "2#bc", "0#"}},
		{"string with only digits", []string{"123", "4567", "0"}},
		{"digits followed by separator", []string{"12#", "5#", "#34"}},
		{"separator in the middle", []string{"ab#cd", "e#f#g"}},
		{"separator after a number-like string", []string{"10", "#", "2#"}},

		// Whitespace and control characters.
		{"spaces", []string{" ", "  ", "a b", " a "}},
		{"newlines and tabs", []string{"a\nb", "\t", "\r\n", "line1\nline2\n"}},
		{"null byte", []string{"a\x00b", "\x00", ""}},
		{"punctuation", []string{"!@$%^&*()", ",.;:'\"", "[]{}<>|\\/"}},

		// Lengths with different amounts of digits in the prefix.
		{"length 9 (one digit)", []string{strings.Repeat("a", 9)}},
		{"length 10 (two digits)", []string{strings.Repeat("a", 10)}},
		{"length 99 (two digits)", []string{strings.Repeat("a", 99)}},
		{"length 100 (three digits)", []string{strings.Repeat("a", 100)}},
		{"maximum length allowed (199)", []string{strings.Repeat("x", 199)}},
		{"long string full of separators", []string{strings.Repeat("#", 150)}},
		{"long string full of digits", []string{strings.Repeat("1", 150)}},
		{"long string that looks like encoded strings", []string{strings.Repeat("5#hello", 25)}},

		// Constraints: up to 99 strings, up to 199 characters each.
		{"maximum number of strings allowed (99)", manyStrings(99, 5)},
		{"maximum size allowed (99 strings of 199 characters)", manyStrings(99, 199)},
		{"every possible byte value", []string{allByteValues()}},
		{"every possible byte value split in several strings", []string{allByteValues()[:100], allByteValues()[100:200], allByteValues()[200:]}},

		// Follow up: any possible set of characters.
		{"non-ASCII characters", []string{"ñandú", "日本語", "😀", "héllo wörld"}},
		{"non-ASCII characters with separators", []string{"日#本", "#😀#"}},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			// Copy the input to detect if the function modifies it.
			input := slices.Clone(tt.strs)

			got := EncodeAndDecode(input)

			if !slices.Equal(input, tt.strs) {
				t.Errorf("EncodeAndDecode modified its input")
			}
			if !slices.Equal(got, tt.strs) {
				t.Errorf("EncodeAndDecode(%q) = %q, want the same list", tt.strs, got)
			}
		})
	}
}

// TestDecodeIsIndependentBetweenCalls checks that decode doesn't keep state
// between calls, as two different messages are decoded by the same function.
func TestDecodeIsIndependentBetweenCalls(t *testing.T) {
	first := []string{"Hello", "World"}
	second := []string{"", "#", "3#abc"}

	encodedFirst := encode(first)
	encodedSecond := encode(second)

	if got := decode(encodedFirst); !slices.Equal(got, first) {
		t.Errorf("decode(first) = %q, want %q", got, first)
	}
	if got := decode(encodedSecond); !slices.Equal(got, second) {
		t.Errorf("decode(second) = %q, want %q", got, second)
	}
	// Decoding the first message again must give the same result.
	if got := decode(encodedFirst); !slices.Equal(got, first) {
		t.Errorf("decode(first) the second time = %q, want %q", got, first)
	}
}

// TestEncodeIsDeterministic checks that the same input always produces the
// same encoded string, and that different inputs produce different ones.
func TestEncodeIsDeterministic(t *testing.T) {
	a := []string{"ab", "c"}
	b := []string{"a", "bc"}

	first, again := encode(a), encode(a)
	if first != again {
		t.Errorf("encode gave different results for the same input: %q and %q", first, again)
	}
	if other := encode(b); first == other {
		t.Errorf("encode(%q) and encode(%q) are the same (%q), so they could not be told apart when decoding", a, b, first)
	}
}
