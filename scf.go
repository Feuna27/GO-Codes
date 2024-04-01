// user input using scanf
package main

import (
	"fmt"
)

func main() {

	var name string // var fname,lname string
	fmt.Scanln(&name)
	fmt.Printf("enter your name")
	//fmt.Scanf("%s", &name)//                  //fmt.Scanf("%s %s",&fname,&lname)
	fmt.Println(name)
}
