package main

import "fmt"

func BinarySearch(arr []int, value int) int {
	low := 0
	high := len(arr) - 1

	for low <= high {
		mid := (low + high) / 2
		if arr[mid] == value {
			return mid
		}
		if arr[mid] < value {
			low = mid + 1
		}
		if arr[mid] > value {
			high = mid - 1
		}
	}
	return -1
}

func RecursiveBinarySearch(arr []int, value int, low int, high int) int {
	if low > high {
		return -1
	}

	mid := (low + high) / 2

	if arr[mid] == value {
		return mid
	}
	if arr[mid] < value {
		return RecursiveBinarySearch(arr, value, mid+1, high)
	}
	if arr[mid] > value {
		return RecursiveBinarySearch(arr, value, low, mid-1)
	}
	return -1
}

func main() {
	arr := []int{1, 2, 3, 4, 10, 12, 17, 28, 39, 60}
	pos1 := BinarySearch(arr, 12)
	pos2 := RecursiveBinarySearch(arr, 17, 0, len(arr)-1)
	fmt.Println(pos1)
	fmt.Println(pos2)
}
