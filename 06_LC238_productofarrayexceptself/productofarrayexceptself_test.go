package productofarrayexceptself

import (
	"slices"
	"testing"
)

func TestProductExceptSelf(t *testing.T) {
	tests := []struct {
		name string
		nums []int
		want []int
	}{
		{"leetcode example 1", []int{1, 2, 3, 4}, []int{24, 12, 8, 6}},
		{"leetcode example 2", []int{-1, 1, 0, -3, 3}, []int{0, 0, 9, 0, 0}},
		{"minimum length", []int{2, 3}, []int{3, 2}},
		{"minimum length with a negative", []int{1, -1}, []int{-1, 1}},
		{"minimum length with a zero", []int{0, 4}, []int{4, 0}},
		{"minimum length with two zeros", []int{0, 0}, []int{0, 0}},
		{"all ones", []int{1, 1, 1, 1}, []int{1, 1, 1, 1}},
		{"all the same value", []int{5, 5, 5}, []int{25, 25, 25}},
		{"all negative", []int{-1, -2, -3, -4}, []int{-24, -12, -8, -6}},
		{"mixed signs", []int{-2, 3, -4}, []int{-12, 8, -6}},
		{"one zero in the middle", []int{2, 0, 3, 4}, []int{0, 24, 0, 0}},
		{"one zero at the start", []int{0, 2, 3}, []int{6, 0, 0}},
		{"one zero at the end", []int{2, 3, 0}, []int{0, 0, 6}},
		{"two zeros", []int{0, 2, 0, 3}, []int{0, 0, 0, 0}},
		{"values at the constraint bounds", []int{-30, 30}, []int{30, -30}},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			// Copy the input to detect if the function modifies it.
			nums := slices.Clone(tt.nums)

			got := ProductExceptSelf(nums)

			if !slices.Equal(nums, tt.nums) {
				t.Errorf("ProductExceptSelf modified its input: got %v, want %v", nums, tt.nums)
			}
			if !slices.Equal(got, tt.want) {
				t.Errorf("ProductExceptSelf(%v) = %v, want %v", tt.nums, got, tt.want)
			}
		})
	}
}

// TestProductExceptSelfLargeInput uses the maximum allowed length (10^5).
// Every element is 1 except the first (2) and the last (3), so the products
// stay within 32 bits as the constraints require.
func TestProductExceptSelfLargeInput(t *testing.T) {
	const n = 100_000

	nums := make([]int, n)
	for i := range nums {
		nums[i] = 1
	}
	nums[0], nums[n-1] = 2, 3

	want := make([]int, n)
	for i := range want {
		want[i] = 6 // 2 * 3
	}
	want[0], want[n-1] = 3, 2

	got := ProductExceptSelf(nums)

	if !slices.Equal(got, want) {
		t.Errorf("ProductExceptSelf(large input) returned a wrong result (length %d, want %d)", len(got), len(want))
	}
}
