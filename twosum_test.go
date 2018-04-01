package gleet

// Given an array of integers, return indices of the two numbers such
// that they add up to a specific target.  You may assume that each
// input would have exactly one solution, and you may not use the same
// element twice.
// Example:
// Given nums = [2, 7, 11, 15], target = 9,
// Because nums[0] + nums[1] = 2 + 7 = 9,
// return [0, 1].

import (
	"testing"
)

type testCase struct {
	nums []int
	target int
	want []int
}

func TestTwoSum(t *testing.T) {
	testCases := []testCase {
		testCase {
			nums: []int {2, 7, 11, 15},
			target: 9,
			want: []int {0, 1},
		},
		testCase {
			[]int {1, 3, 5, 2},
			3,
			[]int {0, 3},
		},
		testCase {
			[]int {100, 200, 300, 400, 500, 600, 700, 800, 900, 9, 8, 7, 6, 5, 4, 3, 2, 1},
			202,
			[]int {1, 16},
		},
	}
	for _, c := range testCases {
		ans := twoSum(c.nums, c.target)
		if ans != nil && len(ans) == len(c.want) {
			for i, e := range c.want {
				if e != ans[i] {
					t.Errorf("case %v, got %v", c, ans)
				}
			}
		} else {
			t.Errorf("case %v, got %v", c, ans)
		}
	}
}
