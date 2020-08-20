package main

import (
	"fmt"
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

	// // arrays
	// var numbers = [4]int{1, 2, 3, 4}

	// fmt.Println(numbers)
	// fmt.Println(len(numbers))
	// fmt.Println(numbers[2])

	// // slices, it's an abstraction built on top of Go's array type.
	// // Create a slice by not adding array length
	// var numbers = []int{1, 2, 7, 9, 3, 4}
	// fmt.Println(numbers)

	// // append the slice
	// numbers = append(numbers, 5)
	// fmt.Println(numbers)

	// // slice the slice LOL
	// numbers = append(numbers[2:4])
	// fmt.Println(numbers)

	// // make allocates and initialises an object, slice, map, array
	// letters := make([]string, 3, 3)

	// letters[0] = "A"
	// letters[1] = "B"
	// letters[2] = "C"

	// letters = append(letters, "D")

	// fmt.Println(letters)
	// // to get the total capicity of the slice
	// fmt.Println(cap(letters))

	// //sort arrays
	// isSorted := sort.IntsAreSorted(numbers)
	// fmt.Println(isSorted)
	// sort.Ints(numbers)
	// fmt.Println(numbers)

	//maps
	// new => allocates and does not initialize a memory
	scores := make(map[string]int)
	scores["rex"] = 50
	scores["rex2"] = 51
	scores["jon"] = 22
	fmt.Println(scores)
	fmt.Println(scores["rex2"])

	// delete a key and value pair
	delete(scores, "rex2")
	fmt.Println(scores)

	// iterate over the map using range
	for k, v := range scores {
		fmt.Printf("Score if %v is %v \n", k, v)
	}

}
