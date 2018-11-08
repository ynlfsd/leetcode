
package codes

type StockSpanner struct {
	 price []int
	 i int
     max int
}


func Constructor() StockSpanner {
	var s StockSpanner
	s.i=0
	s.max=0
	s.price=make([]int,10000)
	return s
}


func (this *StockSpanner) Next(price int) int {
	if this.i==0{
		this.max=price
		this.price[this.i]=price
		this.i++
		return 1
	}
	this.price[this.i]=price
	if this.max<price{
		this.i++
		this.max=price
		return this.i
	}
	this.i++

	return findmax(this.price,this.i)
}
func findmax(s []int,i int) int{
	count:=1
	j:=i-1
	for ;i>=2;i--{
		if s[j]<s[i-2]{
			break
		}
		count++
	}


	return count
}