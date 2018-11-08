package codes


func longestPalindrome(s string) string {

	l:=len(s)
	if l<2{return s}
	end:=l
	for end>1{
		for i:=0;i<l;i++{
			if i+end>l{break}
			if isPalin(s[i:i+end]){
				return s[i:i+end]
			}

		}
		end--
	}


	return s[0:1]
}
func isPalin(s string)bool{
	l:=len(s)
	for i:=0;i<l/2;i++{
		if s[i]!=s[l-i-1]{
			return false
		}
	}


	return true
}