package main

import "fmt"

func insertionSort(numbers []int) {
	for i := 1; i < len(numbers); i++ {
		orderNum := numbers[i]
		j := i - 1
		for j >= 0 && numbers[j] > orderNum {
			numbers[j+1] = numbers[j]
			j = j - 1
		}
		numbers[j+1] = orderNum
	}
}
func main() {
	arreglo := []int{5, 2, 9, 1, 3, 7}
	insertionSort(arreglo)
	fmt.Println(arreglo)
}
