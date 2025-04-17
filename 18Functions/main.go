package main

import "fmt"

func main() {
	greeter();
	fmt.Println("Welcome to Functions in golang");

	result := adder(2,3);

	fmt.Println("Result is : ", result);

	// fmt.Println("Result is : ", proadder(2,3,4,5));

	proRes , myMsg := proadder(2,3,4,5);
	fmt.Println("Result is : ", proRes);
	fmt.Println(myMsg);


	//NOTE:WE can't do this 
	// func trying(){
	// 	fmt.Println("Trying to create a function inside a function")
	// }
}

func greeter(){
	fmt.Println("Hello, World!")
}


func adder(a int , b int) int {
	return a+b;
}

func proadder(values ...int) (int, string) {
	total := 0;

	for _,val := range values {
		total += val;
	}

	return total , "ye bi hota hai bhaiya";
}