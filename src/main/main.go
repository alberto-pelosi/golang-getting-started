package main

import (
	"fmt"
	"math"
	"math/cmplx"
	"math/rand"
	"time"
)

var packageVariable bool

func main() {
	rand.Seed(time.Now().UnixNano())
	fmt.Println("My favorite number is", rand.Intn(1000))
	sq()
	pi()
	fmt.Printf("add %d \n", add(1, 3))
	fmt.Printf("add2 %d \n", add2(3, 5))
	add3(7, 9)
	a, b := swapz("ciao", "mondo")
	fmt.Println(a, b)
	c, d := nakedReturn(5)
	fmt.Println("nakedReturn", c, d)
	variableDefinition()
	variableInitialization()
	basicTypes()

}

func sq() {
	fmt.Printf("Now you have %g problems.\n", math.Sqrt(7))
}

func pi() {
	fmt.Println(math.Pi)
}

func add(x int, y int) int {
	return x + y
}

func add2(x, y int) int {
	return x + y
}

func add3(x, y int) {
	fmt.Printf("add3 %d", x+y)

}

func swapz(x, y string) (string, string) {
	return y, x
}

func nakedReturn(sum int) (x, y int) {
	x = sum * 4 / 9
	y = sum - x
	return
}

func variableDefinition() {
	var funcVariable bool
	fmt.Println(packageVariable, funcVariable)
}

func variableInitialization() {
	var i, b = false, "lol"
	var k, l int = 1, 2
	shortDeclaration := 3
	fmt.Println(i, b)
	fmt.Println(k, l, shortDeclaration)
}

func basicTypes() {
	const k = "lol"
	var (
		b   bool    = true
		i   int     = 123
		i8  int8    = 127
		i16 int16   = 255
		i32 int32   = 4000000
		i64 int64   = 999999999
		by  byte    = 255
		r   rune    = 4000000
		f32 float32 = 4.33
		f64 float64 = 6321321.333
		cp          = cmplx.Sqrt(-5 + 12i)
	)
	fmt.Println(b, i, i8, i16, i32, i64, by, r, f32, f64, cp, k)

}
