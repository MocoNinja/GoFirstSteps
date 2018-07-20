package main

import (
	"fmt"
)

func printSlice(s []int) {
	fmt.Printf("len=%d cap=%d %v\n", len(s), cap(s), s)
}

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
	fmt.Println(elArray[:])

	// fmt.Println(elArray[-1]) // No valen estas homogayadas
}

func slices() {
	fmt.Println("Los slices son cachitos de arrays, pero más flexibles debido a su tamaño DINÁMICO")

	fmt.Println("En la creación de un slice, se crea automáticamente un array al que se asocia el slice")
	arrayCreado := [3]bool{true, true, false}                                              // Un array
	sliceQueCreaUnArrayComoElAnteriorYLoReferenciaEnEsteSlice := []bool{true, true, false} // Slice que apunta a un array que también se crea
	// fmt.Println(arrayCreado == sliceQueCreaUnArrayComoElAnteriorYLoReferenciaEnEsteSlice)  // Si se tiene linter, se puede ver el error de tipo que da
	fmt.Println(sliceQueCreaUnArrayComoElAnteriorYLoReferenciaEnEsteSlice, arrayCreado)

	fmt.Println("Un slice tiene capacidad (que es el tamaño del array subyacente) y longitud(que es su número de elementos)")
	fmt.Println("MUY IMPORTANTE: UN SLICE PUEDE TENER UNA LONGITUD MENOR O IGUAL A SU CAPACIDAD, PERO NUNCA MAYOR")
	s := []int{2, 3, 5, 7, 11, 13}
	printSlice(s)
	s = s[:0]
	printSlice(s)
	s = s[:4]
	printSlice(s)
	// Drop its first two values.
	s = s[2:]
	printSlice(s)

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
	// var cashito1 []int = elArray[2:6] // declaración tocha
	cashito1 := elArray[2:6]
	fmt.Println(cashito1)
	x := make([]int, 5)
	y := make([]int, 5, 10)
	printSlice(x)
	printSlice(y)
	// fmt.Println(len(x), len(y), x[5], y[8]) // No puede accederse a un índice fuera del slice aunque entre dentro del array
	printSlice(y[5:10])
}

func sliceFunctions() {
	slice1 := []int{1, 2, 3}
	slice2 := append(slice1, 4, 5, 6)
	fmt.Println(slice2)
}

func main() {
	arrays()
	slices()
	sliceFunctions()
}
