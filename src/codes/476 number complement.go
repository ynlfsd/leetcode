package codes

func findComplement(num int) int {
	i:=1
	ans:=num
	for i<num{
		ans^=i
		i<<=1
	}

	return ans
}
