package quick_sort_test

import (
	"fmt"
)

type People struct {
}

func (p *People) ShowA() {
	fmt.Printf("ShowA\n")
	p.ShowB()
}

func (p *People) ShowB() {
	fmt.Printf("ShowB\n")
}

type Teach struct {
	People
}

func (t * Teach) ShowC()  {
	fmt.Println("teacher showB")
}

func main() {
	p := Teach{}
	p.ShowC()
}
