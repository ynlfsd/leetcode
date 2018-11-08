package codes

import "fmt"

func GenerateMatrix(n int) [][]int {

	m:=make([][]int,n)
	for t:=0;t<n;t++{
		m[t]=make([]int,n)
	}
	num:=1
	i,j:=0,0
	si,sj:=0,0
	ei,ej:=n,n
	lr,ud:=true,true
	for {

		if lr&&ud{
			for ;j<ej;j++{
				fmt.Printf("i %d , j %d\n",i,j)

				m[i][j]=num
				num++
			}
			j--
			i++
			si++
			lr=!lr
		}else if  !lr&&ud{
			for ;i<ei;i++{
				fmt.Printf("i %d , j %d\n",i,j)
				m[i][j]=num
				num++
			}
			i--
			j--
			ej--
			ud=!ud
		}else if !lr&&!ud{
			for ;j>=sj;j--{
				fmt.Printf("i %d , j %d\n",i,j)
				m[i][j]=num
				num++
			}
			j++
			i--
			ei--

			lr=!lr
		}else {
			for ;i>=si;i--{
				fmt.Printf("i %d , j %d\n",i,j)
				m[i][j]=num
				num++
			}
			i++
			j++
			sj++
			ud=!ud
		}
	if si==ei||sj==ej{break}

	}
	return m
}
