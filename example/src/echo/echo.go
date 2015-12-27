package main

import (
       "fmt"
       "os"
)

func main () {
	var s, sep string

	for i := 1; i < len(os.Args); i++ {
		s += sep + os.Args[i]
		sep = " "
	}

	fmt.Print("For loop ... ")
	fmt.Println(s)

	s = ""
	sep = ""

	for _, arg := range os.Args[1:] {
		s += sep + arg
		sep = " "
	}

	fmt.Print("Range ... ")
	fmt.Println(s)
}