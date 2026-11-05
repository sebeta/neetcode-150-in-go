package groupanagrams

import (
	"slices"
	"strings"
	"testing"
)

// normalize returns a copy of groups where each group is sorted and the
// groups themselves are sorted, so results can be compared regardless of
// the order in which the answer was returned.
func normalize(groups [][]string) [][]string {
	out := make([][]string, 0, len(groups))
	for _, g := range groups {
		c := slices.Clone(g)
		slices.Sort(c)
		out = append(out, c)
	}
	slices.SortFunc(out, func(a, b []string) int {
		return strings.Compare(strings.Join(a, ","), strings.Join(b, ","))
	})
	return out
}

func TestGroupAnagrams(t *testing.T) {
	tests := []struct {
		name string
		strs []string
		want [][]string
	}{
		{
			"leetcode example 1",
			[]string{"eat", "tea", "tan", "ate", "nat", "bat"},
			[][]string{{"bat"}, {"nat", "tan"}, {"ate", "eat", "tea"}},
		},
		{
			"leetcode example 2: single empty string",
			[]string{""},
			[][]string{{""}},
		},
		{
			"leetcode example 3: single letter",
			[]string{"a"},
			[][]string{{"a"}},
		},
		{
			"multiple empty strings are anagrams of each other",
			[]string{"", "", ""},
			[][]string{{"", "", ""}},
		},
		{
			"empty string mixed with other strings",
			[]string{"", "a", ""},
			[][]string{{"", ""}, {"a"}},
		},
		{
			"no anagrams at all",
			[]string{"abc", "def", "ghi"},
			[][]string{{"abc"}, {"def"}, {"ghi"}},
		},
		{
			"all strings are anagrams",
			[]string{"abc", "bca", "cab", "acb", "bac", "cba"},
			[][]string{{"abc", "bca", "cab", "acb", "bac", "cba"}},
		},
		{
			"duplicate strings stay in the same group",
			[]string{"abc", "abc", "bca"},
			[][]string{{"abc", "abc", "bca"}},
		},
		{
			"same letters, different counts are not anagrams",
			[]string{"aab", "abb", "aba", "bab"},
			[][]string{{"aab", "aba"}, {"abb", "bab"}},
		},
		{
			"different lengths are not anagrams",
			[]string{"a", "aa", "aaa", "a"},
			[][]string{{"a", "a"}, {"aa"}, {"aaa"}},
		},
		{
			"repeated letters",
			[]string{"aabbcc", "abcabc", "ccbbaa", "abc"},
			[][]string{{"aabbcc", "abcabc", "ccbbaa"}, {"abc"}},
		},
		{
			"single letters",
			[]string{"a", "b", "a", "c", "b"},
			[][]string{{"a", "a"}, {"b", "b"}, {"c"}},
		},
		{
			"input order does not matter",
			[]string{"nat", "bat", "tan", "tab"},
			[][]string{{"nat", "tan"}, {"bat", "tab"}},
		},
		{
			"longest allowed string (100 letters)",
			[]string{strings.Repeat("ab", 50), strings.Repeat("ba", 50), strings.Repeat("a", 100)},
			[][]string{{strings.Repeat("ab", 50), strings.Repeat("ba", 50)}, {strings.Repeat("a", 100)}},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			strs := slices.Clone(tt.strs)

			got := GroupAnagrams(strs)

			if !slices.Equal(strs, tt.strs) {
				t.Errorf("GroupAnagrams modified its input: got %q, want %q", strs, tt.strs)
			}

			gotN, wantN := normalize(got), normalize(tt.want)
			if len(gotN) != len(wantN) || !slices.EqualFunc(gotN, wantN, slices.Equal[[]string]) {
				t.Errorf("GroupAnagrams(%q) = %q, want %q (in any order)", tt.strs, got, tt.want)
			}
		})
	}
}
