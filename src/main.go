package main

import (
	"fmt"
	"math"
)

func main() {
	// helloWorld()
	// variablesDeclaration()
	//operators()
	//challenge_one()
	//fmtUsage()
	normalFunctionName("hi")
	moreArgs(5, 6, "7")
	fmt.Println(returnValue(5))

	value1, value2 := doubleReturn(5)

	fmt.Println(value1, value2)

	// Descartar value 2
	value1, _ = doubleReturn(10)

	fmt.Println(value1)

}

func normalFunctionName(message string) {
	fmt.Println(message)
}

// A y B son int
func moreArgs(a, b int, c string) {
	fmt.Println(a, b, c)
}

// THis one return two values
func doubleReturn(a int) (c, d int) {
	return a, a * 2
}

func returnValue(a int) int {
	return a * a
}

func fmtUsage() {

	var hello string = "Hello"
	var world string = "World"

	fmt.Println(hello, world)
	fmt.Println(hello, world)

	names := "Platzi"
	number := 500

	fmt.Printf("%s has more than %d \n", names, number)
	fmt.Printf("%v has more than %v \n", names, number)

	message := fmt.Sprintf("%s has more than %d", names, number)
	fmt.Println(message)

	fmt.Printf("hello: %T \n", hello)
	fmt.Printf("number: %T \n", number)

}

func circleArea(radio float32) float32 {
	var circulo float32 = math.Pi * (radio * radio)

	fmt.Println("Circulo ->", circulo)

	return circulo
}

func rectangleArea(base, high int) int {
	result := base * high
	fmt.Println("Rectangulo ->", result)
	return result
}

func trapezeArea(base_a, base_b, h int) int {
	return h * ((base_a * base_b) / 2)
}

func challenge_one() {
	const radio float32 = 5.6869
	var circulo float32 = math.Pi * (radio * radio)

	fmt.Println("Circulo ->", circulo)

	const base int = 5
	const altura int = 10

	result := base * altura

	fmt.Println("Rectangulo ->", result)

	const base_a = 5
	const base_b = 6
	const h = 9

	result = h * ((base_a * base_b) / 2)

	fmt.Println("Trapecio ->", result)

}

func operators() {
	fmt.Println("Lets do some operations")

	x := 10
	y := 20

	var sum int = x + y
	fmt.Println("Suma:", sum)

	mult := x * y
	fmt.Println("Mult", mult)

	rest := x - y
	fmt.Println("Resta", rest)

	div := x / y
	fmt.Println("Div", div)

	mod := x % y
	fmt.Println("Mod", mod)

	// Incremental
	x++
	fmt.Println("Incremental", x)

	// Decremental
	x--
	fmt.Println("Incremental", x)

}

func helloWorld() {
	fmt.Println("Hello!")
}

func variablesDeclaration() {
	// Declaracion de constantes
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
