package threesum

import (
	"math/rand"
	"slices"
	"testing"
)

// checkStatement fails the test if the input doesn't follow the constraints
// of the problem (leetcode15.md):
//   - 3 <= nums.length <= 3000
//   - -10^5 <= nums[i] <= 10^5
func checkStatement(t *testing.T, nums []int) {
	t.Helper()

	if len(nums) < 3 || len(nums) > 3000 {
		t.Fatalf("test setup error: nums.length is %d, want between 3 and 3000", len(nums))
	}
	for _, n := range nums {
		if n < -100_000 || n > 100_000 {
			t.Fatalf("test setup error: %d is out of the range [-100000, 100000]", n)
		}
	}
}

// expectedTriplets is an independent way of finding the answer, used to
// check the test cases and to get the expected result of big inputs. It looks
// at every pair of distinct values (a <= b), takes the third value that makes
// the sum zero (c >= b), and uses the number of times each value appears to
// check that the triplet can be formed without using the same element twice.
// Each triplet is sorted, and so is the result.
func expectedTriplets(nums []int) [][]int {
	freq := make(map[int]int)
	for _, n := range nums {
		freq[n]++
	}
	values := make([]int, 0, len(freq))
	for v := range freq {
		values = append(values, v)
	}
	slices.Sort(values)

	var result [][]int
	for i, a := range values {
		for _, b := range values[i:] {
			c := -a - b
			if c < b {
				continue
			}
			need := map[int]int{a: 0, b: 0, c: 0}
			need[a]++
			need[b]++
			need[c]++
			ok := true
			for v, count := range need {
				if freq[v] < count {
					ok = false
					break
				}
			}
			if ok {
				result = append(result, []int{a, b, c})
			}
		}
	}

	return normalize(result)
}

// normalize returns a copy of triplets where each triplet is sorted and the
// triplets themselves are sorted, so results can be compared regardless of
// the order in which the answer was returned.
func normalize(triplets [][]int) [][]int {
	out := make([][]int, 0, len(triplets))
	for _, tr := range triplets {
		c := slices.Clone(tr)
		slices.Sort(c)
		out = append(out, c)
	}
	slices.SortFunc(out, slices.Compare[[]int])

	return out
}

func sameTriplets(a, b [][]int) bool {
	return slices.EqualFunc(a, b, slices.Equal[[]int])
}

func TestThreeSum(t *testing.T) {
	tests := []struct {
		name string
		nums []int
		want [][]int
	}{
		{"leetcode example 1", []int{-1, 0, 1, 2, -1, -4}, [][]int{{-1, -1, 2}, {-1, 0, 1}}},
		{"leetcode example 2", []int{0, 1, 1}, [][]int{}},
		{"leetcode example 3", []int{0, 0, 0}, [][]int{{0, 0, 0}}},

		// Minimum length.
		{"minimum length without a solution", []int{1, 2, 3}, [][]int{}},
		{"minimum length with a solution", []int{-1, 0, 1}, [][]int{{-1, 0, 1}}},
		{"minimum length in a different order", []int{1, -1, 0}, [][]int{{-1, 0, 1}}},

		// No solution.
		{"all positive numbers", []int{1, 2, 3, 4}, [][]int{}},
		{"all negative numbers", []int{-1, -2, -3}, [][]int{}},
		{"pairs that sum zero but no triplet", []int{-1, 1, -2, 2}, [][]int{}},

		// The same element can't be used more than once.
		{"a single zero can't be used three times", []int{0, 1, 2}, [][]int{}},
		{"a value that must appear twice appears once", []int{-2, 1, 3}, [][]int{}},
		{"a value that appears twice is used twice", []int{-2, 1, 1}, [][]int{{-2, 1, 1}}},

		// Duplicated values must not produce duplicated triplets.
		{"all zeros", []int{0, 0, 0, 0, 0}, [][]int{{0, 0, 0}}},
		{"one extra copy of a value", []int{-2, 1, 1, 1}, [][]int{{-2, 1, 1}}},
		{"repeated values with one triplet", []int{-1, -1, -1, 2, 2, 2}, [][]int{{-1, -1, 2}}},
		{"the same triplet from different positions", []int{-1, 0, 1, 1, -1, 0}, [][]int{{-1, 0, 1}}},
		{"duplicates on both sides of the middle value", []int{-1, -1, 0, 1, 1}, [][]int{{-1, 0, 1}}},

		// Several triplets.
		{"two triplets", []int{-2, 0, 1, 1, 2}, [][]int{{-2, 0, 2}, {-2, 1, 1}}},
		{
			"symmetric values",
			[]int{-3, -2, -1, 0, 1, 2, 3},
			[][]int{{-3, 0, 3}, {-3, 1, 2}, {-2, -1, 3}, {-2, 0, 2}, {-1, 0, 1}},
		},
		{
			"nine triplets",
			[]int{-1, 0, 1, 2, -1, -4, -2, -3, 3, 0, 4},
			[][]int{
				{-4, 0, 4}, {-4, 1, 3}, {-3, -1, 4}, {-3, 0, 3}, {-3, 1, 2},
				{-2, -1, 3}, {-2, 0, 2}, {-1, -1, 2}, {-1, 0, 1},
			},
		},

		// Order of the input.
		{"sorted descending", []int{4, 2, 0, -2, -4}, [][]int{{-4, 0, 4}, {-2, 0, 2}}},
		{"unsorted", []int{3, -1, 0, -2, 2, 1}, [][]int{{-2, -1, 3}, {-2, 0, 2}, {-1, 0, 1}}},

		// Values at the constraint bounds.
		{"minimum and maximum values", []int{-100000, 0, 100000}, [][]int{{-100000, 0, 100000}}},
		{"two minimum values and a big one", []int{-50000, -50000, 100000}, [][]int{{-50000, -50000, 100000}}},
		{"values out of reach of each other", []int{-100000, -100000, 100000}, [][]int{}},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			checkStatement(t, tt.nums)

			// Check that the expected result of the test case is right.
			if want := normalize(tt.want); !sameTriplets(want, expectedTriplets(tt.nums)) {
				t.Fatalf("test setup error: the expected result %v is not the right answer for %v", tt.want, tt.nums)
			}

			got := ThreeSum(slices.Clone(tt.nums))

			if !sameTriplets(normalize(got), normalize(tt.want)) {
				t.Errorf("ThreeSum(%v) = %v, want %v (in any order)", tt.nums, got, tt.want)
			}
		})
	}
}

// TestThreeSumRandomSmallInputs compares the result with an independent
// calculation for many small random inputs, where the values are repeated
// often. It uses a fixed seed so a failure can be reproduced.
func TestThreeSumRandomSmallInputs(t *testing.T) {
	r := rand.New(rand.NewSource(1))

	for range 500 {
		n := 3 + r.Intn(10) // 3 to 12 elements
		nums := make([]int, n)
		for i := range nums {
			nums[i] = r.Intn(11) - 5 // -5 to 5
		}
		checkStatement(t, nums)

		want := expectedTriplets(nums)
		got := ThreeSum(slices.Clone(nums))

		if !sameTriplets(normalize(got), want) {
			t.Fatalf("ThreeSum(%v) = %v, want %v (in any order)", nums, got, want)
		}
	}
}

// TestThreeSumLargeInput uses the maximum allowed length (3000). A solution
// that checks every possible triplet (O(n^3)) takes much longer here.
func TestThreeSumLargeInput(t *testing.T) {
	r := rand.New(rand.NewSource(2))

	random := func(n, limit int) []int {
		nums := make([]int, n)
		for i := range nums {
			nums[i] = r.Intn(2*limit+1) - limit
		}

		return nums
	}

	repeated := make([]int, 0, 3000)
	for range 1000 {
		repeated = append(repeated, -1, 0, 1)
	}

	tests := []struct {
		name string
		nums []int
	}{
		{"all zeros", make([]int, 3000)},
		{"-1, 0 and 1 repeated", repeated},
		{"small values with many triplets", random(3000, 100)},
		{"values in the whole range", random(3000, 100_000)},
		{"only positive values", func() []int {
			nums := random(3000, 100_000)
			for i := range nums {
				if nums[i] <= 0 {
					nums[i] = 1 - nums[i]
				}
			}

			return nums
		}()},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			checkStatement(t, tt.nums)

			want := expectedTriplets(tt.nums)
			got := ThreeSum(slices.Clone(tt.nums))

			if !sameTriplets(normalize(got), want) {
				t.Errorf("ThreeSum(large input) returned %d triplets, want %d", len(got), len(want))
			}
		})
	}
}
