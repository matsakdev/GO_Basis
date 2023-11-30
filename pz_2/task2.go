package main

import (
	"fmt"
	"math"
)

type Person struct {
	Name    string
	Address string
	Age     int
}

type Worker struct {
	Person
	Position   string
	Experience int
	Salary     int
}

type WorkersArray struct {
	Workers []Worker
}

type PersonsArray struct {
	Persons []Person
}

func printPersonInfo(p Person) {
	fmt.Printf("Прізвище: %s\nАдреса: %s\nВік: %d\n", p.Name, p.Address, p.Age)
}

func printWorkerInfo(w Worker) {
	fmt.Printf("Прізвище: %s\nАдреса: %s\nВік: %d\n"+
		"Позиція: %s\nДосвід: %d\nЗарплата: %d\n", w.Name, w.Address, w.Age, w.Position, w.Experience, w.Salary)
}

func findMaxSalaryForOddElements(workers ...Worker) int {
	maxVal := math.MinInt64
	for i, worker := range workers {
		if i%2 != 0 && worker.Salary > maxVal {
			maxVal = worker.Salary
		}
	}
	return maxVal
}

func (persons *PersonsArray) getSortedArray() {
	personsArray := persons.Persons
	for i := 0; i < len(personsArray); i++ {
		minIndex := i
		for j := i + 1; j < len(personsArray); j++ {
			if (personsArray)[j].Age < (personsArray)[minIndex].Age {
				minIndex = j
			}
		}
		personsArray[i], personsArray[minIndex] = personsArray[minIndex], personsArray[i]
	}
}

func (workers *WorkersArray) getSortedArray() {
	workersArray := workers.Workers
	for i := 0; i < len(workersArray); i++ {
		minIndex := i
		for j := i + 1; j < len(workersArray); j++ {
			if (workersArray)[j].Salary < (workersArray)[minIndex].Salary {
				minIndex = j
			}
		}
		workersArray[i], workersArray[minIndex] = workersArray[minIndex], workersArray[i]
	}
}

type RootSolver interface {
	Solve(a, b, epsilon float64) float64
}

type MathInterface interface {
	SquareFunction(x float64) float64
}

type DegreeStruct struct {
	degree float64
}

type LineSegmentFunctionStruct struct {
	point1 float64
}

func halfDivisionMethod(a, b, epsilon float64, function func(float64) float64) float64 {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Panic occurred:", r)
			return
		}
	}()
	for math.Abs(b-a) > epsilon {
		c := (a + b) / 2
		if function(c)*function(a) < 0 {
			b = c
		} else {
			a = c
		}
	}
	return (a + b) / 2
}

func (function DegreeStruct) SquareFunction(x float64) float64 {
	x = math.Pow(x, function.degree)
	return 0.25*x*x*x + x - 1.2502
}

func (function LineSegmentFunctionStruct) SquareFunction(x float64) float64 {
	x = math.Abs(function.point1 - x)
	return 0.25*x*x*x + x - 1.2502
}

func task2() {
	person1 := Person{
		Name:    "ss ff",
		Age:     19,
		Address: "Address 1",
	}
	person2 := Person{
		Name:    "ss ff",
		Age:     23,
		Address: "Address 2",
	}
	person3 := Person{
		Name:    "nii oao",
		Age:     5,
		Address: "Address 3",
	}

	persons := PersonsArray{
		Persons: []Person{
			person1, person2, person3,
		},
	}

	workers := WorkersArray{
		Workers: []Worker{
			{
				Person:     person1,
				Position:   "a1",
				Experience: 5,
				Salary:     1000,
			},
			{
				Person:     person2,
				Position:   "b2",
				Experience: 19,
				Salary:     1950,
			},
			{
				Person:     person3,
				Position:   "c4",
				Experience: 8,
				Salary:     2900,
			},
		},
	}

	fmt.Println("\nPerson 1 info:")
	printPersonInfo(person1)
	fmt.Println("\nWorker 3 info:")
	printWorkerInfo(workers.Workers[2])

	fmt.Println("\nPersons before sorting:")
	fmt.Println(persons)

	persons.getSortedArray()
	fmt.Println("\nSorted Persons (by Age):")
	fmt.Println(persons)

	fmt.Println("\nWorkers before sorting:")
	fmt.Println(workers)

	workers.getSortedArray()
	fmt.Println("\nSorted Workers (by Salary):")
	fmt.Println(workers)

	fmt.Println("\nMax Salary For Odd Workers:")
	maxSalary := findMaxSalaryForOddElements(workers.Workers[0], workers.Workers[1], workers.Workers[2], workers.Workers[2])
	fmt.Println(maxSalary)

	a := 1.5
	b := 0.9
	epsilon := 0.0001

	basicFunc := DegreeStruct{degree: 1.0}
	lineSegmentFunc := LineSegmentFunctionStruct{point1: 1.2}

	var i MathInterface

	i = &basicFunc

	fmt.Println("\nHalf division method for degree function (a || b)^DegreeStruct.degree")
	result := halfDivisionMethod(a, b, epsilon, i.SquareFunction)
	fmt.Printf("\n%.5f\n", result)

	fmt.Println("\nHalf division method for line segment | (a || b) - LineSegmentFunctionStruct.point1 |")
	i = &lineSegmentFunc
	result = halfDivisionMethod(a, b, epsilon, i.SquareFunction)
	fmt.Printf("\n%.5f\n", result)
}
