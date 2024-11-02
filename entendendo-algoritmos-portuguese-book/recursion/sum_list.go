package recursion

// SumListWithoutRecursion sums a list an return it result
// Time: O(n) | Space: O(1)
func SumListWithoutRecursion(list []int) int {
	total := 0
	for _, v := range list {
		total += v
	}
	return 0
}

// SumListWithRecursion sums a list recursively and return it result
// Time: O(n) | Space: O(n)
func SumListWithRecursion(list []int) int {
	if len(list) < 1 {
		return 0
	}
	if len(list) < 2 {
		return list[0]
	}
	return list[0] + SumListWithRecursion(list[1:])
}

// CountListItems sums the number of items recursively in a list and return it
func CountListItems(list []int) int {
	result := 1
	defer func() {
		// recovers when the slice is empty
		if r := recover(); r != nil {
			result = 0
		}
	}()
	return result + CountListItems(list[1:])
}

// Max finds the greater value in the list recursively and returns it
func Max(list []int) int {
	if len(list) < 1 {
		return 0
	}
	if len(list) < 2 {
		return list[0]
	}
	if v := Max(list[1:]); list[0] < v {
		return v
	}
	return list[0]
}
