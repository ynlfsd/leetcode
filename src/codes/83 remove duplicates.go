package codes

func deleteDuplicates(head *ListNode) *ListNode {
	if head==nil{return nil}
	ans:=head
	for head.Next!=nil{
		if head.Val==head.Next.Val{
			head.Next=head.Next.Next
		}else{
			break
		}
	}
	deleteDuplicates(head.Next)
	return ans
}
