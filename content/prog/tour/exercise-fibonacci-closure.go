// +build OMIT

package main

import "fmt"

// fibonacci est une fonction qui renvoie
// une fonction retournant un int.
func fibonacci() func() int {
}

func main() {
	f := fibonacci()
	for i := 0; i < 10; i++ {
		fmt.Println(f())
	}
}
