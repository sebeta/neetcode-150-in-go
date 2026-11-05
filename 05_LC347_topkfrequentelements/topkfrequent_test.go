package topkfrequent

import (
	"math/rand"
	"slices"
	"testing"
)

func TestTopKFrequent(t *testing.T) {
	tests := []struct {
		name string
		nums []int
		k    int
		want []int // expected elements, sorted ascending (the answer can be returned in any order)
	}{
		{"leetcode example 1", []int{1, 1, 1, 2, 2, 3}, 2, []int{1, 2}},
		{"leetcode example 2", []int{1}, 1, []int{1}},
		{"leetcode example 3", []int{1, 2, 1, 2, 1, 2, 3, 1, 3, 2}, 2, []int{1, 2}},
		{"k is 1", []int{4, 4, 4, 5, 5, 6}, 1, []int{4}},
		{"k equals number of unique elements", []int{1, 2, 3, 3, 2, 3}, 3, []int{1, 2, 3}},
		{"all elements are the same", []int{7, 7, 7, 7}, 1, []int{7}},
		{"all elements distinct, k is all of them", []int{5, 3, 1}, 3, []int{1, 3, 5}},
		{"negative numbers", []int{-1, -1, -2, -2, -2, 3}, 2, []int{-2, -1}},
		{"zero is the most frequent", []int{0, 0, 0, 1, 1, 2}, 1, []int{0}},
		{"values at the constraint bounds", []int{10000, -10000, 10000, -10000, 10000, 0}, 2, []int{-10000, 10000}},
		{"unsorted input", []int{3, 1, 2, 3, 2, 3, 1, 3, 2, 2}, 2, []int{2, 3}},
		{"most frequent elements at the end", []int{9, 8, 8, 7, 7, 7, 6, 6, 6, 6}, 2, []int{6, 7}},
		{"most frequent elements at the start", []int{6, 6, 6, 6, 7, 7, 7, 8, 8, 9}, 2, []int{6, 7}},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			// Copy the input to detect if the function modifies it.
			nums := slices.Clone(tt.nums)

			got := TopKFrequent(nums, tt.k)

			if !slices.Equal(nums, tt.nums) {
				t.Errorf("TopKFrequent modified its input: got %v, want %v", nums, tt.nums)
			}

			sorted := slices.Clone(got)
			slices.Sort(sorted)
			if !slices.Equal(sorted, tt.want) {
				t.Errorf("TopKFrequent(%v, %d) = %v, want %v (in any order)", tt.nums, tt.k, got, tt.want)
			}
		})
	}
}

// TestTopKFrequentLargeInput uses an input close to the maximum allowed size
// (10^5 elements) in shuffled order. Value v appears exactly v times, so the
// most frequent elements are the largest values.
func TestTopKFrequentLargeInput(t *testing.T) {
	const maxValue = 440 // 1 + 2 + ... + 440 = 97020 elements

	var nums []int
	for v := 1; v <= maxValue; v++ {
		for range v {
			nums = append(nums, v)
		}
	}
	rand.New(rand.NewSource(1)).Shuffle(len(nums), func(i, j int) {
		nums[i], nums[j] = nums[j], nums[i]
	})

	got := TopKFrequent(nums, 3)

	slices.Sort(got)
	want := []int{438, 439, 440}
	if !slices.Equal(got, want) {
		t.Errorf("TopKFrequent(large input, 3) = %v, want %v (in any order)", got, want)
	}
}
