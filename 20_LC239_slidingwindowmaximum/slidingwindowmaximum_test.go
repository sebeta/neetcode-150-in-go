package slidingwindowmaximum

import (
	"math/rand"
	"slices"
	"testing"
)

// checkStatement fails the test if the input doesn't follow the constraints
// of the problem (leetcode239.md):
//   - 1 <= nums.length <= 10^5
//   - -10^4 <= nums[i] <= 10^4
//   - 1 <= k <= nums.length
func checkStatement(t *testing.T, nums []int, k int) {
	t.Helper()

	if len(nums) < 1 || len(nums) > 100_000 {
		t.Fatalf("test setup error: nums.length is %d, want between 1 and 100000", len(nums))
	}
	for _, n := range nums {
		if n < -10_000 || n > 10_000 {
			t.Fatalf("test setup error: %d is out of the range [-10000, 10000]", n)
		}
	}
	if k < 1 || k > len(nums) {
		t.Fatalf("test setup error: k is %d, want between 1 and %d", k, len(nums))
	}
}

// bruteForce is an independent way of finding the answer: it looks for the
// maximum of every window from scratch. It is too slow for big inputs, so it
// is only used to check the test cases and small random inputs.
func bruteForce(nums []int, k int) []int {
	result := make([]int, 0, len(nums)-k+1)
	for start := 0; start+k <= len(nums); start++ {
		result = append(result, slices.Max(nums[start:start+k]))
	}

	return result
}

// byBlocks calculates the same answer in O(n) with a different idea: nums is
// split in blocks of k elements, and each window is made of the end of one
// block and the start of the next one. It keeps, for each position, the
// maximum from the start of its block up to it (prefix) and from it up to the
// end of its block (suffix). It is used to get the expected result of big
// inputs.
func byBlocks(nums []int, k int) []int {
	n := len(nums)
	prefix := make([]int, n)
	suffix := make([]int, n)
	for i := range nums {
		if i%k == 0 {
			prefix[i] = nums[i]
		} else {
			prefix[i] = max(prefix[i-1], nums[i])
		}
	}
	for i := n - 1; i >= 0; i-- {
		if i == n-1 || (i+1)%k == 0 {
			suffix[i] = nums[i]
		} else {
			suffix[i] = max(suffix[i+1], nums[i])
		}
	}

	result := make([]int, 0, n-k+1)
	for start := 0; start+k <= n; start++ {
		result = append(result, max(suffix[start], prefix[start+k-1]))
	}

	return result
}

func TestMaxSlidingWindow(t *testing.T) {
	tests := []struct {
		name string
		nums []int
		k    int
		want []int
	}{
		{"leetcode example 1", []int{1, 3, -1, -3, 5, 3, 6, 7}, 3, []int{3, 3, 5, 5, 6, 7}},
		{"leetcode example 2", []int{1}, 1, []int{1}},

		// The size of the window.
		{"k is 1: every number is its own maximum", []int{4, 2, 7}, 1, []int{4, 2, 7}},
		{"k is the length: a single window", []int{1, 3, 2}, 3, []int{3}},
		{"k is the length minus one", []int{1, 3, 2, 5}, 3, []int{3, 5}},
		{"k is 2", []int{7, 2, 4}, 2, []int{7, 4}},
		{"single element", []int{5}, 1, []int{5}},

		// Order of the numbers.
		{"all numbers are equal", []int{2, 2, 2, 2}, 2, []int{2, 2, 2}},
		{"ascending numbers", []int{1, 2, 3, 4, 5}, 3, []int{3, 4, 5}},
		{"descending numbers", []int{5, 4, 3, 2, 1}, 3, []int{5, 4, 3}},
		{"zigzag", []int{1, 5, 1, 5, 1, 5}, 2, []int{5, 5, 5, 5, 5}},

		// The maximum leaves the window.
		{"the maximum leaves the window", []int{9, 1, 1, 1, 8}, 3, []int{9, 1, 8}},
		{"the maximum leaves and a much smaller one stays", []int{10, 1, 2, 3, 0}, 2, []int{10, 2, 3, 3}},
		{"the maximum is in the middle of the window", []int{1, 5, 1, 1, 1}, 3, []int{5, 5, 1}},
		{"the maximum is in the last position of the window", []int{1, 2, 3, 1}, 2, []int{2, 3, 3}},

		// Repeated maximums.
		{"the maximum is repeated next to each other", []int{4, 4, 1, 4}, 2, []int{4, 4, 4}},
		{"the maximum is repeated and one of them leaves", []int{3, 1, 3, 1}, 3, []int{3, 3}},
		{"two equal maximums at the ends of the window", []int{6, 2, 6, 1, 1}, 3, []int{6, 6, 6}},

		// Negative numbers.
		{"negative numbers", []int{-1, -3, -2, -5}, 2, []int{-1, -2, -2}},
		{"all negative numbers and k is 1", []int{-4, -1, -9}, 1, []int{-4, -1, -9}},
		{"negative and positive numbers", []int{-3, 4, -2, 0, -7}, 3, []int{4, 4, 0}},
		{"zeros", []int{0, 0, -1, 0}, 2, []int{0, 0, 0}},

		// Values at the constraint bounds.
		{"maximum and minimum values", []int{10000, -10000, 10000}, 2, []int{10000, 10000}},
		{"minimum values", []int{-10000, -10000}, 1, []int{-10000, -10000}},
		{"minimum and maximum values", []int{-10000, 10000}, 2, []int{10000}},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			checkStatement(t, tt.nums, tt.k)

			// Check that the expected result of the test case is right.
			if want := bruteForce(tt.nums, tt.k); !slices.Equal(want, tt.want) {
				t.Fatalf("test setup error: the expected result is %v, but the right answer for (%v, %d) is %v", tt.want, tt.nums, tt.k, want)
			}

			if got := MaxSlidingWindow(slices.Clone(tt.nums), tt.k); !slices.Equal(got, tt.want) {
				t.Errorf("MaxSlidingWindow(%v, %d) = %v, want %v", tt.nums, tt.k, got, tt.want)
			}
		})
	}
}

// TestMaxSlidingWindowRandomSmallInputs compares the result with a brute
// force calculation for many small random inputs, with repeated numbers. It
// uses a fixed seed so a failure can be reproduced.
func TestMaxSlidingWindowRandomSmallInputs(t *testing.T) {
	r := rand.New(rand.NewSource(1))

	for range 3000 {
		n := 1 + r.Intn(15)
		nums := make([]int, n)
		for i := range nums {
			nums[i] = r.Intn(9) - 4 // -4 to 4
		}
		k := 1 + r.Intn(n)
		checkStatement(t, nums, k)

		want := bruteForce(nums, k)
		if got := MaxSlidingWindow(slices.Clone(nums), k); !slices.Equal(got, want) {
			t.Fatalf("MaxSlidingWindow(%v, %d) = %v, want %v", nums, k, got, want)
		}
	}
}

// TestMaxSlidingWindowLargeInput uses the maximum allowed length (10^5). A
// solution that looks for the maximum of every window from scratch (O(n*k))
// takes far too long with a big k.
func TestMaxSlidingWindowLargeInput(t *testing.T) {
	const n = 100_000
	r := rand.New(rand.NewSource(2))

	random := make([]int, n)
	ascending := make([]int, n)
	descending := make([]int, n)
	equal := make([]int, n)
	sawtooth := make([]int, n)
	for i := range n {
		random[i] = r.Intn(20_001) - 10_000
		ascending[i] = -10_000 + i/5 // -10000 up to 10000
		descending[i] = 10_000 - i/5 // 10000 down to -10000
		equal[i] = 7
		sawtooth[i] = i % 1000
	}

	tests := []struct {
		name string
		nums []int
		k    int
	}{
		{"random numbers, k is 1", random, 1},
		{"random numbers, small k", random, 3},
		{"random numbers, medium k", random, 1_000},
		{"random numbers, half of the length", random, n / 2},
		{"random numbers, k is the length minus one", random, n - 1},
		{"random numbers, k is the length", random, n},
		{"ascending numbers, medium k", ascending, 1_000},
		{"descending numbers, medium k", descending, 1_000},
		{"descending numbers, big k", descending, n / 2},
		{"all numbers are equal", equal, 5_000},
		{"sawtooth numbers", sawtooth, 2_500},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			checkStatement(t, tt.nums, tt.k)

			want := byBlocks(tt.nums, tt.k)
			got := MaxSlidingWindow(slices.Clone(tt.nums), tt.k)

			if !slices.Equal(got, want) {
				t.Errorf("MaxSlidingWindow(large input, %d) is wrong: it has %d numbers, want %d", tt.k, len(got), len(want))
			}
		})
	}
}
