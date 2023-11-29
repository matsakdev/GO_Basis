package main

import (
	"fmt"
	"sort"
)

func (array OneDimensionalArray) Task1() OneDimensionalArray {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Panic occurred:", r)
			return
		}
	}()

	result := OneDimensionalArray{}
	for _, num := range array.arr {
		if !isFibonacci(num) {
			result.arr = append(result.arr, num)
		}
	}
	return result
}

func (array TwoDimensionalArray) Task1() TwoDimensionalArray {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Panic occurred:", r)
			return
		}
	}()

	cols := len(array.arr[0])

	for col := 0; col < cols; col++ {
		sort.Slice(array.arr, func(i, j int) bool {
			return array.arr[i][col] > array.arr[j][col]
		})
	}

	return array
}

func isFibonacci(num int) bool {
	a, b := 0, 1
	for b <= num {
		if b == num {
			return true
		}
		a, b = b, a+b
	}
	return false
}

type Array struct{}
type OneDimensionalArray struct {
	Array
	arr []int
}

type TwoDimensionalArray struct {
	Array
	arr [][]int
}

type OneDimensionalArrayInterface interface {
	Task1() OneDimensionalArray
}

type TwoDimensionalArrayInterface interface {
	Task1() TwoDimensionalArray
}

func (array Array) Task1() {
	fmt.Println("Пуста структура")
}

func task1Demonstration() {
	oneDimArray := OneDimensionalArray{
		arr: []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13, 14, 21},
	}
	fmt.Println("Одновимірний масив до видалення чисел Фібоначчі:", oneDimArray.arr)
	oneDimArray = oneDimArray.Task1()
	fmt.Println("Одновимірний масив після видалення чисел Фібоначчі:", oneDimArray.arr)

	twoDimArray := TwoDimensionalArray{
		arr: [][]int{
			{5, 4, 3, 2, 1},
			{10, 9, 8, 7, 6},
			{15, 14, 13, 12, 11},
		},
	}
	fmt.Println("Двовимірний масив до сортування стовпців за спаданням:")
	for _, row := range twoDimArray.arr {
		fmt.Println(row)
	}
	twoDimArray = twoDimArray.Task1()
	fmt.Println("Двовимірний масив після сортування стовпців за спаданням:")
	for _, row := range twoDimArray.arr {
		fmt.Println(row)
	}

	oneDimArray2 := OneDimensionalArray{
		arr: []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13, 14, 21},
	}

	twoDimArray2 := TwoDimensionalArray{
		arr: [][]int{
			{5, 4, 3, 2, 1},
			{10, 9, 8, 7, 6},
			{15, 14, 13, 12, 11},
		},
	}

	var odInterface OneDimensionalArrayInterface

	odInterface = oneDimArray2

	fmt.Println("[Interface] Одновимірний масив до видалення чисел Фібоначчі:", oneDimArray2)

	odInterface.Task1()

	fmt.Println("[Interface] Одновимірний масив після видалення чисел Фібоначчі:", oneDimArray2)

	var tdInterface TwoDimensionalArrayInterface

	tdInterface = twoDimArray2

	fmt.Println("[Interface] Двовимірний масив до сортування стовпців за спаданням:", tdInterface)

	tdInterface.Task1()

	fmt.Println("[Interface] Двовимірний масив після сортування стовпців за спаданням:", tdInterface)
}
