package main

//import "fmt"

// sum sums all elements of the list recursively and return
func sum(list []int) int {
	if len(list) == 0 {
		return 0
	}
	return list[0] + sum(list[1:])
}

//func main() {
//	fmt.Println(sum([]int{2, 4, 6}))
//}
