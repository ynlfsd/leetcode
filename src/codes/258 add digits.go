package codes

func addDigits(num int) int {
	if num<10{return num}
	tmp:=0
	for num>=0{
		tmp+=num%10
		num/=10
	}
	return addDigits(tmp)
}
