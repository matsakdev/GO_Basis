package main

func findAxisOnWhichPointLies(x float64, y float64) string {
	if x == 0 && y == 0 {
		return "Точка A знаходиться в початку координат (вісь X та вісь Y)."
	} else if x == 0 {
		return "Точка A знаходиться на осі Y."
	} else if y == 0 {
		return "Точка A знаходиться на осі X."
	} else {
		return "Точка A не знаходиться на жодній з осей."
	}
}
