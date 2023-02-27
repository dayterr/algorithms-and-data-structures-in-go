package main

import "fmt"

func generateBinary(k int, arr *[]int) {
	if k < 1 {
		fmt.Println(*arr)
		return
	}
	(*arr)[k-1] = 0
	generateBinary(k - 1, arr)
	(*arr)[k-1] = 1
	generateBinary(k - 1, arr)
}

func main() {
	generateBinary(3, &[]int{0, 0, 0})
	generateBinary(5, &[]int{0, 0, 0, 0, 0})
}
