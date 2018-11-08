package codes
type ListNode struct{
	Val int
	Next *ListNode
}
func deleteDuplicates(head *ListNode) *ListNode {
	ans:=head
	for head.Next!=nil{
		tmp:=head.Val
		for head.Next!=nil&&tmp==head.Next.Val{
			head.Next=head.Next.Next
			head=head.Next
		}

	}
	return ans
}
