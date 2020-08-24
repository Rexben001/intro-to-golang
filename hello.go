package main

import (
	"fmt"
	"os"
	"strconv"
	"time"
	"errors"
)

// // create a struct
type User struct {
	Name  string
	Email string
	Age   int
}

// create functions

func multiply(a int, b int) int {
	fmt.Printf("The multiplication of %v and %v = %v \n", a, b, a*b)
	return a * b
}

func loopAdd(values ...int) (int, int, string) {
	sum := 0
	for i := range values {
		sum = sum + values[i]
	}
	length := len(values)
	name := "Got here"
	return sum, length, name
}

// methods
func (u User) CallUser() string {
	s := fmt.Sprintf("My name is %s, I'm %s years of age, you can reach trhough %s ", u.Name, strconv.Itoa(u.Age), u.Email)
	// s := []string{"My name is", u.Name, "I'm", u.Age, "years of age\n"}
	return s
}

//push data to channel with a 4 second delay
func data1(ch chan string) {
	time.Sleep(4 * time.Second)
	ch <- "from data1()"
}

//push data to channel with a 2 second delay
func data2(ch chan string) {
	time.Sleep(2 * time.Second)
	ch <- "from data2()"
}

//function accepts a filename and tries to open it.
func fileopen(name string) {
	f, er := os.Open(name)

	//er will be nil if the file exists else it returns an error object
	if er != nil {
		fmt.Println(er)
		return
	} else {
		fmt.Println("file opened", f.Name())
	}
}

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

	// //maps
	// // new => allocates and does not initialize a memory
	// scores := make(map[string]int)
	// scores["rex"] = 50
	// scores["rex2"] = 51
	// scores["jon"] = 22
	// fmt.Println(scores)
	// fmt.Println(scores["rex2"])

	// // delete a key and value pair
	// delete(scores, "rex2")
	// fmt.Println(scores)

	// // iterate over the map using range
	// for k, v := range scores {
	// 	fmt.Printf("Score if %v is %v \n", k, v)
	// }

	// //structs
	// tob := User{"rob", "rob@vv.cc", 77}
	// // print out the values
	// fmt.Printf("%v\n", tob)
	// // print out the values and keys
	// fmt.Printf("%+v\n", tob)

	// var sam = new(User)
	// sam.Name = "sam"
	// fmt.Printf("%+v\n", sam)

	// //loops
	// var numbers = []int{1, 2, 7, 9, 3, 4}
	// for i := 0; i < len(numbers); i++ {
	// 	fmt.Println(numbers[i])
	// }
	// // use range
	// for i := range numbers {
	// 	fmt.Println(numbers[i])
	// }

	// fmt.Println(multiply(2, 3))

	// res, leng, name := loopAdd(2, 3, 4, 6, 78)
	// fmt.Println(res, leng, name)

	// // defer() => used to defer the execution of a function call until the function that contains the defer statement completes execution
	// defer multiply(3, 4)
	// fmt.Println("Got here first")

	// // stacking defer => multiple defer statements, LIFO, starts from bottom to top
	// defer multiply(30, 4)
	// defer multiply(300, 4)
	// defer multiply(3000, 4)
	// fmt.Println("Got here first")

	// // methods
	// user := User{"Ben", "rex@bmm.co", 30}
	// fmt.Println(user.CallUser())

	// // Concurrency is achieved in Go using Goroutines and Channels.
	// // Goroutine is a function which can run concurrently with other functions

	// // the for loop didn't wait for multiply to complete execution before it started
	// go multiply(3, 4)
	// for i := 0; i < 5; i++ {
	// 	fmt.Println("In main")
	// }

	//channels
	//the main() will wait on the statement receiving channel data until it gets the data
	// To create a channel 	ch := make(chan int)
	// To send data to a channel     ch <- x
	// To receive data from a channel    y := <- ch

	// 	//Select => can be viewed as a switch statement which works on channels. Here the case statements will be a channel operation.//creating channel variables for transporting string values
	// 	chan1 := make(chan string)
	// 	chan2 := make(chan string)

	// 	//invoking the subroutines with channel variables
	// 	go data1(chan1)
	// 	go data2(chan2)

	// 	//Both case statements wait for data in the chan1 or chan2.
	// 	//chan2 gets data first since the delay is only 2 sec in data2().
	// 	//So the second case will execute and exits the select block
	// 	select {
	// 	case x := <-chan1:
	// 		fmt.Println(x)
	// 	case y := <-chan2:
	// 		fmt.Println(y)
	// 	}

	// Mutex is the short form for mutual exclusion.
	// Mutex is used when you don't want to allow a resource to be accessed by multiple subroutines at the same time.
	// Mutex has 2 methods - Lock and Unlock. Mutex is contained in sync package.

	// Error Handling
	fileopen("invalid.txt")

}
