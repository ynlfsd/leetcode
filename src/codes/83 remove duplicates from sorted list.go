package codes

func _deleteDuplicates(head *ListNode) *ListNode {
	if head==nil{return nil}
	ans:=head
	for head!=nil{
		tmp:=head.Val
		for head.Next!=nil&&tmp==head.Next.Val{
			head.Next=head.Next.Next

		}
		head=head.Next
	}
	return ans
}
