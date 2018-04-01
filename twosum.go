package gleet

// Given an array of integers, return indices of the two numbers such
// that they add up to a specific target.  You may assume that each
// input would have exactly one solution, and you may not use the same
// element twice.

func twoSum(nums []int, target int) []int {
	result := []int{}
	p := make(map[int]int)		// position of the remainder's partner
	for i, v := range nums {
		delta := target - v
		if so, ok := p[delta]; ok {
			// this value has a partner seen earlier
			result = []int{so, i}
			break
		} else {
			p[v] = i
		}
	}
	return result
}
