package codes

func mergeTwoLists(l1 *ListNode, l2 *ListNode) *ListNode {

	if l1==nil{return l2}
	if l2==nil{return l1}
	if l1.Val>l2.Val{
		l1,l2=l2,l1
	}
	ans:=l1
	for l1.Next!=nil&&l2!=nil{
		if l1.Next.Val>=l2.Val{
			tmp:=l1.Next
			l1.Next=l2
			l2=l2.Next
			l1.Next.Next=tmp

		}else{
			l1=l1.Next
		}
	}
	if l2!=nil{
		l1.Next=l2
	}
	return ans
}
