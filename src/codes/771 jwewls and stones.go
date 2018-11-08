package codes

func numJewelsInStones(J string, S string) int {
	ans:=0
	m:=make(map[rune]int)
	for _,i:=range J{
		m[i]=1
	}
	for _,i:=range S{
		if v,ok:=m[i];ok{
			ans+=v
		}
	}
	return ans
}
