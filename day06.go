/*
//io.WriteString() function

package main

import (
	"fmt"
	"io"
	"os"
)

func main() {
	var dd int = 17
	var mm int = 04
	var yy int = 2021
	var str string
	str = fmt.Sprintf("%02d-%02d-%04d", dd, mm, yy)
	io.WriteString(os.Stdout.str)

}*/

/*
// fmt.Scan()
package main

import "fmt"

func main() {
	var str string
	var a int
	var b float:float64

	fmt.Printf("enter the values: ")
	fmt.Scan(&str,&a,&float)

	fmt.Printf("result: %s\n", )
}
*/

/*
// uint type cant be stored in int type
package main

import "fmt"

func main() {
	var a uint = 123
	var b int = 0
	b = a
	fmt.Printf("a=%d,b=%d\n", a, b)
}
*/

/*
//the above issue can be resolved by using b=uint(a)//
package main

import "fmt"

func main() {
	var a int = 123
	var b uint = 0
	b = uint(a)
	fmt.Printf("a=%d,b=%d\n", a, b)
}
*/
/*
package main

import "fmt"

func main() {
	var a int = -123
	var b uint = 0
	b = uint(a)
	fmt.Printf("a=%d,b=%d\n", a, b)
}
*/
/*
// square root of a number//
package main

import (
	"math"
)

func main() {
	var a float64
	var b float64
	b = math.Sqrt(a)
	print(b)

}
*/

/*
package main

import (

	"math"

)

	func main() {
		var x int = 225
		var r float32                                               //not possible//
		r = math.Sqrt(float64(x))
		r = float32(math.Sqrt(float64(x)))

}
*/

/*
//correct code of square root //
package main

import (
	"fmt"
	"math"
)

func main() {
	var x int = 225
	var r float64
	r = math.Sqrt(float64(x))
	fmt.Printf("square root of %d is %.2f\n", x, r)
}*/

// EXPLICIT TYPE CONVERSION//
package main

import "fmt"

func main() {
	var x int = 42
	var y float64 = float64(x)
	var z uint = uint(y)

	fmt.Printf("value of x is %d and type is %T\n", x, x)
	fmt.Printf("value of y is %.2f and type is %T\n", y, y)
	fmt.Printf("value of z is %d and type is %T\n", z, z)
}
