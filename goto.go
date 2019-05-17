package main

func main() {
loop:
	for i := 0; i < 10; i++ {
		println(i)
	}
	goto loop
}
//goto不能跳转到其他函数或者内层代码