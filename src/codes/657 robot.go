package codes

func judgeCircle(moves string) bool {
	ud:=0

	rl:=0

	for _,i:=range moves{
		switch i{
		case 'R':rl++
		case 'L':rl--
		case 'U':ud++
		case 'D':ud--
		}
	}
	if ud==0&&rl==0{
		return true
	}
	return false
}
