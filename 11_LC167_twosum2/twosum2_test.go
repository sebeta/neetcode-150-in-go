package twosum2

import (
	"slices"
	"testing"
)

// checkStatement fails the test if the input doesn't follow the statement of
// the problem (leetcode167.md), so a wrong test case is found right away:
//   - 2 <= numbers.length <= 3 * 10^4
//   - -1000 <= numbers[i] <= 1000
//   - numbers is sorted in non-decreasing order
//   - -1000 <= target <= 1000
//   - there is exactly one solution (a pair of different positions)
func checkStatement(t *testing.T, numbers []int, target int) {
	t.Helper()

	if len(numbers) < 2 || len(numbers) > 30_000 {
		t.Fatalf("test setup error: numbers.length is %d, want between 2 and 30000", len(numbers))
	}
	if target < -1000 || target > 1000 {
		t.Fatalf("test setup error: target is %d, want between -1000 and 1000", target)
	}
	for _, n := range numbers {
		if n < -1000 || n > 1000 {
			t.Fatalf("test setup error: %d is out of the range [-1000, 1000]", n)
		}
	}
	if !slices.IsSorted(numbers) {
		t.Fatalf("test setup error: numbers is not sorted in non-decreasing order")
	}

	// Count the pairs of different positions that add up to target.
	freq := make(map[int]int)
	for _, n := range numbers {
		freq[n]++
	}
	pairs := 0
	for a, fa := range freq {
		b := target - a
		fb, ok := freq[b]
		if !ok {
			continue
		}
		switch {
		case a < b:
			pairs += fa * fb
		case a == b:
			pairs += fa * (fa - 1) / 2
		}
	}
	if pairs != 1 {
		t.Fatalf("test setup error: there are %d solutions, want exactly 1", pairs)
	}
}

func TestTwoSum2(t *testing.T) {
	tests := []struct {
		name    string
		numbers []int
		target  int
		want    []int // 1-indexed: [index1, index2] with index1 < index2
	}{
		{"leetcode example 1", []int{2, 7, 11, 15}, 9, []int{1, 2}},
		{"leetcode example 2", []int{2, 3, 4}, 6, []int{1, 3}},
		{"leetcode example 3", []int{-1, 0}, -1, []int{1, 2}},

		// Minimum length.
		{"minimum length", []int{1, 2}, 3, []int{1, 2}},
		{"minimum length with equal values", []int{3, 3}, 6, []int{1, 2}},

		// Position of the answer.
		{"answer at the start", []int{1, 2, 5, 9}, 3, []int{1, 2}},
		{"answer at the end", []int{1, 2, 5, 9, 20}, 29, []int{4, 5}},
		{"answer in the middle", []int{1, 4, 6, 10, 20}, 16, []int{3, 4}},
		{"answer at the first and last positions", []int{-4, -1, 0, 2, 4}, 0, []int{1, 5}},
		{"answer with a number and the next one", []int{1, 2, 3, 4, 5}, 9, []int{4, 5}},

		// The same element can't be used twice.
		{"same value in two different positions", []int{1, 3, 3, 7}, 6, []int{2, 3}},
		{"half of the target appears only once", []int{1, 5, 8, 9}, 10, []int{1, 4}},
		{"duplicates that are not part of the answer", []int{1, 1, 2, 5}, 7, []int{3, 4}},

		// Negative numbers and zero.
		{"negative numbers", []int{-10, -5, -3, -1}, -8, []int{2, 3}},
		{"negative and positive numbers", []int{-4, -1, 0, 2, 4}, 0, []int{1, 5}},
		{"two zeros", []int{0, 0, 3, 4}, 0, []int{1, 2}},
		{"zero and another number", []int{0, 2, 5}, 5, []int{1, 3}},
		{"negative target with a zero", []int{-1000, 0, 500}, -1000, []int{1, 2}},

		// Values at the constraint bounds.
		{"minimum and maximum values", []int{-1000, 1000}, 0, []int{1, 2}},
		{"target at the maximum", []int{-500, 0, 1000}, 1000, []int{2, 3}},
		{"target at the minimum", []int{-600, -400, 1}, -1000, []int{1, 2}},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			checkStatement(t, tt.numbers, tt.target)

			// Copy the input to detect if the function modifies it.
			numbers := slices.Clone(tt.numbers)

			got := TwoSum2(numbers, tt.target)

			if !slices.Equal(numbers, tt.numbers) {
				t.Errorf("TwoSum2 modified its input: got %v, want %v", numbers, tt.numbers)
			}
			if !slices.Equal(got, tt.want) {
				t.Errorf("TwoSum2(%v, %d) = %v, want %v", tt.numbers, tt.target, got, tt.want)
			}
		})
	}
}

// TestTwoSum2LargeInput uses the maximum allowed length (3 * 10^4). Most of
// the elements are 1, and the only pair that adds up to the target is the
// one that is wanted.
func TestTwoSum2LargeInput(t *testing.T) {
	const n = 30_000

	filled := func(count, value int) []int {
		s := make([]int, count)
		for i := range s {
			s[i] = value
		}

		return s
	}

	tests := []struct {
		name    string
		numbers []int
		target  int
		want    []int
	}{
		{
			"answer at the end",
			append(filled(n-2, 1), 499, 501),
			1000,
			[]int{n - 1, n},
		},
		{
			"answer at the first and last positions",
			append(append([]int{-1000}, filled(n-2, 1)...), 1000),
			0,
			[]int{1, n},
		},
		{
			"answer in the middle",
			append(append(filled(15_000, 1), 400, 600), filled(14_998, 1000)...),
			1000,
			[]int{15_001, 15_002},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			checkStatement(t, tt.numbers, tt.target)

			if got := TwoSum2(tt.numbers, tt.target); !slices.Equal(got, tt.want) {
				t.Errorf("TwoSum2(large input, %d) = %v, want %v", tt.target, got, tt.want)
			}
		})
	}
}
