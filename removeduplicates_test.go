package gleet
import (
	"testing"
)

// Given a sorted array, remove the duplicates in-place such that
// each element appear only once and return the new length.  Do not
// allocate extra space for another array, you must do this by
// modifying the input array in-place with O(1) extra memory.
// Example:
// Given nums = [1,1,2],
// Your function should return length = 2, with the first two elements
// of nums being 1 and 2 respectively.
// It doesn't matter what you leave beyond the new length.

type rmDupCase struct {
	nums []int
	want int
}

func TestRmDups(t *testing.T) {
	tCases := []rmDupCase {
		rmDupCase{
			nums: []int{1, 1, 2},
			want: 2,
		},
		rmDupCase{
			nums: []int{0, 0, 2, 3, 4, 4, 4, 4, 5, 5, 6, 6},
			want: 6,
		},
	}
	for _, c := range tCases {
		got := removeDuplicates(c.nums)
		if got != c.want {
			t.Errorf("case %v: want %d, got %d", c.nums, c.want, got)
		} else {
			var p int
			for i := 0; i < got; i++ {
				v := c.nums[i]
				if i > 0 &&  v == p {
					t.Errorf("duplicate nums[%d] = %d, p = %d", i, v, p)
				}
				p = v
			}
		}
	}
}
