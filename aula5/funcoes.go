package main

import "fmt"

func somar(n1 int8, n2 int8) int8 {
	return n1 + n2
}

func calculosMatematicos(n1, n2 int8) (int8, int8) {

	soma := n1 + n2
	subtracao := n1 - n2
	return soma, subtracao

}

func main() {
	soma := somar(10, 3)
	fmt.Println(soma)

	var f = func() {
		fmt.Println("Função F")
	}

	f()

	resultadoSoma, _ := calculosMatematicos(4, 5)
	fmt.Println(resultadoSoma)

	_, resultadoSubtracao := calculosMatematicos(4, 5)
	fmt.Println(resultadoSubtracao)
}
