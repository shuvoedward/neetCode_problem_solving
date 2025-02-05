package arrayhashing

import "fmt"
func twoSum(nums []int, target int) []int {
	for i := 0; i<len(nums); i++{
		for j:=i+1; j < len(nums); j++{

			if nums[i] + nums[j] == target {
				return []int{i, j}
			}
		}
	}
	return []int{}
}

func twoPass(nums []int, target int) []int{
	indices := make(map[int]int)

	for i, n := range(nums){
		indices[n] = i;
	}

	for i, n := range(nums){
		diff := target - n;

		if j, found := indices[diff]; found && j!= i {
			return []int{i,j}
		}
	}

	return []int{}
}

func onePass(nums []int, target int) []int{
	indices := make(map[int]int)

	for i, n := range(nums){
		diff := target - n

		if j, found := indices[diff]; found {
			return []int{j,i}
		}
		indices[n] = i
	}

	return []int{}
}

func main(){
	nums := []int{3, 4, 5, 7}
	target:= 7
	
	// result := twoSum(nums, target)
	// fmt.Println(result)
	// fmt.Println(twoPass(nums, target))
	fmt.Println(onePass(nums, target))
}