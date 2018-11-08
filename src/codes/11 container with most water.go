package codes

func maxArea(height []int) int {
	ans:=make([]int,len(height))
	for i:=1;i<len(height);i++{
		ans[i]=findmax(ans[i-1],calc(height[:i+1]))
	}

	return ans[len(ans)-1]
}
func calc(h []int)int{
	ans:=0
	pos:=len(h)-1
	for i:=0;i<pos;i++{
		var tmp int
		if h[pos]<h[i]{
			tmp=h[pos]
		}else{
			tmp=h[i]
		}
		ans=findmax(ans,tmp*(pos-i))

	}

	return ans
}
func findmax(a,b int)int{
	if a>b{
		return a
	}
	return b
}

