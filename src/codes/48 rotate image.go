package codes

func rotate(matrix [][]int)  {
	l:=len(matrix)
	if l<2 {return}
	for i:=0;i<l-1;i++{
		for j:=i+1;j<l;j++{
			if matrix[i][j]!=matrix[j][i]{
				matrix[i][j],matrix[j][i]=matrix[j][i],matrix[i][j]
			}

		}
	}
	for i:=0;i<l;i++{
		for j:=0;j<l/2;j++{
			if matrix[i][j]!=matrix[i][l-1-j]{
				matrix[i][j],matrix[i][l-1-j]=matrix[i][l-1-j],matrix[i][j]
			}

		}
	}
}