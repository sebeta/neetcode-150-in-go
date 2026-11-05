// Package encodeanddecodestrings solves the "Encode and decode strings" problem.
package encodeanddecodestrings

import (
	"strconv"
	"strings"
)

// EncodeAndDecode encodes strs into a single string and decodes it back,
// simulating the trip from Machine 1 (encode) to Machine 2 (decode). The
// result has the same strings as the input, in the same order.
//
// Time complexity: O(n), where n is the total number of characters in strs.
// Space complexity: O(n), for the encoded string and the decoded list.
//
// NOTE: I didn't like the solution from the video, so I decided to use the
// hints proposed in neetcode.io for this exercise.
func EncodeAndDecode(strs []string) []string {
	return decode(encode(strs))
}

// encode joins strs into a single string, writing each one as
// "<length>#<content>" (for example, ["Hello", "World"] becomes
// "5#Hello5#World"). Since the length comes first, the content can contain
// any character, including "#" and digits.
//
// Time complexity: O(n), where n is the total number of characters in strs.
// Space complexity: O(n), for the encoded string.
func encode(strs []string) string {
	// Using strings.Builder is more optimal and idiomatic in Go than
	// filling a string with `+=`, which copies the whole string on every
	// iteration.
	var sb strings.Builder
	for _, str := range strs {
		sb.WriteString(strconv.Itoa(len(str)))
		sb.WriteByte('#')
		sb.WriteString(str)
	}

	return sb.String()
}

// decode splits a string created by encode back into the original strings.
// It reads the digits until the first "#" to get the length, takes that many
// characters as the content, and repeats from the character after it. The
// content is never inspected, so it can contain "#" without confusing the
// decoding. It assumes encoded is valid.
//
// Time complexity: O(n), where n is the length of encoded.
// Space complexity: O(n), for the decoded list.
func decode(encoded string) []string {
	var strs []string
	var length string
	for i := 0; i < len(encoded); {
		ch := encoded[i]
		if ch == '#' {
			num, _ := strconv.Atoi(length)
			str := encoded[i+1 : i+1+num]
			strs = append(strs, str)
			i += num + 1
			length = ""
			continue
		}
		length += string(ch)
		i++
	}

	return strs
}
