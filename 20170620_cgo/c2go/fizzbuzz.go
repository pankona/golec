package main

import (
	"C"
	"fmt"
	"strconv"
)

//export fizzbuzz
func fizzbuzz(n int) *C.char {
	retStr := ""
	if n%5 == 0 && n%3 == 0 {
		retStr = fmt.Sprintf("FizzBuzz")
	} else if n%5 == 0 {
		retStr = fmt.Sprintf("Buzz")
	} else if n%3 == 0 {
		retStr = fmt.Sprintf("Fizz")
	} else {
		retStr = fmt.Sprintf("%s", strconv.Itoa(n))
	}
	return C.CString(retStr)
}

func main() {}
