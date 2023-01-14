package main

import "fmt"

func main() {
	sequence := []int{2, 3, 4}
	count := 0

	for i := 0; i < len(sequence); i++ {
		if sequence[i]%2 != 0 {
			count++
		}
	}
	fmt.Println("Number of odd numbers:", count)
}
