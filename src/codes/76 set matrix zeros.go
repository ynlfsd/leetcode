package codes

func setZeroes(matrix [][]int)  {
	mi:=make(map[int]bool)
	mj:=make(map[int]bool)
	for i,ma:=range matrix{
		for j,m:=range ma{
			if m==0{
				mi[i]=true
				mj[j]=true
			}
		}
	}
	for k,_:=range mi{
		serZeroesH(matrix,k)
	}
	for k,_:=range mj{
		serZeroesV(matrix,k)
	}
}

func serZeroesH(matrix [][]int ,x int){
	for i:=0;i<len(matrix[x]);i++{
		matrix[x][i]=0
	}

}
func serZeroesV(matrix [][]int ,x int){
	for i:=0;i<len(matrix);i++{
		matrix[i][x]=0
	}

}