package main

import "fmt"

func towersOfHanoi(n, from, to, aux int) {
	if n == 1 {
		fmt.Printf("Move disk 1 from peg %d to peg %d\n", from, to)
		return
	}

	towersOfHanoi(n - 1, from, aux, to)
	fmt.Printf("Move disk %d from peg %d to peg %d\n", n, from, to)

	towersOfHanoi(n - 1, aux, to, from)
}

func main() {
	towersOfHanoi(3, 1, 3, 2)
}
