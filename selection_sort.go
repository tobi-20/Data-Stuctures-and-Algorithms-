package main

import "fmt"

func main() {
	specimen := []int{1, 4, 5, 3, 2, 3, 4, 5, 7, 8}
	fmt.Println(SelectionSort(specimen))
}
func FindSmallest(arr []int) int {
	smallest := arr[0]
	smallestIndex := 0

	for i, v := range arr {
		if v < smallest {
			smallest = v
			smallestIndex = i
		}

	}
	return smallestIndex
}

func SelectionSort(arr []int) []int {
	copiedArr := make([]int, len(arr))
	copy(copiedArr, arr)
	newArr := []int{}
	for len(copiedArr) > 0 {
		smallest := FindSmallest(copiedArr)
		newArr = append(newArr, copiedArr[smallest])
		copiedArr = append(copiedArr[:smallest], copiedArr[smallest+1:]...)
	}

	return newArr
}
