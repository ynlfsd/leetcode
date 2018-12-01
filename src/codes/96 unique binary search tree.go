package codes

func NumTrees(n int) int {
	if n==0||n==1{return 1}

	m:=make([]int,n+1)
	m[0]=1
	m[1]=1

	for i:=range m{
		if i<2{continue}
		tmp:=0
		for j:=i-1;j>=0;j--{
			tmp+=m[j]*m[i-1-j]
		}
		m[i]=tmp
	}
	return m[n]
}
