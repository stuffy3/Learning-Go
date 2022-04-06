package main

import (
	"bufio"
	"fmt"
	"os"
)
// prints the text of each line that appears more than once in the standard input
//preceded by its count
func main () {
	counts := make(map[string]int)
	input := bufio.NewScanner(os.Stdin)
	for input.Scan() {
		counts[input.Text()]++
	}
	for line, n := range counts {
		if n < 1 {
			fmt.Printf("%d\t%s\n", n, line)
		}
	}
}

// func main() {
// 	myProgram := os.Args[0]

// 	fmt.Println(os.Args[1:1], myProgram)
// }


// func main() {
// 	s, sep := "", ""
// 	for _, arg := range os.Args[1:] {
// 		s += sep + arg
// 		sep = " "
// 	}
// 	fmt.Println(s)
// }

//prints Hello
// func main() {
// 	fmt.Println("Hello")
// }