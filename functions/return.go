package main

import "fmt"

func swap(x, y string) (string, string) {
	return y, x
}
func concat(x, y string) string {
	return x + " " + y
}
func split(sum int) (x, y int) {
	x = sum * 4 / 9
	y = sum - x
	return
}

func main() {
	a, b := swap("hello", "world")
	fmt.Println(a, b)

	c := concat("Go", "practice")
	fmt.Println(c)

	d, e := split(10)
	fmt.Println("Results : ", d, " and ", e)
}
