package codes

func PlusOne(digits []int) []int {
	l:=len(digits)
	t:=0
	for i:=l-1;i>=0;i--{
		digits[i]+=1+t
		t=0
		if digits[i]==10{
			digits[i]=0
			t=1
		}else{break}
	}
	if t==1{
		return append([]int{1},digits...)
	}
	return digits
}
