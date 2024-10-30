package main

func binarySearch(list []int, item, low, righ int) *int {
	for low < righ {
		middle := int((low + righ) / 2)
		guess := list[middle]
		if guess == item {
			return &middle
		}
		if item < guess {
			righ = middle - 1
		} else {
			low = middle + 1
		}
	}
	return nil
}

func exponentialSearch(list []int, item int) *int {
	right := 0
	if list[right] == item {
		return &right
	}
	n := len(list)
	right = 1
	for right < n && list[right] < item {
		right = right * 2
	}
	if list[right] == item {
		return &right
	}
	return binarySearch(list, item, int(right/2), right)
}

//func main() {
//	ar1 := []int{1, 2, 3, 4, 5}
//	println(*exponentialSearch(ar1, 4))
//}
