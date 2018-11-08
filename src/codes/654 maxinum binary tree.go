package codes

func constructMaximumBinaryTree(nums []int) *TreeNode {
	if len(nums)==0{return nil}
	if len(nums)==1{return &TreeNode{Val:nums[0]}}
	i,max:=0,nums[0]
	for j,k:=range nums{
		if max<k{
			i=j
			max=k
		}
	}
	t:=&TreeNode{Val:max}
	t.Left=constructMaximumBinaryTree(nums[:i])
	t.Right=constructMaximumBinaryTree(nums[i+1:])
	return t
}
