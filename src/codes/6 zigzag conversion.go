package codes

import "bytes"

func Convert(s string, numRows int) string {
	if numRows==1{return s}
	if len(s)<numRows{return s}
	var b bytes.Buffer
	//b:=[]byte{}

	for i:=0;i<numRows;i++{
		tmp:=i
		pos:=numRows*2-1-i*2-1
		switch tmp {
		case 0:
			for{
				b.WriteByte(s[tmp])

				tmp+=pos
				if tmp>=len(s){break}
			}
		case numRows-1:
			pos=numRows*2-2-pos
			for{
				b.WriteByte(s[tmp])

				tmp+=pos
				if tmp>=len(s){break}
			}
		default:
			for{
				b.WriteByte(s[tmp])

				tmp=tmp+pos
				pos=numRows*2-2-pos
				if tmp>=len(s){break}
			}
		}

	}
	return b.String()
}

