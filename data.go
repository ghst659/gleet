package gleet
import (
	"fmt"
	"strings"
	"strconv"
)

type ListNode struct {
	Val int
	Next *ListNode
}

func Slice2List(ary []int) (h *ListNode) {
	if ary != nil && len(ary) > 0 {
		var p *ListNode
		for i, v := range ary {
			n := new(ListNode)
			n.Val = v
			if i == 0 {
				h = n
			} else {
				p.Next = n
			}
			p = n
		}
	}
	return
}

func List2Slice(h *ListNode) (a []int) {
	a = []int{}
	for c := h; c != nil; c = c.Next {
		a = append(a, c.Val)
	}
	return
}

func List2String(h *ListNode) string {
	str := func(ary []int) []string {
		r := make([]string, len(ary))
		for i, v := range ary {
			r[i] = strconv.Itoa(v)
		}
		return r
	}
	return fmt.Sprintf("(%s)", strings.Join(str(List2Slice(h)), " -> "))
}
