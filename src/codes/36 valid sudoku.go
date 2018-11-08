package codes

func IsValidSudoku(board [][]byte) bool {
	for i:=0;i<7;i+=3{
		for j:=0;j<7;j+=3{
			if !validsquare(board,i,j){
				return false
			}
		}
	}
	for i:=0;i<9;i++{
		for j:=0;j<9;j++{
			if board[i][j]=='.'{continue}
			if !validline(board,i,j){
				return false
			}
		}
	}
	return true
}
func validline(line [][]byte,i,j int)bool{
	for l,k:=range line[i]{
		if l==j{continue}
		if k==line[i][j]{return false}
	}
	for l:=0;l<9;l++{
		if l==i{continue}
		if line[l][j]==line[i][j]{return false}
	}
	return true
}
func validsquare(b [][]byte,h,v int)bool{
	m:=make(map[byte]bool)
	for i:=0;i<3;i++{
		for j:=0;j<3;j++{
			if b[h+i][v+j]=='.'{continue}
			if _,ok:=m[b[h+i][v+j]];!ok{
				m[b[h+i][v+j]]=true
			}else {return false}
		}
	}


	return true
}
