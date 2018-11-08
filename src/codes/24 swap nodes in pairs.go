package codes

func SwapPairs(head *ListNode) *ListNode {
	if head==nil||head.Next==nil{
		return head
	}
	ans:=head.Next

	head.Next=SwapPairs(head.Next.Next)

	ans.Next=head
	return ans

}
