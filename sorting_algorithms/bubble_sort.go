package sorting_algorithms

import ()

func Bubble_Sort(arr []int) *[]int {
	for k := 0; k < len(arr); k++ {
		for i := 0; i < len(arr)-1; i++ {
			if arr[i] > arr[i+1] {
				num := arr[i]
				arr[i] = arr[i+1]
				arr[i+1] = num
			}
		}
	}
	return &arr
}
