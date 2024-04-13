/*package main

import "fmt"

func main() {
	student := make(map[int]string)

	student[1] = "harry"
	student[2] = "sejal"
	student[3] = "shrey"

	fmt.Println(student)
}
*/
/*
package main

import "fmt"

func main() {
	places := map[string]string("nepal":"kathmandu", "US":"Washington", "Norway":"uolo")
	for country:= range places{

	fmt.Println(country)
	}

}*/

/*
//map of using structure
package main

import "fmt"

type vertex struct {
	Lat, Long float64
}

var m = map[string]vertex{
	"bell labs": vertex{
		40.68, -74.399,
	},
	"google": vertex{
		37.5, -122.08,
	},
	"nagpur": vertex{
		45.5, 72.5,
	},
}
func main() {
	fmt.Println(m)
}
var m["nagpur"]=vertex{45.5,72.5}
fmt.Println(m)
*/
/*
//pointer and address
package main

import "fmt"

func main() {
	var intData = 20
	var intPointer *int

	intPointer = &intData

	fmt.Println("what intData stores:", intData)
	fmt.Println("address of intdata:", &intData)
	fmt.Println("what intpointer stores:", intPointer)

	*intPointer = 10

	fmt.Println("what intData stores:", intData)

}*/
/*
package main

import "fmt"

func swap(x int, y int) (int, int) {
	return y, x

}
func main() {
	var a int = 10
	var b int = 20
	a, b = swap(a, b)
	fmt.Println(a, b)
}*/
/*
package main

import "fmt"

func swap(x *int, y *int) {
	t := *x
	*x = *y
	*y = t

}
func main() {
	a, b := 10, 20
	swap(&a, &b)
	fmt.Println(a, b)
}*/

// go proram that maked an http get request to an API
package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
)

func main() {
	apiURL := "https://api.example.com/data"

	response, err := http.Get(apiURL)
	if err != nil {
		fmt.Printf("error making get request: %s\n", err)
		return
	}
	defer response.Body.Close()

	body.err := ioutil.ReadAll(response.Body)
	if err != nil {
		fmt.Printf("error reading response body: %s\n", err)
		return
	}
	fmt.Println(string(body))
}
