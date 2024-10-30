package sorting

// searchMinor returns the index of the lowest value in the list
func searchMinor(list []int) *int {
	low := list[0]
	lowIndex := 0
	for i := 1; i < len(list); i++ {
		if list[i] < low {
			low = list[i]
			lowIndex = i
		}
	}
	return &lowIndex
}

// SelectSort teratively finds the smallest element in the list, appends it to a new list, and removes it from the original list until sorted. 
func SelectSort(list []int) []int {
	new := make([]int, 0, len(list))
	for range list {
		lowIndex := searchMinor(list)
		new = append(new, list[*lowIndex])
		list = append(list[:*lowIndex], list[*lowIndex+1:]...)
	}
	return new
}

