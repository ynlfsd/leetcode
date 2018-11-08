package codes

func FairCandySwap(A []int, B []int) []int {
	sum:=0
	suma:=0
	sumb:=0
	var ans []int
	m:=make(map[int]int)
	for _,k:=range A{
		suma+=k
		m[k]=k
	}
	for _,k:=range B{
		sumb+=k
	}
	sum=(sumb-suma)/2



	for _,k:=range B{
		if _,ok:=m[k-sum];ok{
			ans=[]int{k-sum,k}

			break
		}
	}
	return ans
}
