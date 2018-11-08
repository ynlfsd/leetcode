package codes

import (
	"math"
	"strconv"
)

func Reverse(x int) int {
	biao:=1
	if x<0{
		x=-x
		biao=-1
	}
	s:=strconv.Itoa(x)

	l:=len(s)
	b:=make([]byte,l)
	for i:=l-1;i>=0;i--{
		b[l-1-i]=s[i]
	}

	n,_:=strconv.Atoi(string(b))
	if n>math.MaxInt32||n<math.MinInt32{return 0}



	return n*biao
}
