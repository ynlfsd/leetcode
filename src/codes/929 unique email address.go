package codes

import (
	"bytes"
	"strings"
)

func numUniqueEmails(emails []string) int {
	m:=make(map[string]int)
	ans:=0
	for _,k:=range emails{
		ans+=checkEmails(k,m)
	}
	return ans
}
func checkEmails(s string,m map[string]int)int {

	at := strings.Index(s, "@")
	plus := strings.Index(s[:at], "+")
	if plus==-1{plus=at}
	sbyte:=[]byte(s[:plus])
	tmp:=bytes.Split(sbyte,[]byte{'.'})
	sbyte=nil
	for i:=0;i<len(tmp);i++{
		sbyte=append(sbyte,tmp[i]...)
	}
	sbyte=append(sbyte,[]byte(s[at:])...)
	if _,ok:=m[string(sbyte)];ok{
		return 0
	}else{
		m[string(sbyte)]=1
		return 1
	}

}