package main

import "fmt"

func main() {
	fmt.Println("Welcome to Pointers of golang");

	// var ptr *int  ;
	// fmt.Println("Value of pointer is : ",ptr);
	// fmt.Println("Address of pointer is : ",&ptr);
	// fmt.Println("Value of pointer is : ",*ptr);

	// var mynum *int = 2;
	mynum := 23;
	 var ptr1 = &mynum;
	 fmt.Println("Value of pointer is : ",ptr1);
	 fmt.Println("Value of pointer is : ",*ptr1);
	 fmt.Println("Value of pointer is : ",&ptr1);
	 fmt.Println("NEw Value of pointer is : ",*ptr1 + 2);

}