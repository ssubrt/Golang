package main

import (
	"fmt"
	"sort"
)

func main(){
	fmt.Println("Welcome to slices of golang");

	var fruitsList = []string{"Yoyo"};
	fmt.Printf("Type of fruitsList is : %T\n",fruitsList);
	fruitsList = append(fruitsList, "apple", "banana", "mango");
	fmt.Println("Fruit list is : ",fruitsList);
	fruitsList = append(fruitsList[1:3]);
	fmt.Println("Fruit list is : ",fruitsList);

	highScores := make([]int,4);
	highScores[0] = 160;
	highScores[1] = 101;
	highScores[2] = 12;
	highScores[3] = 103;
	// highScores[4] = 104;
	highScores = append(highScores, 555,165);
	fmt.Println("High scores are : ",highScores);
	fmt.Println("sizes of  scores are : ",(len(highScores)));

	sort.Ints(highScores);
	fmt.Println("High scores are : ",highScores);

	// how to remove a value from slice based on index
	var courses = []string{"reactjs", "js", "python", "c++", "java"};
	fmt.Println("Courses are : ",courses);
	var index  int = 2;
	courses = append(courses[:index], courses[index+1:]...);
	fmt.Println("Courses are : ",courses);
}