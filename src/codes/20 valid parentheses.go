package codes

func IsValid(s string) bool {
	if len(s)%2!=0{return false}
	b:=[]byte{}
	for _,v:=range []byte(s){
		switch v{
		case '(','{','[':
			b=append(b,v)
		case ')','}',']':
			if len(b)==0{
				return false
			}
			if !(b[len(b)-1]==v-1||b[len(b)-1]==v-2){
				return false
			}
			b=b[:len(b)-1]
		}
	}
	if len(b)>0{
		return false
	}


	return true
}
