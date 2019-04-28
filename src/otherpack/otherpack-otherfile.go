package otherpack

import (
	"fmt"
	"runtime"
	"time"
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
	sqrt(2)
	switchstatement()
	switchstatement2()
	switchstatement3()
	deferstatement()
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

//square root with for loop
func sqrt(x float64) float64 {
	z := float64(1)
	i := 0
	for i < 10 {
		z -= (z*z - x) / (2 * z)
		i++
		fmt.Printf("iteration %d, z is equals to %f \n", i, z)
	}
	return z
}

//no need of break
func switchstatement() {

	switch os := runtime.GOOS; os {
	case "windows":
		fmt.Println("case windows")
	default:
		fmt.Println(os)
	}
}

func switchstatement2() {
	fmt.Println("When's Saturday?")
	today := time.Now().Weekday()
	fmt.Println(today)
	switch time.Saturday {
	case today + 0:
		fmt.Println("Today.")
	case today + 1:
		fmt.Println("Tomorrow.")
	case today + 2:
		fmt.Println("In two days.")
	default:
		fmt.Println("Too far away.")
	}
}

//switch without condition always evaluated
func switchstatement3() {
	t := time.Now()
	switch {
	case t.Hour() < 12:
		fmt.Println("Good morning!")
	case t.Hour() < 17:
		fmt.Println("Good afternoon.")
	default:
		fmt.Println("Good evening.")
	}
}

func deferstatement() {
	fmt.Println("counting")

	for i := 0; i < 10; i++ {
		defer fmt.Println(i)
	}

	fmt.Println("done")
}
