package main

import "fmt"

type Describer interface {
	Describe()
}

type People struct {
	name string
	age  int
}

func (p People) Describe() {
	fmt.Printf("name%s,age%d \n", p.name, p.age)
}

func main() {
	var d1 Describer = People{name: "000", age: 1}
	d1.Describe()
	p1 := People{name: "111", age: 1}
	d1 = p1
	d1.Describe()
	p2 := People{name: "222", age: 1}
	d1 = p2
	d1.Describe()
}
