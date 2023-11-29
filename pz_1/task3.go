package main

func getGradeName(grade int) string {

	switch grade {
	case 1:
		return "Погано"
	case 2:
		return "Незадовільно"
	case 3:
		return "Задовільно"
	case 4:
		return "Добре"
	case 5:
		return "Відмінно"
	default:
		return "Некоректна відмітка."
	}
}
