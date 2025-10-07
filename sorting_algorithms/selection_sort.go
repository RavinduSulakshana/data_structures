package sorting_algorithms

import "fmt"

func Selection_Sort(arr *[]int) *[]int {
	n := len(*arr)
	for i := 0; i < n-1; i++ {
		min_index := i
		for j := i + 1; j < n; j++ {
			if (*arr)[min_index] > (*arr)[j] {
				min_index = j
			}
		}
		(*arr)[i], (*arr)[min_index] = (*arr)[min_index], (*arr)[i]
	}
	fmt.Printf("Selection sort \n")
	return arr
}
