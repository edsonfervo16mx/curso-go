package main

import "fmt"

func main(){
	mensaje()
}

func mensaje(){
	fmt.Println("Este es un mensaje!!!")

	fmt.Println(operacion(10,8,"+"))
	fmt.Println(operacion(10,8,"-"))
	fmt.Println(operacion(10,8,"*"))
	fmt.Println(operacion(10,8,"/"))
}

func operacion (n1 float32, n2 float32, op string) float32{
	var resultado float32

	if(op == "+"){
		resultado = n1 + n2
	}
	if(op == "-"){
		resultado = n1 - n2
	}
	if(op == "*"){
		resultado = n1 * n2
	}
	if(op == "/"){
		resultado = n1 / n2
	}
	return resultado
}