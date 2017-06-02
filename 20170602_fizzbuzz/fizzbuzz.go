package main

import (
	"fmt"
	"log"
	"os"
	"strconv"
)

func fizzbuzz(n int) string {
	if n%5 == 0 && n%3 == 0 {
		return fmt.Sprintf("FizzBuzz")
	} else if n%5 == 0 {
		return fmt.Sprintf("Buzz")
	} else if n%3 == 0 {
		return fmt.Sprintf("Fizz")
	}
	return fmt.Sprintf("%s", strconv.Itoa(n))
}

func main() {
	if len(os.Args) < 2 {
		log.Println("not enough arguments")
		os.Exit(1)
	}
	n, err := strconv.Atoi(os.Args[1])
	if err != nil {
		log.Println(err)
		os.Exit(1)
	}

	for i := 1; i <= n; i++ {
		fmt.Println(fizzbuzz(i))
	}
}
