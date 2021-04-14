package main

import "fmt"

func main() {
	// Declaracion de constantes
	fmt.Println("Hello!")
	// 	name	Type = value
	const pi float64 = 3.14 // Nunca cambiara en el tiempo
	const pi2 = 3.15

	fmt.Println(pi, pi2)

	fmt.Println("pi:", pi)
	fmt.Println("pi2:", pi2)

	// Variables

	// Enteras

	// This will create a new variable and give the value of 12
	// Look, the key is on : = , the : before = means created if this
	// does not exist
	base := 12

	// This one is a variable of type int with 12 as value
	var altura int = 12

	// I create this one, but the value will come later
	var area int

	// Go didn't work if we don't use all te variables, compilation error
	fmt.Println(base, altura, area)

	// Zero values

	var a int     // 0 by default
	var b float64 // 0 by default
	var c string  // Empty
	var d bool    // false

	fmt.Println(a, b, c, d)

	//Area cuadrado
	const baseCuadrado = 10
	areaCuadrado := baseCuadrado * baseCuadrado
	fmt.Println("Area cuadrado:", areaCuadrado)
}
