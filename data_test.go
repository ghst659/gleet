package gleet
import (
	"testing"
)

type listCase struct {
	Values []int
	Text string
}

func TestList(t *testing.T) {
	tCases := []listCase {
		listCase {
			Values: []int {10, 8, 6, 4, 2, 0},
			Text: "(10 -> 8 -> 6 -> 4 -> 2 -> 0)",
		},
		listCase {
			Values: []int {},
			Text: "()",
		},
		listCase {
			Values: []int {84},
			Text: "(84)",
		},
	}
	for _, c := range tCases {
		h := Slice2List(c.Values)
		n := h
		for _, v := range c.Values {
			if n.Val != v {
				t.Errorf("Expect %v, got %v", v, n.Val)
			}
			n = n.Next
		}
		txt := List2String(h)
		if txt != c.Text {
			t.Errorf("case %v, expect %v, text %v\n", c.Values, c.Text, txt)
		}
	}
}
