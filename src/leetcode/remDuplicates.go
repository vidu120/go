package main

//this will be our structure for our node
type ListNode struct {
	Val int
	Next *ListNode
}

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
//this is the function to implement
func DeleteDuplicates(head *ListNode) *ListNode {
	if head == nil{
		return nil
	}
	//store the head of the list
	store := head
	var prevNode = head

	for head != nil{
		for head != nil && head.Val == prevNode.Val{
			head = head.Next
		}
		prevNode.Next = head
		prevNode = head
	}
	return store
}


//
//func main() {
//	node1 := ListNode{Val: 1 , Next: nil}
//	node2 := ListNode{Val: 1 , Next: nil}
//	node3 := ListNode{Val: 2 , Next: nil}
//	node1.Next = &node2
//	node2.Next = &node3
//	newHead := DeleteDuplicates(&node1)
//	for newHead != nil{
//		fmt.Println(newHead.Val)
//		newHead = newHead.Next
//	}
//}