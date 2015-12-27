package main

import (
       "fmt"
       "os"
       "bufio"
)

/*
 * Print lines that appear more than once.
 */

func main () {
     counts := make(map[string]int)
     input := bufio.NewScanner(os.Stdin)
     
     for input.Scan() {
     	 counts[input.Text()]++
     }

     // NOTE ignore errors from input.Err()

     for line, n := range counts {
     	 if n > 1 {
	    fmt.Printf("%d\t%s\n", n, line)
	 }
     }
}