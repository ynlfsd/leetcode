package codes

func countBits(num int) []int {
	ans:=[]int{0,1}
	if num<2{return ans[:num+1]}
	i:=2
	for{
		ans=append(ans,add1(ans)...)
		if i*2>num{break}
		i*=2
	}

	return ans[:num+1]
}
func add1(n []int)[]int{
	tmp:=make([]int,len(n))
	for i:=0;i<len(n);i++{
		tmp[i]=n[i]+1
	}
	return tmp
}

