// Package topkfrequent solves the "Top K frequent elements" problem.
package topkfrequent

// TopKFrequent returns the k most frequent elements in nums, in any order.
// It counts how many times each number appears, groups the numbers by that
// frequency (bucket sort), and then takes numbers from the highest frequency
// down until it has k of them. It assumes the answer is unique.
//
// Time complexity: O(n), where n is the length of nums.
// Space complexity: O(n), for the frequency map and the buckets.
//
// NOTE: This is my solution but I used some help from AI. This is
// more optimal than the one from the video.
func TopKFrequent(nums []int, k int) []int {
	if len(nums) == k {
		return nums
	}

	freqPerNum := make(map[int]int)
	for _, num := range nums {
		freqPerNum[num]++
	}

	// Keep the numbers using its frequency as position, considering
	// that the maximum frequency is the length of nums
	numsByFreq := make([][]int, len(nums)+1)
	for num, freq := range freqPerNum {
		numsByFreq[freq] = append(numsByFreq[freq], num)
	}

	top := make([]int, 0, k)
	// len(top) < 0 can be used since the answer is unique by constraint
	for freq := len(numsByFreq) - 1; freq > 0 && len(top) < k; freq-- {
		top = append(top, numsByFreq[freq]...)
	}

	return top
}
