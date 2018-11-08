package codes

func postorderTraversal(root *TreeNode) []int {
	if root==nil{return []int{}}

	ans:=postorderTraversal(root.Left)
	ans=append(ans,postorderTraversal(root.Right)...)
	ans=append(ans,root.Val)
	return ans
}
