package main

import (
	"errors"
	"fmt"
)

var errIncorrectN = errors.New("n should be equal to the length of the array")

func checkIfArrSorted(arr []int, n int) (bool, error) {
	if n == 1 {
		return true, nil
	}

	if n > len(arr) {
		return false, errIncorrectN
	}

	if arr[n - 1] < arr[n - 2] {
		return false, nil
	}
	return checkIfArrSorted(arr, n - 1)
}

func main() {
	fmt.Println(checkIfArrSorted([]int{1, 3, 5}, 3))
	fmt.Println(checkIfArrSorted([]int{1, 5, 3}, 3))
	fmt.Println(checkIfArrSorted([]int{1}, 1))
}
