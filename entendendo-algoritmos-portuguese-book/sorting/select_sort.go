package sorting

// searchMinor returns the index of the lowest value in the list
func searchMinor(list []int) int {
	low := list[0]
	lowIndex := 0
	for i := 1; i < len(list); i++ {
		if list[i] < low {
			low = list[i]
			lowIndex = i
		}
	}
	return lowIndex
}

// SelectSort teratively finds the smallest element in the list, appends it to a new list, and removes it from the original list until sorted.
func SelectSort(list []int) []int {
	new := make([]int, len(list))
	copy(new, list)
	for i := range new {
		lowIndex := searchMinor(new[i:])
		new[i], new[i+lowIndex] = new[i+lowIndex], new[i]
	}
	return new
}
