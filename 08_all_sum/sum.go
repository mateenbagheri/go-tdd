package main

func sum(numbers []int) int {
	var totalSum int
	for _, num := range numbers {
		totalSum += num
	}
	return totalSum
}
