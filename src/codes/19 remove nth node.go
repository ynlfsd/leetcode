package codes


func removeNthFromEnd(head *ListNode, n int) *ListNode {

	i:=1
	ans:=head
	tmp:=head
	for head.Next!=nil{
		if i>n{
			tmp=tmp.Next
		}
			i++

		head=head.Next
	}
	if i>n+1{
		tmp.Next=tmp.Next.Next
		return ans
	}

	return ans.Next
}
