package longestconsecutive

import (
	"math/rand"
	"testing"
)

func TestLongestConsecutive(t *testing.T) {
	tests := []struct {
		name string
		nums []int
		want int
	}{
		{"leetcode example 1", []int{100, 4, 200, 1, 3, 2}, 4},
		{"leetcode example 2", []int{0, 3, 7, 2, 5, 8, 4, 6, 0, 1}, 9},
		{"leetcode example 3", []int{1, 0, 1, 2}, 3},

		// Small inputs.
		{"empty slice", []int{}, 0},
		{"nil slice", nil, 0},
		{"single element", []int{5}, 1},
		{"all elements are the same", []int{7, 7, 7}, 1},
		{"two consecutive elements", []int{1, 2}, 2},
		{"two consecutive elements in reverse order", []int{2, 1}, 2},
		{"two elements that are not consecutive", []int{1, 3}, 1},
		{"no consecutive elements", []int{10, 20, 30, 40}, 1},

		// Order of the input.
		{"sorted ascending", []int{1, 2, 3, 4, 5}, 5},
		{"sorted descending", []int{5, 4, 3, 2, 1}, 5},
		{"unsorted sequence", []int{4, 2, 1, 3}, 4},
		{"element that joins two sequences", []int{1, 2, 4, 5, 3}, 5},

		// Duplicates.
		{"duplicates inside a sequence", []int{1, 2, 2, 3, 3, 4}, 4},
		{"many duplicates of the same element", []int{1, 1, 1, 1, 2, 2, 2}, 2},

		// Several sequences.
		{"two sequences of the same length", []int{1, 2, 3, 10, 11, 12}, 3},
		{"longest sequence is the second one", []int{1, 2, 10, 11, 12, 13}, 4},
		{"longest sequence is the first one", []int{1, 2, 3, 4, 10, 11}, 4},
		{"gap of one in the middle", []int{1, 2, 3, 5, 6, 7, 8}, 4},
		{"longest sequence is not at the extremes", []int{50, 1, 2, 3, 4, 5, 100}, 5},

		// Negative numbers and zero.
		{"negative numbers", []int{-3, -2, -1}, 3},
		{"sequence that crosses zero", []int{-2, -1, 0, 1, 2}, 5},
		{"negatives and positives with a gap", []int{-1, 0, 2, 3}, 2},
		{"negative numbers in any order", []int{-1, -3, -2, -10}, 3},

		// Values at the constraint bounds.
		{"minimum and maximum values", []int{-1000000000, 1000000000}, 1},
		{"sequence at the maximum value", []int{1000000000, 999999999, 999999998}, 3},
		{"sequence at the minimum value", []int{-1000000000, -999999999}, 2},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := LongestConsecutive(tt.nums); got != tt.want {
				t.Errorf("LongestConsecutive(%v) = %d, want %d", tt.nums, got, tt.want)
			}
		})
	}
}

// TestLongestConsecutiveLargeInput uses the maximum allowed length (10^5).
// A solution that is not O(n) (for example, one that counts up from every
// element) takes a very long time here.
func TestLongestConsecutiveLargeInput(t *testing.T) {
	const n = 100_000
	r := rand.New(rand.NewSource(1))

	t.Run("one long sequence in random order", func(t *testing.T) {
		nums := make([]int, n)
		for i := range nums {
			nums[i] = i
		}
		r.Shuffle(len(nums), func(i, j int) { nums[i], nums[j] = nums[j], nums[i] })

		if got := LongestConsecutive(nums); got != n {
			t.Errorf("LongestConsecutive(large input) = %d, want %d", got, n)
		}
	})

	t.Run("one long sequence in descending order", func(t *testing.T) {
		nums := make([]int, n)
		for i := range nums {
			nums[i] = n - i
		}

		if got := LongestConsecutive(nums); got != n {
			t.Errorf("LongestConsecutive(large input) = %d, want %d", got, n)
		}
	})

	t.Run("no consecutive elements", func(t *testing.T) {
		nums := make([]int, n)
		for i := range nums {
			nums[i] = i * 2 // 0, 2, 4, ... every gap is 2
		}
		r.Shuffle(len(nums), func(i, j int) { nums[i], nums[j] = nums[j], nums[i] })

		if got := LongestConsecutive(nums); got != 1 {
			t.Errorf("LongestConsecutive(large input) = %d, want 1", got)
		}
	})

	t.Run("many duplicates", func(t *testing.T) {
		nums := make([]int, n)
		for i := range nums {
			nums[i] = i % 1000 // 0..999 repeated 100 times
		}
		r.Shuffle(len(nums), func(i, j int) { nums[i], nums[j] = nums[j], nums[i] })

		if got := LongestConsecutive(nums); got != 1000 {
			t.Errorf("LongestConsecutive(large input) = %d, want 1000", got)
		}
	})
}
