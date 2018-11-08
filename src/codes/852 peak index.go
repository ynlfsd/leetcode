package codes

func peakIndexInMountainArray(A []int) int {
	l:=len(A)
	l2:=l/2
	if A[l2]>A[l2+1]&&A[l2]>A[l2-1]{
		return l2
	}else if A[l2]>A[l2-1]{
		return peakIndexInMountainArray(A[l2+1:])
	}else{
		return peakIndexInMountainArray(A[:l2])
	}
}
