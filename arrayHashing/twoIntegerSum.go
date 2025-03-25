package arrayhashing

func twoSum(nums []int, target int) []int {
	for i := 0; i < len(nums); i++ {
		for j := i + 1; j < len(nums); j++ {

			if nums[i]+nums[j] == target {
				return []int{i, j}
			}
		}
	}
	return []int{}
}

func twoPass(nums []int, target int) []int {
	indices := make(map[int]int)

	for i, n := range nums {
		indices[n] = i
	}

	for i, n := range nums {
		diff := target - n

		if j, found := indices[diff]; found && j != i {
			return []int{i, j}
		}
	}

	return []int{}
}

func onePass(nums []int, target int) []int {
	indices := make(map[int]int)

	for i, n := range nums {
		diff := target - n

		if j, found := indices[diff]; found {
			return []int{j, i}
		}
		indices[n] = i
	}

	return []int{}
}
