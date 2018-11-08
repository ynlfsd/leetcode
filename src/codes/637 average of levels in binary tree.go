package codes
type TreeNode struct {
	    Val int
	    Left *TreeNode
	    Right *TreeNode
}
func averageOfLevels(root *TreeNode) []float64 {
	ret:=[]float64{}
	counts:=[]int{}
	ret,counts=add(root,0,ret,counts)
	l:=len(ret)

	for i:=0;i<l;i++{
		ret[i]/=float64(l)
	}
	return ret
}
func add(tr *TreeNode,level int,ret []float64,counts []int)([]float64,[]int){
	if len(ret)<=level{
		ret=append(ret,0)
		counts=append(counts,0)
	}
	ret[level]+=float64(tr.Val)
	counts[level]+=1
	if tr.Left!=nil{
		r,c:=add(tr.Left,level+1,ret,counts)
		ret=append(ret,r...)
		counts=append(counts,c...)
	}
	if tr.Right!=nil{
		r,c:=add(tr.Left,level+1,ret,counts)
		ret=append(ret,r...)
		counts=append(counts,c...)
	}
	return ret,counts
}