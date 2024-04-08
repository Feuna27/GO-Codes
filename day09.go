/*//to calc sum of positive numbers//

package main

import "fmt"

func sum(number int) int {
	//condition to break recursion
	if number <= 0 {
		return 0
	} else {
		return number + sum(number-1)
	}
}

func main() {
	var num int
	fmt.Println("enter the number:\n")
	fmt.Scanln(&num)
	var result = sum(num)
	fmt.Println("Sum: ", result)

}
*/

/*
package main

import "fmt"

func sum(number int) int {
	//condition to break recursion
	if number <= 0 {
		return 0
	} else {
		return number + sum(number-1)
	}
}

func main() {
	var num int
	fmt.Println("enter the number:\n")
	fmt.Scanln(&num)
	if num < 0 {                                                   //usingif else for negative number//
		fmt.Println("invalid input try again!")
	} else {
		var result = sum(num)
		fmt.Println("Sum: ", result)

	}

}*/

/*
//factorial

package main

import "fmt"

func factorial(num int) int {
	if num == 0 {
		return 1
	} else {
		return num * factorial(num-1)

	}

}

func main() {

	var number = 3
	var result = factorial(number)
	fmt.Println("the factorial of 3 is ", result)

}*/

/*
// anonymous function
package main

import "fmt"

func main() {
	//anonymous func
	var greet = func() {
		fmt.Println("heyyyoooo")
	}

	//function call
	var welcome = greet
	greet()
	welcome()
}*/

/*
// anonymous function with parameters
package main

import "fmt"

func main() {

	//anonymous function with arguments
	var sum = func(n1, n2 int) {
		result := n1 + n2
		fmt.Println("sum is ", result)
	}
	fmt.Printf("\ntype %T\n", &sum)
	sum(5, 3)
}*/
/*
// return value from anonymous function
package main

import "fmt"

func main() {
	var sum = func(n1, n2 int) int {
		sum := n1 + n2
		return sum
	}

	result := sum(5, 3)
	fmt.Println("sum is", result)
}
*/

//anonymous function as arguments to other functions

package main

import "fmt"

var sum = 0

func findSquare(num int) int {
	square := num * num
	return square
}

func main() {
	sum := func(number1 int, number2 int) int {
		return number1 + number2
	}

	result := findSquare(sum(6, 9))
	fmt.Println("result is: ", result)

}
