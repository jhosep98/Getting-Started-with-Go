/*
Go tiene varios tipos diferentes para representar números. Generalmente dividimos los números en dos tipos diferentes: integers and floating-point numbers.
Numeros ->  A diferencia del sistema decimal de base 10 que usamos para representar números, las computadoras usan un sistema binario de base 2.*/

package main

import "fmt"

var x uint    // uint8 es de  0 - 255
var y int     //En general, si está trabajando con enteros, solo debe usar el "int" Tipo.
var z float64 // o float32 (En general, debemos seguir float64 cuando trabajamos con números de coma flotante.)

func main() {
	x = 256  // entero sin signo
	y = -21  // entero con signo o 21
	z = 21.5 // decimal
	fmt.Println(x)
	fmt.Println(y)
	fmt.Println(z)
}
