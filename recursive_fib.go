package main

import "fmt"
import "os"
import "strconv"

func main(){
	arg := os.Args[1]
	// os.Args[1] is first argument passed, second would be [2] and so on.
	n, err := strconv.Atoi(arg)
	// strconv.Atoi takes string and spits out int AND err. 
	// It also spits err, which HAS to be declared
	if err == nil && n <= 20 {
	// if Atoi works, err is nil (null or None equivalent). 
	// n is just a limiter so my CPU doesn't die trying a big n fib
		fib_string := ""
		for i := 0; i < n; i++{
			fib_string += strconv.Itoa(fib(i)) + " "
			// Itoa is just flipped Atoi, so it converts int into string
			// For some reason this does not care about err
			// Probably because int to a string is guaranteed to work? not sure.
			// This really isn't necessary, but I wanted it to display a nice string of numbers.
		}
		fmt.Println(fib_string)
	} else{
		fmt.Println("\nHold your horses, pick a NUMBER below 20\n")
	}
}

func fib(i int) int{

	if i <= 1 {
		return i
	} else {
		return fib(i-1)+ fib(i-2)
	}
	// classic recursive fib algo. 
}