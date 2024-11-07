package main

import "fmt"

func main() {
	//aritmeticos
	soma := 1 + 2
	subtracao := 1 - 2
	divisao := 10 / 4
	multiplicacao := 10 * 5
	restoDaDivisao := 10 % 2

	fmt.Println(soma, subtracao, divisao, multiplicacao, restoDaDivisao)

	var num1 int16 = 10
	var num2 int32 = 25

	soma2 := num1 + int16(num2)
	fmt.Println(soma2)
	//fim aritmeticos

	//atribuicao
	var variavel1 string = "String"
	variavel2 := "String2"

	fmt.Println(variavel1, variavel2)
	//fim atribuicao

	//relacionais
	fmt.Println(1 > 2)
	fmt.Println(1 < 2)
	fmt.Println(1 >= 2)
	fmt.Println(1 <= 2)
	fmt.Println(1 == 2)
	fmt.Println(1 != 2)
	//fim relacionais

	//logicos
	fmt.Println("------------------------------------------")
	verdadeiro, falso := true, false
	fmt.Println(verdadeiro && verdadeiro)
	fmt.Println(verdadeiro || falso)
	fmt.Println(!verdadeiro)
	fmt.Println(!falso)
	//fim logicos

	//operadores unarios
	numero := 10
	numero++
	numero += 33
	fmt.Println(numero)

	numero--
	numero -= 12
	fmt.Println(numero)

	numero *= 3
	numero /= 10
	numero %= 3
	fmt.Println(numero)
	//fim unarios

}
