// Package containerwithmostwater solves the "Container with most water" problem.
package containerwithmostwater

// MaxArea returns the maximum amount of water a container can store, where a
// container is formed by two of the vertical lines in heights and the x-axis.
// The area of a container is the height of its shorter line multiplied by the
// distance between both lines. It returns 0 if no container can hold water.
//
// It uses two pointers, one at the start and one at the end of heights, which
// is the widest container. Then it keeps moving towards the center the
// pointer of the shorter line: moving the taller one can only give a narrower
// container with the same height limit, so it can't be better.
//
// Time complexity: O(n), where n is the length of heights.
// Space complexity: O(1), since only two pointers are used.
//
// NOTE: This is my solution. It is the same as the most optimal from the video.
func MaxArea(heights []int) int {
	var maxArea int
	l := 0                // Left
	r := len(heights) - 1 // Right
	for l < r {
		area := min(heights[l], heights[r]) * (r - l)
		maxArea = max(maxArea, area)

		// Move the pointer of the sorter line
		if heights[l] <= heights[r] {
			l++
		} else {
			r--
		}
	}

	return maxArea
}
