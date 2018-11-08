package codes

func sortArrayByParityII(A []int) []int {
	l:=len(A)
	if l==0{return A}
	even:=0
	odd:=1
	for ;even<l;even+=2{
		if A[even]%2!=0{
			for ;odd<l;odd+=2{
				if A[odd]%2==0{
					A[even],A[odd]=A[odd],A[even]
					break
				}
			}
		}
	}
	return A

}
