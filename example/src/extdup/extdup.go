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
     files := os.Args[1:]
     
     if len(files) == 0 {
     	countLines(os.Stdin, counts)
     } else {
       for _, arg := range files {
       	   f, err := os.Open(arg)
	   if err != nil {
	      fmt.Fprintf(os.Stderr, "extdup: %v\n", err)
	      continue
	   }
	   countLines(f, counts)
	   f.Close()
	}
     }

     for line, n := range counts {
	    if n > 1 {
	    	    fmt.Printf("%d\t%s\n", n, line)
	    }
     }
}

func countLines(f *os.File, counts map[string]int) {
     input := bufio.NewScanner(f)
     for input.Scan() {
     	 counts[input.Text()]++
     }

     // NOTE: ignoring errors from input.Err()
}
