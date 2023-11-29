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
		"Позиція: %s\nДосвід: %d\nВік: %d\nЗарплата", w.Name, w.Address, w.Age, w.Position, w.Experience, w.Salary)
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

func halfDivisionMethod(a, b, epsilon float64) float64 {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Panic occurred:", r)
			return
		}
	}()
	for math.Abs(b-a) > epsilon {
		c := (a + b) / 2
		if f(c)*f(a) < 0 {
			b = c
		} else {
			a = c
		}
	}
	return (a + b) / 2
}

func f(x float64) float64 {
	return 0.25*x*x*x + x - 1.2502
}

func testMethods() {
	person1 := Person{
		Name:    "ss ff",
		Age:     19,
		Address: "pohui",
	}
	person2 := Person{
		Name:    "ss ff",
		Age:     23,
		Address: "pohui",
	}
	person3 := Person{
		Name:    "nii oao",
		Age:     5,
		Address: "pohui 3",
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

	persons.getSortedArray()

	workers.getSortedArray()

}

//func main2() {
//	// Ініціалізація масиву структур
//	people := []Person{
//		{Name: "Іванов", Address: "Київ", BirthDate: "01.01.1990"},
//		{Name: "Петров", Address: "Львів", BirthDate: "15.05.1985"},
//		{Name: "Сидоров", Address: "Одеса", BirthDate: "20.11.2000"},
//	}
//
//	// 1) Виклик функції з параметрами, що замовчуються
//	printPersonInfo(people[0], 33)
//
//	// 2) Виклик функції зі змінним числом параметрів
//	max := findMaxOdd(4, 8, 15, 7, 25, 12, 10)
//	fmt.Printf("Максимальний елемент на непарних позиціях: %d\n", max)
//
//	// 3) Виклик методу сортування масиву
//	people.selectionSort()
//	fmt.Println("Відсортований масив:")
//	for _, person := range people {
//		fmt.Printf("Прізвище: %s, Адреса: %s, Дата народження: %s\n", person.Name, person.Address, person.BirthDate)
//	}
//
//	// Відрізок, що містить корінь
//	a := 0.0
//	b := 2.0
//	// Точне значення коріння
//	expectedRoot := 1.0001
//
//	// Створюємо об'єкт, який реалізує інтерфейс RootSolver
//	var solver RootSolver
//	solver = &f
//
//	// Знаходимо корінь
//	root := solver.Solve(a, b, 0.0001)
//
//	// Порівнюємо знайдений корінь з точним значенням
//	if math.Abs(root-expectedRoot) > 0.0001 {
//		fmt.Println("Знайдений корінь не відповідає точному значенню")
//	} else {
//		fmt.Printf("Знайдений корінь: %.4f\n", root)
//	}
//}
