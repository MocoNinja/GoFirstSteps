package main

import (
	"fmt"
	"reflect" // For typeof
	"strconv" // Para las representaciones de bases
)

func numbers() {
	// Pruebas generales de declaración, tipos y casteos
	var entero = 20
	enteroTambien := -23
	var otroEnteroMas int = -2

	fmt.Printf("%d, %d, %d tienen los tipos %s, %s, %s...\n", entero, enteroTambien, otroEnteroMas, reflect.TypeOf(entero), reflect.TypeOf(enteroTambien), reflect.TypeOf(otroEnteroMas))

	resultadoInferido := entero - enteroTambien

	fmt.Println("Se ha cogido el tipo de dato de forma automática al operar", resultadoInferido, ", que es de tipo ", reflect.TypeOf(resultadoInferido))

	/*
		var resultadoRaro float32 = entero + otroEnteroMas
		fmt.Printf("Al especificarse operaciones entre enteros, NO puede forzarse el tipo %s", reflect.TypeOf(resultadoRaro))
	*/

	var resultado int = entero - otroEnteroMas
	var resultadoCasteado float32 = float32(resultado)
	resultadtoCasteadoInferido := float32(resultado)
	fmt.Print(resultadtoCasteadoInferido == resultadoCasteado && reflect.TypeOf(resultadtoCasteadoInferido) == reflect.TypeOf(resultadoCasteado))

	// Overview de los distintos tipos de variables
	//	Numéricas
	fmt.Println("Enteros...")
	fmt.Println("Los siguientes tipos no dependen de la máquina")
	var var1 uint8 = 2
	var var2 uint16 = 2
	var var3 uint32 = 2
	var var4 uint64 = 2
	var var5 int8 = 2
	var var6 int16 = 2
	var var7 int32 = 2
	var var8 int64 = 2
	fmt.Print(var1, var2, var3, var4, var5, var6, var7, var8, "\n")
	fmt.Print(reflect.TypeOf(var1), reflect.TypeOf(var2), reflect.TypeOf(var3), reflect.TypeOf(var4), reflect.TypeOf(var5), reflect.TypeOf(var6), reflect.TypeOf(var7), reflect.TypeOf(var8), "\n")
	fmt.Println("Los siguientes tipos sí que dependen de la máquina")
	var var9 uint = 2
	var var10 int = 2 // Es el que se infiere por defecto
	var var11 uintptr = 2
	fmt.Print(var9, var10, var11, "\n")
	fmt.Print(reflect.TypeOf(var9), reflect.TypeOf(var10), reflect.TypeOf(var11), "\n")

	fmt.Println("También hay aliases para algunos tipos de dato...")
	var miByte byte = 2 // Alias de uint8
	var miRune rune = 2 // Alias de int32
	// Más detaleles en https://golang.org/doc/go1#rune, https://stackoverflow.com/questions/19310700/what-is-a-rune# https://blog.golang.org/strings
	fmt.Println(reflect.TypeOf(miByte) == reflect.TypeOf(var1))
	fmt.Println(reflect.TypeOf(miRune) == reflect.TypeOf(var7))

	fmt.Println("Decimales...")
	var comaInferido = 23.23
	otroComaInferido := 23.23
	fmt.Printf("%f->%s\n%f->%s\n", comaInferido, reflect.TypeOf(comaInferido), otroComaInferido, reflect.TypeOf(otroComaInferido))

	var var12 float32 = -0.23 // Se suele denominar 'single precision'
	var var13 float64 = -2    // Se suele denominar 'double precision' Es el recomendado(notar su uso por defecto)
	fmt.Printf("%f->%s\n%f->%s\n", var12, reflect.TypeOf(var12), var13, reflect.TypeOf(var13))

	fmt.Println("También hay complejos...")
	complejo1 := 9i
	var complejo2 complex64 = 23 + 3i
	var complejo3 complex128 = -23 - 1i
	fmt.Printf("%f->%s\n%f->%s\n%f->%s\n", complejo1, reflect.TypeOf(complejo1), complejo2, reflect.TypeOf(complejo2), complejo3, reflect.TypeOf(complejo3))

	fmt.Println("Operaciones...")

	operacion1 := 9 + 10
	operacion2 := 9.2 - 3
	operacion3 := 2 - 3i
	operacion4 := 10 / 5
	operacion5 := 10 / 6
	operacion6 := 10.2 / 5.1
	operacion7 := 4 * 2
	operacion8 := 10 * (1 / 5)
	operacion9 := 10 % 4

	fmt.Println(operacion1, reflect.TypeOf(operacion1), "\n", operacion2, reflect.TypeOf(operacion2), "\n", operacion3, reflect.TypeOf(operacion3), "\n", operacion4, reflect.TypeOf(operacion4), "\n", operacion5, reflect.TypeOf(operacion5), "\n", operacion6, reflect.TypeOf(operacion6), "\n", operacion7, reflect.TypeOf(operacion7), "\n", operacion8, reflect.TypeOf(operacion8), "\n", operacion9, reflect.TypeOf(operacion9))

	fmt.Println("Notar especialmente las operaciones 5 y 8 que pueden sorprender a simple vista!")
}

func strings() {
	fmt.Println("El uso básico del tipo de dato 'string' es bastante típico")
	mensaje1 := "WOLAS"
	var mensaje2 string = "OLIS"
	fmt.Printf("%s->%s\n%s->%s\n", mensaje1, reflect.TypeOf(mensaje1), mensaje2, reflect.TypeOf(mensaje2))

	tocho := "HOLAS " + "MU" + "   " + "sdad" + string(2923) // + 24 // NO RULA
	fmt.Printf(tocho)

	fmt.Println("Notar que ahí no está pillando el número sino el valor del byte que representa. Eso nos lleva a las disintas cosas.")
	len := len(tocho)
	fmt.Println("Sabiendo por ejemplo que podemos sacar la longitud, dando ", len)
	for i := 0; i < len; i++ {
		extraido := tocho[i]
		fmt.Println("Extraido:", extraido, "de tipo", reflect.TypeOf(extraido))
		var extraidoEquivaleA byte = tocho[i]
		fmt.Println("Es lo mismo que un byte??:", extraidoEquivaleA == extraido)
		caracter := string(tocho[i])
		fmt.Println("Ese byte equivale a", caracter)
	}
}

func booleanos() {
	facilito := true
	var lioso bool = false
	fmt.Printf("%t->%s\n%t->%s\n", facilito, reflect.TypeOf(facilito), lioso, reflect.TypeOf(lioso))

	tipico := facilito || lioso
	novedoso := facilito && lioso
	innovador := !facilito

	fmt.Println(tipico, novedoso, innovador)
}

func bases() {
	hexadecimal := 0xfe0fe0
	octal := 0123
	//	binary := 0122121 // Nope :(
	fmt.Printf("%d->%s\n%d->%s\n", hexadecimal, reflect.TypeOf(hexadecimal), octal, reflect.TypeOf(octal))
	valorHex, _ := strconv.ParseInt("fe0fe0", 16, 0)
	fmt.Println(valorHex)
	valorBin, _ := strconv.ParseInt("00000101", 2, 0)
	fmt.Println(valorBin)
	valorBin2, _ := strconv.ParseInt("101", 2, 0)
	fmt.Println(valorBin2)
	for i := 0; i < 32; i++ {
		valorBin3, _ := strconv.ParseInt("00101000", 2, i)
		fmt.Println(valorBin3)
	}
	fmt.Println(strconv.FormatInt(11, 2))
	fmt.Println(strconv.FormatInt(11, 16))
}

func main() {
	// numbers()
	// strings()
	// booleanos()
	bases()
}
