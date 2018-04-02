package gleet
import (
	"testing"
)

// You are given two non-empty linked lists representing two
// non-negative integers. The digits are stored in reverse order and
// each of their nodes contain a single digit. Add the two numbers and
// return it as a linked list.  You may assume the two numbers do not
// contain any leading zero, except the number 0 itself.  Example
// 
// Input: (2 -> 4 -> 3) + (5 -> 6 -> 4)
// Output: 7 -> 0 -> 8
// Explanation: 342 + 465 = 807.

type twoNumCase struct {
	l1 []int
	l2 []int
	want []int
}

func TestAddTwoNum(t *testing.T) {
	tCases := []twoNumCase {
		twoNumCase{
			l1: []int{2, 4, 3},
			l2: []int{5, 6, 4},
			want: []int{7, 0, 8},
		},
		twoNumCase{
			l1: []int{7},
			l2: []int{5},
			want: []int{2, 1},
		},
		twoNumCase{
			l1: []int{0, 9},
			l2: []int{5, 9},
			want: []int{5, 8, 1},
		},
		twoNumCase{
			l1: []int{9},
			l2: []int{5, 9},
			want: []int{4, 0, 1},
		},
	}
	for _, c := range tCases {
		h1 := Slice2List(c.l1)
		h2 := Slice2List(c.l2)
		gotList := addTwoNumbers(h1, h2)
		gotAry := List2Slice(gotList)
		if len(gotAry) == len(c.want) {
			for i, w := range c.want {
				if gotAry[i] != w {
					t.Errorf("unmatched value %d: %v vs %v", i, gotAry[i], w)
				}
			}
		} else {
			t.Errorf("unequal arrays: %v vs %v", gotAry, c.want)
		}
	}
}
