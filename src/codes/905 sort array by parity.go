package codes

func sortArrayByParity(A []int) []int {

	even:=[]int{}
	odd:=[]int{}
	for _,i:=range A{
		if i%2==0{
			even=append(even,i)
		}else{
			odd=append(odd,i)
		}
	}
	return append(even,odd...)

}
func index(A []int)int{
	for v,i:=range A{
		if i%2==0{
			return v
		}
	}
	return -1
}