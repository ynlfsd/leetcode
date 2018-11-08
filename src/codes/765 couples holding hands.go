package codes


func minSwapsCouples(row []int) int{

	count :=0
	for i:=0 ; i < len(row); i+=2{
		if row[i]%2==0{
			if row[i]+1 == row[i+1]{
				continue
			}
			swap(row,i+1)
			count++
		}else {
			if row[i]-1 == row[i+1]{
				continue
			}
			swap(row,i+1)
			count++
		}

	}
	return count
}
func swap(row []int, i int){
	var tmp int
	if row[i]%2==0{
		tmp = row[i]+1
	}else{
		tmp =row[i]-1
	}
	j:=i-1
	i++
	for ;i<len(row);i++{
		if tmp == row[i]{
			row[i]=row[j]
		}
	}
}

