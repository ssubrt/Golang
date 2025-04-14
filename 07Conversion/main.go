package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	fmt.Println("Welcome to ou pizza app");
	fmt.Println("Rate our pizza b/w 1 and 5"); // after this everything will come in new line
	fmt.Print("enter the value : " ); // after this everything will come in same line
	//fmt.Printf("enter the value : " ); // after this everything will come in same line


	reader := bufio.NewReader( os.Stdin);
	input, _  :=  reader.ReadString('\n');

	fmt.Println("Thanks for rating",input);

	numRating, err := strconv.ParseFloat(strings.TrimSpace(input),64);

	if err != nil {
		fmt.Println(err);
	} else{
		fmt.Println("Your rating is", numRating + 1);
	}
}