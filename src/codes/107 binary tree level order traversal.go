package codes
func levelOrderBottom(root *TreeNode) [][]int {
	if root==nil {return nil}

	l:=levelOrderBottom(root.Left)
	r:=levelOrderBottom(root.Right)
	ll:=len(l)
	rr:=len(r)
	if ll>=rr{
		for i:=0;ll-1-i>=0&&rr-1-i>=0;i++{
			l[ll-1-i]=append(l[ll-1-i],r[rr-1-i]...)
		}
		l=append(l,[]int{root.Val})
		return l
	}else{
		for i:=0;ll-1-i>=0&&rr-1-i>=0;i++{
			r[rr-1-i]=append(l[ll-1-i],r[rr-1-i]...)
		}
		r=append(r,[]int{root.Val})
		return r
	}


}
