package main

import (
	"fmt"
)

func arrays() {
	fmt.Println("Un array es como en C, una secuencia ordenada de longitud fija y elementos del mismo tipo")
	var numeros [5]int // [0 0 0 0 0]
	numeros[2] = 69
	fmt.Println(numeros)

	otrosNumeros := [4]int{15, 16, -23, -2}

	inclusoOtraFormaDeDeclararlo := [5]int{
		1,
		2,
		3,
		4,
		5,
		//	5, // Notar la coma final para la facilidad de comentar
	}
	fmt.Println(inclusoOtraFormaDeDeclararlo)
	fmt.Println(otrosNumeros)
	fmt.Printf("El array 'otrosNumeros' tiene %d elemento(s)!\n", len(otrosNumeros))

	fmt.Println("Y hay una sintaxis super maja para acceder a rangos y tal")

	elArray := [10]byte{
		0,
		1,
		2,
		3,
		4,
		5,
		6,
		7,
		8,
		9,
	}

	fmt.Println(elArray[:4])
	fmt.Println(elArray[6:])
	fmt.Println(elArray[6:9])

	// fmt.Println(elArray[-1]) // No valen estas homogayadas
}

func slices() {
	fmt.Println("Los slices son cachitos de arrays")
	elArray := [10]byte{
		0,
		1,
		2,
		3,
		4,
		5,
		6,
		7,
		8,
		9,
	}
	fmt.Println(elArray)
}

func main() {
	arrays()
}
