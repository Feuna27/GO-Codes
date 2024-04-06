/*
//to find the avg of two numbers by using type conversion

package main

import "fmt"

func main() {
	var num1 int
	var num2 int
	//var avg float64//
	num1 = 10
	num2 = 5
	//avg = (float64(num1) + float64(num2)) / 2//
	fmt.Printf("average of %d and %d is %.2f\n", num1, num2, (num1+num2)/2)

}*/

/*
// if statement//
package main

import "fmt"

	func main() {
		if 20 > 18 {
			fmt.Println("20 is greater than 18")
		}
	}

package main

import "fmt"

	func main() {
		if !(20 < 18) {
			fmt.Println("20 is greater than 18")
		}
	}
*/
/*
package main

import "fmt"

func main() {
	x := 8
	if y := 10; x < y {
		fmt.Println("x is less than y")
	}
}*/

/*
package main

import (
	"fmt"
	"os"
)

func main() {
	var name string
	var age int
	if _, err := fmt.Scan(&name, &age); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	fmt.Printf("your name is: %s\n", name)
	fmt.Printf("your age is : %d\n", age)

}*/

/*
//write a program to print biggest number among three numbers

package main

import "fmt"

func main() {
	fmt.Println("Enter 3 numbers :")
	var a, b, c int
	fmt.Scanln(&a, &b, &c)                           //scanln ek hi value count krega pehle

	if a >= b && a >= c {
		fmt.Println(a, "is Largest")
	}
	if b >= a && b >= c {
		fmt.Println(b, "is Largest")
	}
	if c >= a && c >= b {
		fmt.Println(c, "is Largest")
	}
}*/

/*
// if else//
package main

import "fmt"

func main() {
	x := 12
	y := 10
	if x <= y {
		fmt.Println("x is less than y")
	} else {                                                                                  //else ko same line me hi likhna rehta next line me error denga
		fmt.Println("x is greater than or equal to y")
	}
}
*/

/*
// write a program to show whether the number is even or odd
package main

import "fmt"

func main() {
	x := 13
	if x%2 == 0 {
		fmt.Println("x is even number")
	} else {
		fmt.Println("x is odd number")
	}

}


//write the program to print given number in Absolute format.

package main

import (
	"fmt"
	"math"
)

func main() {
	var x float64
	fmt.Print("num: \n")
	fmt.Scan(&x)
	fmt.Printf("abs value is %.2f", math.Abs(x))
}
*/

func add(x, y) {
	z = x + y
	fmt.Println("addition is :", x)
}
