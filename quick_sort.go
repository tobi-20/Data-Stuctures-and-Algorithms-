package main

import "fmt"

func main() {
	fmt.Println(QuickSort([]int{4, 5, 2, 4, 6, 1, 9, 8, 4}))

}

func QuickSort(arr []int) []int {
	if len(arr) < 2 { // length less than 2 means it can be 1 or 0 : [i] or [](empty array )
		return arr
	}
	pivot := arr[0] // pivot :=  arr[len(arr) / 2]is better  used for a sorted or nearly sorted list(further reading)

	less := []int{}
	greater := []int{}
	for _, v := range arr[1:] { // for _, v := range append(arr[:len(arr)/2], arr[(len(arr)/2 + 1):]...) works better for a sorted or nearly sorted list(further reading)
		if v <= pivot {
			less = append(less, v)

		} else {
			greater = append(greater, v)
		}

	}
	return append(append(QuickSort(less), pivot), QuickSort(greater)...) // QuickSort(less)+ [pivot] + QuickSort(greater)
	//there is a better memory efficient manner of doing implementing this algorithm because each recursion creates new slices again and again imagine if len(arr)==1000, but this is the level I am as of this day
}

// //QuickSort Algorithm Outline
// Start with an unsorted list

// Example: B C A E F G

// Check the base case:If the list has 0 or 1 elements, return the list as-is.

// Pick a pivot element

// The pivot can be any element of the list.

// Keep it aside for now.

// Divide the remaining elements into two parts

// Less than or equal to the pivot: goes into the first part.

// Greater than the pivot: goes into the second part.

// The two parts do not have to be equal in size.

// Sort the two parts individually

// Apply the same algorithm recursively on each part.

// Concatenate results

// Merge the sorted first part + pivot + sorted second part.

// The first part is all elements less than the pivot, and the second part is all elements greater than the pivot.

// Return the concatenated list

// This gives a fully sorted list.
