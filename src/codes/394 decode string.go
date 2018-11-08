package codes
import "strings"

func DecodeString(s string) string {
	var num,start []int
	var end int
	i:=0
	//ans:=[]string{}
	for {

		switch {
		case s[i]>='0'&&s[i]<='9':
			num=append(num,int(s[i]-'0'))
		case s[i]=='[':
			start=append(start,i+1)
		case s[i]==']':
			end= i
			ss:=start[len(start)-1]
			start=start[:len(start)-1]
			nn:=num[len(num)-1]
			num=num[:len(num)-1]
			i=len(strings.Repeat(s[ss:end],nn))+ss-4
			s=s[:ss-2]+strings.Repeat(s[ss:end],nn)+s[end+1:]

			//ans=append(ans,strings.Repeat(s[ss:end],nn))
		}
		i++
		if i>=len(s){
			break
		}
	}
	return s
}
func find(s string,i int)int{
	for ;i<len(s);i++{
		if s[i]=='['{
			return i
		}
	}
	return 0
}