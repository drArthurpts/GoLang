package main

import (
	"modulo/auxiliar"
	"fmt"
	"github.com/badoux/checkmail"
)

func main() {
	fmt.Println("Hello, World!")
	auxiliar.Escrever()
	erro := checkmail.ValidateFormat("arthurpatriciosousa@gmail.com")
	fmt.Println(erro)
}