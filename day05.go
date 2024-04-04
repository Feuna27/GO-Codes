/*package main

import "fmt"

func main() {
	integer := 23
	fmt.Print("%T %T\n", integer, &integer)
}*/

/*
package main

import "fmt"

func main() {
	floatingnumber := 1234.57883939
	fmt.Printf("default: %f\n", floatingnumber)
	fmt.Printf(".2f:%.2f\n", floatingnumber)
	fmt.Printf(".5.25f:%.5.25\n", floatingnumber)     //not possible
	fmt.Printf(".4f:%.4f\n", floatingnumber)
}*/

/*
// pointer//
package main

import "fmt"

func main() {
	fmt.Printf("%s*s\n", 40, "text")
	fmt.Printf("%040s", "text") //text is of 4 letter hence it will print 36 spaces and 4 spaces for text//
}
*/

// program to get the ASCII value of a character//
/*package main

import "fmt"

func main() {
	ch := 'A'
	bch := "B"
	fmt.Printf("Char value is: %c", ch)
	fmt.Printf("\nString value is: %s", bch)
}
*/
/*
package main

import "fmt"

func main() {
	var val byte = 'A'
	fmt.Printf("char value: %c", val)
	fmt.Printf("\nASCII value: %d", val)
}
*/

/*
//ABSOLUTE Value of float number

package main

import (
	"fmt"
	"math"
)

func main() {
	var val float64 = -19.35

	fmt.Printf("Absolute value of %f is %f", val, math.Abs(val))  //value must be of float type in Abs function//
}
*/

/*
// largest number between two numbers//
package main

import (
	"fmt"
	"math"
)

func main() {
	var num1 float64 = 10.25
	var num2 float64 = 20.14
	var large float64 = math.Max(num1, num2)  //or var large float64 = 0
	large = math.Max(num1, num2)
	fmt.Printf("Largest number is : %f", large)
}
*/

/*
//to calculate min betw two numbers

package main

import (
	"fmt"
	"math"
)

func main() {
	var num1 float64 = 10.25
	var num2 float64 = 20.14
	var small float64 = math.Max(num1, num2) //or var small float64 = 0
	small = math.Min(num1, num2)
	fmt.Printf("smallest number is : %f", small)
}
*/

/*
// power of a number//
package main

import (
	"fmt"
	"math"
)

func main() {
	var num float64 = 10.25
	var p float64 = 20.14
	var result float64 = 0
	result = math.Pow(num, p)

	fmt.Printf("Result : %f", result)
}

*/

/*
//math.Ceil() function

package main

import (
	"fmt"
	"math"
)

func main() {
	var num1 float64 = 10.25
	var num2 float64 = 20.14

	var result float64 = 0
	result = math.Ceil(num1)
	fmt.Printf("result is: %f", result)
	result = math.Ceil(num2)
	fmt.Printf("\nresult is : %f", result)
}
*/

/*
//math.floor function

package main

import (
	"fmt"
	"math"
)

func main() {
	var num1 float64 = 10.25
	var num2 float64 = 20.14
	var result float64 = 0
	result = math.Floor(num1)
	fmt.Printf("result is : %f", result)
	result = math.Floor(num2)
	fmt.Printf("\n result is : %f", result)
}
*/

/*
// printing without package
package main

func main() {

	println("using println instead of fmt.Println")  //if both th statements needed to be concatenated then print hsould be used in both the statements.
	print("using pritn instead of fmt.Print")
}
*/

//operators:
//arithmetic      relational        logical         bitwise      assignment    miscellaneous //

//assignment
/*
package main

import "fmt"

func main() {

	var x = 10
	x += 3
	fmt.Println(x)
}
*/

/*
package main

import (
	"fmt"
	"math"
)

func main() {
	a := 10.4
	b := 2.3
	fmt.Println(math.Mod(a, b))
}
*/

/*
package main

import "fmt"

func main() {
	a := 5
	b := 8
	fmt.Println(a % b)
}*/

/*
// OR AND FUNCTION
package main

import "fmt"

func main() {
	a := 5
	b := 3
	c := a | b
	fmt.Println(c)
}
*/

/*
//XOR//
package main

import "fmt"

func main() {
	a := 5
	b := 3
	c := a ^ b
	fmt.Println(c)
}
*/

//BITWISE SHIFTING//  Left shift-->value inc   right shift---> value decreases

/*
// comparison/relational operators
package main

import "fmt"

func main() {
	a := 10
	fmt.Println(a > 5)
	fmt.Println(a == 10)
	fmt.Println(a < 5)
}
*/

/*
package main

import "fmt"

func main() {
	a := 1010
	b := 0110
	fmt.Println(a && b)
	fmt.Println(a || b)
	fmt.Println(!a && b)
}*/

/*
// sizeof() operator
package main

import (
	"fmt"
	"unsafe"
)

func main() {
	var num1 int = 10
	var num2 byte = 20

	fmt.Printf("Size of num1: %d", unsafe.Sizeof(num1))
	fmt.Printf("\nSize of num2: %d", unsafe.Sizeof(num2))
}
*/

/*
package main

import (
	"fmt"
	"unsafe"
)

func main() {
	var num1 float64 = 10.2
	var num2 byte = 20

	fmt.Printf("Size of num1: %d", unsafe.Sizeof(num1))
	fmt.Printf("\nSize of num2: %d", unsafe.Sizeof(num2))
}
*/
/*
package main

import (
	"fmt"
	"unsafe"
)

func main() {
	var num1 string = "rajesh"
	var num2 byte = 20

	fmt.Printf("Size of num1: %d", unsafe.Sizeof(num1))
	fmt.Printf("\nSize of num2: %d", unsafe.Sizeof(num2))
}
*/

/*
// find the variable type using reflect.TypefOf() func
package main

import (
	"fmt"
	"reflect"
)

func main() {
	a := 10
	b := 10.20
	c := "hello"
	d := true
	fmt.Println("a = ", a)
	fmt.Println("b = ", b)
	fmt.Println("c = ", c)
	fmt.Println("d = ", d)

	fmt.Println("type of a= ", reflect.TypeOf(a))
	fmt.Println("type of b= ", reflect.TypeOf(b))
	fmt.Println("type of c= ", reflect.TypeOf(c))
	fmt.Println("type of d= ", reflect.TypeOf(d))

}
*/

/*
// reflect.ValueOf().Kind() function
package main

import (
	"fmt"
	"reflect"
)

func main() {
	a := 10
	b := 10.20
	c := "hello"
	d := true

	fmt.Println("a = ", a)
	fmt.Println("b = ", b)
	fmt.Println("c = ", c)
	fmt.Println("d = ", d)

	fmt.Println("type of a= ", reflect.ValueOf(a).Kind())
	fmt.Println("type of b= ", reflect.ValueOf(b).Kind())
	fmt.Println("type of c= ", reflect.ValueOf(c).Kind())
	fmt.Println("type of d= ", reflect.ValueOf(d).Kind())

}
*/

/*
//defer keyword  ---->  reverse order me print krega!!!

package main

import (
	"fmt"
)

func main() {

	defer fmt.Println("hello")
	defer fmt.Println("hello")
	fmt.Println("good morning")

}
*/

/*
//area of circle

package main

import (
	"fmt"
)

func main() {
	var radius float32
	var area float32
	const PI = 3.14

	fmt.Printf("enter radius: ")
	fmt.Scanf("%f", &radius)

	area = PI * radius * radius

	fmt.Printf("Area of circle is : %f", area)          //%.2f %g dono use kr sakte//

}
*/

/*
//perimeter of circle

package main

import (
	"fmt"
)

func main() {
	var radius float32
	var perimeter float32

	fmt.Printf("enter radius: ")
	fmt.Scanf("%f", &radius)

	perimeter = 2 * 3.14 * radius

	fmt.Printf("perimeter of circle is : %f", perimeter) //%.2f %g dono use kr sakte//

}
*/

/*
//fahrenheit to celsius

package main

import (
	"fmt"
)

func main() {
	var ftemp float32
	var ctemp float32

	fmt.Printf("enter temperature in Farenheit: ")
	fmt.Scanf("%f", &ftemp)
	ctemp = (ftemp - 32) / 1.8

	fmt.Printf("temperature in celsius is : %f", ctemp) //%.2f %g dono use kr sakte//

}
*/

/*
//celsius to fahrenheit

package main

import (
	"fmt"
)

func main() {
	var ftemp float32
	var ctemp float32

	fmt.Printf("enter temperature in celsius: ")
	fmt.Scanf("%f", &ctemp)
	ftemp = (ctemp * 1.8) + 32

	fmt.Printf("temperature in celsius is : %f", ftemp) //%.2f %g dono use kr sakte//

}
*/
