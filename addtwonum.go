package gleet

// You are given two non-empty linked lists representing two
// non-negative integers. The digits are stored in reverse order and
// each of their nodes contain a single digit. Add the two numbers and
// return it as a linked list.  You may assume the two numbers do not
// contain any leading zero, except the number 0 itself.  Example
// 
// Input: (2 -> 4 -> 3) + (5 -> 6 -> 4)
// Output: 7 -> 0 -> 8
// Explanation: 342 + 465 = 807.

func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
    var result *ListNode
	var previous *ListNode
	for carry := 0; l1 != nil || l2 != nil || carry != 0; l1, l2 = next(l1, l2) {
		sum := val(l1) + val(l2) + carry
		digit := sum % 10
		carry = sum / 10
		current := new(ListNode)
		current.Val = digit
		if result == nil {
			result = current
		} else {
			previous.Next = current
		}
		previous = current
	}
	return result
}

func val(n *ListNode) int {
	if n != nil {
		return n.Val
	} else {
		return 0
	}
}

func next(l1 *ListNode, l2 *ListNode) (n1 *ListNode, n2 *ListNode) {
	if l1 != nil {
		n1 = l1.Next
	}
	if l2 != nil {
		n2 = l2.Next
	}
	return
}
