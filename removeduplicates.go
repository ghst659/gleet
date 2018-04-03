package gleet

// Given a sorted array, remove the duplicates in-place such that
// each element appear only once and return the new length.  Do not
// allocate extra space for another array, you must do this by
// modifying the input array in-place with O(1) extra memory.
// Example:
// Given nums = [1,1,2],
// Your function should return length = 2, with the first two elements
// of nums being 1 and 2 respectively.
// It doesn't matter what you leave beyond the new length.

func removeDuplicates(nums []int) int {
	var p int
	n := len(nums)
	var j int
	for i := 0; i < n; i++ {
		if i == 0 {
			p = nums[i]
			j = i + 1
		} else if nums[i] == p {
			// same as last element
		} else {
			// new element
			p = nums[i]
			nums[j] = nums[i]
			j++
		}
	}
	return j
}
