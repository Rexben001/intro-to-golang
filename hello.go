package main

import "fmt"

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

	// //use bufio to read inputs
	// reader := bufio.NewReader(os.Stdin)
	// fmt.Print("Enter your full name: ")
	// myname, _ := reader.ReadString('\n')

	// readertwo := bufio.NewReader(os.Stdin)
	// fmt.Print("Enter your full rating: ")
	// rating, _ := readertwo.ReadString('\n')
	// // convert string to number and remove extra space
	// ratingTwo, _ := strconv.ParseFloat(strings.TrimSpace(rating), 64)

	// // handle conditionals
	// fmt.Printf("Hello %v, \n we have received your rating %v. \n Thanks. \n\n Rating was recorded at %v", myname, ratingTwo, time.Now().Format(time.Stamp))
	// fmt.Println()

	// if ratingTwo == 4 {
	// 	fmt.Println("Wait now")
	// } else if ratingTwo > 4 {
	// 	fmt.Println("Opo juuu")
	// } else {
	// 	fmt.Println("Thanks for banking with us")
	// }

	// // pointers
	// var total float64 = 33.2

	// // & reference to the memory location
	// totalRef := &total

	// fmt.Println(total, totalRef)

	// // * gets the value instead of the memory reference
	// fmt.Println(*totalRef)

	// arrays
	var numbers = [4]int{1, 2, 3, 4}

	fmt.Println(numbers)
	fmt.Println(len(numbers))
	fmt.Println(numbers[2])

}
