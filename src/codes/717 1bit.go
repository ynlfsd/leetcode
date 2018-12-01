package codes

func isOneBitCharacter(bits []int) bool {
	l:=len(bits)
	if l==0{return false}
	if l==1{return true}
	ans:=false
	for i:=0;i<l;{
		if bits[i]==0{
			i++
			ans=true
		}else{
			i+=2
			ans=false
		}
	}
	return ans
}