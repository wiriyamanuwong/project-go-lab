package basicfunc

import "fmt"

func factorial(x uint) uint {
	if x == 0 {
		return 1
	}
	return x * factorial(x-1)
}

func DoSomething() {
	fmt.Println("Hello, world!")
}
