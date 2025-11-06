package main

import "fmt"

func main() {
	arr := []int{1, 4, 6, 9, 54}
	item := 34
	result := BinarySearch(arr, item)
	if result == -1 {
		fmt.Println("Value not present")
	} else {
		fmt.Println(result)
	}

}

func BinarySearch(arr []int, item int) int {
	low := 0
	high := len(arr) - 1

	for low <= high {
		mid := (low + high) / 2
		guess := arr[mid]
		if guess == item {
			return mid
		}
		if guess < item {
			low = mid + 1
		} else {
			high = mid - 1
		}
	}
	return -1
}
