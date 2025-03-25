package main

import (
	"fmt"
	"neetcode/arrayhashing"
)

func main() {
	// result := arrayhashing.SequenceBruteForce([]int{2, 20, 4, 10, 3, 4, 5})
	// fmt.Println("result is ", result)
	// result := arrayhashing.SequenceHashmap([]int{2, 20, 4, 10, 3, 4, 5})
	result := arrayhashing.SequenceSet([]int{2, 20, 4, 10, 3, 4, 5})
	fmt.Println("result is ", result)

}
