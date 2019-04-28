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
