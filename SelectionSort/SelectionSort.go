package main

import "fmt"

func selectionSort(arr []int) {
	n := len(arr)

	for i := 0; i < n-1; i++ {
		minInd := i

		for j := i + 1; j < n; j++ {
			if arr[j] < arr[minInd] {
				minInd = j
			}
		}

		if minInd != i {
			arr[i], arr[minInd] = arr[minInd], arr[i]
		}
	}
}

func main() {
	arr := []int{64, 25, 12, 22, 11}
	fmt.Println("Array original:", arr)

	selectionSort(arr)

	fmt.Println("Array ordenado:", arr)
}
