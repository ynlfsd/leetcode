package codes

import "sort"

func combinationSum2(candidates []int, target int) [][]int {

	l:=len(candidates)
	if l<1{return [][]int{}}
	sort.Ints(candidates)



	return combination2(candidates,target)
}
func combination2(candidates []int, target int)[][]int{
	ans:=[][]int{}
	l:=len(candidates)
	if l<1{return ans}

	t:=0
	for i:=0;i<l;i++{
		if i!=0&&candidates[i]==candidates[i-1]{
			continue
		}
		if candidates[i]==target{
			ans=append(ans,[]int{target})
			break
		}else if candidates[i]<target{

			ans=append(ans,combination2(candidates[i+1:],target-candidates[i])...)
			for ;t<len(ans);t++{
				ans[t]=append(ans[t],candidates[i])
			}
		}else{break}


	}

	return ans


}