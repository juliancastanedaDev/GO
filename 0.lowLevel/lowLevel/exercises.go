package lowLevel

import "fmt"

//Print message hello world
func Greet() {
	fmt.Print("Hello world")
}


func AddNumbers() {
	num1 := 3
	num2 := 6
	var add int

	add = num1 + num2

	fmt.Print(add)
}
