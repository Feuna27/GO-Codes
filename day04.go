/*package main

import (
	"fmt"
)

func main() {
	var name string = "Feuna"
	var age int = 21
	var branch string = "CYSE"
	var colname string = "RCOEM"
	fmt.Println("Name: ", name, "   Age: ", age, "    Branch: ", branch, "   Collegename: ", colname)
                                                // new line ke liye \nBranch//
}*/

/*
package main

import "fmt"

func main() {
	var a string = "feuna"
	fmt.Println("My name is", "'", a, "'")
	fmt.Printf("my name is '%s' ", a)
}
*/

/*
package main

import "fmt"

func main() {
	const (
		name = "GFG"
		dept = "CS"
	)
	fmt.Printf("%s is a %s portal", name, dept)
}*/

//to print scientific nottation int float decimal value
/*package main

import "fmt"

func main() {
	var str = "GFG"
	fmt.Printf("the string is %s", str)
	var num1 int = 20
	fmt.Printf("the decimal value is %d \n", num1)
	var num2 float32 = 7.7
	fmt.Printf("the floating point is %g \n", num2)               //%v can be use as it is universal format specifier//
	var num3 int = 14
	fmt.Printf("the binary value is %b \n", num3)
	var num4 = 4 + 1i
	fmt.Printf("Scientific notation of num4 : %e \n", num4)
}
*/
// l values and R values jo left aur right side me likh sakte hai     literals are always r values

/*
// HEX FORMAT                   //0O17-->OCTAL        //0X--> HEXADECIMAL

package main

import "fmt"

func main() {
	var a int = 21
	var b int = 15
	println(15 == 017) //true
	println(15 == 0xF) //true
	a++
	fmt.Printf("Value of a is %d\n", a)
                                //   a++ cant be written at the place of a it should be written above the printf function//
	a--
	fmt.Printf("Value of a is %d\n", a)

	b++
	fmt.Printf("Value of a is %d\n", b)

	b--                                                       //--b would not work
	fmt.Printf("Value of a is %d\n", b)

}
*/

/*
//binary number//

package main

import (
	"fmt"
)

func main() {
	var a int = 15
	fmt.Printf("binary %0b ", a)
}
*/

//format specifiers  v--default format, d--decimal, g--floating point number,b--binary, o-octal, t-for boolean, s--string//
/*
package main

import (
	"fmt"
)

func main() {
	var mydata float32 = 3.142
	fmt.Printf("the floating point number is %g", mydata)
}
*/
/*
package main

import (
	"fmt"
)

func main() {
	var i, j string = "hii", "feuna"
	fmt.Print(i, "\n") //fmt.Printf("%s\n%s",i,j)
	fmt.Print(j, "\n")
	fmt.Printf("%s %s", i, j)
}
*/

/*
package main

import (
	"fmt"
)

func main() {
	var i = 15.5
	var txt = "hello world"
	fmt.Printf("%v\n", i)
	fmt.Printf("%#v\n", i)
	fmt.Printf("%v%%\n", i)  //agar % me value chahiye tab //
	fmt.Printf("%T\n", i)    //type//
	fmt.Printf("%v\n", txt)  //string//
	fmt.Printf("%#v\n", txt) //"string"//
	fmt.Printf("%T\n", txt)
	fmt.Printf("'%s", txt)
}
*/
/*
package main

import (
	"fmt"
)

func main() {
	var i = 15
	fmt.Printf("%b\n", i)
	fmt.Printf("%d\n", i) //"%q\n" ---> to print in quotes
	fmt.Printf("%+d\n", i)
	fmt.Printf("%o\n", i)
	fmt.Printf("%O\n", i)
	fmt.Printf("%x\n", i)
	fmt.Printf("%X\n", i)
	fmt.Printf("%#x\n", i)
	fmt.Printf("%4d\n", i)
	fmt.Printf("%-4d\n", i) //space chodne ke liye - sign
	fmt.Printf("%04d\n", i)
}
*/
