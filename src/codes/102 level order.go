package codes

func levelOrder(root *TreeNode) [][]int {
	if root==nil {return nil}

	l:=levelOrder(root.Left)
	r:=levelOrder(root.Right)
	ll:=len(l)
	rr:=len(r)

	if ll>=rr{
		for i:=0;i<ll&&i<rr;i++{
			l[i]=append(l[i],r[i]...)
		}
		l=append([][]int{{root.Val}},l...)
		return l
	}else{
		for i:=0;i<ll&&i<rr;i++{
			r[i]=append(l[i],r[i]...)
		}
		r=append([][]int{{root.Val}},r...)
		return r
	}
}
