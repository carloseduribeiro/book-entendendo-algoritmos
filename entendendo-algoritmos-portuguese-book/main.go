package main

import (
	"entendendo-algoritmos/sorting"
	"fmt"
)

func main() {
	list := []int{6, 1, 3, 5, 2, 5}
	fmt.Printf("select sort result:\t%v\n", sorting.SelectSort(list))
	fmt.Printf("quick sort result:\t%v\n", sorting.QuickSort(list))
}
