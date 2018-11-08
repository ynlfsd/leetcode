package codes

//func isMatch(s string, p string) bool {
//	ls:=len(s)
//	if ls==0 {return true}
//	if strings.Contains(p,".*"){return true}
//	if strings.Contains(p,s){return true}
//
//
//	return false
//}
//
//
//func ismatch(s string , p string)bool{
//	ls:=len(s)
//	if ls==1{return true}
//	lp:=len(p)
//
//	for i:=1;i<ls;i++{
//		if i>=lp{return false}
//		pnext:=i
//		for p[i]=='*'&&i+1<lp{
//			t:=1
//			for i+t<lp&&p[i+t]==p[i-1]{
//				pnext++
//			}
//		}
//		for p[i]=='*'&&s[i-1]==s[i]{
//			i++
//		}
//	}
//
//
//}