package codes

func uniquePaths(m int, n int) int {
	if m==0||n==0{return 0}
	if m==1||n==1{return 1}
	var rec [100][100]int
	for i:=0;i<m;i++{
		rec[i][0]=1
	}
	for i:=0;i<n;i++{
		rec[0][i]=1
	}
	for i:=1;i<m;i++{
		for j:=1;j<n;j++{
			rec[i][j]=rec[i-1][j]+rec[i][j-1]
		}
	}
	return rec[m-1][n-1]
}
