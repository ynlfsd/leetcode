package codes

import (
	"fmt"
	"math"
)

func GrayCode(n int) []int {
	if n==0{return []int{0}}
	m:=make([]int,int(math.Pow(2,float64(n+1))-1))
	t:=1
	s:=2

	for s-1<len(m){
		mark:=false
		for i:=0;i<s;i+=2{
			if mark{
				m[s-1+i]=t
				m[s-1+i+1]=0
				mark=false
			}else{
				m[s-1+i]=0
				m[s-1+i+1]=t
				mark=true
			}
		}
		t*=2
		s*=2
	}
	fmt.Println(m)
	ans:=dfs80(m,0)
	return ans

}
func dfs80(a []int,i int)[]int{
	t:=(i+1)*2
	if t>len(a){return []int{a[i]}}
	l:=dfs80(a,t-1)
	r:=dfs80(a,t)
	for j:=range l{
		l[j]+=a[i]
		r[j]+=a[i]
	}
	return append(l,r...)

}
