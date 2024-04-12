/*package main

import "fmt"

func main() {
	primeNumbers := []int{2, 3, 5, 7}
	numbers := []int{1, 2, 3, 6, 8}

	copy(numbers, primeNumbers)
	fmt.Println("Numbers:", numbers)
	fmt.Println("length:", len(numbers))
	//for loop which iterates through slice
	for i := 0; i < len(numbers); i++ {
		fmt.Println(numbers[i])
	}
}
*/
/*
package main

import "fmt"

func main() {
	arr := [10]int{10, 20, 30, 40, 50, 60, 70, 80, 90, 100}
	intslice := arr[2:5]

	fmt.Println("slice elements: ")
	for index, ele := range intslice {
		fmt.Printf("index=%d,element=%d\n", index, ele)
	}

}
*/

/*
//sort striungs of slice

package main

import (
	"fmt"
	"sort"
)

func main() {
	slice := []string{"honesty", "is", "the", "best", "policy"}
	sort.Strings(slice)

	fmt.Println("sorted slice: ")
	for _, item := range slice {
		fmt.Printf("%s \t", item)
	}
}*/

/*
package main

import "fmt"

func main() {

	subjectMarks := map[string]float32{"golang": 85, "java": 80, "python": 81}
	fmt.Println(subjectMarks)
}*/
/*
package main

import "fmt"

func main() {
	flowercolor := map[string]string{"sunflower": "yellow", "jasmine": "white", "hibiscus": "red"}

	fmt.Println(flowercolor["lily"])
}*/

/*
package main

import "fmt"

func main() {
	capital := map[string]string{"nepal": "kathmandu", "US": "New york"}
	fmt.Println("initial map:", capital)

	capital["US"] = "Washington DC"
	fmt.Println("updated:", capital)
}
*/

/*
package main

import "fmt"

func main() {
	students := map[int]string{1: "john", 2: "lily"}
	fmt.Println("initial map:", students)

	students[3] = "robin"
	students[5] = "julie"
	fmt.Println("updated:", students)
}*/

/*
// delete element
package main

import "fmt"

func main() {
	percentage := map[string]string{"nepal": "kathmandu", "US": "New york"}
	fmt.Println("initial map:", percentage)

	delete(percentage, "nepal")
	fmt.Println("updated:", percentage)
}*/

package main

import "fmt"

func main() {
	squarednumber := map[int]int{2: 4, 3: 5, 4: 9, 5: 20}
	for number, squared := range squarednumber {
		fmt.Printf("square of %d is %d\n", number, squared)
	}
}
