package main

import "fmt"

func main(){
	//El slices no es mas que un array dinamico, sin declaracion de tamaño al inicio
	peliculas := [] string{
		"Norbit",
		"Frozen",
		"Que pasó ayer?",
		"Batman"}

	//agregamos un nuevo valor en el array
	peliculas = append(peliculas,"La liga de la justicia")
	peliculas = append(peliculas,"Superman")
	peliculas = append(peliculas,"Acuaman")

	var cantidad_elementos int 
	cantidad_elementos = len(peliculas)

	fmt.Println(peliculas)

	fmt.Println(cantidad_elementos)

	//pedir un rango de elementos
	fmt.Println(peliculas[0:3])

}