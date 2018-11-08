package codes

func AddStrings(num1 string, num2 string) string {
	l1:=len(num1)
	l2:=len(num2)

	n1:=[]byte(num1)
	n2:=[]byte(num2)
	if l2>l1{
		l1,l2=l2,l1
		n1,n2=n2,n1
	}
	var t byte =0
	i,j:=l1-1,l2-1
	for j>=0{
		tmp:=n1[i]-'0'+n2[j]-'0'+t
		t=0
		if tmp>9{
			n1[i]=tmp-10+'0'
			t=1
		}else{
			n1[i]=tmp+'0'
		}
		i--
		j--
	}
	if t!=0{
		for i>=0{
			tmp:=n1[i]-'0'+t
			t=0
			if tmp>9{
				n1[i]=tmp-10+'0'
				t=1
			}else{

				n1[i]=tmp+'0'
				break
			}
			i--
		}
	}
	if t==1{return "1"+string(n1)}
	return string(n1)

}
