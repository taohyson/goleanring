package quick_sort_test

import (
	"fmt"
)

var G  = 7

func main() {
	y := func() int {
		fmt.Printf("G: %d, G的地址:%p\n", G, &G)
		G += 1
		return G
	}
	fmt.Println(y(), y)
	fmt.Println(y(), y)
	fmt.Println(y(), y) //y的地址

	// 影响全局变量G，注意z的匿名函数是直接执行，所以结果不变
	z := func() int {
		fmt.Printf("G2: %d, G2的地址:%p\n", G, &G)
		G += 1
		return G
	}()
	fmt.Printf("%T%d", z, &z)
	fmt.Println(z, &z)
	fmt.Println(z, &z)
}


