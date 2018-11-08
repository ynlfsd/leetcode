package codes

func mergeKLists(lists []*ListNode) *ListNode {
		if len(lists)==2{return mergeTwoLists(lists[0],lists[1])}
		for i:=0;i<len(lists)-1;i++{
			lists[i+1]=mergeTwoLists(lists[i],lists[i+1])
		}
		return lists[len(lists)-1]
}