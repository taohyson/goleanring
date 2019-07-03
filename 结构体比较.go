package quick_sort_test

import "fmt"

func main() {
	sn := struct {
		age  int
		name string
	}{age: 1, name: "a"}

	sn2 := struct {
		age  int
		name string
	}{age: 1, name: "a"}
	if sn == sn2 {
		fmt.Println("sn==sn2")
	}
	sm1 := struct {
		age int
		m   map[string]string
	}{age: 11, m: map[string]string{"a": "1"}}
	sm2 := struct {
		age int
		m   map[string]string
	}{age: 11, m: map[string]string{"a": "1"}}

	//但是结构体属性中有不可以比较的类型，如map,slice。
	if sm1 == sm2 {
		fmt.Println("sm1 == sm2")
	}
}
