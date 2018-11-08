package codes

type ListNode struct{
	Val int
	Next *ListNode
}

func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
	head:=l1
	t:=0
	for {
		l1.Val=l1.Val+l2.Val+t
		t=0
		if l1.Val>9{
			l1.Val-=10
			t=1
		}
		if l1.Next!=nil&&l2.Next!=nil{
			l1=l1.Next
			l2=l2.Next
		}else{
			break
		}

	}
	if l1.Next==nil{
		l1.Next=l2.Next
	}
	if t==1{
		addone(l1)
	}
	return head
}

func addone(l *ListNode){
	if l.Next==nil{
		l.Next=&ListNode{1,nil}
	}else{
		l.Next.Val+=1
		if l.Next.Val>9{
			l.Next.Val-=10
			addone(l.Next)
		}
	}

}