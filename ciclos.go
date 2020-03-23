package main

import "fmt"

func main()  {
	for i := 0; i<= 15 ; i++{
		fmt.Println(i," este es el numero actual")
	}

	fmt.Println("******************************")

	for j:= 0;j<= 15;j++{
		if j%2 == 0{
			fmt.Println("El numero es par: ",j)
		}else{
			fmt.Println("El numero es impar: ",j)
		}
	}

}