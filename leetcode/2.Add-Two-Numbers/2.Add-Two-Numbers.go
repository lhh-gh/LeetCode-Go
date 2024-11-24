package leetcode

// ListNode 链表节点定义
type ListNode struct {
	Val  int
	Next *ListNode
}

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */

func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
	head := &ListNode{Val: 0}               //定义头节点
	n1, n2, carry, current := 0, 0, 0, head //n1 和 n2 用于存储两个链表当前节点的值 carry 用于存储进位，current 用于追踪结果链表的当前节点。
	for l1 != nil || l2 != nil || carry != 0 {
		if l1 == nil {
			n1 = 0
		} else {
			n1 = l1.Val
			l1 = l1.Next
		}
		if l2 == nil {
			n2 = 0
		} else {
			n2 = l2.Val
			l2 = l2.Next
		}
		current.Next = &ListNode{Val: (n1 + n2 + carry) % 10} //创建一个新的链表节点，其值为 (n1 + n2 + carry) % 10，即当前位的和加上进位后取模10的结果，这个值就是结果链表当前位的值。
		current = current.Next                                //将 current 移动到结果链表的下一个节点。
		carry = (n1 + n2 + carry) / 10                        //计算新的进位值，即 (n1 + n2 + carry) 除以10的商。
	}
	return head.Next
}
