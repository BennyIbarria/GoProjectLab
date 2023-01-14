package main

import (
	"reflect"
	"testing"
)

func TestInsertionSort(t *testing.T) {
	numbers := []int{5, 2, 9, 1, 3, 7}
	insertionSort(numbers)
	expected := []int{1, 2, 3, 5, 7, 9}
	if !reflect.DeepEqual(numbers, expected) {
		t.Errorf("Expected %v but got %v", expected, numbers)
	}
}
