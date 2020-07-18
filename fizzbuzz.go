package main

import "fmt"
import "os"
import "strconv"

func fizzbuzz(i int){
	fizz := 3
	buzz := 5
	for n := 1; n <= i; n++{
		if n % fizz == 0 && n % buzz == 0 {
			fmt.Println("FizzBuzz")
		} else if n % fizz == 0 {
			fmt.Println("Fizz")
		} else if n % buzz == 0{
			fmt.Println("Buzz")
		} else {
			fmt.Println(n)
		}
	}
}

func main(){
	arg := os.Args[1]
	n, err := strconv.Atoi(arg)
	if err == nil {
		fizzbuzz(n)
	}
}