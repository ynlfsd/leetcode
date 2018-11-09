package codes

func maxProfit(prices []int) int {
	l:=len(prices)
	if l<2{return 0}
	buy:=prices[l-1]
	max:=0
	for i:=l-2;i>=0;i--{
		if buy<prices[i]{
			buy=prices[i]
			continue
			}
		if max<buy-prices[i]{max=buy-prices[i]}
	}
	return max
}