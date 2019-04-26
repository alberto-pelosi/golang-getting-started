package main

import (
    "fmt"
    "math"
    "math/rand"
    "time"
)

var packageVariable bool

func main() {
    rand.Seed(time.Now().UnixNano())
    fmt.Println("My favorite number is", rand.Intn(1000))
    sq()
    pi()
    fmt.Printf("add %d \n", add(1,3))
    fmt.Printf("add2 %d \n", add2(3,5))
    add3(7, 9)
    a, b := swapz("ciao", "mondo")
    fmt.Println(a, b)
    c, d := nakedReturn(5)
    fmt.Println("nakedReturn", c, d)
    variableDefinition()
    
    
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
    fmt.Printf("add3 %d", x + y)
   
}

func swapz(x, y string) (string, string) {
    return y, x
}

func nakedReturn(sum int) (x, y int) {
    x = sum * 4 / 9
    y = sum -x
    return
}

func variableDefinition() {
	var funcVariable bool
	fmt.Println(packageVariable, funcVariable)
}

