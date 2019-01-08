package main

import "fmt"

func Add(x, y int) int {
	return x + y
}

func main() {
	fmt.Println(Add(42, 18))
	c, python, java := true, false, "no!"
	var i int
	fmt.Println(i, c, python, java)
}
