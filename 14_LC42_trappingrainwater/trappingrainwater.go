// Package trappingrainwater solves the "Trapping rain water" problem.
package trappingrainwater

// Trap returns how much rain water an elevation map can trap, where each
// heights[i] is the height of a bar of width 1. The water above a bar reaches
// the height of the shorter of the two tallest bars on its left and on its
// right, minus the height of the bar itself (or 0 if that is not positive).
//
// It uses two pointers, l and r, that start at the ends of heights and move
// towards the center. At each step it looks at the shorter of the two bars
// l and r, because that side is the bottleneck: the other side has a bar at
// least as tall, so the water above the shorter bar only depends on the
// tallest bar seen so far on its own side (maxL or maxR). The water above
// that bar is added to the total, and its pointer moves one step inwards. A
// bar that is taller than all the previous ones on its side only raises
// maxL or maxR and adds no water.
//
// For example, with [3, 0, 2]: l=0 (3) is taller than r=2 (2), so the right
// side is handled first: maxR=2 and no water is added; then r=1 (0) is
// shorter than l (3), so maxR=2 gives 2-0=2 units of water above it.
//
// Time complexity: O(n), where n is the length of heights, since each step
// moves a pointer and it never goes back.
// Space complexity: O(1), since only two pointers and two maximums are used.
//
// NOTE: This exercise blew my mind. I couldn't get the O(n) answer without
// a lot of help from AI. I know this is the most optimal solution.
func Trap(heights []int) int {
	var water int
	l, r := 0, len(heights)-1 // Left and right
	maxL, maxR := 0, 0
	for l < r {
		if heights[l] < heights[r] {
			maxL = max(maxL, heights[l])
			water += maxL - heights[l]
			l++
		} else {
			maxR = max(maxR, heights[r])
			water += maxR - heights[r]
			r--
		}
	}

	return water
}
