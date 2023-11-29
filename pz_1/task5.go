package main

func findProductOfPositiveNumbers(arr [][]int) int {
	line := arr[1]
	product := 1
	for _, num := range line {
		if num > 0 {
			product *= num
		}
		if num == 0 {
			return 0
		}
	}
	return product
}
