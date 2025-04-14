package main

import (
	"fmt"
	"math/rand"
	"time"
)


func main(){
	fmt.Println("Welcome to switch case in golang");

	var num int = 10;

	switch num {
	case 1:
		fmt.Println("Number is 1");
	case 6:
		fmt.Println("Number is 6");	
    case 10:
		fmt.Println("Number is 10");
	} 

	//NOTE :- Old method of random number generation
	rand.Seed(time.Now().UnixNano());
	diceNum := rand.Intn(6) + 1;  // here rando numer genarates bwtween 0 to 5 and then + 1 making the range 0 to 6
	fmt.Println("Dice number is : ", diceNum);


	// NOTE :- New method  of random number generation

	randSrc := rand.NewSource(time.Now().UnixNano())
	randGen := rand.New(randSrc)
		
	diceNumber := randGen.Intn(6) + 1;
	fmt.Println("Dice number is : ", diceNumber);


	switch diceNum {
	case 1:
		fmt.Println("You rolled a 1!")
	case 2:
		fmt.Println("You rolled a 2!")
	case 3:
		fmt.Println("You rolled a 3!");
		fallthrough // It is like continue keyword in c++ but in this it includes current value also and rint next value also
	case 4:
		fmt.Println("You rolled a 4!")
	case 5:
		fmt.Println("You rolled a 5!")
	case 6:
		fmt.Println("You rolled a 6!")
	default:
		fmt.Println("Invalid roll!")
	}

	switch diceNumber {
	case 1:
		fmt.Println("You rolled a 1!")
	case 2:
		fmt.Println("You rolled a 2!")
	case 3:
		fmt.Println("You rolled a 3!")
	case 4:
		fmt.Println("You rolled a 4!")
	case 5:
		fmt.Println("You rolled a 5!")
	case 6:
		fmt.Println("You rolled a 6!")
	default:
		fmt.Println("Invalid roll!")
	}
		
		
}