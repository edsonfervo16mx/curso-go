package main

import "fmt"

type Gorra struct{
	marca string
	color string
	precio float32
	plana bool
}

func main(){

	var gorra_negra = Gorra{
		marca: "Nike",
		color: "Negro",
		precio: 34.87,
		plana: false}

	fmt.Println(gorra_negra)
	//
	fmt.Println(gorra_negra.marca)
}