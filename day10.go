/*
// print odd numbers using golang closure
package main

import "fmt"

func calculate() func() int {
	num := 1
	//return inner function
	return func() int {
		num = num + 2
		return num
	}
}

func main() {
	//call the outer function
	odd := calculate()
	//call the ineer function
	fmt.Println(odd())
	fmt.Println(odd())
	fmt.Println(odd())

	//call the outer function again
	odd2 := calculate()
	fmt.Println(odd2())

}
*/

/*
package main

import "fmt"

//outer function

func greet() func() string {
	//variable defined outside the inner function
	name := "john"
	//return a nested anonymous function
	return func() string {
		name = "hi" + name
		return name
	}
}

func main() {
	//call the outer function
	message := greet()
	//call the inner function
	fmt.Println(message())
}*/

/*
// closure helps in data isolation
package main

import "fmt"

func displayNumbers() func() int {
	number := 0
	return func() int {
		number++
		return number
	}
}

func main() {
	//num1 := displayNumbers()
	fmt.Println(displayNumbers())
	fmt.Println(displayNumbers()())
	fmt.Println(displayNumbers()())
	//num2 := displayNumbers()
	fmt.Println(displayNumbers()())
	fmt.Println(displayNumbers()())

}

*/

//fibonacci series in anonymous function

package main

import "fmt"

func main() {
	var fib func(n int) int
	fib = func(n int) int {
		if n < 2 {
			return 0
		}
		return fib(n-1) + fib(n-2)
	}
	fmt.Println(fib(7))
}
