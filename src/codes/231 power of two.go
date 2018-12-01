package codes

func isPowerOfTwo(n int) bool {
	if n<0{return false}
	if n==0||n==1{return true}
	for n%2==0{
		n/=2
	}
	if n==1{return true}
	return false
}
