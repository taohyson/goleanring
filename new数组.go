package main

import "fmt"

func main1() {
	list := make([]int,0)//new 不能用slice
	list = append(list, 1)
	fmt.Println(list)
}

func main() {
	s1 := []int{1, 2, 3}
	//s2 := []int{4, 5}
	//s1 = append(s1, s2)//append
	fmt.Println(s1)
}
