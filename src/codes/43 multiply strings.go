package codes

import (
	"strconv"
)

func Multiply(num1 string, num2 string) string {
	mul:=func(s1 string,n2 byte)string{


		n2-='0'
		if n2==0{return "0"}
		if n2==1{return s1}
		n1:=[]byte(s1)
		t:=byte(0)
		for i:=len(n1)-1;i>=0;i--{
			tmp:=(n1[i]-'0')*n2+t
			n1[i]=tmp%10+'0'
			t=tmp/10
		}
		if t!=0{return strconv.Itoa(int(t))+string(n1)}
		return string(n1)
	}
	l1:=len(num1)
	l2:=len(num2)

	n1:=[]byte(num1)
	n2:=[]byte(num2)
	if l2>l1{
		l1,l2=l2,l1
		n1,n2=n2,n1
	}
	s:="0"
	for i:=l2-1;i>=0;i--{
		tmp:=make([]byte,l2-1-i)
		for i:=range tmp{
			tmp[i]='0'
		}


		s=addStrings(s,mul(string(append(n1,tmp...)),n2[i]))

	}
	return s
}

func addStrings(num1 string, num2 string) string {
	//fmt.Println(num1)
	//fmt.Println(num2)
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