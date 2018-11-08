package codes

func mirrorReflection(p int, q int) int {
	if q==p{return 1}
	if q==0{return 0}
	t:=q
	for q<<1%t!=0{
		q=q<<1%t

	}
	q=q<<1%t

	if q==0{return 0}
	return 2
}
