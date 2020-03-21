package main

import "fmt"

func main(){
	fmt.Println(gorras(45))
}

func gorras(pedido float32) (string, float32){

	precio := func() float32{
		return pedido * 7
	}

	return "El monto del pedido es: ", precio()
}