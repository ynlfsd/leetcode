package codes

func MaxProfit3(prices []int) int {
	l:=len(prices)
	if l<2{return 0}
	max:=make([]int,l)
	min:=prices[0]
	for i:=1;i<l;i++{
		max[i]=find2max(max[i-1],prices[i]-min)
		if min>prices[i]{min=prices[i]}
	}
	mi:=make([]int,l)
	m:=prices[l-1]
	ans:=max[l-1]
	for i:=l-2;i>=0;i--{
		mi[i]=find2max(mi[i+1],m-prices[i])
		ans=find2max(ans,max[i]+mi[i])
		if m<prices[i]{m=prices[i]}
	}



	return ans

}

func find2max(a,b  int)int{
	if a<b{
		return b
	}

	return a
}