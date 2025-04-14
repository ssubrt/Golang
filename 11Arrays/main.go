package main

import "fmt"

func main(){
	fmt.Println("Welcome to arrays of golang");

	var fruitlist [] string = []string{"apple", "banana", "mango"}
	fmt.Println("Fruit list is : ",fruitlist);
	// Accessing elements of array
	fmt.Println("Fruit list 1 is : ",fruitlist[0]);
	fmt.Println("Length of fruit list is : ",len(fruitlist));

	var fruitlist2 [4] string = [4]string{"apple", "banana", "mango", "orange"}
	fmt.Println("Fruit list 2 is : ",fruitlist2);

	// fruitlist[4] = "kiwi";
	fruitlist[1] = "grapes";
	fmt.Println("Fruit list is : ",fruitlist);

	var vegetablelist [4] string;
	vegetablelist[0] = "potato";
	vegetablelist[1] = "onion";	
	vegetablelist[2] = "tomato";

	fmt.Println("Vegetable list is : ",vegetablelist);
	fmt.Println("Length of vegetable list is : ",len(vegetablelist));
}