package main

import (
       "fmt"
       "os"
       "io/ioutil"
       "strings"
)

/*
 * Print lines that appear more than once.
 */

func main () {
     counts := make(map[string]int)
     for _, filename := range os.Args[1:] {
     	 data, err := ioutil.ReadFile(filename)

	 if err != nil {
	    fmt.Fprintf(os.Stderr, "filedup: %v\n", err)
	    continue
	 }
	 for _, line := range strings.Split(string(data), "\n") {
	     counts[line]++
	 }
     }

     for line, n := range counts {
     	 if n > 1 {
	    fmt.Printf("%d\t%s\n", n, line)
	 }
     }
}