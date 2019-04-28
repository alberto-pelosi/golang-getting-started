package otherpack

import (
	"fmt"
)

func Extpackfunc() {
	fmt.Println("Extpackfunc")
}

func Extpackcaller() {
	forloop()
}

func forloop() {
	sum := 0
	for i := 0; i < 10; i++ {
		sum += 1
	}
	fmt.Println(sum)
}
