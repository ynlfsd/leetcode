package codes

func inorderTraversal(root *TreeNode) []int {
	if root==nil{return []int{}}
	ans:=[]int{}
	ans=append(ans,inorderTraversal(root.Left)...)
	ans=append(ans,root.Val)
	ans=append(ans,inorderTraversal(root.Right)...)
	return ans
}
