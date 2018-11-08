package codes

import "fmt"

func PredictPartyVictory(senate string) string {
	r:=0

	d:=0

	m:=map[byte]string{'R':"Radiant",'D':"Dire"}
	l:=len(senate)
	if l==1{return m[senate[0]]}
	b:=[]byte(senate)
	for j:=0;j<3;j++{
		for i:=0;i<l;i++{
			if b[i]=='R'{
				if d==0{
					r++

				}else{
					d--
					b[i]=' '
				}
			}else if b[i]=='D'{
				if r==0{
					d++

				}else{
					r--
					b[i]=' '
				}
			}
		}
	}

fmt.Println(r)
	fmt.Println(d)
	if r==0&&d==0{return m[senate[0]]}
	if r>0{return "Radiant"}else{return "Dire"}


}


