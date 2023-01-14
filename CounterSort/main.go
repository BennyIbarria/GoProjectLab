package main

import (
	"fmt"
	"os"
)


func readArray() []int {
	var n int
	fmt.Fprint(os.Stdout, "Capture the number of elements in the array: ")
	fmt.Fscan(os.Stdin, &n)

	arr := make([]int, n)
	for i := 0; i < n; i++ {
		fmt.Fprint(os.Stdout, "Enter the element number ", i+1, " to the array: ")
		fmt.Fscan(os.Stdin, &arr[i])
	}
	return arr
}
func countingSort(array []int) {
	maximunLenght := 0
	for _, v := range array {
		if v > maximunLenght {
			maximunLenght = v
		}
	}
	counts := make([]int, maximunLenght+1)
	for _, v := range array {
		counts[v]++
	}
	index := 0
	for i, c := range counts {
		for j := 0; j < c; j++ {
			array[index] = i
			index++
		}
	}
}
func main() {
	arrayEglo := readArray()
	countingSort(arrayEglo)
	fmt.Println("Sorted array: ", arrayEglo)
}
