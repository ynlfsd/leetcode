package codes


func isMonotonic(A []int) bool {
	for i:=0;i<len(A)-1;i++{
		if A[i]==A[i+1]{continue}
		if A[i+1]>A[i]{
			return up(A[i+1:])
		}else{
			return down(A[i+1:])
		}
	}




	return true
}
func up(A []int)bool{
	for i:=0;i<len(A)-1;i++{
		if A[i]>A[i+1]{return false}
	}

	return true
}

func down(A []int)bool{
	for i:=0;i<len(A)-1;i++{
		if A[i]<A[i+1]{return false}
	}

	return true
}