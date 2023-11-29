package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
)

func printStudentsWithTwos(students []STUDENT) {
	found := false

	for _, student := range students {
		if allTwos(student.SES) {
			fmt.Printf("Прізвище: %sНомер групи: %d\n--------------\n", student.NAME, student.GROUP)
			found = true
		}
	}

	if !found {
		fmt.Println("Студентів з оцінками 2 немає.")
	}
}

func allTwos(arr [5]int) bool {
	for _, val := range arr {
		if val != 2 {
			return false
		}
	}
	return true
}

func describeStudents() {
	var STUD1 [10]STUDENT

	for i := 0; i < 3; i++ {
		reader := bufio.NewReader(os.Stdin) // fmt.Scan works incorrect
		fmt.Printf("Введіть дані для студента #%d:\n", i+1)
		fmt.Print("Прізвище та ініціали: ")
		name, _ := reader.ReadString('\n')
		STUD1[i].NAME = name
		fmt.Print("Номер групи: ")
		fmt.Scan(&STUD1[i].GROUP)
		fmt.Print("Успішність (введіть 5 оцінок через пробіл): ")
		fmt.Scan(&STUD1[i].SES[0], &STUD1[i].SES[1], &STUD1[i].SES[2], &STUD1[i].SES[3], &STUD1[i].SES[4])
	}

	// Сортування за алфавітом
	sort.Slice(STUD1[:], func(i, j int) bool {
		return STUD1[i].NAME < STUD1[j].NAME
	})

	// Виведення прізвищ і номерів груп для студентів з оцінками 2
	fmt.Println("\nСтуденти з оцінками 2:")
	printStudentsWithTwos(STUD1[:])
}
