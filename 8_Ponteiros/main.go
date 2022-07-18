package main

import "fmt"

func main() {
	var var1 int = 10
	var var2 int = var1

	fmt.Println(var1, var2)

	var1++
	fmt.Println(var1, var2)

	var var3 int
	var var4 *int

	var3++
	fmt.Println(var3, var4)

	var4 = &var3

	fmt.Println(var3, var4)
	fmt.Println(var3, *var4)

	var3++
	fmt.Println(var3, var4)
	fmt.Println(var3, *var4)
}
