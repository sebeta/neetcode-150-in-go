package trappingrainwater

import (
	"math/rand"
	"testing"
)

// checkStatement fails the test if the input doesn't follow the constraints
// of the problem (leetcode42.md):
//   - 1 <= height.length <= 2 * 10^4
//   - 0 <= height[i] <= 10^5
func checkStatement(t *testing.T, height []int) {
	t.Helper()

	if len(height) < 1 || len(height) > 20_000 {
		t.Fatalf("test setup error: height.length is %d, want between 1 and 20000", len(height))
	}
	for _, h := range height {
		if h < 0 || h > 100_000 {
			t.Fatalf("test setup error: %d is out of the range [0, 100000]", h)
		}
	}
}

// bruteForce is an independent way of finding the answer: for each bar it
// looks for the tallest bar on its left and on its right, and the water above
// it is the shorter of those two minus its own height. It is too slow for
// big inputs, so it is only used to check the test cases and small random
// inputs.
func bruteForce(height []int) int {
	total := 0
	for i := range height {
		leftMax, rightMax := 0, 0
		for j := 0; j <= i; j++ {
			leftMax = max(leftMax, height[j])
		}
		for j := i; j < len(height); j++ {
			rightMax = max(rightMax, height[j])
		}
		total += min(leftMax, rightMax) - height[i]
	}

	return total
}

// withPrefixMax calculates the same answer in O(n) with two arrays that keep
// the tallest bar to the left and to the right of each position. It is used
// to get the expected result of big inputs.
func withPrefixMax(height []int) int {
	n := len(height)
	leftMax := make([]int, n)
	rightMax := make([]int, n)
	for i := range height {
		leftMax[i] = height[i]
		if i > 0 {
			leftMax[i] = max(leftMax[i-1], height[i])
		}
	}
	for i := n - 1; i >= 0; i-- {
		rightMax[i] = height[i]
		if i < n-1 {
			rightMax[i] = max(rightMax[i+1], height[i])
		}
	}

	total := 0
	for i := range height {
		total += min(leftMax[i], rightMax[i]) - height[i]
	}

	return total
}

func TestTrap(t *testing.T) {
	tests := []struct {
		name   string
		height []int
		want   int
	}{
		{"leetcode example 1", []int{0, 1, 0, 2, 1, 0, 1, 3, 2, 1, 2, 1}, 6},
		{"leetcode example 2", []int{4, 2, 0, 3, 2, 5}, 9},

		// Too small to trap water.
		{"single bar of zero height", []int{0}, 0},
		{"single bar", []int{5}, 0},
		{"two bars", []int{3, 5}, 0},
		{"two bars of the same height", []int{4, 4}, 0},

		// No water is trapped.
		{"all bars with zero height", []int{0, 0, 0, 0}, 0},
		{"all bars with the same height", []int{2, 2, 2}, 0},
		{"ascending heights", []int{1, 2, 3, 4}, 0},
		{"descending heights", []int{4, 3, 2, 1}, 0},
		{"a mountain", []int{0, 1, 2, 3, 2, 1, 0}, 0},
		{"a hole at the start is not a container", []int{0, 3, 3}, 0},
		{"a hole at the end is not a container", []int{3, 3, 0}, 0},

		// A single valley.
		{"valley of width one", []int{3, 0, 3}, 3},
		{"valley of width two", []int{3, 0, 0, 3}, 6},
		{"valley with a smaller base", []int{2, 0, 2}, 2},
		{"left wall is taller", []int{3, 0, 2}, 2},
		{"right wall is taller", []int{2, 0, 3}, 2},
		{"bars of different height inside the valley", []int{5, 4, 3, 2, 3, 4, 5}, 9},
		{"the water level is set by the lower wall", []int{5, 0, 1, 0, 2}, 5},

		// Several valleys.
		{"two valleys", []int{3, 0, 3, 0, 3}, 6},
		{"two valleys separated by a tall wall", []int{3, 0, 3, 10, 3, 0, 3}, 6},
		{"two wide valleys", []int{4, 0, 0, 4, 0, 0, 4}, 16},
		{"valleys inside a valley", []int{5, 1, 3, 1, 2, 0, 4}, 13},

		// Values at the constraint bounds.
		{"maximum height", []int{100000, 0, 100000}, 100000},
		{"two valleys with the maximum height", []int{100000, 0, 100000, 0, 100000}, 200000},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			checkStatement(t, tt.height)

			// Check that the expected result of the test case is right.
			if want := bruteForce(tt.height); want != tt.want {
				t.Fatalf("test setup error: the expected result is %d, but the right answer for %v is %d", tt.want, tt.height, want)
			}

			if got := Trap(tt.height); got != tt.want {
				t.Errorf("Trap(%v) = %d, want %d", tt.height, got, tt.want)
			}
		})
	}
}

// TestTrapRandomSmallInputs compares the result with a brute force
// calculation for many small random inputs. It uses a fixed seed so a
// failure can be reproduced.
func TestTrapRandomSmallInputs(t *testing.T) {
	r := rand.New(rand.NewSource(1))

	for range 1000 {
		n := 1 + r.Intn(15) // 1 to 15 bars
		height := make([]int, n)
		for i := range height {
			height[i] = r.Intn(11) // 0 to 10
		}
		checkStatement(t, height)

		want := bruteForce(height)
		if got := Trap(height); got != want {
			t.Fatalf("Trap(%v) = %d, want %d", height, got, want)
		}
	}
}

// TestTrapLargeInput uses the maximum allowed length (2 * 10^4). A solution
// that looks for the tallest bars again for every position (O(n^2)) takes
// much longer here. Some results are also very big, so they check for
// overflow.
func TestTrapLargeInput(t *testing.T) {
	const n = 20_000
	r := rand.New(rand.NewSource(2))

	wideValley := make([]int, n)
	wideValley[0], wideValley[n-1] = 100_000, 100_000

	sawtooth := make([]int, n)
	for i := range sawtooth {
		if i%2 == 0 {
			sawtooth[i] = 100_000
		}
	}

	random := func(limit int) []int {
		height := make([]int, n)
		for i := range height {
			height[i] = r.Intn(limit + 1)
		}

		return height
	}

	bowl := make([]int, n) // 100000 at the ends, 50000 in the middle
	for i := range bowl {
		bowl[i] = 100_000 - min(i, n-1-i)*5
	}

	ascending := make([]int, n)
	for i := range ascending {
		ascending[i] = i
	}

	cases := []struct {
		name   string
		height []int
		want   int // -1 means: use the O(n) calculation
	}{
		{"very wide valley", wideValley, 100_000 * (n - 2)},
		{"sawtooth of maximum height", sawtooth, 100_000 * (n/2 - 1)},
		{"random small heights", random(10), -1},
		{"random big heights", random(100_000), -1},
		{"bowl shape", bowl, -1},
		{"ascending heights", ascending, 0},
	}

	for _, tt := range cases {
		t.Run(tt.name, func(t *testing.T) {
			checkStatement(t, tt.height)

			want := withPrefixMax(tt.height)
			if tt.want != -1 && tt.want != want {
				t.Fatalf("test setup error: the expected result is %d, but the right answer is %d", tt.want, want)
			}
			if got := Trap(tt.height); got != want {
				t.Errorf("Trap(large input) = %d, want %d", got, want)
			}
		})
	}
}
