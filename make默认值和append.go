package quick_sort_test

import "fmt"

func main() {
	s := make([]int, 5)
	s = append(s, 1, 2, 3)
	fmt.Println(s)
	b := make([]int, 0)
	b = append(b, 1, 2, 3)
	fmt.Println(b)//[1 2 3]
}
