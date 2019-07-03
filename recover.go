

package quick_sort_test

import (
"fmt"
)
func div(a, b int) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Printf("捕获到异常：%s\n", r)
		}
	}()
	if b < 0 {
		panic("除数需要大于0")
	}
	fmt.Println("余数为：", a/b)

}
func fun3() (t int) {
	t = 5
	defer func() {
		t = t + 5
		fmt.Println(t)
	}()
	return t
}
func main() {
	i:=fun3()
	fmt.Println(i)
}
