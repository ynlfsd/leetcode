package codes

func reverseWords(s string) string {
	l:=len(s)
	if l<2{return s}
	b:=[]byte(s)
	start:=0
	for i,k:=range b{
		if k==' '{
			reverse(b[start:i])
			start=i+1
		}
	}
	reverse(b[start:l+1])
	return string(b)
}
func reverse(b []byte){
	i:=0
	j:=len(b)-1
	for i<j{
		b[i],b[j]=b[j],b[i]
		i++
		j--
	}
}
