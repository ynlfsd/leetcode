package codes

import (
	"strings"
	"fmt"
)

func MostCommonWord(paragraph string, banned []string) string {
	m:=make(map[string]int)
	var max int=0
	var re string
	for _,s:= range banned{
		m[s]=-1
	}
	paragraph=strings.ToLower(paragraph)
	sp:=strings.Split(paragraph," ")
	for i:=0;i<len(sp);i++{
		sp[i]=change(sp[i])
		//fmt.Println(sp[i])
	}


	for _,s:=range sp{
		if k,ok:=m[s];ok{
			if k!=-1{m[s]=k+1}
		}else{
			m[s]=1
		}
		if max<m[s]{
			max=m[s]
			re=s
		}
	}
	//paragraph=fmt.Println(paragraph)
	return re
}
func change(s string)string{
	fmt.Println(len(s))
	if []rune(s[len(s)-1:])[0]<'a'{
		return s[:len(s)-1]
	}
	return s
}