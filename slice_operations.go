package main

import (
	"fmt"
)

func main() {
	// Example slice for testing
	numbers := []int{3, 1, 4, 1, 5, 9, 2, 6}

	// Test FindMax
	max := FindMax(numbers)
	fmt.Printf("Maximum value: %d\n", max)

	// Test RemoveDuplicates
	unique := RemoveDuplicates(numbers)
	fmt.Printf("After removing duplicates: %v\n", unique)

	// Test ReverseSlice
	reversed := ReverseSlice(numbers)
	fmt.Printf("Reversed: %v\n", reversed)

	// Test FilterEven
	evenOnly := FilterEven(numbers)
	fmt.Printf("Even numbers only: %v\n", evenOnly)
}

// FindMax returns the maximum value in a slice of integers.
// If the slice is empty, it returns 0.
func FindMax(numbers []int) int {
	// TODO: Implement this function
	if len(numbers) == 0 {
		return 0
	}
	max := numbers[0]
	for i := range numbers {
		if max < numbers[i] {
			max = numbers[i]
		}
	}
	return max
}

// RemoveDuplicates returns a new slice with duplicate values removed,
// preserving the original order of elements.
func RemoveDuplicates(numbers []int) []int {
	// TODO: Implement this function
	nums := append([]int(nil), numbers...)

    for i := 0; i < len(nums); i++ {
        for j := i + 1; j < len(nums)-1; j++{
            if nums[i] == nums[j] {
                nums = append(nums[:j], nums[j+1:]...)
            }
        }
    }
    return nums
}

// ReverseSlice returns a new slice with elements in reverse order.
func ReverseSlice(slice []int) []int {
	// TODO: Implement this function
	nums := []int{}
	for i := range slice {
		nums = append(nums, slice[len(slice)-1-i])
	}
	return nums
}

// FilterEven returns a new slice containing only the even numbers
// from the original slice.
func FilterEven(numbers []int) []int {
	// TODO: Implement this function
	nums := []int{}
	for i:= range numbers {
	    if (numbers[i]%2 == 0) {
	        nums = append(nums, numbers[i])
	    }
	}
	return nums
}
