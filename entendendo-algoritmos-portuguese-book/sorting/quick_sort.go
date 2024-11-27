package sorting

func QuickSort(list []int) []int {
	if len(list) < 2 {
		return list
	}
	pivot := len(list) - 1
	var lowers []int
	var highers []int
	for i := range list {
		if i == pivot {
			continue
		}
		if list[i] < list[pivot] {
			lowers = append(lowers, list[i])
		} else {
			highers = append(highers, list[i])
		}
	}
	result := QuickSort(lowers)
	result = append(result, list[pivot])
	return append(result, QuickSort(highers)...)
}
