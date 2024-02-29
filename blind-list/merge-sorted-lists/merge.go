package merge_sorted_lists

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
func mergeTwoLists(list1 *ListNode, list2 *ListNode) *ListNode {
	var head, current *ListNode
	for list1 != nil && list2 != nil {
		var node *ListNode
		if list1.Val < list2.Val {
			node = list1
			list1 = list1.Next
		} else {
			node = list2
			list2 = list2.Next
		}
		if current == nil {
			head = node
			current = node
		} else {
			current.Next = node
			current = current.Next
		}
	}
	if list1 != nil {
		if current == nil {
			head = list1
		} else {
			current.Next = list1
		}
	}
	if list2 != nil {
		if current == nil {
			head = list2
		} else {
			current.Next = list2
		}
	}
	return head

}
