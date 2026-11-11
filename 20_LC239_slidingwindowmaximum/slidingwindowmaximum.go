// Package slidingwindowmaximum solves the "Sliding window maximum" problem.
package slidingwindowmaximum

// MaxSlidingWindow returns the maximum of every window of k consecutive
// numbers of nums, from the leftmost window to the rightmost one.
//
// It slides a window from position l to position r, and keeps the position of
// the maximum seen in it (maxPos). Each number that enters the window on the
// right is compared with that maximum. When the window has k numbers, its
// maximum is saved and l moves one position. If the maximum is no longer in
// the window (maxPos < l), r goes back to l and the window is read again to
// find the new maximum.
//
// For example, with nums=[9,1,2,3] and k=2: the first window [9,1] has the
// maximum 9. When the window moves to [1,2], the 9 is no longer inside, so the
// window is read again and the maximum is 2. The last window [2,3] has the
// maximum 3, so it returns [9,2,3].
//
// Time complexity: O(n*k) in the worst case, where n is the length of nums.
// When the maximum leaves the window, the window is read again, which costs
// up to k steps. This happens at every step with descending numbers. If the
// maximum rarely leaves the window, it is close to O(n).
// Space complexity: O(1), not counting the result.
//
// NOTE: This is my solution. It runs in O(n*k) and I'm kind of proud of it,
// even if I didn't achieve the optimal O(n) time complexity. I'll keep it.
func MaxSlidingWindow(nums []int, k int) []int {
	var maxNums []int
	var maxPos int
	for l, r := 0, 0; r <= len(nums)-1; r++ {
		if maxPos >= l && nums[r] > nums[maxPos] {
			maxPos = r
		} else if maxPos < l {
			r = l
			maxPos = l
		}

		if r-l == k-1 {
			maxNums = append(maxNums, nums[maxPos])
			l++
		}
	}

	return maxNums
}

// MaxSlidingWindow2 returns the maximum of every window of k consecutive
// numbers of nums, from the leftmost window to the rightmost one.
//
// It keeps a list of candidates: the positions of the numbers in the current
// window that can still be the maximum of some window, from the biggest number
// to the smallest one. When a new number enters, the candidates that are
// smaller than it are removed from the end of the list, since the new number
// will stay in the window longer than them, so they can never be the maximum
// again. Then the new position is added at the end. If the first candidate is
// not in the window anymore, it is removed from the start of the list. Once
// the first window is complete, the first candidate is the maximum of the
// window. Positions are kept instead of values to know when a candidate
// leaves the window.
//
// For example, with nums=[1,3,-1,-3,5,3,6,7] and k=3: after reading 1, 3 and
// -1, the candidates are 3 and -1 (the 1 was removed by the 3), so the first
// maximum is 3. When 5 enters, it removes -3, -1 and 3, so it is the only
// candidate and the maximum of that window.
//
// Time complexity: O(n), where n is the length of nums. Every position enters
// the list once and leaves it at most once.
// Space complexity: O(k), since the list never has more than k candidates
// (not counting the result).
//
// NOTE: This is the optimal solution, provided by AI.
func MaxSlidingWindow2(nums []int, k int) []int {
	result := make([]int, 0, len(nums)-k+1)
	var candidates []int // Positions, from the biggest number to the smallest
	for i, num := range nums {
		// The numbers smaller than num can't be the maximum anymore
		for len(candidates) > 0 && nums[candidates[len(candidates)-1]] <= num {
			candidates = candidates[:len(candidates)-1]
		}
		candidates = append(candidates, i)

		// The first candidate is not in the window anymore
		if candidates[0] <= i-k {
			candidates = candidates[1:]
		}

		// The first window is complete from position k-1
		if i >= k-1 {
			result = append(result, nums[candidates[0]])
		}
	}

	return result
}
