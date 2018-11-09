package codes

func AddBinary(a string, b string) string {
	la:=len(a)
	lb:=len(b)
	ba:=[]byte(a)
	bb:=[]byte(b)
	if la<lb{
		la,lb=lb,la
		ba,bb=bb,ba
	}
	t:=byte(0)
	i,j:=la-1,lb-1
	for j>=0{
		tmp:=ba[i]+bb[j]-'0'*2+t
		t=0
		if tmp>=2{
			ba[i]=tmp-2+'0'
			t=1
		}else{
			ba[i]=tmp+'0'
		}
		i--
		j--
	}
	if t!=0{
		for i>=0{
			tmp:=ba[i]+t-'0'
			t=0
			if tmp==2{
				ba[i]=0
				t=1
			}else{
				ba[i]=tmp+'0'
			}
			if t==0{break}
			i--
		}

	}
	if t!=0{
		return "1"+string(ba)
	}else{
		return string(ba)
	}
}