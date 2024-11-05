package main

import (
	"errors"
	"fmt"
)

func main() {

	//numeros inteiros
	var numero1 int8 = 111
	fmt.Println(numero1)

	var numero2 int16 = 11111
	fmt.Println(numero2)

	var numero3 int32 = 1111143242
	fmt.Println(numero3)

	var numero4 int64 = 1111143242234324234
	fmt.Println(numero4)

	numero5 := 12342134213
	fmt.Println(numero5)
	//fim numeros inteiros

	//numeros reais
	var numero6 float32 = 123433333666666666633333333333333536666.333765756756
	fmt.Println(numero6)

	var numero7 float64 = 1234333336643122414312341241234214213421342134213421341234214124312432314123432421421432134214312334213342134214123343123123123127777777777777777777777777777777777777777777777777777777775555555555555555555555531231231231231231231231231266666665464564564564564564565464566633333333333333536666.333765756756
	fmt.Println(numero7)

	numero8 := 123123.54
	fmt.Println(numero8)
	//numeros reais

	var str string = "texto"
	fmt.Println(str)

	str2 := "Texto 2"
	fmt.Println(str2)

	char := 'B'
	fmt.Println(char)

	var bool1 bool = true
	fmt.Println(bool1)

	var bool2 bool
	fmt.Println(bool2)

	var erro error
	fmt.Println(erro)

	var erro2 error = errors.New("Erro interno")
	fmt.Println(erro2)

}
