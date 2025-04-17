package main

import "fmt"


// Note:- Defer is used to delay the execution of a function until the surrounding function returns.
// NOTE :- Defers folows LIFO order
func main() {
	defer fmt.Println("Welcome to Defer in golang")
	fmt.Println("Hello World");

	defer fmt.Println("Hello ")
	defer fmt.Println("world")
	defer fmt.Println("3");
	fmt.Println("Hi");
	mydefer()
}

func mydefer(){
	defer fmt.Println("Checking defer in golang")
	for i:=0; i<5 ; i++{
		defer fmt.Println(i); // LIFO order
	}
	defer fmt.Println("End of mydefer function")
}