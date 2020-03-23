package main

import (
	"fmt"
	"os"
	"strconv"
)

func main(){
	fmt.Println("CONDICION IF")
	//enviar paramentros de la consola
	fmt.Println(os.Args)
	fmt.Println("Hola " + os.Args[1] + "bienvenido")

	puntos,err := strconv.Atoi(os.Args[3])
	fmt.Println(err)

	//para evitar el uso de la variable de error
	/*puntos,_ := strconv.Atoi(os.Args[3])*/

	fmt.Println(puntos)

	edad := 25

	if edad >= 18 && edad != 25 && edad != 99{
		fmt.Println("Eres mayor de edad")
	}else if edad == 25 {
		fmt.Println("Tienes 25 años")
	}else if edad == 99{
		fmt.Println("Pronto moriras")
	}else{
		fmt.Println ("Eres menor de edad")
	}

}