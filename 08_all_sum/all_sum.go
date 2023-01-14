package main

func allSum(numbers ...[]int) []int {
	all_sum := []int{}
	for _, nums := range numbers {
		all_sum = append(all_sum, sum(nums))
	}
	return all_sum
}
