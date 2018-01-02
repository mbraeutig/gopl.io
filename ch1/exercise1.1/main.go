// Copyright © 2016 Alan A. A. Donovan & Brian W. Kernighan.
// License: https://creativecommons.org/licenses/by-nc-sa/4.0/

//Echo1 prints its command-line arguments.
package main

import (
	"fmt"
	"os"
)

func main() {
	fmt.Printf("Command:%s\n", os.Args[0])

	var s, sep string

	for i := 1; i < len(os.Args); i++ {
		s += sep + os.Args[i]
		sep = " "
	}
	fmt.Println(s)
}
