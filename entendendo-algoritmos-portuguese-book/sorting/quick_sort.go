package sorting

func QuickSort(list []int) []int {
	if len(list) < 2 {
		return list
	}
	pivot := len(list) - 1
	lowers := []int{}
	for i := range list[1:] {
		if list[i] <= pivot {
			lowers = append(lowers, pivot)
		}
	}
	righers := []int{}
	for j := range list[1:] {
		if list[j] > pivot {
			righers = append(righers, pivot)
		}
	}
	result := QuickSort(lowers)
	result = append(result, pivot)
	return append(result, QuickSort(righers)...)
}
