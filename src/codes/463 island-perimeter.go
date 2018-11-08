package codes

func islandPerimeter(grid [][]int) int {
	l:=len(grid)
	if l==0{return 0}
	w:=len(grid[0])
	ans:=0
	for i:=0;i<w;i++{
		for j:=0;j<l;j++{
			if j==0&&grid[i][j]==1{ans++}
			if j>0&&grid[i][j]==1&&grid[i][j-1]==0{ans++}
			if j==l-1&&grid[i][j]==1{ans++}
			if j<l-1&&grid[i][j]==1&&grid[i][j+1]==0{ans++}
			if i==0&&grid[i][j]==1{ans++}
			if i>0&&grid[i][j]==1&&grid[i-1][j]==0{ans++}
			if i==w-1&&grid[i][j]==1{ans++}
			if i<w-1&&grid[i][j]==1&&grid[i+1][j]==0{ans++}
		}
	}
	return ans
}
