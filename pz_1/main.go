package main

import (
	"fmt"
)

func main() {
	// Task 1
	fmt.Printf("[Task 1]\n")
	equalTriangleResult := isTriangleEquilateral(1, 1, 1)
	notEqualTriangleResult := isTriangleEquilateral(1, 2, 1)
	fmt.Printf("%t\n%t\n", equalTriangleResult, notEqualTriangleResult)

	// Task 2
	fmt.Printf("[Task 2]\n")
	point1 := findAxisOnWhichPointLies(0.0, 0.5)
	point2 := findAxisOnWhichPointLies(0.0, 0.0)
	point3 := findAxisOnWhichPointLies(0.9, 0.0)
	point4 := findAxisOnWhichPointLies(0.9, 0.8)
	fmt.Println(point1)
	fmt.Println(point2)
	fmt.Println(point3)
	fmt.Println(point4)

	// Task 3
	fmt.Printf("[Task 3]\n")
	grade1 := getGradeName(1)
	grade2 := getGradeName(5)
	grade3 := getGradeName(7)
	grade4 := getGradeName(0)

	fmt.Println(grade1)
	fmt.Println(grade2)
	fmt.Println(grade3)
	fmt.Println(grade4)

	// Task 4
	fmt.Printf("[Task 4]\n")
	findNumbersWithDigitsSum(3, 5)

	// Task 5
	fmt.Printf("[Task 5]\n")

	matrix := [][]int{
		{1, -2, 3},
		{-5, 4, 2},
		{7, -8, 9},
	}

	product := findProductOfPositiveNumbers(matrix)
	fmt.Printf("Добуток додатних чисел другого рядка: %d\n", product)

	// Task 6
	fmt.Printf("[Task 6]\n")
	describeStudents()

}
