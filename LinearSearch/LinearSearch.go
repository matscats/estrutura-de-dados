package main

import "fmt"

func LinearSearch(arr []int, value int) (int, int) {
	for i := 0; i < len(arr); i++ {
		if arr[i] == value {
			return value, i
		}
	}
	return -1, -1
}

func main() {
	arr := []int{1, 2, 3, 4, 5, 6, 7, 8, 9}
	value, index := LinearSearch(arr, 11)
	fmt.Println("(valor, posição): ", value, ",", index)
}
