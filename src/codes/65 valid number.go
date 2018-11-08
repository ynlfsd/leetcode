package codes

import (
	"fmt"
	"strconv"
	"strings"
)

func IsNumber(s string) bool {
	s=strings.TrimSpace(s)
	l:=len(s)
	mark:=false
	if l==0{return false}
	if s[0]=='.'{
		s=s[1:]
		mark=true
		if l-1==0{return false}
		if s[0]<'0'||s[0]>'9'{return false}
	}else if s[l-1]=='.'{
		s=s[:l-1]
		mark=true
		if l-1==0{return false}
		if s[0]<'0'||s[0]>'9'{return false}
	}else{
		if strings.Index(s,"e")!=strings.LastIndex(s,"e"){return false}
		if strings.Index(s,"e")==0||strings.Index(s,"e")==len(s)-1{return false}
		s=strings.Replace(s,"e","",1)
	}
	if !mark{
		s=strings.Replace(s,".","",1)
	}

	fmt.Println(s)

	_,ok:=strconv.Atoi(s)
	fmt.Println(ok)
	return ok==nil
}
