// structure syntax:  type person struct{ --------}  collection of diff data types
/*
//accessing structure members
package main

import "fmt"

type Books struct {
	title   string
	author  string
	subject string
	book_id int
}

func main() {
	var Book1 Books
	var Book2 Books

	Book1.title = "go program"
	Book1.author = "mahesh"
	Book1.subject = "tutorial"
	Book1.book_id = 64589

	Book2.title = "telecom"
	Book2.author = "sara"
	Book2.subject = "telecom tutorial"
	Book2.book_id = 345589

	fmt.Printf("Book 1 title : %s\n", Book1.title)
	fmt.Printf("Book 1 author : %s\n", Book1.author)
	fmt.Printf("Book 1 subject : %s\n", Book1.subject)
	fmt.Printf("Book 1 book_id : %d\n", Book1.book_id)

	fmt.Printf("Book 2 title : %s\n", Book2.title)
	fmt.Printf("Book 2 author : %s\n", Book2.author)
	fmt.Printf("Book 2 subject : %s\n", Book2.subject)
	fmt.Printf("Book 2 book_id : %d\n", Book2.book_id)

}
*/

/*
package main

import "fmt"

type Books struct {
	title   string
	author  string
	subject string
	book_id int
}

func main() {
	var Book1 Books
	var Book2 Books

	Book1.title = "go program"
	Book1.author = "mahesh"
	Book1.subject = "tutorial"
	Book1.book_id = 64589

	Book2.title = "telecom"
	Book2.author = "sara"
	Book2.subject = "telecom tutorial"
	Book2.book_id = 345589

	printBook(Book1)
	printBook(Book2)
}

func printBook(book Books) {
	fmt.Printf("Book title : %s\n", book.title)
	fmt.Printf("Book author : %s\n", book.author)
	fmt.Printf("Book subject : %s\n", book.subject)
	fmt.Printf("Book book_id : %d\n", book.book_id)
}*/

/*
package main

import "fmt"

type Rectangle func(int, int) int
type rectanglePara struct {
	length  int
	breadth int
	color   string
	rect    Rectangle
}

func main() {
	result := rectanglePara{
		length:  10,
		breadth: 20,
		color:   "red",
		rect: func(length int, breadth int) int {
			return length * breadth
		},
	}
	fmt.Println("colour :", result.color)
	fmt.Println("area:", result.rect(result.length, result.breadth))
}
*/

// slice
/*
package main

import "fmt"

func main() {
	numbers := []int{1, 2, 3, 4, 5}
	fmt.Println("Numbers: ", numbers)
}*/
/*
package main

import "fmt"

func main() {
	var numbers = make([]int, 5, 6)
	printSlice(numbers)
}
func printSlice(x []int) {
	fmt.Printf("len=%d cap=%d slice=%v\n", len(x), cap(x), x)
}
*/

package main

func main() {

}
