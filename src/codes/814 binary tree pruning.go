package codes

func pruneTree(root *TreeNode) *TreeNode {
	if root==nil{return nil}
	//if root.Right==nil&&root.Left==nil{
	//	if root.Val==1{
	//		return root
	//	}else{
	//		return nil
	//	}
	//}
	root.Left=pruneTree(root.Left)
	root.Right=pruneTree(root.Right)
	if root.Right==nil&&root.Left==nil&&root.Val==0{return nil}
	return root
}
