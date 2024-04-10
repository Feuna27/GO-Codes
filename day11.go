/*
//golang code to check given year is a leap year or not

package main
import "fmt"
func main() {
	var year int=0

	fmt.Printf("enter number: ")
	fmt.Scanf("%d",&year)

	if ((year%4==0 && year%100!=0) || (year%4 ==0)&&
}

*/

/*
//nested if else statements(else if)

package main

import "fmt"

func main() {
	x := 8
	y := 10

	if x != y {
		if x < y {
			fmt.Println("x is less than y")
		} else if x > y {
			fmt.Println("x is graeter than y")
		}
	} else {
		fmt.Println("x is equal to y")
	}
}
*/
/*
package main

import "fmt"

func main() {
	var time int = 0
	fmt.Printf("enter time: ")
	fmt.Scanf("%d", &time)

	if time < 10 {
		fmt.Println("good morning")
	} else if 10 < time && time < 20 {
		fmt.Println("good day")
	} else {
		fmt.Println("good evening")
	}

}*/
/*  switch cases
package main

import "fmt"

func main() {
	thisMonth := 5
	switch thisMonth {
	case 1:
		fmt.Println("january")
	case 2:
		fmt.Println("feb")
	case 3:
		fmt.Println("march")
	case 4:
		fmt.Println("april")
	case 5:
		fmt.Println("may")
	}
}*/

/*
package main

import (
	"fmt"
	"time"
)

func main() {
	today := time.Now()
	switch today.Day() {
	case 2:
		fmt.Println("today is 2:go clean ur house")
	case 10:
		fmt.Println("today is 10:go clean ur car")
	default:
		fmt.Println("no info")
	}
}*/

/*
//fallthrough  --> transfers the execution to the next case

package main

import "fmt"

func main() {
	x := 3
	switch x {
	case 1:
		fmt.Println("1")
		fallthrough
	case 3:
		fmt.Println("3")
		fallthrough
	case 5:
		fmt.Println("5")
	}
}*/

/*
package main

import "fmt"

func main() {
	var x interface{}
	switch i := x.(type) {
	case nil:
		fmt.Printf("type of x : %T", i)
	case int:
		fmt.Printf("x is int")
	case float64:
		fmt.Printf("x is float64")
	case func(int) float64:
		fmt.Printf("x is func(int)")
	case bool, string:
		fmt.Printf("x is bool or string")
	default:
		fmt.Printf("dont know")
	}
}*/

/*
//goto : transfer control to another part of code within func//
package main

import "fmt"

func main() {
	i := 0
loop:
	fmt.Println(i)
	i++
	if i < 5 {
		goto loop
	}
	fmt.Println("loop ends")
}*/
/*
package main

import (
	"fmt"
	"strings"
)

func main() {
	var str string = "hello"
	var substr string = "wor"

	if strings.Contains(str, substr) {
		fmt.Printf("String(%s) contains sub-string(%s)", str, substr)
	} else {
		fmt.Printf("string (%s) does not contain substring (%s)", str, substr)
	}
}*/

// uppercase
package main

import (
	"fmt"
	"strings"
)

func main() {
	var str string = "hello world"
	var result string
	var ind int = 0

	result = strings.ToUpper(str)
	fmt.Println("string in uppercase : ", result)
	result = strings.ToLower(str)
	fmt.Println("string in lowercase : ", result)
	ind = strings.Index(str, "w")
	fmt.Println("index is : ", ind)
}
