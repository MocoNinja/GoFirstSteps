package main

import (
	"fmt"
	"time"
)

func fors() {
	fmt.Println("En GO no hay while, solo hay for. Puede usarse de distintas maneras")

	for i := 0; i < 5; i++ {
		fmt.Println("Esto es una iteración de un for tradicional")
	}

	fmt.Println("Puede usarse como un while")
	i := 18
	for i < 20 { // Esta sintaxis es la versión bonita de : ; i < 20;
		fmt.Println("Esto es como un while")
		i++
	}

	fmt.Println("Y un bucle infinito puede expresarse concisamente como: ")
	for {
		if i == 22 {
			break
		} else {
			i++
			fmt.Println(i)
		}
	}

}

func ifs() {
	esperable := true
	inusual := false
	if esperable {
		fmt.Println("El if básico no presenta sorpresas")
	} else if esperable && inusual {
		fmt.Println("Este caso no lo verás")
	} else {
		fmt.Println("Y con eso está cubierto todo lo que no sea feamente anidado")
	}

	fmt.Println("Como cualquier lenguaje que se precie, existen los switch para evitar un montón de if/else if/else...")
	nombre := ""
	fmt.Print("Inserta tu nombre: ")
	// fmt.Scanf("%s", &nombre) // Paso de meter nada en mis pruebas, salu2
	nombre = "Javier"
	switch nombre {
	case "Javier":
		fmt.Println("Eres guay")
	case "Docker":
		fmt.Println("Notar la ausencia de breaks")
	default:
		fmt.Println("Error de sintaxis")
	}

	fmt.Println("También puede omitirse la opción del switch")

	t := time.Now()
	switch {
	case t.Hour() < 12:
		fmt.Println("Cereales")
	default:
		fmt.Println("водка") // Notar el soporte para otros idiomas!!
	}
}

func otherIf(lim int) {
	if i := 2; lim < i {
		fmt.Println("El número que has introducido es una mierda")
	}
}
func main() {
	// fors()
	ifs()
	otherIf(1)
	otherIf(4)
}
