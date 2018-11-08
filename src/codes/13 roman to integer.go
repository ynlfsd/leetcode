package codes

import "strings"

func romanToInt(s string) int {
	ans:=0
	m:=make(map[int32]int)
	m['I']=1
	m['V']=5
	m['X']=10
	m['L']=50
	m['C']=100
	m['D']=500
	m['M']=1000
	for _,i:=range s{
		ans+=m[i]
	}

	if strings.Contains(s,"IV")||strings.Contains(s,"IX"){
		ans-=2
	}

	if  strings.Contains(s,"XL")||strings.Contains(s,"XC"){
		ans-=20
	}


	if strings.Contains(s,"CD")||strings.Contains(s,"CM"){
		ans-=200
	}




	return ans
}