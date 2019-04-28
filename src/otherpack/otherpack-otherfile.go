package otherpack

import (
	"fmt"
)

func Extpackfunc() {
	fmt.Println("Extpackfunc")
}

func Extpackcaller() {
	forloop()
	forloop2()
	forloop3()
	ifstatement()
	ifstatement2()
	ifstatement3()
}

func forloop() {
	sum := 0
	for i := 0; i < 10; i++ {
		sum += 1
	}
	fmt.Println(sum)
}

//init and post statements are optional
func forloop2() {
	sum := 1
	for sum < 100 {
		sum += 1
	}
	fmt.Println(sum)
}

// for is like while
func forloop3() {
	sum := 1
	for sum < 100 {
		sum += 1
	}
	fmt.Println(sum)
}

// no ()
func ifstatement() {
	sum := 1
	if sum == 1 {
		fmt.Println("sum is equals to ", sum)
	}
}

// if with statement before conditions
func ifstatement2() {
	if v := 3; v == 3 {
		fmt.Println("v is equals to", v)
	}
}

// v is visibile also in else statement
func ifstatement3() {
	if v := 3; v == 3 {
		fmt.Println("v is equals to", v)
	} else {
		fmt.Println("v is equals to", v+v)
	}
}
