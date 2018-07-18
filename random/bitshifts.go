package main

import (
	"fmt"
	"strconv"
)

// Thx https://medium.com/learning-the-go-programming-language/bit-hacking-with-go-e0acee258827

func c00lprint(number int64, base int) {
	representacion := strconv.FormatInt(number, base)
	fmt.Printf("El número %d tiene una representación de %s\n", number, representacion)
}

func operadores() {
	var mask1 int64 = 5
	var mask2 int64 = 2

	c00lprint(mask1, 2)
	c00lprint(mask2, 2)

	y := mask1 & mask2 // Esto coge el tipo guay
	o := mask1 | mask2
	xor := mask1 ^ mask2
	andnot := mask1 &^ mask2

	c00lprint(y, 2)
	c00lprint(o, 2)
	c00lprint(xor, 2)
	c00lprint(andnot, 2)
}
func main() {
	operadores()
}
