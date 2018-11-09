package codes

import "fmt"

func MaxProfit2(prices []int) int {
	l:=len(prices)
	if l<1{return 0}
	max:=0
	buy:=prices[0]
	for i:=1;i<l;i++{
		if buy>prices[i]{
			buy=prices[i]
			continue
		}
		sell:=prices[i]
		j:=i+1
		fmt.Println(i)
		for ;j<l;j++{
			if sell>prices[j]{break}
			sell=prices[j]
		}
		i=j

		max+=sell-buy
		if i<l{

			buy=prices[i]
		}

	}
	return max
}
