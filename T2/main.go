package main

import (
	"fmt"
	"os"
)

func readArray() []int {
	var n int
	fmt.Fprint(os.Stdout, "Capture the number of elements to content in the array: ")
	fmt.Fscan(os.Stdin, &n)
	arr := make([]int, n)
	for i := 0; i < n; i++ {
		fmt.Fprint(os.Stdout, "Capture the element number ", i+1, " for the array: ")
		fmt.Fscan(os.Stdin, &arr[i])
	}
	return arr
}
func oddNumbers(readArray []int) {
	maxCounter := 0
	for i := 0; i <= len(readArray)-1; i++ {
		if readArray[i]%2 != 0 {
			maxCounter++
		}
	}
	fmt.Println("number of odd numbers are: ", maxCounter)
}
func main() {
	newArr := readArray()
	oddNumbers(newArr)

}
