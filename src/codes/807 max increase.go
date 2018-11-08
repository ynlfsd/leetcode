package codes

func maxIncreaseKeepingSkyline(grid [][]int) int {
	v:=len(grid)
	h:=len(grid[0])
	maxv:=make([]int,v)
	maxh:=make([]int,h)
	for i,k:=range grid{
		for _,j:=range k{
			if maxv[i]<j{maxv[i]=j}
		}
	}
	for i:=0;i<h;i++{
		for j:=0;j<v;j++{
			if maxh[i]<grid[j][i]{maxh[i]=grid[j][i]}
		}
	}
	ans:=0
	for i,k:=range grid{
		for s,j:=range k{
			ans+=min(maxv[i],maxh[s])-j
		}
	}
	return ans


}
func mmin(a,b int)int{
	if a>b{return b}
	return a
}