package codes

import (
	"math"
	"strconv"
	"strings"
)

func myAtoi(str string) int {

	str=strings.TrimSpace(str)
	if len(str)==0{return 0}
	var ans int
	switch str[0] {
	case '-':ans=matio(str[1:])*(-1)
	case '+':ans=matio(str[1:])
	default:
		ans=matio(str)
	}

	if ans>math.MaxInt32{return math.MaxInt32}
	if ans<math.MinInt32{return math.MinInt32}

	return ans
}
func matio(s string)int {
	i:=0
	for ;i<len(s);i++{
		if s[i]>'9'||s[i]<'0'{
			break
		}
	}
	if i==0{return 0}
	ans,_:=strconv.Atoi(s[:i])
	return ans
}