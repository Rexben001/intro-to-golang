package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
	"time"
)

func main() {

	// // := it guesses the dataype and use it
	// score := 40
	// fmt.Println(score)

	// // %v gets the variable
	// // %T gets the variable datatype
	// fmt.Printf("%v, %T", score, score)
	// fmt.Println()

	// // create two variable the same time
	// var name, age string = "Ben", "20"
	// fmt.Println(name, age)

	// // create standalone variable together
	// var (
	// 	job     = "techician"
	// 	country = "US"
	// )

	// fmt.Println(job, country)

	// // get user inputs default
	// var myString string
	// fmt.Scanln(&myString)
	// fmt.Println(myString)

	//use bufio to read inputs
	reader := bufio.NewReader(os.Stdin)
	fmt.Print("Enter your full name: ")
	myname, _ := reader.ReadString('\n')

	readertwo := bufio.NewReader(os.Stdin)
	fmt.Print("Enter your full rating: ")
	rating, _ := readertwo.ReadString('\n')
	// convert string to number and remove extra space
	ratingTwo, _ := strconv.ParseFloat(strings.TrimSpace(rating), 64)

	// handle
	fmt.Printf("Hello %v, \n we have received your rating %v. \n Thanks. \n\n Rating was recorded at %v", myname, ratingTwo, time.Now().Format(time.Stamp))
	fmt.Println()

	if ratingTwo == 4 {
		fmt.Println("Wait now")
	} else if ratingTwo > 4 {
		fmt.Println("Opo juuu")
	} else {
		fmt.Println("Thanks for banking with us")
	}

}
