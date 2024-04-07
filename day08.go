/*
package main

import "fmt"

func greet() {
	fmt.Println("good morning")
}

func main() {
	greet()

}*/

/*
package main

import "fmt"

func addNumbers() {
	n1 := 12
	n2 := 8

	sum := n1 + n2
	fmt.Println("sum: ", sum)
}

func main() {
	addNumbers()
}*/

/*
package main

import "fmt"

func addNumbers(n1 int, n2 int) {
	sum := n1 + n2
	fmt.Println("sum: :", sum)
}

func main() {
	var x int
	var y int
	fmt.Printf("\n x and y : \n")
	fmt.Scanln(&x, &y)
	addNumbers(x, y)
}
*/
/*
package main

import "fmt"

func addNumbers(n1 int, n2 int) int {
	sum := n1 + n2
	return sum
}

func main() {
	result := addNumbers(21, 13)

	fmt.Println(" sum :", result)
}*/

/*
package main

import "fmt"

func addNumbers(n1 int, n2 int) int {
	sum := n1 + n2

	fmt.Println("after return statement")
	return sum
}
func main() {
	result := addNumbers(21, 13)

	fmt.Println(" sum :", result)

}*/

/*
package main

import "fmt"

func calculate(n1 int, n2 int) (int, int) {
	sum := n1 + n2
	difference := n1 - n2

	//return two values
	return sum, difference

}
func main() {
	sum, difference := calculate(21, 13)
	fmt.Println("Sum:", sum, "Difference", difference)
	fmt.Scanln(&sum, &difference)
	fmt.Println("sum:", sum, "Difference", difference)

}*/

/*
package main

import "fmt"
//function definition
func getSquare(num int) {
	square := num * num
	fmt.Printf("square of %d is %d\n", num, square)

}

//main function
func main() {
	//call function 3 times
	getSquare(3)
	getSquare(5)
	getSquare(10)

}*/

/*
//local variable  //galat code

package main

import "fmt"

func addNumbers() {
	var sum int
	sum = 5 + 9
}
func main() {
	var sum int = 0
	addNumbers()
	fmt.Println("sum is ", sum)
}*/

/*
//recursion

package main

import "fmt"

func countDown(number int) {

	//display the number
	fmt.Println(number)

	//recursive call by decreasing number
	countDown(number - 1)
}

func main() {
	fmt.Println("countdown starts: ")
	countDown(3)
}*/

/*
//to stop the countdown
package main

import "fmt"

func countDown(number int) {
	if number > 0 {
		fmt.Println(number)
		//recursive call
		countDown(number - 1)
	} else {
		//ends the recursive function
		fmt.Println("countdown stop")
	}
}

func main() {
	fmt.Println("countdown starts")
	countDown(3)
}*/
