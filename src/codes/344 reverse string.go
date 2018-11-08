package codes

func reverseString(s string) string {
	l:=len(s)
	if l<2{return s}
	b:=[]byte(s)
	start:=0
	end:=l-1
	for start<end{
		b[start],b[end]=b[end],b[start]
		start++
		end--
	}
	return string(b)
}
