package main

import (
	"fmt"
	"log"
	"os"
	"strconv"
)

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

	fmt.Printf("input = %d\n", n)

	// ...

}
