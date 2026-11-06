package containerwithmostwater

import (
	"math/rand"
	"testing"
)

// checkStatement fails the test if the input doesn't follow the constraints
// of the problem (leetcode11.md):
//   - 2 <= height.length <= 10^5
//   - 0 <= height[i] <= 10^4
func checkStatement(t *testing.T, height []int) {
	t.Helper()

	if len(height) < 2 || len(height) > 100_000 {
		t.Fatalf("test setup error: height.length is %d, want between 2 and 100000", len(height))
	}
	for _, h := range height {
		if h < 0 || h > 10_000 {
			t.Fatalf("test setup error: %d is out of the range [0, 10000]", h)
		}
	}
}

// bruteForce is an independent way of finding the answer: it tries every
// pair of lines. It is too slow for big inputs, so it is only used to check
// the test cases and for small random inputs.
func bruteForce(height []int) int {
	best := 0
	for i := 0; i < len(height); i++ {
		for j := i + 1; j < len(height); j++ {
			area := min(height[i], height[j]) * (j - i)
			best = max(best, area)
		}
	}

	return best
}

func TestMaxArea(t *testing.T) {
	tests := []struct {
		name   string
		height []int
		want   int
	}{
		{"leetcode example 1", []int{1, 8, 6, 2, 5, 4, 8, 3, 7}, 49},
		{"leetcode example 2", []int{1, 1}, 1},

		// Minimum length.
		{"two equal lines", []int{3, 3}, 3},
		{"two lines, the second one is taller", []int{1, 2}, 1},
		{"two lines, the first one is taller", []int{2, 1}, 1},
		{"two lines with zero height", []int{0, 0}, 0},
		{"two lines, only the first has height", []int{5, 0}, 0},
		{"two lines, only the second has height", []int{0, 5}, 0},

		// Position of the answer.
		{"tall lines at the ends", []int{5, 1, 1, 1, 5}, 20},
		{"tall lines in the middle", []int{1, 2, 10, 10, 2, 1}, 10},
		{"tall lines are far from the ends", []int{1, 1, 9, 1, 1, 1, 9, 1, 1}, 36},
		{"the widest container is not the best", []int{2, 3, 10, 5, 7, 8, 9}, 36},
		{"the two tallest lines are not the best pair", []int{6, 1, 1, 1, 1, 1, 1, 10, 10, 6}, 54},

		// Order of the heights.
		{"ascending heights", []int{1, 2, 3, 4, 5}, 6},
		{"descending heights", []int{5, 4, 3, 2, 1}, 6},
		{"all lines with the same height", []int{4, 4, 4, 4}, 12},
		{"a valley", []int{6, 3, 1, 3, 6}, 24},
		{"a mountain", []int{1, 3, 6, 3, 1}, 6},

		// Zeros.
		{"all zeros", []int{0, 0, 0}, 0},
		{"a single line with height", []int{0, 3, 0}, 0},
		{"two lines with height separated by zeros", []int{0, 5, 0, 5, 0}, 10},

		// Values at the constraint bounds.
		{"two lines with the maximum height", []int{10000, 10000}, 10000},
		{"maximum height with a gap", []int{10000, 1, 10000}, 20000},
		{"maximum height next to zero", []int{10000, 0, 0, 0, 0}, 0},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			checkStatement(t, tt.height)

			// Check that the expected result of the test case is right.
			if want := bruteForce(tt.height); want != tt.want {
				t.Fatalf("test setup error: the expected result is %d, but the right answer for %v is %d", tt.want, tt.height, want)
			}

			if got := MaxArea(tt.height); got != tt.want {
				t.Errorf("MaxArea(%v) = %d, want %d", tt.height, got, tt.want)
			}
		})
	}
}

// TestMaxAreaRandomSmallInputs compares the result with a brute force
// calculation for many small random inputs. It uses a fixed seed so a
// failure can be reproduced.
func TestMaxAreaRandomSmallInputs(t *testing.T) {
	r := rand.New(rand.NewSource(1))

	for range 1000 {
		n := 2 + r.Intn(14) // 2 to 15 lines
		height := make([]int, n)
		for i := range height {
			height[i] = r.Intn(21) // 0 to 20
		}
		checkStatement(t, height)

		want := bruteForce(height)
		if got := MaxArea(height); got != want {
			t.Fatalf("MaxArea(%v) = %d, want %d", height, got, want)
		}
	}
}

// TestMaxAreaLargeInput uses the maximum allowed length (10^5). A solution
// that tries every pair of lines (O(n^2)) takes very long here. The results
// are also the biggest possible ones, so they check for overflow.
func TestMaxAreaLargeInput(t *testing.T) {
	const n = 100_000

	allEqual := make([]int, n)
	for i := range allEqual {
		allEqual[i] = 10_000
	}

	tallAtEnds := make([]int, n)
	tallAtEnds[0], tallAtEnds[n-1] = 10_000, 10_000

	tallInTheMiddle := make([]int, n)
	tallInTheMiddle[40_000], tallInTheMiddle[60_000] = 10_000, 10_000

	ascending := make([]int, n)
	for i := range ascending {
		ascending[i] = i / 10 // 0 to 9999
	}

	tests := []struct {
		name   string
		height []int
		want   int
	}{
		{"all lines with the maximum height", allEqual, 10_000 * (n - 1)},
		{"tall lines only at the ends", tallAtEnds, 10_000 * (n - 1)},
		{"tall lines only in the middle", tallInTheMiddle, 10_000 * 20_000},
		{"ascending heights", ascending, bestAscending(n)},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			checkStatement(t, tt.height)

			if got := MaxArea(tt.height); got != tt.want {
				t.Errorf("MaxArea(large input) = %d, want %d", got, tt.want)
			}
		})
	}
}

// bestAscending returns the answer for the heights i/10 (0, 0, ..., 1, 1,
// ...), which never decrease. In that case the best container always has the
// last line as its right side (moving the right side to the end makes the
// container wider and its height can't get lower), so it is enough to try
// every left line with the last one.
func bestAscending(n int) int {
	best := 0
	j := n - 1
	for i := 0; i < j; i++ {
		best = max(best, min(i/10, j/10)*(j-i))
	}

	return best
}
