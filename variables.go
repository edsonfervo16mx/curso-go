package main 

import "fmt"

func main(){
	var suma int = 8 + 9
	var resta int = 6 - 4
	var nombre string = "edsonfer"

	var apellidos string = "Ventura Oropeza"
	apellidos = "De la Cruz"

	pais := "Mexico" // El lenguaje infiere el tipo de variable

	var estado bool = true
	var precio float32 = 19.81

	const year int = 2020 // si se le intenta cambiar el valor a una constante marca error


	fmt.Println("Hola mundo desde Go")
	fmt.Println(suma)
	fmt.Println(resta)
	fmt.Println("Mi nombre es: "+nombre + apellidos + "Soy de : " + pais)
	fmt.Println(estado)
	fmt.Println(precio)

	fmt.Println(year)
}