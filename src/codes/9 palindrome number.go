package codes

func isPalindrome(x int) bool {
	if x<0||(x/10!=0&&x%10==0){return false}
	ans:=0
	for ans<x{
		ans=ans*10+x%10
		x/=10
	}


	return ans==x||x==ans/10
}
