package codes

func invertTree(root *TreeNode) *TreeNode {
	if root==nil{return root}
	root.Right,root.Left=root.Left,root.Right
	invertTree(root.Right)
	invertTree(root.Left)
	return root
}
