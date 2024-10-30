package main

//import "fmt"

// BinarySearch receives an ordered slice and an item. If the item is in the slice, it position is returned. Otherwise it returns nil.
func BinarySearch(list []int, item int) *int {
	low := 0
	righ := len(list) - 1
	for low <= righ {
		middle := int((low + righ) / 2)
		guess := list[middle]
		if guess == item {
			return &middle
		}
		if guess > item {
			righ = middle - 1
		} else {
			low = middle + 1
		}
	}
	return nil
}

//func main() {
//	list := []int{1, 3, 5, 7, 9}
//	value1 := BinarySearch(list, 3)
//	value2 := BinarySearch(list, -1)
//	fmt.Printf("%d\n", *value1) // must return 1
//	fmt.Printf("%v\n", value2)  // must return nil
//}
