package codes

import (
	"strings"
	"fmt"
)

func IsPalindrome(s string) bool {
	if len(s)==0||len(s)==1 {return true}
	s=strings.ToLower(s)

	k:=len(s)-1
	for i:=0;i<len(s);i++{
		for s[i]<'a'{
			i++
			if i>len(s)-1{
				fmt.Println("here")
				return true
			}
		}

		for s[k]<'a'{
			k--
		}

		if s[i]!=s[k]{

			return false
		}
		k--
		if i>k{
			break
		}
	}

	return true
}
func isdigit(b byte)bool{
	if b>='0'&&b<='9'{
		return true
	}
	if b>='a'&&b<='z'{
		return true
	}
	return false
}