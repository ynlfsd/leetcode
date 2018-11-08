package codes

func climbStairs(n int) int {
	ret:=make([]int,50)
	ret[0]=0
	ret[1]=1
	for i:=2;i<50;i++{
		ret[i]=ret[i-1]+ret[i-2]
	}
	return ret[n]
}
