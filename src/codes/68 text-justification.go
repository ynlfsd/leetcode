package codes

import (
	"strings"
)

func FullJustify(words []string, maxWidth int) []string {
	p,v:=findpos(words,maxWidth)
	s:=[]string{}
	for p!=0{
		tmp:=strings.Join(words[:p]," ")
		space:=strings.Repeat(" ",v)
		s= append(s, tmp+space+words[p])
		words=words[p+1:]
		p,v=findpos(words,maxWidth)
	}
	s= append(s, strings.Join(words," "))
	return s
}
func findpos(words []string,max int)(int,int){
	t:=0
	if len(words[0])==max{return 0,0}
	for p,w:=range words{
		if t+len(w)>=max{
			return p-1,max-t+1
		}else{
			t+=len(w)+1
		}
	}
	return 0,0
}
