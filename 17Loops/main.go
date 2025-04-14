package main

import "fmt"


func main() {
	fmt.Println("Welcome to Loops in golang");

	days := []string{"Monday", "Tuesday", "Wednesday", "Thursday", "Friday", "Saturday", "Sunday"};

	fmt.Println("Days are : ",days);

	// 1st Loop :=  for loop
	for d:=0; d<len(days); d++{
		fmt.Println("Day is : ",days[d]);
	}

	// 2nd Loop :=  for loop with range
	for i := range days {
		fmt.Println("Day is : ",days[i]);
	}

	// 3rd Loop :=  for loop with range and value or for Each loop
	for index, day := range days{
		fmt.Printf("Day %v  and index is : %v\n", day, index);
	}


	// 4th Loop :=  Kind of While loop
	roughValue := 1

	for roughValue < 10 {
		if roughValue == 2 {
			goto lco;
		} else if roughValue == 5 {
			roughValue++; // This is needed othervise loops just stop here not ended
			continue;
		} else if roughValue == 7 {
			break;
		}
		fmt.Println("Rough value is : ", roughValue);
		roughValue++;
	}


	lco: // label for loop
	fmt.Println("This is label") 
	



}