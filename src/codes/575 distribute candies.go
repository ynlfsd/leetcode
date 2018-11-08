package codes

func distributeCandies(candies []int) int {
	l:=len(candies)
	if l==2{return 1}
	m:=make(map[int]bool)
	sis:=1

	for i:=1;i<l;i++{
		if _,ok:=m[candies[i]];!ok{
			m[candies[i]]=true

			sis++
		}

		if sis>=l/2{
			return sis
		}
	}
	return sis
}
