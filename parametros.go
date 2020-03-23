package main

import "fmt"

func main(){
	pantalon("rojo", "largo", "sin bolsillos", "nike")
}

func pantalon(caracteristicas ... string){
	//es como hacer un foreach, es un recorrido de las caracteristicas
	for _,caracteristicas := range caracteristicas{
		fmt.Println(caracteristicas)	
	}
	
}


