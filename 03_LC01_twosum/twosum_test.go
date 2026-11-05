package twosum

import (
	"slices"
	"testing"
)

func TestTwoSum(t *testing.T) {
	tests := []struct {
		name   string
		nums   []int
		target int
		want   []int // expected indices, sorted ascending (the answer can be returned in any order)
	}{
		{"leetcode example 1", []int{2, 7, 11, 15}, 9, []int{0, 1}},
		{"leetcode example 2", []int{3, 2, 4}, 6, []int{1, 2}},
		{"leetcode example 3", []int{3, 3}, 6, []int{0, 1}},
		{"answer at the end", []int{1, 5, 9, 14, 20}, 34, []int{3, 4}},
		{"answer at the start", []int{4, 6, 10, 20, 30}, 10, []int{0, 1}},
		{"answer not adjacent", []int{1, 8, 3, 5, 2}, 7, []int{3, 4}},
		{"negative numbers", []int{-1, -2, -3, -4, -5}, -8, []int{2, 4}},
		{"negative and positive", []int{-3, 4, 3, 90}, 0, []int{0, 2}},
		{"zeros", []int{0, 4, 3, 0}, 0, []int{0, 3}},
		{"zero as one of the numbers", []int{5, 0, 8}, 5, []int{0, 1}},
		{"duplicate values form the answer", []int{1, 3, 3, 7}, 6, []int{1, 2}},
		{"must not reuse same element", []int{3, 2, 4}, 6, []int{1, 2}}, // using 3 twice would be wrong
		{"unsorted input", []int{15, 11, 7, 2}, 9, []int{2, 3}},
		{"large values", []int{1000000000, 2, -1000000000, 7}, 0, []int{0, 2}},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			// Copy the input so the function can't hide a mutation from the checks below.
			nums := slices.Clone(tt.nums)

			got := TwoSum(nums, tt.target)

			if len(got) != 2 {
				t.Fatalf("TwoSum(%v, %d) = %v, want exactly 2 indices", tt.nums, tt.target, got)
			}
			i, j := got[0], got[1]
			if i < 0 || i >= len(nums) || j < 0 || j >= len(nums) {
				t.Fatalf("TwoSum(%v, %d) = %v, index out of range", tt.nums, tt.target, got)
			}
			if i == j {
				t.Fatalf("TwoSum(%v, %d) = %v, the same element was used twice", tt.nums, tt.target, got)
			}
			if nums[i]+nums[j] != tt.target {
				t.Fatalf("TwoSum(%v, %d) = %v, but %d + %d = %d", tt.nums, tt.target, got, nums[i], nums[j], nums[i]+nums[j])
			}

			sorted := slices.Clone(got)
			slices.Sort(sorted)
			if !slices.Equal(sorted, tt.want) {
				t.Errorf("TwoSum(%v, %d) = %v, want %v (in any order)", tt.nums, tt.target, got, tt.want)
			}
		})
	}
}
