package main

import (
	"fmt"
	"modulo/auxiliar"

	"github.com/badoux/checkmail"
)

func main() {
	println("Escrevendo da main")
	auxiliar.Escrever()

	error := checkmail.ValidateFormat("devbookgmail.com")
	fmt.Println(error)
}
