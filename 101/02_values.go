package main

import "fmt"

func main(){
	fmt.Println("hello" + "world" + " !!!") // "helloworld !!!"

	fmt.Println("1+1=", 1+1) // "1+1= 2"
	fmt.Println("7.0/3.0=", 7.0/3.0) // "2.33.."
	fmt.Println(true && false) // "false"
	fmt.Println(true || false) // "true"
	fmt.Println(!true) // "false"
	fmt.Println(!!true && !true)
}