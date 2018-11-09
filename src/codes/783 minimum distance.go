package codes

func MinDiffInBST(root *TreeNode) int {
	loop :=[]int{}
	loop=dfs(root,loop)
	println(len(loop))
	//min:=loop[1]-loop[0]
	//for i:=1;i<len(loop)-1;i++{
	//	if min>loop[i+1]-loop[i]{
	//		min=loop[i+1]-loop[i]
	//	}
	//}
	return 1
}

func dfs(leaf *TreeNode,loop []int )[]int{

	if leaf.Left!=nil{
		print("left")
		println(len(loop))
		loop=dfs(leaf.Left,loop)
		print("left")
		println(len(loop))

	}
	loop=append(loop,leaf.Val)
	println(len(loop))
	if leaf.Right!=nil{
		loop=dfs(leaf.Right,loop)
	}
return loop
}

