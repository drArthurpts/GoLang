package main

import "fmt"

func somar(n1 int, n2 int) int {
	return n1 + n2
}

func calculosMatematicos(n1, n2 int) (int8, int8) {
	soma := n1 + n2
	subtracao := n1 - n2
	return int8(soma), int8(subtracao)
}

func main() {
	resultado := somar(10, 10)
	fmt.Println(resultado)

	var f = func(txt string) {
		fmt.Println(txt)
	}

	f("Texto da função f")

	calculoSoma, calculoSub  := calculosMatematicos(10,15)
	fmt.Println(calculoSoma, calculoSub)
}