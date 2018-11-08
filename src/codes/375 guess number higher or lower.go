package codes

import "fmt"

func GetMoneyAmount(n int) int {
	if n==1{return 0}
	if n==2{return 1}
	ans:=make([][]int,n+1)
	for i:=range ans{
		ans[i]=make([]int,n+1)
	}

	for i:=1;i<n;i++{
		for j:=1;j<n+1;j++{
			if j+i>n{break}
			ma:=j+ans[j+1][j+i]
			for k:=j+1;k<j+i;k++{
				tmp:=max(k+ans[j][k-1],k+ans[k+1][j+i])
				if ma>tmp{ma=tmp}
			}
			ans[j][j+i]=ma
		}
	}
	for _,i:=range ans{
		fmt.Println(i)
	}

	return ans[1][n]



}
func max(a,b int)int{
	if a>b{return a}
	return b
}