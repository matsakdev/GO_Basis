package main

func isTriangleEquilateral(a int, b int, c int) bool {
	isFirstPairEqual := a == b
	isSecondPairEqual := b == c

	return isFirstPairEqual && isSecondPairEqual
}
