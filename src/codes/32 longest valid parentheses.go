package codes

import "fmt"

func LongestValidParentheses(s string) int {
	l:=len(s)
	if l<2 {return 0}
	max:=0
	ret:=make([]int,l)
	tmp:=0
	for i,k:=range s{

		switch k{
		case '(':
			tmp++
		case ')':
			if tmp>0{
				if ret[i-1]==0{
					if i-2>0{
						ret[i]=ret[i-2]+2
					}else{ret[i]=2}


				}else {
					if i-ret[i-1]-2>0{
						ret[i]=ret[i-ret[i-1]-2]+ret[i-1]+2
					}else{ret[i]=ret[i-1]+2}


				}
				tmp--
			}
		}
		if max<ret[i]{
			max=ret[i]
		}
	}
	fmt.Println(ret)
	return  max
}

